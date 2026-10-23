# 11 · Manual lifecycle

## Что тренируем

Проверку HTTP contract глазами клиента.

## Что должно получиться

Запустите API и выполните полный сценарий из `scripts/curl-demo.sh`.

## Как рассуждать

Следите не только за JSON body, но и за status, `Location`,
`WWW-Authenticate` и пустым body для `204`.

## Подсказка

```bash
make run
# второй терминал
make curl-demo
```

## Частая ошибка

Копировать token с кавычками или отправлять его без префикса `Bearer`.

## Как проверить

Alice создаёт Note, Bob получает `403`, Bob list пуст, Alice может обновить и
удалить Note.
