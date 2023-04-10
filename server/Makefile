postgres:
	docker run --name nckh_db -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:15-alpine 
createdb:
	docker exec -it nckh_db createdb research_management --username=root --owner=root
dropdb:
	docker exec -it nckh_db dropdb research_management
migrate:
	migrate create -ext sql -dir api/db/migration -seq add_library
migrateup:
	migrate -path api/db/migration -database "postgres://thopham:Lk8UZozxGk0BfPfAdfjsjUsptsYtMcKu@dpg-cgkogb8rddleudvr9ee0-a.singapore-postgres.render.com/nckhdb" -verbose up
migratedown:
	migrate -path api/db/migration -database "postgres://thopham:Lk8UZozxGk0BfPfAdfjsjUsptsYtMcKu@dpg-cgkogb8rddleudvr9ee0-a.singapore-postgres.render.com/nckhdb" -verbose down
sqlcinit:
	docker run --rm -v "$$(pwd):/src" -w /src kjconroy/sqlc:1.17.0 /workspace/sqlc init
sqlc:
	docker run --rm -v "$$(pwd):/src" -w /src kjconroy/sqlc:1.17.0 /workspace/sqlc generate
run:
	go run main.go