migratedown:
    migrate -path migration/aql -database postgres://postgres:postgres@localhost:5432/postgres -sourcefile users.sql

migrateup:
    migrate -path migration/sql -database postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable

sqlc:
    sqlc generate

test: 
    go test -v -cover ./

.PHONY: postgres createdb dropdb migrateup migratedown sqlc