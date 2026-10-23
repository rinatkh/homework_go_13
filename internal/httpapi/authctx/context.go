// Package authctx хранит только request-scoped identity для HTTP transport.
// Приватный тип ключа исключает коллизии с ключами других пакетов.
package authctx

import "context"

type userIDKey struct{}

func WithUserID(ctx context.Context, userID int64) context.Context {
	// TODO: верни новый context с userID. Не используй строковый ключ.
	_ = userID
	return ctx
}

func UserID(ctx context.Context) (int64, bool) {
	// TODO: безопасно прочитай положительный int64 из context.
	_ = ctx
	return 0, false
}
