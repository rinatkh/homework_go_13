# 05 · JWT Manager

## Что тренируем

Token issuing/verifying за infrastructure contract.

## Что должно получиться

Реализуйте `Issue` и `Verify` в `internal/auth/token.go`.

## Как рассуждать

`Issue` создаёт HS256 token с `user_id`, `iat` и `exp`. `Verify` проверяет
signature, expiration, ожидаемый алгоритм и положительный `user_id`.

## Подсказка

Используйте `jwt.ParseWithClaims`, `jwt.WithExpirationRequired`,
`jwt.WithTimeFunc` и `jwt.WithValidMethods`.

## Частая ошибка

Доверять алгоритму из token header без allow-list или забыть проверку `exp`.

## Как проверить

```bash
go test ./internal/auth -v
```
