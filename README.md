# Тестовое задание для НЛМК

## задание
Написать http сервер который возвращает текущее время. Формат и временная зона передается в параметрах. Если параметров нет: используется временная зона сервера и формат по умолчанию.

Написать тесты для проверки работы сервера.
Для реализации можно использовать gin
Плюсом будет использование swagger.
 go-swagger/go-swagger: Swagger 2.0 implementation for go
 
## Запуск в Docker
`make all`

запрос:

`localhost:8080/api/v1/time?format=&tz=`

документация:

`localhost:8080/swagger/index.html`

## Обновить документацию
`go get -u github.com/swaggo/swag/cmd/swag`
`make docs`

## Тесты
`make test`




