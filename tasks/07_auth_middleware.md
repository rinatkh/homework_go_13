# 07 · Auth Middleware

## Что тренируем

Middleware signature, Bearer parsing и остановку handler chain.

## Что должно получиться

Реализуйте `bearerToken` и `Auth.RequireUser` в
`internal/httpapi/middleware/auth.go`.

## Как рассуждать

Middleware выполняет общий transport flow:

1. читает `Authorization`;
2. требует `Bearer <token>`;
3. вызывает `TokenVerifier.Verify`;
4. при ошибке возвращает `401` и завершает функцию;
5. записывает `userID` в context;
6. вызывает `next.ServeHTTP`.

## Подсказка

Для разбора header удобны `strings.Fields` и `strings.EqualFold`.

## Частая ошибка

Вызвать `next` после записи `401` или забыть `next` в успешной ветке.

## Как проверить

```bash
go test ./internal/httpapi/middleware -v
```
