- chmp nho ve em

- cd ~/Projects/keikibook
- mkdir keikibook
- cd keikibook
- mkdir -p db/migration

build app golang
- go build server.go

- ./api --help
- ./api --version 

run commands
- ./api peppers
- ./api cheese
- ./api pineapple

migrate 
- curl -L https://packagecloud.io/golang-migrate/migrate/gpgkey | apt-key add -
- echo "deb https://packagecloud.io/golang-migrate/migrate/ubuntu/ $(lsb_release -sc) main" > /etc/apt/sources.list.d/migrate.list
- sudo apt-get update
- sudo apt-get install -y migrate
- migrate -version

init 
- migrate create -ext sql -dir db/migration -seq init_schema

Run the migration 
- migrate -path db/migration -database "postgresql://root:1234@localhost:5432/hospital?sslmode=disable" -verbose up

Rollback
- migrate -path db/migration -database "postgresql://root:1234@localhost:5432/hospital?sslmode=disable" -verbose down

- Make run makefile
- make migratedown
- make migrateup