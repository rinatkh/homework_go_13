// Package memory содержит потокобезопасные in-memory реализации repository.
// Это infrastructure detail: usecase видит только интерфейсы user.Repository
// и note.Repository.
package memory

import (
	"context"
	"sync"

	"github.com/rinatkh/homework_go_13/internal/user"
)

type UserRepository struct {
	mu      sync.RWMutex
	nextID  int64
	byEmail map[string]user.User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{nextID: 1, byEmail: make(map[string]user.User)}
}

func (r *UserRepository) Create(_ context.Context, value user.User) (user.User, error) {
	// TODO:
	// Защити map и nextID через Lock. Email — ключ в normalized lowercase.
	// При duplicate верни user.ErrAlreadyExists.
	_ = r
	_ = value
	return user.User{}, user.ErrTODO
}

func (r *UserRepository) GetByEmail(_ context.Context, email string) (user.User, error) {
	// TODO: используй RLock и верни user.ErrNotFound для неизвестного email.
	_ = r
	_ = email
	return user.User{}, user.ErrTODO
}
