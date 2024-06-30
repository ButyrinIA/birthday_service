## Описание
Этот проект представляет собой сервис для управления пользователями и напоминания о днях рождения. 
Он включает в себя регистрацию пользователей, аутентификацию, подписку на уведомления о днях рождения 
и отправку этих уведомлений по электронной почте.

# Структура проекта
Проект состоит из нескольких пакетов:

- main - основной файл для запуска приложения.
- config - пакет для работы с конфигурационными файлами.
- auth - пакет для аутентификации пользователей с использованием JWT.
- database - пакет для инициализации подключения к базе данных.
- handlers - пакет с обработчиками HTTP-запросов.
- models - пакет с определением моделей данных.
- notifier - пакет для отправки уведомлений по электронной почте.
- repository - пакет для взаимодействия с базой данных.
- routes - пакет для определения маршрутов HTTP.
- service - пакет с бизнес-логикой приложения.

## Установка и настройка
# Требования
- Go 1.16 или выше
- PostgreSQL 13 или выше
# Настройка
- Склонируйте репозиторий:

- Создайте файл конфигурации config.yaml в корне проекта со следующим содержимым:

Копировать код 

```
DB_HOST: "localhost"
DB_PORT: "5432"
DB_USER: "db_user"
DB_PASSWORD: "db_password"
DB_NAME: "birthday_service"
JWT_SECRET: "jwt_secret"

```

# Настройте Docker для запуска PostgreSQL:

Копировать код

```
version: "3.3"

networks:
  net:
    driver: bridge

services:
  db:
    image: postgres:13
    environment:
      POSTGRES_DB: birthday_service
      POSTGRES_USER: db_user
      POSTGRES_PASSWORD: db_password
    ports:
      - "5432:5432"
```

# Запустите Docker-контейнер с базой данных:

```docker-compose up -d```

# Инициализируйте базу данных и выполните миграции:

```make test-migration-up```

# Запустите приложение:

```go run main.go```

# Команды Makefile
```
Создать новую миграцию:

make migration-create name=<migration_name>

Применить миграции:

make test-migration-up

Откатить миграции:

make test-migration-down
```

# Использование
- Регистрация пользователя

```curl.exe -X POST http://localhost:8080/register -H "Content-Type: application/json" -d '{\"username\": \"test4\", \"password\": \"123\", \"email\": \"testuser@example.com\", \"birthday\": \"2004-06-29\", \"is_subscribed\": true}'```

- Аутентификация пользователя
  ```curl.exe -X POST http://localhost:8080/login -H "Content-Type: application/json" -d '{\"username\": \"test3\", \"password\": \"123\"}' ```    

- Вывод: {"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTk3MDE2MzIsInVzZXJuYW1lIjoidGVzdDMifQ.HEYVrM-o6UrGGK1T8WI9LibgJtx5cw8G1xI86Wlzspo"}


- Получение сегодняшних дней рождений
  ```curl.exe -X GET -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTk3MDE4MDgsInVzZXJuYW1lIjoidGVzdCJ9.3lfF4aOd_daS9UyDq3JIgv8oF-R-EOdy91Q7jS46vnU" http://localhost:8080/birthdays```
- Вывод:  [{"id":14,"username":"test3","email":"testuser@example.com","birthday":"2024-06-29T00:00:00Z","is_subscribed":true},{"id":15,"username":"test4","email":"testuser@example.com","birthday":"2004-06-29T00:00:00Z","is_subscribed":true}]

- Подписка на уведомления
  ```PS C:\Users\pgs22\GolandProjects\rutube> curl.exe -X POST http://localhost:8080/subscribe -H "Content-Type: application/json" -d '{\"id\": 2}'
```
- Отписка от уведомлений
```curl.exe -X POST http://localhost:8080/unsubscribe -H "Content-Type: application/json" -d '{\"id\": 2}'
```

#  Автор
Бутырин Иван Алексеевич
# Контактная информация
+79272607422

vanya.1.butyrin@gmail.com
