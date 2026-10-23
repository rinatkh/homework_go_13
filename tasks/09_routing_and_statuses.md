# 09 · Protected routing и status mapping

## Что тренируем

Границу между public и protected endpoints, а также `401` и `403`.

## Что должно получиться

Изучите `internal/httpapi/router.go` и handlers. После предыдущих задач все
`/notes` routes должны проходить `RequireUser`, а `/register` и `/login`
оставаться public.

## Как рассуждать

Middleware формирует `401` до Handler. `note.ErrForbidden` появляется в UseCase
и превращается Handler-ом в `403`.

## Подсказка

UseCase не должен возвращать `http.StatusForbidden`.

## Частая ошибка

Обернуть middleware-ом `/login`: клиент ещё не может предъявить token.

## Как проверить

```bash
go test ./internal/httpapi -run ProtectedRoutesAndOwnership -v
```
