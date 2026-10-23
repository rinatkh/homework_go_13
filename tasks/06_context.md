# 06 · Auth identity в request.Context

## Что тренируем

Request-scoped data и private context key.

## Что должно получиться

Реализуйте `WithUserID` и `UserID` в
`internal/httpapi/authctx/context.go`.

## Как рассуждать

Middleware создаёт новый context для конкретного request. Handler извлекает
положительный `int64`. Отсутствующее или неверное значение даёт `ok == false`.

## Подсказка

Используйте уже объявленный private `userIDKey`.

## Частая ошибка

Использовать строку `"userID"` или хранить identity в поле общего Handler.

## Как проверить

```bash
go test ./internal/httpapi/middleware -v
```
