migrateup:
	migrate -path api/db/migration -database "postgres://root:secret@localhost:5432/research_management?sslmode=disable" -verbose up
migratedown:
	migrate -path api/db/migration -database "postgres://root:secret@localhost:5432/research_management?sslmode=disable" -verbose down
sqlcinit:
	docker run --rm -v "$$(pwd):/src" -w /src kjconroy/sqlc:1.17.0 /workspace/sqlc init
sqlc:
	docker run --rm -v "$$(pwd):/src" -w /src kjconroy/sqlc:1.17.0 /workspace/sqlc generate
