// Package references - содержит веб методы для обработки запросов к таблицам справочников
package references

import (
	"encoding/json"
	"errors"
	"net/http"
	"paintings-shop/internal/databases"
	"paintings-shop/internal/setup"
	"paintings-shop/packages/shared"
	"paintings-shop/packages/signinupout"
	"strconv"
)

// Authors - обработчик для работы со справочником авторы
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
// 	идентичным по структуре Author (см. файл models.go в пакете databases)
//
// DELETE
//
// 	ожидается заголовок ItemID с индексом элемента, который нужно удалить
func Authors(w http.ResponseWriter, req *http.Request) {

	role, auth := signinupout.AuthGeneral(w, req)

	if !auth {
		return
	}

	var err error

	switch {
	case req.Method == http.MethodGet:

		if setup.ServerSettings.CheckRoleForRead(role, "Authors") {

			PageStr := req.Header.Get("Page")
			LimitStr := req.Header.Get("Limit")
			ItemID := req.Header.Get("ItemID")

			if PageStr == "" && LimitStr == "" {

				if ItemID != "" {

					ID, err := strconv.Atoi(ItemID)

					if shared.HandleInternalServerError(w, err) {
						return
					}

					var au databases.Author

					au, err = databases.PostgreSQLSingleAuthorSelect(ID, setup.ServerSettings.SQL.ConnPool)

					if err != nil {
						if errors.Is(err, databases.ErrAuthorNotFound) {
							shared.HandleOtherError(w, err.Error(), err, http.StatusBadRequest)
							return
						}

						if shared.HandleInternalServerError(w, err) {
							return
						}
					}

					shared.WriteObjectToJSON(w, au)

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

				var au databases.AuthorsResponse

				au, err = databases.PostgreSQLAuthorsSelect(Page, Limit, setup.ServerSettings.SQL.ConnPool)

				if err != nil {
					if errors.Is(err, databases.ErrLimitOffsetInvalid) {
						shared.HandleOtherError(w, err.Error(), err, http.StatusBadRequest)
						return
					}

					if errors.Is(err, databases.ErrAuthorNotFound) {
						shared.HandleOtherError(w, err.Error(), err, http.StatusBadRequest)
						return
					}

					if shared.HandleInternalServerError(w, err) {
						return
					}
				}

				shared.WriteObjectToJSON(w, au)

			}

		} else {
			shared.HandleOtherError(w, shared.ErrForbidden.Error(), shared.ErrForbidden, http.StatusForbidden)
		}

	case req.Method == http.MethodPost:

		if setup.ServerSettings.CheckRoleForChange(role, "Authors") {

			// Читаем тело запроса в структуру
			var au databases.Author

			err = json.NewDecoder(req.Body).Decode(&au)

			if shared.HandleOtherError(w, "Invalid JSON in request body", err, http.StatusBadRequest) {
				return
			}

			au, err = databases.PostgreSQLAuthorsChange(au, setup.ServerSettings.SQL.ConnPool)

			if shared.HandleInternalServerError(w, err) {
				return
			}

			shared.WriteObjectToJSON(w, au)

		} else {
			shared.HandleOtherError(w, shared.ErrForbidden.Error(), shared.ErrForbidden, http.StatusForbidden)
		}

	case req.Method == http.MethodDelete:

		if setup.ServerSettings.CheckRoleForDelete(role, "Authors") {

			ItemID := req.Header.Get("ItemID")

			if ItemID != "" {

				ID, err := strconv.Atoi(ItemID)

				if shared.HandleInternalServerError(w, err) {
					return
				}

				err = databases.PostgreSQLAuthorsDelete(ID, setup.ServerSettings.SQL.ConnPool)

				if err != nil {
					if errors.Is(databases.ErrNoDeleteIfLinksExist, err) {
						shared.HandleOtherError(w, err.Error(), err, http.StatusBadRequest)
						return
					}

					if shared.HandleInternalServerError(w, err) {
						return
					}
				}

				shared.HandleSuccessMessage(w, "Автор успешно удален")

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
