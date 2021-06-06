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

// Terms - обработчик для работы со справочником условия продажи
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
// 	идентичным по структуре Terms (см. файл models.go в пакете databases)
//
// DELETE
//
// 	ожидается заголовок ItemID с индексом элемента, который нужно удалить
func Terms(w http.ResponseWriter, req *http.Request) {

	role, auth := signinupout.AuthGeneral(w, req)

	if !auth {
		return
	}

	var err error

	switch {
	case req.Method == http.MethodGet:

		if setup.ServerSettings.CheckRoleForRead(role, "Terms") {

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

					var term databases.Term

					term, err = databases.PostgreSQLSingleTermSelect(ID, dbc)

					if err != nil {
						if errors.Is(err, databases.ErrTermNotFound) {
							shared.HandleOtherError(w, err.Error(), err, http.StatusBadRequest)
							return
						}

						if shared.HandleInternalServerError(w, err) {
							return
						}
					}

					shared.WriteObjectToJSON(false, w, term)

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

				var Terms databases.TermsResponse

				Terms, err = databases.PostgreSQLTermsSelect(Page, Limit, dbc)

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

				shared.WriteObjectToJSON(false, w, Terms)

			}

		} else {
			shared.HandleOtherError(w, shared.ErrForbidden.Error(), shared.ErrForbidden, http.StatusForbidden)
		}

	case req.Method == http.MethodPost:

		if setup.ServerSettings.CheckRoleForChange(role, "Terms") {

			// Читаем тело запроса в структуру
			var term databases.Term
			term.Currency = databases.Currency{}

			err = json.NewDecoder(req.Body).Decode(&term)

			if shared.HandleOtherError(w, "Invalid JSON in request body", err, http.StatusBadRequest) {
				return
			}

			// Создаём соединение с базой данных с ролью пользователя
			dbc := setup.ServerSettings.SQL.Connect(w, role)
			if dbc == nil {
				return
			}
			defer dbc.Close()

			term, err = databases.PostgreSQLTermsChange(term, dbc)

			if shared.HandleInternalServerError(w, err) {
				return
			}

			shared.WriteObjectToJSON(false, w, term)

		} else {
			shared.HandleOtherError(w, shared.ErrForbidden.Error(), shared.ErrForbidden, http.StatusForbidden)
		}

	case req.Method == http.MethodDelete:

		if setup.ServerSettings.CheckRoleForDelete(role, "Terms") {

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

				err = databases.PostgreSQLTermsDelete(ID, dbc)

				if err != nil {
					if errors.Is(databases.ErrNoDeleteIfLinksExist, err) {
						shared.HandleOtherError(w, err.Error(), err, http.StatusBadRequest)
						return
					}

					if shared.HandleInternalServerError(w, err) {
						return
					}
				}

				shared.HandleSuccessMessage(w, "Условие успешно удалено")

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
