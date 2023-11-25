postgres: 
	docker run --name post12 -dp 5432:5432 -e POSTGRES_PASSWORD=secret postgres:12-alpine

createdb: 
	docker exec -it post12 createdb --username=postgres --owner=postgres mini_bank

dropdb: 
	docker exec -it post12 dropdb mini_bank

migrateup:
	migrate -path db/migrations -database "postgres://postgres:secret@localhost:5432/mini_bank?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migrations -database "postgres://postgres:secret@localhost:5432/mini_bank?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migrations -database "postgres://postgres:secret@localhost:5432/mini_bank?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migrations -database "postgres://postgres:secret@localhost:5432/mini_bank?sslmode=disable" -verbose down 1

sqlcgenerate: 
	docker run --rm -v "${pwd}:/src" -w /src kjconroy/sqlc generate

test: 
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb  -destination db/mock/store.go github.com/buddhimaaushan/mini_bank/db Store

.PHONY: postgres createdb dropdb migrateup migratedown migrateup1 migratedown1 sqlc test server mock