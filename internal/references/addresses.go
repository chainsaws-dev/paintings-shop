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

// Addresses - обработчик для работы со справочником адреса
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
func Addresses(w http.ResponseWriter, req *http.Request) {

	role, auth := signinupout.AuthGeneral(w, req)

	if !auth {
		return
	}

	var err error

	switch {
	case req.Method == http.MethodGet:

		if setup.ServerSettings.CheckRoleForRead(role, "Addresses") {

			PageStr := req.Header.Get("Page")
			LimitStr := req.Header.Get("Limit")
			ItemID := req.Header.Get("ItemID")

			if PageStr == "" && LimitStr == "" {

				if ItemID != "" {

					ID, err := strconv.Atoi(ItemID)

					if shared.HandleInternalServerError(setup.ServerSettings.Lang, w, req, err) {
						return
					}

					var Addr databases.Address

					Addr, err = databases.PostgreSQLSingleAddressSelect(ID, setup.ServerSettings.SQL.ConnPool)

					if err != nil {
						if errors.Is(err, databases.ErrAddressNotFound) {
							shared.HandleOtherError(setup.ServerSettings.Lang, w, req, err.Error(), err, http.StatusBadRequest)
							return
						}

						if shared.HandleInternalServerError(setup.ServerSettings.Lang, w, req, err) {
							return
						}
					}

					shared.WriteObjectToJSON(setup.ServerSettings.Lang, w, req, Addr)

				} else {
					shared.HandleOtherError(setup.ServerSettings.Lang, w, req, shared.ErrHeadersNotFilled.Error(), shared.ErrHeadersNotFilled, http.StatusBadRequest)
					return
				}

			} else {

				Page, err := strconv.Atoi(PageStr)

				if shared.HandleInternalServerError(setup.ServerSettings.Lang, w, req, err) {
					return
				}

				Limit, err := strconv.Atoi(LimitStr)

				if shared.HandleInternalServerError(setup.ServerSettings.Lang, w, req, err) {
					return
				}

				var Addr databases.AddressesResponse

				Addr, err = databases.PostgreSQLAddressesSelect(Page, Limit, setup.ServerSettings.SQL.ConnPool)

				if err != nil {
					if errors.Is(err, databases.ErrLimitOffsetInvalid) {
						shared.HandleOtherError(setup.ServerSettings.Lang, w, req, err.Error(), err, http.StatusBadRequest)
						return
					}

					if errors.Is(err, databases.ErrAddressNotFound) {
						shared.HandleOtherError(setup.ServerSettings.Lang, w, req, err.Error(), err, http.StatusBadRequest)
						return
					}

					if shared.HandleInternalServerError(setup.ServerSettings.Lang, w, req, err) {
						return
					}
				}

				shared.WriteObjectToJSON(setup.ServerSettings.Lang, w, req, Addr)

			}

		} else {

			shared.HandleOtherError(setup.ServerSettings.Lang, w, req, shared.ErrForbidden.Error(), shared.ErrForbidden, http.StatusForbidden)
		}

	case req.Method == http.MethodPost:

		if setup.ServerSettings.CheckRoleForChange(role, "Addresses") {

			// Читаем тело запроса в структуру
			var Addr databases.Address
			Addr.Country = databases.Country{}

			err = json.NewDecoder(req.Body).Decode(&Addr)

			if shared.HandleOtherError(setup.ServerSettings.Lang, w, req, shared.ErrInvalidRequestJSON.Error(), err, http.StatusBadRequest) {
				return
			}

			Addr, err = databases.PostgreSQLAddressChange(Addr, setup.ServerSettings.SQL.ConnPool)

			if shared.HandleInternalServerError(setup.ServerSettings.Lang, w, req, err) {
				return
			}

			shared.WriteObjectToJSON(setup.ServerSettings.Lang, w, req, Addr)

		} else {
			shared.HandleOtherError(setup.ServerSettings.Lang, w, req, shared.ErrForbidden.Error(), shared.ErrForbidden, http.StatusForbidden)
		}

	case req.Method == http.MethodDelete:

		if setup.ServerSettings.CheckRoleForDelete(role, "Addresses") {

			ItemID := req.Header.Get("ItemID")

			if ItemID != "" {

				ID, err := strconv.Atoi(ItemID)

				if shared.HandleInternalServerError(setup.ServerSettings.Lang, w, req, err) {
					return
				}

				err = databases.PostgreSQLAddressesDelete(ID, setup.ServerSettings.SQL.ConnPool)

				if err != nil {
					if errors.Is(databases.ErrNoDeleteIfLinksExist, err) {
						shared.HandleOtherError(setup.ServerSettings.Lang, w, req, err.Error(), err, http.StatusBadRequest)
						return
					}

					if shared.HandleInternalServerError(setup.ServerSettings.Lang, w, req, err) {
						return
					}
				}

				shared.HandleSuccessMessage(setup.ServerSettings.Lang, w, req, shared.MsgEntryDeleted)

			} else {
				shared.HandleOtherError(setup.ServerSettings.Lang, w, req, shared.ErrHeadersNotFilled.Error(), shared.ErrHeadersNotFilled, http.StatusBadRequest)
				return
			}

		} else {
			shared.HandleOtherError(setup.ServerSettings.Lang, w, req, shared.ErrForbidden.Error(), shared.ErrForbidden, http.StatusForbidden)
		}

	default:
		shared.HandleOtherError(setup.ServerSettings.Lang, w, req, shared.ErrNotAllowedMethod.Error(), shared.ErrNotAllowedMethod, http.StatusMethodNotAllowed)
	}

}
