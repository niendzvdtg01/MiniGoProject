-include .env
export

CONN_STRING=postgresql://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSLMODE)
MIGRATION_DIR=internal/db/migration

# Run server
server:
	go run cmd/main.go

# Database backup
import-db:
	docker exec -i postgres-db psql -U root -d master-golang < ./backupdb-master-golang.sql

export-db:
	docker exec postgres-db pg_dump -U root -d master-golang > ./backupdb-master-golang.sql

# Migration
migrate-create:
	migrate create -ext sql -dir $(MIGRATION_DIR) $(NAME)

migrate-up:
	migrate -path $(MIGRATION_DIR) -database "$(CONN_STRING)" up

migrate-down:
	migrate -path $(MIGRATION_DIR) -database "$(CONN_STRING)" down 1

migrate-force:
	migrate -path $(MIGRATION_DIR) -database "$(CONN_STRING)" force $(VERSION)

migrate-drop:
	migrate -path $(MIGRATION_DIR) -database "$(CONN_STRING)" drop

migrate-goto:
	migrate -path $(MIGRATION_DIR) -database "$(CONN_STRING)" goto $(VERSION)

.PHONY: server import-db export-db \
        migrate-create migrate-up migrate-down \
        migrate-force migrate-drop migrate-goto