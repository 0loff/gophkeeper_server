// Утилита для работы с ключаим параметров, передаваемых с помощью контекста
package utils

import "context"

type contextKey string

// Ключи параметров передаваемых с помощью контекста
var (
	ContextKeyUserID    = contextKey("uid")
	ContextKeyAuthToken = contextKey("auth token")
)

// Вывод ключа параметра передаваемого с помощью контекста в формате String
func (c contextKey) String() string {
	return string(c)
}

// Получение ключа ID пользователя
func GetUserIDFromContext(ctx context.Context) (string, bool) {
	UserID, ok := ctx.Value(ContextKeyUserID).(string)
	return UserID, ok
}

// Get Auth token from context
func GetAuthTokenFromContext(ctx context.Context) (string, bool) {
	AuthToken, ok := ctx.Value(ContextKeyAuthToken).(string)
	return AuthToken, ok
}
