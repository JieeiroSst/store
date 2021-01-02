postgres:
#    docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=1234 -d postgres:12-alpine

createdb:
    docker exec -it postgres12 createdb --username=root --owner=root hospital

dropdb:
    docker exec -it postgres12 dropdb hospital

create_migration:
	migrate create -ext sql -dir db/migration -seq init_schema

migrateup:
#    migrate -path db/migration -database "postgresql://root:1234@localhost:5432/store?sslmode=disable" -verbose up

migratedown:
#    migrate -path db/migration -database "postgresql://root:1234@localhost:5432/store?sslmode=disable" -verbose down

run_server:
	go-reload server.go

.PHONY: postgres createdb dropdb migrateup migratedown create_migration run_server