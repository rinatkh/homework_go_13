// Package middleware содержит HTTP boundary для authentication.
package middleware

import (
	"encoding/json"
	"net/http"
)

type TokenVerifier interface {
	Verify(token string) (int64, error)
}

type Auth struct {
	verifier TokenVerifier
}

func NewAuth(verifier TokenVerifier) *Auth {
	return &Auth{verifier: verifier}
}

// RequireUser аутентифицирует запрос один раз и передаёт identity дальше через
// request context. Ownership конкретной Note остаётся ответственностью UseCase.
func (m *Auth) RequireUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO:
		// 1. Разбери Authorization: Bearer <token>.
		// 2. Проверь token через m.verifier.
		// 3. При любой auth-ошибке верни 401 и НЕ вызывай next.
		// 4. Положи userID в request context через authctx.WithUserID.
		// 5. В успешной ветке вызови next.ServeHTTP.
		_ = m
		_ = next
		_ = r
		http.Error(w, "TODO: RequireUser", http.StatusNotImplemented)
	})
}

func bearerToken(header string) (string, bool) {
	// TODO: нужен ровно формат Bearer <non-empty token>.
	_ = header
	return "", false
}

func unauthorized(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("WWW-Authenticate", "Bearer")
	w.WriteHeader(http.StatusUnauthorized)
	_ = json.NewEncoder(w).Encode(map[string]string{"error": "authentication required"})
}
