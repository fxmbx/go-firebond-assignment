postgres:
	docker run --name postgres12  --network firebond-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root firebond_db

createtestdb:
	docker exec -it postgres12 createdb --username=root --owner=root test_db

dropdb:
	docker exec -it postgres12 dropdb firebond_db
droptestdb:
	docker exec -it postgres12 dropdb test_db

createmigration:
	migrate -help
	migrate create -ext sql -dir db/migration -seq [ADDNAME]

migrateup:
	migrate -path db/migration -database "postgresql://root:root@localhost:5432/firebond_db?sslmode=disable" -verbose up

migratetestdbup:
	migrate -path db/migration -database "postgresql://root:root@localhost:5432/test_db?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migration -database "postgresql://root:root@localhost:5432/firebond_db?sslmode=disable" -verbose up 1
migratedown:
	migrate -path db/migration -database "postgresql://root:root@localhost:5432/firebond_db?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://root:root@localhost:5432/firebond_db?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate
test:
	go test -v -cover ./...
server:
	go run main.go
mock:
	mockgen -package mockdb -destination db/mock/store.go  go-firebond-assignment/db/sqlc Store
builddocker:
	docker build -t firebond_assessment:latest .
dockerrun:
	docker run --name firebond_assessment --network firebond-network -p 8080:8080 -e GIN_MODE=release -e DB_SOURCE="postgresql://root:root@postgres12:5432/firebond_db?sslmode=disable" firebond_assessment:latest 

.PHONY: postgres createdb  dropdb migrateup migratetestdbup migratedown migrateup1 migratedown1 sqlc server mock builddocker dockerrun
