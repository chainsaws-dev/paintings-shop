package signinupout

import "paintings-shop/packages/authentication"

// SessionsResponse - структура возвращаемая в ответ на запрос сессий
type SessionsResponse struct {
	Sessions
	Total  int
	Offset int
	Limit  int
}

// Sessions - структура описывающая список активных сессий
type Sessions []authentication.ActiveToken
