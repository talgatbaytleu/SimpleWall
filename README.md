# SimpleWall — микросервисная платформа для публикации и взаимодействия с контентом

SimpleWall — это backend-платформа, позволяющая пользователям делиться фотографиями, взаимодействовать с контентом и отслеживать активность.

## Основные возможности:
- Регистрация и публикация постов с описанием
- Редактирование и удаление своих постов
- Просмотр всех постов
- Лайки и дизлайки для любого поста
- Отображение количества лайков и списка лайкеров
- Комментирование постов
- Редактирование и удаление своих комментариев
- Просмотр конкретного комментария или списка комментариев к посту
- Получение списка всех постов пользователя с подсчетом лайков и комментариев

Проект построен на основе микросервисной архитектуры и не включает frontend, предоставляя API для работы с данными.

## Структура проекта
Приложение состоит из 7 микросервисов, каждый из которых выполняет свою функцию:

- **Gateway** — маршрутизация запросов (трехслойная архитектура).
- **Auth** — аутентификация пользователей с JWT-токенами (трехслойная архитектура).
- **Post** — управление постами (трехслойная архитектура).
- **Like** — обработка лайков и дизлайков (трехслойная архитектура).
- **Comment** — управление комментариями (трехслойная архитектура).
- **Wall** — сборка ленты новостей (hexagonal архитектура).
- **Notification** — уведомления (трехслойная архитектура).

## Используемые технологии
- **PostgreSQL** — используется две базы данных:
  - `sw_users_auth` для аутентификации.
  - `sw_posts_db` для хранения постов, лайков и комментариев.
- **Redis** — кэширование для быстрой загрузки ленты (Wall).
- **Kafka** — передача сообщений (`comment-kafka-notification`).
- **Прототип S3** — собственная (очень плохая, но рабочая) реализация хранилища бинарных файлов.
- **Docker** — все сервисы разворачиваются в контейнерах.
- **Git** — разработка велась через отдельные ветки, имитируя реальную работу в команде.

Этот стек технологий обеспечивает высокую масштабируемость, отказоустойчивость и скорость работы платформы.
