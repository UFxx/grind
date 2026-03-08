include internal/config/.env

migrations_dir=internal/migrations

dev:
	air

migration.new:
	goose -dir $(migrations_dir) create -s $(n) sql

migration.up:
	goose -dir $(migrations_dir) $(DB_DRIVER) $(DSN) up

migration.down:
	goose -dir $(migrations_dir) $(DB_DRIVER) $(DSN) down

migration.up-to:
	goose -dir $(migrations_dir) $(DB_DRIVER) $(DSN) up-to $(to)

migration.down-to:
	goose -dir $(migrations_dir) $(DB_DRIVER) $(DSN) down-to $(to)

migration.reset:
	goose -dir $(migrations_dir) $(DB_DRIVER) $(DSN) reset

migration.status:
	goose -dir $(migrations_dir) $(DB_DRIVER) $(DSN) status

prepare-configs:
	cp internal/config/.env.example internal/config/.env
