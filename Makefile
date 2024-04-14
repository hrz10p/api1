postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine	

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root y.kuanyshDB

migrateup:
	migrate -path migrations -database "postgresql://root:secret@localhost:5432/y.kuanyshDB?sslmode=disable" -verbose up

migratedown:
	migrate -path migrations -database "postgresql://root:secret@localhost:5432/y.kuanyshDB?sslmode=disable" -verbose down



