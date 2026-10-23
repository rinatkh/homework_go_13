# Домашняя работа · Урок 13

## Authentication, Authorization и Middleware

Нужно достроить starter Notes API так, чтобы пользователь мог
зарегистрироваться, войти, получить JWT и работать только со своими заметками.

## Порядок выполнения

1. [`01_user_and_errors.md`](../tasks/01_user_and_errors.md) — модель User и
   application errors.
2. [`02_memory_user_repository.md`](../tasks/02_memory_user_repository.md) —
   конкурентно безопасный memory repository.
3. [`03_register.md`](../tasks/03_register.md) — validation, bcrypt hash и
   создание User.
4. [`04_login.md`](../tasks/04_login.md) — Compare password и выдача token.
5. [`05_jwt_manager.md`](../tasks/05_jwt_manager.md) — Issue/Verify, method и
   expiration.
6. [`06_context.md`](../tasks/06_context.md) — request-scoped userID.
7. [`07_auth_middleware.md`](../tasks/07_auth_middleware.md) — Bearer parsing,
   401 и вызов `next`.
8. [`08_note_ownership.md`](../tasks/08_note_ownership.md) — OwnerID и
   `ErrForbidden`.
9. [`09_routing_and_statuses.md`](../tasks/09_routing_and_statuses.md) — public
   и protected routes, 401/403/404.
10. [`10_tests.md`](../tasks/10_tests.md) — unit и HTTP tests.
11. [`11_manual_lifecycle.md`](../tasks/11_manual_lifecycle.md) — Alice/Bob curl
    lifecycle.
12. [`12_architecture_questions.md`](../tasks/12_architecture_questions.md) —
    ответы о границах ответственности.

Каждый файл использует один формат:

```text
Что тренируем
Что должно получиться
Как рассуждать
Подсказка
Частая ошибка
Как проверить
```

## Проверки

До выполнения TODO:

```bash
make check-start
make todo
```

После выполнения:

```bash
make fmt-check
make vet
make deps
make test
make test-race
```

Финальная архитектурная проверка:

- User/Note UseCases не импортируют `net/http` или JWT package;
- Auth Middleware не проверяет ownership;
- OwnerID не принимается из JSON;
- invalid token даёт 401, valid token для чужой Note — 403;
- memory repositories проходят race detector.
