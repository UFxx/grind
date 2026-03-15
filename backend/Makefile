migrations_dir=internal/migrations

define run_with_env
	@set -a; . ./.env; set +a; $(1)
endef

dev:
	air

migration.new:
	goose -dir $(migrations_dir) create -s $(n) sql

migration.up:
	$(call run_with_env,goose -dir $(migrations_dir) "$$DB_DRIVER" "$$DSN" up)

migration.down:
	$(call run_with_env,goose -dir $(migrations_dir) "$$DB_DRIVER" "$$DSN" down)

migration.up-to:
	$(call run_with_env,goose -dir $(migrations_dir) "$$DB_DRIVER" "$$DSN" up-to $(to))

migration.down-to:
	$(call run_with_env,goose -dir $(migrations_dir) "$$DB_DRIVER" "$$DSN" down-to $(to))

migration.reset:
	$(call run_with_env,goose -dir $(migrations_dir) "$$DB_DRIVER" "$$DSN" reset)

migration.status:
	$(call run_with_env,goose -dir $(migrations_dir) "$$DB_DRIVER" "$$DSN" status)

prepare-configs:
	@cp --update=none .env.example .env
	@cp --update=none config.yml.example config.yml
