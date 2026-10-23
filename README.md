# Homework 13 — защищённый Notes API

Нужно завершить authentication и ownership authorization в уже разделённом
Notes API. Starter содержит структуру, DTO, handlers, contracts и tests. Ваша
задача — реализовать TODO по порядку, не перенося HTTP и JWT внутрь UseCase.

PostgreSQL в этом задании не используется. Данные хранятся в потокобезопасных
memory repositories.

## Целевая схема

```text
POST /register, POST /login
  → AuthHandler → AuthUseCase → UserRepository
                         ├→ PasswordHasher
                         └→ TokenIssuer

Authorization: Bearer <token>
  → AuthMiddleware → request.Context(userID)
  → NoteHandler → Note UseCase → ownership → NoteRepository
```

## API contract

Public:

| Method | Path | Успех |
|---|---|---:|
| `GET` | `/health` | `200` |
| `POST` | `/register` | `201` |
| `POST` | `/login` | `200` |

Protected:

| Method | Path | Успех |
|---|---|---:|
| `POST` | `/notes` | `201` |
| `GET` | `/notes` | `200` |
| `GET` | `/notes/{id}` | `200` |
| `PUT` | `/notes/{id}` | `200` |
| `DELETE` | `/notes/{id}` | `204` |

## Request examples

Register:

```json
{
  "email": "alice@example.com",
  "password": "secret123"
}
```

Response `201 Created`:

```json
{
  "id": 1,
  "email": "alice@example.com"
}
```

PasswordHash не возвращается.

Login response `200 OK`:

```json
{
  "token": "ey..."
}
```

Protected request:

```http
GET /notes HTTP/1.1
Authorization: Bearer ey...
```

## Status codes

| Ситуация | Status |
|---|---:|
| invalid JSON/ID/input | `400` |
| invalid credentials или token | `401` |
| доступ к чужой Note | `403` |
| Note отсутствует | `404` |
| email уже занят | `409` |
| unexpected error | `500` |

## Порядок работы

Читайте файлы `tasks/` строго по номеру:

```text
01 User и auth errors
02 Memory User Repository
03 Register
04 Login
05 JWT Manager
06 Context
07 Auth Middleware
08 Note ownership
09 Protected routing и error mapping
10 Tests
11 Manual lifecycle
12 Architecture questions
```

Каждый файл содержит:

- что тренируем;
- что должно получиться;
- как рассуждать;
- подсказку;
- частую ошибку;
- как проверить.

## Стартовая проверка

```bash
make check-start
```

Starter компилируется и проходит `go vet`. Полный набор tests до реализации
TODO намеренно падает:

```bash
make test
```

Это нормально. Ошибка теста описывает следующий недостающий contract.

## Что разрешено менять

- тела функций с `TODO`;
- небольшие private helpers рядом с ними;
- свои дополнительные tests;
- комментарии с архитектурными ответами.

## Что не менять

- module path;
- публичные signatures;
- готовые tests и ожидаемые status codes;
- JSON field names;
- routes;
- dependency direction.

Не добавляйте PostgreSQL, ORM, сторонний router, global user state или mocking
framework.

## Acceptance criteria

- `make ci` проходит;
- password никогда не хранится plaintext;
- response и JWT не содержат PasswordHash;
- register duplicate возвращает `409`;
- invalid credentials и token возвращают `401`;
- protected endpoint без token возвращает `401`;
- token с правильной подписью передаёт `userID` через context;
- каждый пользователь получает только свои Notes;
- чужая Note даёт `403` и не изменяется;
- отсутствующая Note даёт `404`;
- middleware не проверяет ownership;
- UseCase не импортирует `net/http` или JWT package;
- map и nextID защищены mutex;
- `go test -race ./...` не находит гонок;
- PostgreSQL не используется.

## Проверка после решения

```bash
make fmt
make test-unit
make test
make test-race
make vet
make ci
```

Ручной lifecycle:

```bash
make run
# второй терминал
make curl-demo
```

## Bonus

Необязательная задача: добавить role `admin`, которая позволяет читать все
Notes. Основной ownership contract менять нельзя. Сначала полностью завершите
обязательные пункты.
