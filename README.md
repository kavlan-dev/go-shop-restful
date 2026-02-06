<div align="center">
  <img src="https://raw.githubusercontent.com/kavlan-dev/golang-shop-restful/main/assets/logo.png" alt="Golang Shop RESTful API" width="200">
  <h1>Golang Shop RESTful API</h1>
  <p>RESTful API для управления интернет-магазином</p>

  <div>
    <img src="https://img.shields.io/badge/Go-1.25+-00ADD8?style=for-the-badge&logo=go" alt="Go Version">
    <img src="https://img.shields.io/badge/PostgreSQL-12+-336791?style=for-the-badge&logo=postgresql" alt="PostgreSQL">
    <img src="https://img.shields.io/badge/License-MIT-green?style=for-the-badge" alt="License">
    <img src="https://img.shields.io/badge/Status-Active-brightgreen?style=for-the-badge" alt="Status">
  </div>

  <br>

  <div>
    <a href="https://github.com/kavlan-dev/golang-shop-restful/issues">Report Bug</a> •
    <a href="https://github.com/kavlan-dev/golang-shop-restful/issues">Request Feature</a> •
    <a href="https://github.com/kavlan-dev/golang-shop-restful/discussions">Discussions</a>
  </div>
</div>

---

**Профессиональное RESTful API** для управления интернет-магазином, написанное на Go с использованием фреймворка Gin. Этот проект предоставляет полноценный функционал для управления интернет-магазином с поддержкой:

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

- **CRUD операций для продуктов** с пагинацией и фильтрацией
- **Системы аутентификации и авторизации** на основе JWT с контролем доступа по ролям
- **Управления корзиной покупок** с проверкой наличия товаров
- **Регистрации и управления пользователями** с валидацией данных
- **Контроля доступа на основе ролей** (customer/admin)
- **Обработки ошибок** с стандартными HTTP-статусами
- **Логирования** всех важных событий
- **CORS поддержки** для фронтенд-интеграции
- **Чистой архитектуры** с четким разделением слоев

## Требования

<div style="display: flex; flex-wrap: wrap; gap: 10px; align-items: center;">
  <img src="https://img.shields.io/badge/Go-1.25+-00ADD8?style=flat-square&logo=go" alt="Go">
  <img src="https://img.shields.io/badge/PostgreSQL-12+-336791?style=flat-square&logo=postgresql" alt="PostgreSQL">
  <img src="https://img.shields.io/badge/Git-F05032?style=flat-square&logo=git" alt="Git">
</div>

- **Go 1.25+** - Язык программирования
- **PostgreSQL 12+** - Реляционная база данных
- **Git** - Система контроля версий

**Важно**: Убедитесь, что у вас установлены все необходимые зависимости перед запуском проекта.

## Установка

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

1. **Шаблон конфигурации**: В проекте доступен шаблон конфигурационного файла `config/config.example.yaml`. Скопируйте его в `config/config.yaml` и настройте параметры:

```bash
cp config/config.example.yaml config/config.yaml
```

2. **Структура конфигурационного файла**:

```yaml
# Server Configuration
server:
  host: localhost
  port: 8080

# Database Configuration
database:
  host: localhost
  user: myuser
  password: pass
  name: mydb
  port: 5432

# JWT Configuration
jwt:
  secret: your-very-secure-secret-key

# CORS Configuration
cors:
  allow_origins:
    - "http://localhost:3000"
    - "https://your-frontend.com"
```

3. **Важно**: Файл `config/config.yaml` добавлен в `.gitignore`, чтобы избежать коммита чувствительных данных (паролей, секретных ключей) в репозиторий.

Приложение будет автоматически загружать конфигурацию из файла `config/config.yaml` при запуске.

## Запуск

```bash
go run main.go
```

Сервер будет запущен на порту `8080` по умолчанию. Для production-окружения рекомендуется использовать:

```bash
go build -o shop-api
./shop-api
```

**Важно**: При первом запуске приложение автоматически создаст учетную запись администратора, если она не существует. Параметры администратора настраиваются в конфигурационном файле `config/config.yaml`:

```yaml
# Admin
admin:
  username: "admin"
  password: "admin123"
  email: "admin@email.com"
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

### Регистрация пользователя

```bash
curl -X POST http://localhost:8080/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "password": "testpass123",
    "email": "test@example.com"
  }'
```

### Аутентификация пользователя

```bash
curl -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "password": "testpass123"
  }'
```

### Создание продукта

```bash
curl -X POST http://localhost:8080/api/products \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -d '{
    "title": "Новый продукт",
    "description": "Описание продукта",
    "price": 1000.99,
    "category": "Электроника",
    "stock": 10
  }'
```

### Получение списка продуктов с пагинацией

```bash
curl -X GET "http://localhost:8080/api/products?limit=10&offset=0" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

### Обновление продукта

```bash
curl -X PUT http://localhost:8080/api/products/1 \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -d '{
    "title": "Обновленный продукт",
    "price": 1200.50
  }'
```

### Добавление продукта в корзину

```bash
curl -X POST http://localhost:8080/api/cart/1 \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

### Получение текущей корзины

```bash
curl -X GET http://localhost:8080/api/cart \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

### Очистка корзины

```bash
curl -X DELETE http://localhost:8080/api/cart \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

### Повышение пользователя до админа

```bash
curl -X POST http://localhost:8080/api/admin/users/42/promote \
  -H "Authorization: Bearer ADMIN_JWT_TOKEN"
```

### Понижение админа до обычного пользователя

```bash
curl -X POST http://localhost:8080/api/admin/users/42/downgrade \
  -H "Authorization: Bearer ADMIN_JWT_TOKEN"
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

```
.
├── main.go                  # Точка входа приложения
├── config/                  # Конфигурационные файлы
│   └── config.example.yaml  # Шаблон конфигурации
├── internal/                # Основной исходный код
│   ├── config/              # Работа с конфигурацией
│   ├── database/            # Подключение к базе данных
│   ├── handlers/            # HTTP обработчики
│   ├── middleware/          # Middleware (аутентификация, CORS)
│   ├── models/              # Модели данных
│   ├── services/            # Бизнес-логика
│   └── utils/               # Утилиты (JWT, логирование)
├── go.mod                   # Модуль Go
├── go.sum                   # Контрольные суммы зависимостей
└── README.md                # Документация
```

## Технологии

- **Фреймворк**: [Gin](https://github.com/gin-gonic/gin) - высокопроизводительный HTTP фреймворк
- **ORM**: [GORM](https://gorm.io/) - работа с PostgreSQL
- **Логирование**: [Zap](https://github.com/uber-go/zap) - высокопроизводительное логирование
- **JWT**: [golang-jwt/jwt](https://github.com/golang-jwt/jwt) - аутентификация
- **Конфигурация**: [Viper](https://github.com/spf13/viper) - управление конфигурацией
- **База данных**: PostgreSQL - реляционная база данных
- **CORS**: [gin-contrib/cors](https://github.com/gin-contrib/cors) - middleware для CORS

## Архитектура

Проект следует принципам чистой архитектуры с четким разделением слоев:

1. **Handlers** - обработка HTTP-запросов и ответов
2. **Services** - бизнес-логика и работа с данными
3. **Models** - определение структур данных
4. **Database** - работа с базой данных
5. **Utils** - вспомогательные функции

## Безопасность и контроль доступа

- Все чувствительные данные (пароли, JWT секреты) хранятся в конфигурационном файле, который исключен из git
- Используется bcrypt для хеширования паролей
- JWT токены имеют ограниченное время жизни (24 часа) и включают информацию о роли пользователя
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

### Как назначить роль при регистрации

```bash
curl -X POST http://localhost:8080/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "username": "adminuser",
    "password": "adminpass123",
    "email": "admin@example.com",
    "role": "admin"
  }'
```

Если роль не указана, по умолчанию назначается роль "customer".

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

## Лицензия

Этот проект лицензирован по лицензии MIT. См. файл [LICENSE](LICENSE) для получения дополнительной информации.
