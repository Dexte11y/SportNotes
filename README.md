# **REST API Для учета персональных тренировок на Golang**
Небольшой веб-сервер, реализующий REST API для учёта медперсонала, а также пациентов и их приёмов. Цель создания- отработка навыков программирования на языке Go и пополнение портфолио.

# **В данном проекте реализовано следующее:**
> - Работа с фреймворком gin-gonic/gin.
> - Работа со структурными тегами.
> - Подход Чистой Архитектуры в построении структуры приложения. Техника внедрения зависимости.
> - Работа с БД Postgres. Запуск из Docker. Генерация файлов миграций.
> - Конфигурация приложения с помощью библиотеки spf13/viper. Работа с переменными окружения.
> - Работа с БД, используя библиотеку sqlx.
> - Регистрация и аутентификация. Работа с JWT. Middleware.
> - Написание SQL запросов.
> - Использование Swagger для документации REST API.
> - Использование GoMock для написания unit тестирования.
# **Для запуска приложения:**
make build && make run
**Если запуск выполняется впервые, следует применить миграции к базе данных:**
make migrate
