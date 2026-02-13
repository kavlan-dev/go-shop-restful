<div align="center">
  <img src="https://raw.githubusercontent.com/kavlan-dev/golang-shop-restful/main/assets/logo.png" alt="Golang Shop RESTful API" width="200">
  <h1>Golang Shop RESTful API</h1>
  <p>Современное RESTful API для управления интернет-магазином с поддержкой аутентификации, корзины покупок и управления продуктами</p>

  <div>
    <img src="https://img.shields.io/badge/Go-1.25+-00ADD8?style=for-the-badge&logo=go" alt="Go Version">
    <img src="https://img.shields.io/badge/PostgreSQL-12+-336791?style=for-the-badge&logo=postgresql" alt="PostgreSQL">
    <img src="https://img.shields.io/badge/License-MIT-green?style=for-the-badge" alt="License">
    <img src="https://img.shields.io/badge/Status-Active-brightgreen?style=for-the-badge" alt="Status">
  </div>

  <br>

  <div>
    <a href="https://github.com/kavlan-dev/golang-shop-restful/issues">Сообщить об ошибке</a>
    <a href="https://github.com/kavlan-dev/golang-shop-restful/issues">Запросить функцию</a>
    <a href="https://github.com/kavlan-dev/golang-shop-restful/discussions">Обсуждения</a>
  </div>
</div>

---

## О проекте

**Golang Shop RESTful API** - это современное, высокопроизводительное RESTful API для управления интернет-магазином, написанное на языке Go с использованием фреймворка Gin. Проект предоставляет полноценный функционал для создания и управления онлайн-магазином с поддержкой:

<div style="display: flex; flex-wrap: wrap; gap: 10px;">
  <img src="https://img.shields.io/badge/Feature-CRUD_Products-blue?style=flat-square" alt="CRUD Products">
  <img src="https://img.shields.io/badge/Feature-JWT_Auth-green?style=flat-square" alt="JWT Auth">
  <img src="https://img.shields.io/badge/Feature-RBAC-red?style=flat-square" alt="RBAC">
  <img src="https://img.shields.io/badge/Feature-Shopping_Cart-orange?style=flat-square" alt="Shopping Cart">
  <img src="https://img.shields.io/badge/Feature-Validation-yellow?style=flat-square" alt="Validation">
  <img src="https://img.shields.io/badge/Feature-Logging-purple?style=flat-square" alt="Logging">
  <img src="https://img.shields.io/badge/Feature-CORS-lightgrey?style=flat-square" alt="CORS">
  <img src="https://img.shields.io/badge/Feature-Clean_Architecture-brightgreen?style=flat-square" alt="Clean Architecture">
</div>

## Основные возможности

- **Полноценный CRUD** для управления продуктами с пагинацией и фильтрацией
- **Система аутентификации и авторизации** на основе JWT с контролем доступа по ролям
- **Управление корзиной покупок** с проверкой наличия товаров
- **Регистрация и управление пользователями** с валидацией данных
- **Контроль доступа на основе ролей** (customer/admin)
- **Обработка ошибок** с стандартными HTTP-статусами
- **Логирование** всех важных событий
- **Поддержка CORS** для фронтенд-интеграции
- **Чистая архитектура** с четким разделением слоев

## Зачем нужен этот проект?

Этот проект решает несколько ключевых проблем, с которыми сталкиваются разработчики при создании интернет-магазинов:

1. **Готовое решение** - Предоставляет полноценное API для интернет-магазина, которое можно сразу использовать или адаптировать под свои нужды
2. **Безопасность** - Встроенная система аутентификации и авторизации с JWT и RBAC
3. **Производительность** - Написано на Go с использованием высокопроизводительных библиотек
4. **Масштабируемость** - Чистая архитектура позволяет легко расширять функциональность
5. **Готовый фронтенд** - Включает базовый фронтенд на Bootstrap для быстрого старта

## Преимущества

- **Высокая производительность** благодаря использованию Go и фреймворка Gin
- **Безопасность** с JWT аутентификацией и RBAC
- **Гибкость** благодаря чистой архитектуре
- **Готовность к производству** с Docker контейнеризацией
- **Полная документация** с примерами использования

## Требования

<div style="display: flex; flex-wrap: wrap; gap: 10px; align-items: center;">
  <img src="https://img.shields.io/badge/Go-1.25+-00ADD8?style=flat-square&logo=go" alt="Go">
  <img src="https://img.shields.io/badge/PostgreSQL-12+-336791?style=flat-square&logo=postgresql" alt="PostgreSQL">
  <img src="https://img.shields.io/badge/Git-F05032?style=flat-square&logo=git" alt="Git">
  <img src="https://img.shields.io/badge/Docker-2496ED?style=flat-square&logo=docker" alt="Docker">
</div>

Для работы с проектом вам потребуется:

- **Go 1.25+** - Язык программирования
- **PostgreSQL** - Реляционная база данных
- **Git** - Система контроля версий
- **Docker и Docker Compose** (опционально, для контейнеризации)

**Важно**: Убедитесь, что у вас установлены все необходимые зависимости перед запуском проекта.

## Установка и запуск

### Вариант 1: Локальная установка (без Docker)

1. Клонируйте репозиторий:

```bash
git clone https://github.com/kavlan-dev/golang-shop-restful.git
cd golang-shop-restful
```

2. Установите зависимости:

```bash
go mod download
```

## Конфигурация

Проект использует переменные окружения для конфигурации. Создайте файл `.env` на основе шаблона `.env.example`:

```bash
cp .env.example .env
```

### Основные переменные окружения:

```env
# Окружение (dev/prod)
ENV=dev

# Сервер
SERVER_HOST=0.0.0.0
SERVER_PORT=8080

# База данных PostgreSQL
DATABASE_HOST=localhost
DATABASE_USER=postgres
DATABASE_PASSWORD=postgres
DATABASE_NAME=mydb
DATABASE_PORT=5432

# JWT Конфигурация
JWT_SECRET=your-very-secure-secret-key

# CORS (разделенные запятыми)
CORS_ALLOW_ORIGINS=http://localhost:3000,https://your-frontend.com

# Администратор по умолчанию
ADMIN_USERNAME=admin
ADMIN_PASSWORD=admin
ADMIN_EMAIL=admin@example.com
```

**Важно**: Файл `.env` добавлен в `.gitignore` для защиты чувствительных данных.

3. Настройте конфигурацию:

```bash
cp .env.example .env
```

Отредактируйте файл `.env` согласно вашим настройкам базы данных и другим параметрам.

## Запуск

### Локальный запуск (без Docker)

```bash
go run cmd/app/main.go
```

### Сборка и запуск

```bash
go build -o shop-api ./cmd/app
./shop-api
```

### Запуск с Docker (рекомендуется)

```bash
docker-compose up --build
```

Сервер будет доступен на `http://localhost:8080`.

### Вариант 2: Запуск с Docker Compose (рекомендуется)

1. Убедитесь, что у вас установлены Docker и Docker Compose
2. Создайте файл `.env` на основе шаблона:

```bash
cp .env.example .env
```

3. Запустите контейнеры:

```bash
docker-compose up -d --build
```

Это запустит:
- Бэкенд API на порту 8080
- PostgreSQL базу данных
- Фронтенд приложение на порту 8000

## Доступ к приложению

После успешного запуска:
- **API**: `http://localhost:8080`
- **Фронтенд**: `http://localhost:8000`

## Docker и Docker Compose

Проект поддерживает контейнеризацию с помощью Docker и Docker Compose.

### Требования

- Docker или Podman
- Docker Compose или Podman Compose

### Запуск с Docker Compose

```bash
# Создайте .env файл на основе шаблона
cp .env.example .env

# Запустите контейнеры
docker-compose up --build
```

### Конфигурация Docker

- **Dockerfile**: Многоэтапная сборка для оптимизации размера образа
- **docker-compose.yml**: Конфигурация для запуска приложения, PostgreSQL и фронтенда
- **Тома**: Данные PostgreSQL сохраняются в томе `pgdata` для сохранения данных между перезапусками
- **Сервисы**:
  - `backend`: Бэкенд API на порту 8081
  - `frontend`: Фронтенд на порту 8080
  - `db`: База данных PostgreSQL

### Использование с Podman

```bash
# Запуск с Podman Compose
podman-compose up --build
```

**Важно**: При первом запуске приложение автоматически создаст учетную запись администратора, если она не существует. Параметры администратора настраиваются в переменных окружения:

```env
ADMIN_USERNAME=admin
ADMIN_PASSWORD=admin
ADMIN_EMAIL=admin@example.com
```

Если администратор уже существует, автоматическое создание будет пропущено.

## API Эндпоинты

### Аутентификация

- **POST** `/api/auth/register` - Регистрация нового пользователя
- **POST** `/api/auth/login` - Аутентификация пользователя и получение JWT токена

### Продукты

Все эндпоинты для продуктов требуют JWT аутентификации (передавайте токен в заголовке `Authorization: Bearer <token>`):

#### Доступно всем аутентифицированным пользователям:
- **GET** `/api/products` - Получение списка продуктов (с поддержкой пагинации)
  - Параметры: `limit` (по умолчанию: 100, максимум: 1000), `offset` (по умолчанию: 0)
- **GET** `/api/products/:id` - Получение продукта по ID

#### Только для администраторов (роль: admin):
- **POST** `/api/products` - Создание нового продукта
- **PUT** `/api/products/:id` - Обновление продукта (частичное обновление)
- **DELETE** `/api/products/:id` - Удаление продукта

### Управление пользователями (только для администраторов)

Эндпоинты для управления ролями пользователей:

- **POST** `/api/admin/users/:id/promote` - Повысить пользователя до админа
- **POST** `/api/admin/users/:id/downgrade` - Понизить админа до обычного пользователя

### Корзина

Все эндпоинты корзины требуют JWT аутентификации:

- **GET** `/api/cart` - Получение текущей корзины пользователя
- **POST** `/api/cart/:id` - Добавление продукта в корзину
- **DELETE** `/api/cart` - Очистка корзины пользователя

## Примеры использования

### Быстрый старт

Вот несколько примеров, как использовать API для основных операций:

### 1. Регистрация нового пользователя

```bash
curl -X POST http://localhost:8080/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "username": "newuser",
    "password": "securepassword123",
    "email": "user@example.com"
  }'
```

### 2. Аутентификация и получение JWT токена

```bash
curl -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "newuser",
    "password": "securepassword123"
  }'
```

### 3. Создание нового продукта (только для администраторов)

```bash
curl -X POST http://localhost:8080/api/products \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_ADMIN_JWT_TOKEN" \
  -d '{
    "title": "Новый продукт",
    "description": "Описание продукта",
    "price": 1000.99,
    "category": "Электроника",
    "stock": 50
  }'
```

### 4. Получение списка продуктов с пагинацией

```bash
curl -X GET "http://localhost:8080/api/products?limit=10&offset=0" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

### 5. Обновление продукта

```bash
curl -X PUT http://localhost:8080/api/products/1 \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -d '{
    "title": "Обновленный продукт",
    "price": 1200.50
  }'
```

### 6. Добавление продукта в корзину

```bash
curl -X POST http://localhost:8080/api/cart/1 \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

### 7. Просмотр текущей корзины

```bash
curl -X GET http://localhost:8080/api/cart \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

### 8. Очистка корзины

```bash
curl -X DELETE http://localhost:8080/api/cart \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

### 9. Повышение пользователя до админа

```bash
curl -X POST http://localhost:8080/api/admin/users/42/promote \
  -H "Authorization: Bearer ADMIN_JWT_TOKEN"
```

### 10. Понижение админа до обычного пользователя

```bash
curl -X POST http://localhost:8080/api/admin/users/42/downgrade \
  -H "Authorization: Bearer ADMIN_JWT_TOKEN"
```

### Пример аутентификации

```go
package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
)

type LoginRequest struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

type LoginResponse struct {
    Token string `json:"token"`
}

func main() {
    // Создаем запрос на аутентификацию
    loginData := LoginRequest{
        Username: "admin",
        Password: "admin",
    }
    
    jsonData, _ := json.Marshal(loginData)
    
    // Отправляем запрос
    resp, err := http.Post(
        "http://localhost:8080/api/auth/login",
        "application/json",
        bytes.NewBuffer(jsonData),
    )
    
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()
    
    // Читаем ответ
    body, _ := ioutil.ReadAll(resp.Body)
    
    var loginResp LoginResponse
    json.Unmarshal(body, &loginResp)
    
    fmt.Printf("Получен токен: %s\n", loginResp.Token)
}
```

### Пример работы с продуктами

```go
func getProducts(token string) {
    client := &http.Client{}
    
    req, err := http.NewRequest(
        "GET",
        "http://localhost:8080/api/products?limit=5&offset=0",
        nil,
    )
    
    if err != nil {
        panic(err)
    }
    
    // Добавляем токен аутентификации
    req.Header.Add("Authorization", "Bearer "+token)
    
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()
    
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Printf("Список продуктов: %s\n", string(body))
}
```

## Обработка ошибок

API возвращает стандартные HTTP-статусы и JSON-ответы с информацией об ошибках:

- **400 Bad Request** - Некорректные входные данные
- **401 Unauthorized** - Отсутствие или неверный токен аутентификации
- **403 Forbidden** - Недостаточно прав доступа (например, когда пользователь без роли admin пытается создать продукт)
- **404 Not Found** - Ресурс не найден
- **500 Internal Server Error** - Внутренняя ошибка сервера

### Примеры ответов об ошибках

**Ошибка аутентификации (401):**
```json
{
  "error": "Invalid token"
}
```

**Ошибка доступа (403):**
```json
{
  "error": "Admin access required"
}
```

**Ошибка валидации (400):**
```json
{
  "error": "Invalid product data"
}
```

## Структура проекта

Проект следует принципам чистой архитектуры с четким разделением на слои:

```
.
├── cmd/
│   └── app/
│       └── main.go          # Точка входа приложения
├── frontend/                # Фронтенд приложение (Bootstrap + Vanilla JS)
│   ├── css/                 # Стили (Bootstrap + кастомные)
│   ├── js/                  # JavaScript (модули для работы с API)
│   ├── *.html               # HTML страницы (главная, продукты, корзина, аутентификация)
│   └── Dockerfile           # Docker конфигурация фронтенда (Nginx)
├── internal/                # Основной исходный код (Go)
│   ├── app/                 # Основное приложение и инициализация
│   ├── config/              # Работа с конфигурацией (переменные окружения)
│   ├── handler/             # HTTP обработчики (Gin роуты)
│   ├── middleware/          # Middleware (JWT аутентификация, CORS, RBAC)
│   ├── model/               # Модели данных (структуры Go + GORM модели)
│   ├── router/              # Маршрутизация (определение API эндпоинтов)
│   ├── service/             # Бизнес-логика (сервисы для работы с данными)
│   ├── storage/             # Хранилище данных (интерфейсы и реализации)
│   │   └── postgres/        # Реализация для PostgreSQL (GORM)
│   └── util/                # Утилиты (JWT, логирование, валидация)
├── pgdata/                  # Данные PostgreSQL (том Docker для сохранения данных)
├── .env.example             # Шаблон переменных окружения
├── docker-compose.yml       # Docker Compose конфигурация (backend + frontend + db)
├── Dockerfile               # Docker конфигурация бэкенда (многоэтапная сборка)
├── go.mod                   # Модуль Go (зависимости)
├── go.sum                   # Контрольные суммы зависимостей
└── README.md                # Документация (этот файл)
```

### Архитектурные слои

1. **Handlers** - Обработка HTTP-запросов и ответов (Gin)
2. **Services** - Бизнес-логика и работа с данными
3. **Models** - Определение структур данных и валидация
4. **Storage** - Работа с базой данных (PostgreSQL через GORM)
5. **Utils** - Вспомогательные функции (JWT, логирование)

### Ключевые файлы

- `internal/app/app.go` - Основное приложение и инициализация
- `internal/config/config.go` - Загрузка и парсинг конфигурации
- `internal/router/routers.go` - Определение всех API маршрутов
- `internal/handler/*.go` - HTTP обработчики для разных сущностей
- `internal/service/*.go` - Бизнес-логика
- `internal/storage/postgres/*.go` - Работа с базой данных
- `internal/middleware/*.go` - Middleware (аутентификация, авторизация)

## Технологии и ресурсы

### Бэкенд технологии

- **Фреймворк**: [Gin](https://github.com/gin-gonic/gin) - высокопроизводительный HTTP фреймворк для Go
- **ORM**: [GORM](https://gorm.io/) - работа с PostgreSQL, автоматическая миграция схемы
- **Логирование**: [Zap](https://github.com/uber-go/zap) - высокопроизводительное структурированное логирование
- **JWT**: [golang-jwt/jwt](https://github.com/golang-jwt/jwt) - аутентификация и авторизация
- **База данных**: PostgreSQL - реляционная база данных
- **CORS**: [gin-contrib/cors](https://github.com/gin-contrib/cors) - middleware для CORS
- **Контейнеризация**: Docker и Docker Compose для развертывания
- **Валидация**: Встроенная валидация Gin с использованием тегов `binding`

### Фронтенд технологии

- **CSS Фреймворк**: [Bootstrap](https://getbootstrap.com/) - адаптивные компоненты и стили
- **JavaScript**: Vanilla JavaScript без фреймворков (модульный подход)
- **HTTP Клиент**: Fetch API для взаимодействия с RESTful API
- **Хранение**: LocalStorage для JWT токенов и состояния сессии
- **Веб-сервер**: Nginx для продакшена
- **Темная тема**: Кастомная темная тема с адаптированными Bootstrap компонентами
- **Маршрутизация**: Простая клиентская маршрутизация

### Дополнительные ресурсы

- **Документация Gin**: [https://gin-gonic.com/docs/](https://gin-gonic.com/docs/)
- **Документация GORM**: [https://gorm.io/docs/](https://gorm.io/docs/)
- **JWT спецификация**: [https://jwt.io/](https://jwt.io/)
- **Bootstrap**: [https://getbootstrap.com/docs/5.0/](https://getbootstrap.com/docs/5.0/)
- **PostgreSQL**: [https://www.postgresql.org/docs/](https://www.postgresql.org/docs/)

## Архитектура

Проект следует принципам чистой архитектуры с четким разделением слоев:

1. **Router** - маршрутизация HTTP-запросов и ответов
2. **Handlers** - обработка HTTP-запросов и ответов
3. **Services** - бизнес-логика и работа с данными
4. **Models** - определение структур данных
5. **Database** - работа с базой данных
6. **Utils** - вспомогательные функции

## Безопасность и контроль доступа

- Все чувствительные данные (пароли, JWT секреты) хранятся в конфигурационном файле, который исключен из git
- Используется bcrypt для хеширования паролей
- JWT токены имеют ограниченное время жизни (24 часа) и включают информацию о роли и имени пользователя
- Реализована проверка аутентификации для защищенных маршрутов
- CORS настроен для ограничения доступа только к разрешенным источникам
- **Контроль доступа на основе ролей**: Реализована система RBAC (Role-Based Access Control) с ролями "customer" и "admin"

## Система контроля доступа (RBAC)

Проект реализует систему контроля доступа на основе ролей (Role-Based Access Control) с двумя основными ролями:

- **customer**: Обычный пользователь, может просматривать продукты и управлять своей корзиной
- **admin**: Администратор, имеет полный доступ ко всем операциям с продуктами

### Доступные роли

| Роль | Описание | Доступные операции |
|------|----------|-------------------|
| customer | Обычный пользователь | Просмотр продуктов, управление корзиной |
| admin | Администратор | Все операции с продуктами (создание, редактирование, удаление) |

## Модели данных

### Пользователь (User)
```json
{
  "username": "string (3-50 символов)",
  "password": "string (6-100 символов)",
  "email": "string (email формат)",
  "role": "customer|admin"
}
```

### Продукт (Product)
```json
{
  "title": "string (1-100 символов)",
  "description": "string (до 1000 символов)",
  "price": "number (десятичное число)",
  "category": "string (до 50 символов)",
  "stock": "number (целое число, >= 0)"
}
```

### Корзина (Cart)
```json
{
  "user_id": "number",
  "items": [
    {
      "product_id": "number",
      "quantity": "number",
      "price": "number",
      "product": "Product"
    }
  ]
}
```

## Заметки по разработке

**Технические детали:**

- Проект использует GORM для автоматической миграции базы данных при запуске
- Все модели данных имеют валидацию с использованием тегов `binding`
- JWT токены генерируются с использованием пользовательского ID
- Логирование реализовано с использованием Zap для высокой производительности с поддержкой разных окружений (dev/prod)
- CORS middleware настроен для поддержки фронтенд-приложений

**Важно**:

> В модели ProductUpdateRequest отсутствует поле Title, что означает, что заголовок продукта нельзя изменить через API. Это сделано для предотвращения случайного изменения названия продукта.

**Планируемая функциональность**:

> В коде присутствуют закомментированные модели Order и OrderItem, что указывает на планируемую реализацию системы заказов в будущем.

**Советы по разработке**:

```go
// Пример использования логгера
logger.Debug("Сообщение для отладки")
logger.Info("Информационное сообщение")
logger.Warn("Предупреждение")
logger.Error("Ошибка")
```

**Производительность**:

- Использование Gin обеспечивает высокую производительность HTTP-сервера
- Zap предоставляет высокопроизводительное логирование
- GORM оптимизирован для работы с PostgreSQL

## Фронтенд приложение

Проект включает простое фронтенд приложение с темной темой Bootstrap. См. [frontend/README.md](frontend/README.md) для получения дополнительной информации.

### Особенности фронтенда

- **Темная тема**: Современный темный интерфейс с адаптированными Bootstrap компонентами
- **Аутентификация**: Страницы входа и регистрации с JWT аутентификацией
- **Управление продуктами**: Просмотр, фильтрация, сортировка и пагинация продуктов
- **Корзина покупок**: Полнофункциональная корзина с добавлением и удалением товаров
- **Админ панель**: Возможность добавлять и удалять продукты для пользователей с ролью admin
- **Адаптивный дизайн**: Работает на всех устройствах от мобильных до десктопов

### Доступные страницы

- `/index.html` - Главная страница
- `/products.html` - Каталог продуктов
- `/cart.html` - Корзина покупок
- `/login.html` - Страница входа
- `/register.html` - Страница регистрации

### Запуск фронтенда

Фронтенд автоматически запускается вместе с бэкендом при использовании Docker Compose:

```bash
docker-compose up -d --build
```

После запуска:
- Фронтенд: `http://localhost:8000`

## Лицензия

Этот проект лицензирован по лицензии MIT. См. файл [LICENSE](LICENSE) для получения дополнительной информации.
