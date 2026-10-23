# 08 · Note ownership

## Что тренируем

Authorization как business rule в UseCase.

## Что должно получиться

Реализуйте `Create`, `List`, `Get`, `Update` и `Delete` в
`internal/note/usecase.go`, а также `ListByOwner` в memory repository.

## Как рассуждать

- Create получает owner только из authenticated `userID`.
- List просит repository вернуть Notes этого owner.
- Get/Update/Delete сначала загружают Note.
- При `OwnerID != userID` возвращается `ErrForbidden`.
- Repository mutation не вызывается после forbidden.

## Подсказка

Не передавайте `OwnerID` из request DTO. Клиент не выбирает владельца.

## Частая ошибка

Проверить ownership в Handler или после вызова Repository.Update/Delete.

## Как проверить

```bash
go test ./internal/note -v
go test ./internal/repository/memory -v
```
