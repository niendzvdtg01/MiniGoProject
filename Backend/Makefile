include .env
export 

CONN_STRING=postgresql://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSLMODE)
MIGRATION_DIR = internal/db/migration

#Import database 
import-db:
	docker exec -i postgres-db psql -U root -d master-golang < ./backupdb-master-golang.sql
#Export database 
export-db:
	docker exec -i postgres-db pg_dump -U root -d master-golang > ./backupdb-master-golang.sql   
#migration
server:
	go run cmd/main.go
#create new migration
migrate-create:
	migrate create -ext sql -dir $(MIGRATION_DIR)-seq $(NAME)
#run all pending migration
migrate-up:
	migrate -path $(MIGRATION_DIR) -database "$($CONN_STRING)" up
migrate-down:
	migrate -path $(MIGRATION_DIR) -database "$($CONN_STRING)" down 1
migrate-force:
	migrate -path $(MIGRATION_DIR) -database "$($CONN_STRING)" force $(VERSION)

.PHONY: import-db export-db server migrate-create migrate-up migrate-down migrate-force
