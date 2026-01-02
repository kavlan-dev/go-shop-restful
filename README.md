# Golang Shop RESTful API

Профессиональное RESTful API для управления интернет-магазином, написанное на Go с использованием фреймворка Gin.

## Описание

Этот проект предоставляет полноценный функционал для управления интернет-магазином с поддержкой:

- **CRUD операций для продуктов** с пагинацией и фильтрацией
- **Системы аутентификации и авторизации** на основе JWT
- **Управления корзиной покупок** с проверкой наличия товаров
- **Регистрации и управления пользователями** с валидацией данных
- **Обработки ошибок** с стандартными HTTP-статусами
- **Логирования** всех важных событий

## Требования

- Go 1.20+
- PostgreSQL 12+
- Git

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

## API Эндпоинты

### Аутентификация

- **POST** `/api/auth/register` - Регистрация нового пользователя
- **POST** `/api/auth/login` - Аутентификация пользователя и получение JWT токена
- **GET** `/api/auth/:name` - Получение информации о пользователе по имени

### Продукты

Все эндпоинты для продуктов требуют JWT аутентификации (передавайте токен в заголовке `Authorization: Bearer <token>`):

- **GET** `/api/products` - Получение списка продуктов (с поддержкой пагинации)
  - Параметры: `limit` (по умолчанию: 100, максимум: 1000), `offset` (по умолчанию: 0)
- **POST** `/api/products` - Создание нового продукта
- **GET** `/api/products/:id` - Получение продукта по ID
- **PUT** `/api/products/:id` - Обновление продукта (частичное обновление)
- **DELETE** `/api/products/:id` - Удаление продукта (закомментировано в текущей версии)

### Корзина

- **POST** `/api/cart/add/:id` - Добавление продукта в корзину (требует аутентификации)

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
curl -X POST http://localhost:8080/api/cart/add/1 \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

## Обработка ошибок

API возвращает стандартные HTTP-статусы и JSON-ответы с информацией об ошибках:

- **400 Bad Request** - Некорректные входные данные
- **401 Unauthorized** - Отсутствие или неверный токен аутентификации
- **404 Not Found** - Ресурс не найден
- **500 Internal Server Error** - Внутренняя ошибка сервера

Пример ответа об ошибке:
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

## Безопасность

- Все чувствительные данные (пароли, JWT секреты) хранятся в конфигурационном файле, который исключен из git
- Используется bcrypt для хеширования паролей
- JWT токены имеют ограниченное время жизни (24 часа)
- Реализована проверка аутентификации для защищенных маршрутов

## Лицензия

Этот проект лицензирован по лицензии MIT. См. файл [LICENSE](LICENSE) для получения дополнительной информации.
