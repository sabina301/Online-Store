postgresinit:
	docker run --name online-store -p 5432:5432 -e POSTGRES_PASSWORD=PASSWORD -d postgres:15-alpine

postgres:
	docker exec -it online-store psql -U postgres -d online-store-db

createdb:
	docker exec -it online-store createdb --username=postgres --owner=postgres online-store-db

dropdb:
	docker exec -it online-store dropdb online-store-db

migrateup:
	migrate -path ./schema -database 'postgres://postgres:PASSWORD@localhost:5432/online-store-db?sslmode=disable' up

migratedown:
	migrate -path ./schema -database 'postgres://postgres:PASSWORD@localhost:5432/online-store-db?sslmode=disable' down

.PHONY: postgresinit