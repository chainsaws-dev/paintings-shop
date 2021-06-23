// Package settings - реализует модели данных для хранения настроек сервера и их частичного автозаполнения
package settings

import (
	"database/sql"
	"errors"
	"log"
	"net/http"
	"paintings-shop/packages/databases"
	"paintings-shop/packages/randompassword"
	"paintings-shop/packages/shared"
)

// Список типовых ошибок
var (
	ErrRoleNotFound         = errors.New("роль с указанным именем не найдена")
	ErrDatabaseNotSupported = errors.New("не реализована поддержка базы данных")
	ErrDatabaseOffline      = errors.New("база данных недоступна")
	ErrUsupportedDBType     = errors.New("указан неподдерживаемый тип базы данных")
)

// AutoFillRoles - автозаполняет список ролей для SQL сервера
func (SQLsrv *SQLServer) AutoFillRoles() {

	SQLsrv.Roles = SQLRoles{}

	SQLsrv.Roles = append(SQLsrv.Roles, SQLRole{
		Name:    "guest_role_read_only",
		Desc:    "Гостевая роль",
		Login:   "paintings_shop_guest",
		Pass:    randompassword.NewRandomPassword(20),
		TRules:  GetTRulesForGuest(),
		Default: true,
		Admin:   false,
	})

	SQLsrv.Roles = append(SQLsrv.Roles, SQLRole{
		Name:    "admin_role_CRUD",
		Desc:    "Администратор",
		Login:   "paintings_shop_admin",
		Pass:    randompassword.NewRandomPassword(20),
		TRules:  GetTRulesForAdmin(),
		Default: false,
		Admin:   true,
	})
}

// DropDatabase - автоматизировано удаляет базу и роли
func (SQLsrv *SQLServer) DropDatabase(donech chan bool) {
	switch {
	case SQLsrv.Type == "PostgreSQL":
		// Удаляем базу данных

		dbc, err := databases.PostgreSQLConnect(databases.PostgreSQLGetConnString(SQLsrv.Login, SQLsrv.Pass, SQLsrv.Addr, "", true))
		if err != nil {
			log.Fatalln(err)
		}

		databases.PostgreSQLDropDatabase(SQLsrv.DbName, dbc)

		for _, currole := range SQLsrv.Roles {

			databases.PostgreSQLDropRole(currole.Login, dbc)
		}

		dbc.Close()

		donech <- true

	default:
		log.Fatalln("Указан неподдерживаемый тип базы данных " + SQLsrv.Type)
	}
}

// CreateDatabase - Создаёт базу данных если её нет
func (SQLsrv *SQLServer) CreateDatabase(donech chan bool, CreateRoles bool) {
	switch {
	case SQLsrv.Type == "PostgreSQL":
		// Создаём базу данных
		cs := databases.PostgreSQLGetConnString(SQLsrv.Login, SQLsrv.Pass, SQLsrv.Addr, "", true)
		dbc, err := databases.PostgreSQLConnect(cs)
		if err != nil {
			log.Fatalln(err)
		}
		databases.PostgreSQLCreateDatabase(SQLsrv.DbName, dbc)
		dbc.Close()

		// Заполняем базу данных
		cs = databases.PostgreSQLGetConnString(SQLsrv.Login, SQLsrv.Pass, SQLsrv.Addr, SQLsrv.DbName, false)
		dbc, err = databases.PostgreSQLConnect(cs)
		if err != nil {
			log.Fatalln(err)
		}

		err = databases.PostgreSQLCreateTables(dbc)

		if err != nil {
			if errors.Is(databases.ErrTablesAlreadyExist, err) {
				donech <- false
				return
			}
		}

		placeholder := databases.File{
			FileName: "placeholder.jpg",
			FileSize: 0,
			FileType: "jpg",
			FileID:   "",
		}
		databases.PostgreSQLFileChange(placeholder, dbc)

		if CreateRoles {
			for _, currole := range SQLsrv.Roles {

				databases.PostgreSQLCreateRole(currole.Login, currole.Pass, SQLsrv.DbName, dbc)

				for _, tablerule := range currole.TRules {

					databases.PostgreSQLGrantRightsToRole(currole.Login, tablerule.TName, formRightsArray(tablerule), dbc)
				}
			}
		}

		dbc.Close()

		donech <- true

	default:
		log.Fatalln("Указан неподдерживаемый тип базы данных " + SQLsrv.Type)
	}
}

// FindRoleInRoles - Ищем роль в списке ролей по имени
func FindRoleInRoles(RoleName string, Roles SQLRoles) (SQLRole, error) {
	for _, si := range Roles {
		if si.Name == RoleName {
			return si, nil
		}
	}
	return SQLRole{}, ErrRoleNotFound
}

// GetConnectionString - Формируем строку соединения
func GetConnectionString(SQLsrv *SQLServer, Role string) (string, error) {

	ActiveRole, err := FindRoleInRoles(Role, SQLsrv.Roles)

	if err != nil {
		return "", err
	}

	return databases.PostgreSQLGetConnString(
		ActiveRole.Login,
		ActiveRole.Pass,
		SQLsrv.Addr,
		SQLsrv.DbName,
		false), nil
}

// Connect - открывает соединение с базой данных Postgresql
func (SQLsrv *SQLServer) Connect(w http.ResponseWriter, role string) *sql.DB {

	switch {
	case SQLsrv.Type == "PostgreSQL":
		cs, err := GetConnectionString(SQLsrv, role)

		if shared.HandleOtherError(w, "Роль не найдена", err, http.StatusServiceUnavailable) {
			return nil
		}

		dbc, err := databases.PostgreSQLConnect(cs)

		if shared.HandleOtherError(w, "База данных недоступна", err, http.StatusServiceUnavailable) {
			return nil
		}

		return dbc

	default:
		log.Fatalln("Указан неподдерживаемый тип базы данных " + SQLsrv.Type)
	}

	return nil
}

// ConnectAsAdmin - подключаемся к базе с ролью администратора
func (SQLsrv *SQLServer) ConnectAsAdmin() *sql.DB {
	switch {
	case SQLsrv.Type == "PostgreSQL":
		cs, err := GetConnectionString(SQLsrv, "admin_role_CRUD")

		if err != nil {
			log.Println(err)
			return nil
		}

		dbc, err := databases.PostgreSQLConnect(cs)

		if err != nil {
			log.Println(err)
			return nil
		}

		return dbc

	default:
		log.Fatalln("Указан неподдерживаемый тип базы данных " + SQLsrv.Type)
	}

	return nil
}

// ConnectAsGuest - подключаемся к базе с ролью гостя
func (SQLsrv *SQLServer) ConnectAsGuest() *sql.DB {
	switch {
	case SQLsrv.Type == "PostgreSQL":
		cs, err := GetConnectionString(SQLsrv, "guest_role_read_only")

		if err != nil {
			log.Println(err)
			return nil
		}

		dbc, err := databases.PostgreSQLConnect(cs)

		if err != nil {
			log.Println(err)
			return nil
		}

		return dbc

	default:
		log.Fatalln("Указан неподдерживаемый тип базы данных " + SQLsrv.Type)
	}

	return nil
}

// formRightsArray - формирует массив прав для таблицы
func formRightsArray(rule TRule) []string {
	var result []string

	if rule.SELECT {
		result = append(result, "SELECT")
	}

	if rule.INSERT {
		result = append(result, "INSERT")
	}

	if rule.UPDATE {
		result = append(result, "UPDATE")
	}

	if rule.DELETE {
		result = append(result, "DELETE")
	}

	if rule.REFERENCES {
		result = append(result, "REFERENCES")
	}

	return result
}

// CheckRoleForRead - проверяет роль для разрешения доступа к разделу системы
func (ss WServerSettings) CheckRoleForRead(RoleName string, AppPart string) bool {
	switch {
	case AppPart == "ArtworkTypes":
		return ss.CheckExistingRole(RoleName)
	case AppPart == "Authors":
		return ss.CheckExistingRole(RoleName)
	case AppPart == "Terms":
		return ss.CheckExistingRole(RoleName)
	case AppPart == "Currencies":
		return ss.CheckExistingRole(RoleName)
	case AppPart == "Countries":
		return ss.CheckExistingRole(RoleName)
	case AppPart == "CurrentUser":
		return ss.CheckExistingRole(RoleName)
	case AppPart == "CheckSecondFactor":
		return ss.CheckExistingRole(RoleName)
	case AppPart == "SecondFactor":
		return ss.CheckExistingRole(RoleName)
	case AppPart == "GetQRCode":
		return ss.CheckExistingRole(RoleName)
	case AppPart == "HandleRecipes":
		return ss.CheckExistingRole(RoleName)
	case AppPart == "HandleRecipesSearch":
		return ss.CheckExistingRole(RoleName)
	case AppPart == "HandleShoppingList":
		return ss.CheckExistingRole(RoleName)
	case AppPart == "HandleFiles":
		return ss.CheckExistingRole(RoleName)
	case AppPart == "HandleUsers":
		return checkAdmin(RoleName)
	case AppPart == "HandleSessions":
		return checkAdmin(RoleName)
	default:
		return false
	}
}

// CheckRoleForChange - проверяет роль для разрешения изменений в разделе системы
func (ss WServerSettings) CheckRoleForChange(RoleName string, AppPart string) bool {
	switch {
	case AppPart == "ArtworkTypes":
		return checkAdmin(RoleName)
	case AppPart == "Authors":
		return checkAdmin(RoleName)
	case AppPart == "Terms":
		return checkAdmin(RoleName)
	case AppPart == "Currencies":
		return checkAdmin(RoleName)
	case AppPart == "Countries":
		return checkAdmin(RoleName)
	case AppPart == "CurrentUser":
		return ss.CheckExistingRole(RoleName)
	case AppPart == "CheckSecondFactor":
		return ss.CheckExistingRole(RoleName)
	case AppPart == "SecondFactor":
		return ss.CheckExistingRole(RoleName)
	case AppPart == "GetQRCode":
		return ss.CheckExistingRole(RoleName)
	case AppPart == "HandleFiles":
		return checkAdmin(RoleName)
	case AppPart == "HandleUsers":
		return checkAdmin(RoleName)
	case AppPart == "HandleSessions":
		return checkAdmin(RoleName)
	default:
		return false
	}
}

// CheckRoleForDelete - проверяет роль для разрешения доступа к удалению элементов раздела системы
func (ss WServerSettings) CheckRoleForDelete(RoleName string, AppPart string) bool {
	switch {
	case AppPart == "ArtworkTypes":
		return checkAdmin(RoleName)
	case AppPart == "Authors":
		return checkAdmin(RoleName)
	case AppPart == "Terms":
		return checkAdmin(RoleName)
	case AppPart == "Currencies":
		return checkAdmin(RoleName)
	case AppPart == "Countries":
		return checkAdmin(RoleName)
	case AppPart == "CurrentUser":
		return ss.CheckExistingRole(RoleName)
	case AppPart == "CheckSecondFactor":
		return ss.CheckExistingRole(RoleName)
	case AppPart == "SecondFactor":
		return ss.CheckExistingRole(RoleName)
	case AppPart == "GetQRCode":
		return ss.CheckExistingRole(RoleName)
	case AppPart == "HandleFiles":
		return checkAdmin(RoleName)
	case AppPart == "HandleUsers":
		return RoleName == "admin_role_CRUD"
	case AppPart == "HandleSessions":
		return checkAdmin(RoleName)
	default:
		return false
	}
}

func checkAdmin(RoleName string) bool {
	return RoleName == "admin_role_CRUD"
}

// CheckExistingRole - проверяет что роль это существующая роль
func (ss WServerSettings) CheckExistingRole(RoleName string) bool {

	for _, role := range ss.SQL.Roles {
		if role.Name == RoleName {
			return true
		}
	}

	return false

}
