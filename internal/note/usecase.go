package note

import (
	"context"
)

type UseCase struct {
	repository Repository
}

func NewUseCase(repository Repository) *UseCase {
	return &UseCase{repository: repository}
}

func (uc *UseCase) Create(ctx context.Context, userID int64, input CreateInput) (Note, error) {
	// TODO: validate title и создай Note с OwnerID из authenticated identity.
	// Клиент не должен присылать OwnerID в JSON.
	_ = ctx
	_ = userID
	_ = input
	return Note{}, ErrTODO
}

func (uc *UseCase) List(ctx context.Context, userID int64) ([]Note, error) {
	// TODO: возвращай только Notes текущего пользователя.
	_ = ctx
	_ = userID
	return nil, ErrTODO
}

func (uc *UseCase) Get(ctx context.Context, userID, noteID int64) (Note, error) {
	// TODO: загрузи Note и верни ErrForbidden, если OwnerID != userID.
	_ = ctx
	_ = userID
	_ = noteID
	return Note{}, ErrTODO
}

func (uc *UseCase) Update(ctx context.Context, userID, noteID int64, input UpdateInput) (Note, error) {
	// TODO: сначала ownership, затем validation и Repository.Update.
	_ = ctx
	_ = userID
	_ = noteID
	_ = input
	return Note{}, ErrTODO
}

func (uc *UseCase) Delete(ctx context.Context, userID, noteID int64) error {
	// TODO: загрузи Note, проверь OwnerID и только затем вызывай Delete.
	_ = ctx
	_ = userID
	_ = noteID
	return ErrTODO
}
