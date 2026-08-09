drop table if exists test123;

migrate create -ext sql -dir internal/db/migration -seq users

migrate -path internal/db/migration -database "postgresql://root:1234@localhost:5433/master-golang?sslmode=disable" up