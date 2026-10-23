# 12 · Architecture questions

## Что тренируем

Объяснение причин, а не запоминание расположения файлов.

## Что должно получиться

Добавьте ответы в этот файл или отдельный `answers.md`.

## Вопросы

1. Чем authentication отличается от authorization на примере API?
2. Почему отсутствие token даёт `401`, а чужая Note — `403`?
3. Почему ownership находится в Note UseCase?
4. Что именно делает AuthMiddleware?
5. Когда middleware не должен вызывать `next`?
6. Почему userID передаётся через request context?
7. Почему context key имеет private type?
8. Почему PasswordHasher и TokenIssuer — interfaces потребителя?
9. Почему JWTManager может реализовать два разных маленьких interface?
10. Почему payload JWT нельзя считать секретным?
11. Что изменится при JWT → opaque token?
12. Что изменится при memory → PostgreSQL?
13. Почему Response DTO не содержит PasswordHash?
14. Почему map repository нужен mutex?
15. Может ли API вернуть `404` вместо `403` для чужого ресурса? Какой trade-off?

## Как рассуждать

В каждом ответе назовите responsibility, вход/выход и запрещённую зависимость.
Ссылайтесь на конкретный package или function из проекта.

## Подсказка

Главный путь:

```text
Bearer token → Middleware → userID → Context → Handler
→ UseCase → ownership → Repository
```

## Частая ошибка

Ответ «так требует Clean Architecture» без связи с текущим кодом.

## Как проверить

Попросите себя объяснить схему без README и назвать место каждой проверки.
