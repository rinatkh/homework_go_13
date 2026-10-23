package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrInvalidToken = errors.New("invalid token")
	ErrTODO         = errors.New("TODO: implement JWT manager")
)

type Claims struct {
	UserID int64 `json:"user_id"`
	jwt.RegisteredClaims
}

// JWTManager реализует два маленьких контракта: TokenIssuer для AuthUseCase и
// TokenVerifier для HTTP middleware.
type JWTManager struct {
	secret []byte
	ttl    time.Duration
	now    func() time.Time
}

func NewJWTManager(secret string, ttl time.Duration) *JWTManager {
	return &JWTManager{secret: []byte(secret), ttl: ttl, now: time.Now}
}

func (m *JWTManager) Issue(userID int64) (string, error) {
	// TODO:
	// Создай claims только с user_id и временными claims, установи exp через
	// m.ttl и подпиши HS256 с m.secret. Не помещай password или hash в payload.
	_ = userID
	return "", ErrTODO
}

func (m *JWTManager) Verify(raw string) (int64, error) {
	// TODO:
	// ParseWithClaims должен проверить signature, exp и разрешить только HS256.
	// Любой malformed/expired/invalid token превращай в ErrInvalidToken.
	_ = raw
	return 0, ErrTODO
}
