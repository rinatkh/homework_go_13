package user

import (
	"context"
)

// PasswordHasher и TokenIssuer — небольшие интерфейсы потребителя. UseCase
// зависит от нужного поведения, а не от bcrypt и конкретной JWT-библиотеки.
type PasswordHasher interface {
	Hash(password string) (string, error)
	Compare(hash, password string) error
}

type TokenIssuer interface {
	Issue(userID int64) (string, error)
}

type UseCase struct {
	repository Repository
	hasher     PasswordHasher
	tokens     TokenIssuer
}

func NewUseCase(repository Repository, hasher PasswordHasher, tokens TokenIssuer) *UseCase {
	return &UseCase{repository: repository, hasher: hasher, tokens: tokens}
}

func (uc *UseCase) Register(ctx context.Context, input RegisterInput) (User, error) {
	// TODO:
	// 1. Нормализуй email через TrimSpace + ToLower.
	// 2. Проверь email и минимальную длину password.
	// 3. Убедись через Repository, что email ещё не занят.
	// 4. Получи hash через PasswordHasher.
	// 5. Сохрани User только с PasswordHash, не с plaintext password.
	//
	// UseCase не должен импортировать net/http, bcrypt или JWT library.
	_ = ctx
	_ = input
	return User{}, ErrTODO
}

func (uc *UseCase) Login(ctx context.Context, input LoginInput) (LoginResult, error) {
	// TODO:
	// 1. Найди User по normalized email.
	// 2. Отсутствующий User и неправильный password должны давать одну внешнюю
	//    ошибку ErrInvalidCredentials.
	// 3. Сравни password через PasswordHasher.Compare.
	// 4. Получи token через TokenIssuer и верни LoginResult.
	_ = ctx
	_ = input
	return LoginResult{}, ErrTODO
}
