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

migrate-create:
	migrate create -ext sql -dir $(MIGRATION_DIR)-seq $(NAME)