# 02 · Memory User Repository

## Что тренируем

Repository implementation, map и `sync.RWMutex`.

## Что должно получиться

Реализуйте `Create` и `GetByEmail` в
`internal/repository/memory/user.go`.

## Как рассуждать

HTTP requests выполняются конкурентно. `Create` меняет map и `nextID`, поэтому
нужен `Lock`. `GetByEmail` только читает и использует `RLock`.

Email хранится и ищется в lowercase. Duplicate даёт
`user.ErrAlreadyExists`, неизвестный email — `user.ErrNotFound`.

## Подсказка

Сначала вычислите key, затем выполняйте lookup под соответствующим lock.

## Частая ошибка

Защитить map, но увеличивать `nextID` вне critical section.

## Как проверить

```bash
go test ./internal/repository/memory -v
go test -race ./internal/repository/memory
```
