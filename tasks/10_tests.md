# 10 · Tests

## Что тренируем

Изолированную проверку слоёв через небольшие fakes и `httptest`.

## Что должно получиться

Все готовые tests должны проходить. Добавьте минимум один собственный test на
выбранный edge case: duplicate register, malformed Bearer, expired JWT или
forbidden update.

## Как рассуждать

- AuthUseCase test не требует HTTP и real bcrypt.
- Middleware test не требует login и real JWT.
- Note UseCase test не требует Router.
- Router test проверяет полный public/protected lifecycle.

## Подсказка

Сначала запускайте один failing test через `-run`, затем весь package.

## Частая ошибка

Менять ожидаемый status в готовом test вместо исправления production code.

## Как проверить

```bash
make test-unit
make test-race
```
