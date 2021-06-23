// Package references - содержит веб методы для обработки запросов к таблицам справочников
package references

import (
	"encoding/json"
	"errors"
	"net/http"
	"paintings-shop/packages/databases"
	"paintings-shop/packages/setup"
	"paintings-shop/packages/shared"
	"paintings-shop/packages/signinupout"
	"strconv"
)

// ArtworkTypes - обработчик для работы со справочником типов картин
//
// Аутентификация
//
//  Куки
//  Session - шифрованная сессия
//	Email - шифрованный электронный адрес пользователя
//
//  или
//
//	Заголовки:
//  Auth - Токен доступа
//
//	и
//
//	ApiKey - Постоянный ключ доступа к API *
//
// GET
//
//  Если нужен список элементов:
// 	ожидается заголовок Page с номером страницы
// 	ожидается заголовок Limit с максимумом элементов на странице
//
//	Если нужен один элемент:
//	ожидается заголовок ItemID с индексом элемента
//
// POST
//
// 	тело запроса должно быть заполнено JSON объектом
// 	идентичным по структуре Currency (см. файл models.go в пакете databases)
//
// DELETE
//
// 	ожидается заголовок ItemID с индексом элемента, который нужно удалить
func ArtworkTypes(w http.ResponseWriter, req *http.Request) {

	role, auth := signinupout.AuthGeneral(w, req)

	if !auth {
		return
	}

	var err error

	switch {
	case req.Method == http.MethodGet:

		if setup.ServerSettings.CheckRoleForRead(role, "ArtworkTypes") {

			PageStr := req.Header.Get("Page")
			LimitStr := req.Header.Get("Limit")
			ItemID := req.Header.Get("ItemID")

			// Создаём соединение с базой данных с ролью пользователя
			dbc := setup.ServerSettings.SQL.Connect(w, role)
			if dbc == nil {
				return
			}
			defer dbc.Close()

			if PageStr == "" && LimitStr == "" {

				if ItemID != "" {

					ID, err := strconv.Atoi(ItemID)

					if shared.HandleInternalServerError(w, err) {
						return
					}

					var ArtType databases.ArtworkType

					ArtType, err = databases.PostgreSQLSingleArtworkTypeSelect(ID, dbc)

					if err != nil {
						if errors.Is(err, databases.ErrArtTypeNotFound) {
							shared.HandleOtherError(w, err.Error(), err, http.StatusBadRequest)
							return
						}

						if shared.HandleInternalServerError(w, err) {
							return
						}
					}

					shared.WriteObjectToJSON(false, w, ArtType)

				} else {
					shared.HandleOtherError(w, shared.ErrHeadersNotFilled.Error(), shared.ErrHeadersNotFilled, http.StatusBadRequest)
					return
				}

			} else {

				Page, err := strconv.Atoi(PageStr)

				if shared.HandleInternalServerError(w, err) {
					return
				}

				Limit, err := strconv.Atoi(LimitStr)

				if shared.HandleInternalServerError(w, err) {
					return
				}

				var ArtTypes databases.ArtworkTypesResponse

				ArtTypes, err = databases.PostgreSQLArtworkTypesSelect(Page, Limit, dbc)

				if err != nil {
					if errors.Is(err, databases.ErrLimitOffsetInvalid) {
						shared.HandleOtherError(w, err.Error(), err, http.StatusBadRequest)
						return
					}

					if errors.Is(err, databases.ErrContryNotFound) {
						shared.HandleOtherError(w, err.Error(), err, http.StatusBadRequest)
						return
					}

					if shared.HandleInternalServerError(w, err) {
						return
					}
				}

				shared.WriteObjectToJSON(false, w, ArtTypes)

			}

		} else {
			shared.HandleOtherError(w, shared.ErrForbidden.Error(), shared.ErrForbidden, http.StatusForbidden)
		}

	case req.Method == http.MethodPost:

		if setup.ServerSettings.CheckRoleForChange(role, "ArtworkTypes") {

			// Читаем тело запроса в структуру
			var ArtType databases.ArtworkType

			err = json.NewDecoder(req.Body).Decode(&ArtType)

			if shared.HandleOtherError(w, "Invalid JSON in request body", err, http.StatusBadRequest) {
				return
			}

			// Создаём соединение с базой данных с ролью пользователя
			dbc := setup.ServerSettings.SQL.Connect(w, role)
			if dbc == nil {
				return
			}
			defer dbc.Close()

			ArtType, err = databases.PostgreSQLArtworkTypeChange(ArtType, dbc)

			if shared.HandleInternalServerError(w, err) {
				return
			}

			shared.WriteObjectToJSON(false, w, ArtType)

		} else {
			shared.HandleOtherError(w, shared.ErrForbidden.Error(), shared.ErrForbidden, http.StatusForbidden)
		}

	case req.Method == http.MethodDelete:

		if setup.ServerSettings.CheckRoleForDelete(role, "ArtworkTypes") {

			ItemID := req.Header.Get("ItemID")

			if ItemID != "" {

				ID, err := strconv.Atoi(ItemID)

				if shared.HandleInternalServerError(w, err) {
					return
				}

				// Создаём соединение с базой данных с ролью пользователя
				dbc := setup.ServerSettings.SQL.Connect(w, role)
				if dbc == nil {
					return
				}
				defer dbc.Close()

				err = databases.PostgreSQLArtworkTypesDelete(ID, dbc)

				if err != nil {
					if errors.Is(databases.ErrNoDeleteIfLinksExist, err) {
						shared.HandleOtherError(w, err.Error(), err, http.StatusBadRequest)
						return
					}

					if shared.HandleInternalServerError(w, err) {
						return
					}
				}

				shared.HandleSuccessMessage(w, "Тип картины успешно удалён")

			} else {
				shared.HandleOtherError(w, shared.ErrHeadersNotFilled.Error(), shared.ErrHeadersNotFilled, http.StatusBadRequest)
				return
			}

		} else {
			shared.HandleOtherError(w, shared.ErrForbidden.Error(), shared.ErrForbidden, http.StatusForbidden)
		}

	default:
		shared.HandleOtherError(w, "Method is not allowed", shared.ErrNotAllowedMethod, http.StatusMethodNotAllowed)
	}

}
