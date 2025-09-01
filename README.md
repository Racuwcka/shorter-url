# URL Shortener 🔗

[![Go](https://img.shields.io/badge/Go-1.24-blue)](https://golang.org/)
[![Docker](https://img.shields.io/badge/Docker-ready-blue)](https://www.docker.com/)
[![Postgres](https://img.shields.io/badge/Postgres-16-blue)](https://www.postgresql.org/)
[![Swagger](https://img.shields.io/badge/API-Swagger-green)](./swagger-ui/openapi.yaml)
[![Tests](https://img.shields.io/badge/tests-passing-brightgreen)](#)

Сервис для сокращения ссылок с поддержкой in-memory Cash, базы данных и автогенерации Swagger.  
Проект написан на **Go** придерживаясь **Clean Architecture**.  

---

## ✨ Возможности

- Сокращение длинных ссылок до коротких идентификаторов  
- Получение информации о сокращённой ссылке  
- Получение оригинальной ссылки по короткой  
- Редирект по короткой ссылке

---

## 🛠️ Технологии

- **Go** (чистая архитектура)  
- **PostgreSQL** (через `pgx`)  
- **Docker / Docker Compose** (два файла: dev и prod)  
- **oapi-codegen** (генерация кода из Swagger)  
- **github.com/air-verse/air** (hot reload для разработки)  
- **mockgen** (моки для тестов)  

---

## 📡 API

| Метод | Путь                   | Описание                       |
|-------|------------------------|--------------------------------|
| POST  | `/api/v1/shorten`      | Сократить ссылку               |
| GET   | `/api/v1/shorten`      | Получить список сокращённых    |
| GET   | `/api/v1/original`     | Получить оригинальную ссылку   |
| GET   | `/link/{short_id}`     | Редирект по короткой ссылке    |

📖 Swagger спецификация: [`swagger-ui/openapi.yaml`](./swagger-ui/openapi.yaml)  
Генерация кода:  

```bash
make swagger       # генерация кода из Swagger
make cover         # тесты с покрытием
make gen           # генерация моков
make project/init  # первый запуск проекта (dev-среда)
make docker/dev    # dev-среда
make docker/prod   # prod-среда
make docker/down   # остановить контейнеры
