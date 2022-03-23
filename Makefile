migration:
	migrate create -ext sql -dir db/migration -seq init_schema
	
dropdb:
	migrate -path db/migration -database "postgresql://postgres:password123@localhost:5432/postgres?sslmode=disable" -verbose drop

migrateup:
	migrate -path db/migration -database "postgresql://postgres:password123@localhost:5432/postgres?sslmode=disable"  -verbose up


migratedown:
	migrate -path db/migration -database "postgresql://postgres:password123@localhost:5432/postgres?sslmode=disable" -verbose down

sqlc:
	sqlc generate 

.PHONY: migration dropdb migrateup migratedown
