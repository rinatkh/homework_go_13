# 01 · User и auth errors

## Что тренируем

Разделение internal Model, input и application errors.

## Что должно получиться

Изучите `internal/user/model.go` и `errors.go`. Убедитесь, что `User` хранит
`PasswordHash`, а plaintext password существует только в `RegisterInput` и
`LoginInput`.

## Как рассуждать

User — сохранённое состояние. Password из request нужен только для hashing или
сравнения и не должен сохраняться или возвращаться.

## Подсказка

Response DTO находится отдельно в `internal/httpapi/dto/auth.go` и не содержит
PasswordHash.

## Частая ошибка

Добавить JSON tags внутреннему User и автоматически encode все поля.

## Как проверить

```bash
go test ./internal/user -run Register -v
```
