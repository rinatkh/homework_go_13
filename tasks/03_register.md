# 03 · Register UseCase

## Что тренируем

Business validation, password hashing и dependency interfaces.

## Что должно получиться

Реализуйте `UseCase.Register` в `internal/user/usecase.go`.

## Как рассуждать

Порядок:

1. normalize email;
2. validate email/password;
3. проверить duplicate;
4. вызвать `PasswordHasher.Hash`;
5. сохранить User с hash.

## Подсказка

`Repository.GetByEmail` возвращает `ErrNotFound`, если email свободен. Другие
errors нельзя маскировать под duplicate.

## Частая ошибка

Передать `input.Password` в `PasswordHash` без вызова hasher.

## Как проверить

```bash
go test ./internal/user -run Register -v
```
