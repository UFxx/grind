# Требования
---
```
Go: go1.25.4
make package (если не линукс, просто копируем команды из Makefile)
```

# Запуск проекта для локальной разработки
---
```bash
Клонируем проект
git clone https://github.com/sunsetsavorer/grind

Переходим в папку backend и устанавливаем зависимости
cd backend
go mod download

Устанавливаем goose для миграций и air для хотрелоада
go install github.com/air-verse/air@latest
go install github.com/pressly/goose/v3/cmd/goose@latest

Подготавливаем конфиги, поднимаем миграции, запускаем приложение
make prepare-configs
make migration.up
make dev
```