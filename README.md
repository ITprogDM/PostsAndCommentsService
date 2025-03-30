# Posts And Comments Service

## 📌 Описание проекта

Система для добавления и чтения постов и комментариев с использованием GraphQL

## Основные возможности
- 📝 **Посты**
    - Просмотр списка постов
    - Просмотр отдельного поста с комментариями
    - Возможность запрета комментариев к посту

- 💬 **Комментарии**
    - Иерархическая структура с неограниченной вложенностью
    - Ограничение длины текста (2000 символов)
    - Пагинация для списка комментариев
    - Реализация через GraphQL Subscriptions (асинхронные уведомления)

## 🛠️ Используемые технологии

**PostgreSQL** (в качестве Базы Данных)

**Docker**

**golang-migrate/migrate** (для миграций БД)

**Sqlx** (драйвер для работы с PostgreSQL)

**golang/mock, testify** (для unit-тестирования)

**GraphQL**: gqlgen

**Хранилища** 2 вида:
- Inmemory 
- PostgreSQL

## 🔧 Для запуска сервиса c помощью Docker необходимо:

Заполнить .env файл на основе .env.example. (Укажите в .env POSTGRES_HOST=database); Если хотите использовать Inmemory хранилище
(Указать IN_MEMORY=true)

Запустить сервис:
```
make start-docker
```
Применить миграции БД:
```
make migrate
```

## 🔧 Для локального запуска сервиса необходимо:

Заполнить .env файл на основе .env.example. (Укажите в .env POSTGRES_HOST=localhost); Если хотите использовать Inmemory хранилище 
(Указать IN_MEMORY=true) 

Запустить Базу данных:
```
make start-db
```
Применить миграции БД:
```
make migrate
```
Запустить сервер:
```
make start-local
```
## Для остановки сервиса запущенного в Docker необходимо:
Остановка контейнера
```
make stop-docker
```

## 🐳 Тестирование

Запуск тестов:
```
make test
```
## Примеры запросов

### Создание поста

```
mutation CreatePost {
    CreatePost(
        post: {
            name: "Тест"
            content: "ТестТест"
            author: "Тест1"
            commentsAllowed: true
        }
    ) {
        id
        createdAt
        name
        author
        content
    }
}
```


### Получить список всех постов

```
query GetAllPosts {
    GetAllPosts(page: 1, pageSize: 3) {
        id
        createdAt
        name
        author
        content
    }
}
```


### Получить детальную информацию о конкретном посте
```
query GetPostById {
    GetPostById(id: 1) {
        id
        createdAt
        name
        author
        content
        commentsAllowed
        comments(page: 1, pageSize: 2) {
            id
            createdAt
            author
            content
            post
            replies {
                id
                createdAt
                author
                content
                post
                replyTo
            }
        }
    }
}
```

### Создать комментарий

```
mutation CreateComment {
    CreateComment(input: { author: "Тест1", content: "ТестТестов", post: "1" }) {
        id
        createdAt
        author
        content
        post
        replyTo
    }
}
```


### Подписка на комментарий

```
subscription CommentsSubscription {
    CommentsSubscription(postId: "1") {
        id
        createdAt
        author
        content
        post
        replyTo
    }
}
```

