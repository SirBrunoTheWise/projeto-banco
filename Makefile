postgres:
	docker run --name postgres16 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:16-alpine


createdb:
	docker exec -it postgres16 createdb --username=root --owner=root hunt

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up


migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down


migrateupaws:
	migrate -path db/migration -database "postgresql://root:juFaRdlLPcBmj4AGIGrn@classdb.cn48cgplikcn.us-east-1.rds.amazonaws.com/classdb" -verbose up


migratedownaws:
	migrate -path db/migration -database "postgresql://root:juFaRdlLPcBmj4AGIGrn@classdb.cn48cgplikcn.us-east-1.rds.amazonaws.com/classdb" -verbose down

dropdb:
	docker exec -it postgres16 dropdb simple_bank

sqlc:
	sqlc generate

.PHONY: postgres createdb dropdb migrateup migratedown sqlc
