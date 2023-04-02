postgres:
	docker run --name nckh_db -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:15-alpine 
createdb:
	docker exec -it nckh_db createdb research_management --username=root --owner=root
dropdb:
	docker exec -it nckh_db dropdb research_management
migrateup:
	migrate -path api/db/migration -database "postgresql://u7xfvwpftmunpstzixfb:C966Tikw9TVr5CT3oICZkPpOlMMqPQ@bpjvf0abkz3vn3j7p5qx-postgresql.services.clever-cloud.com:5432/bpjvf0abkz3vn3j7p5qx" -verbose up
migratedown:
	migrate -path api/db/migration -database "postgresql://u7xfvwpftmunpstzixfb:C966Tikw9TVr5CT3oICZkPpOlMMqPQ@bpjvf0abkz3vn3j7p5qx-postgresql.services.clever-cloud.com:5432/bpjvf0abkz3vn3j7p5qx" -verbose down
sqlcinit:
	docker run --rm -v "$$(pwd):/src" -w /src kjconroy/sqlc:1.17.0 /workspace/sqlc init
sqlc:
	docker run --rm -v "$$(pwd):/src" -w /src kjconroy/sqlc:1.17.0 /workspace/sqlc generate
