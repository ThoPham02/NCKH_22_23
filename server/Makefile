migrateup:
	migrate -path db/migration -database "postgres://root:secret@localhost:5432/research_management?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgres://root:secret@localhost:5432/research_management?sslmode=disable" -verbose down
sqlcinit:
	sudo docker run --rm -v "$$(pwd):/src" -w /src kjconroy/sqlc:1.17.0 /workspace/sqlc init
sqlc:
	sudo docker run --rm -v "$$(pwd):/src" -w /src kjconroy/sqlc:1.17.0 /workspace/sqlc generate
