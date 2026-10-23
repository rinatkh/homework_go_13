# 04 · Login UseCase

## Что тренируем

Проверку credentials без раскрытия лишней информации.

## Что должно получиться

Реализуйте `UseCase.Login`.

## Как рассуждать

Найдите User, сравните password с hash через `PasswordHasher.Compare`, затем
получите token через `TokenIssuer`.

Отсутствующий User и неверный password возвращают одну ошибку
`ErrInvalidCredentials`.

## Подсказка

Внутреннюю unexpected repository error возвращайте как есть. Только
`ErrNotFound` превращается в invalid credentials.

## Частая ошибка

Сравнивать `value.PasswordHash == input.Password`.

## Как проверить

```bash
go test ./internal/user -run Login -v
```
