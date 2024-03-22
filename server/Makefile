postgresinit:
	docker run --name online-store -p 5432:5432 -e POSTGRES_PASSWORD=Ertyu55555 -d postgres:15-alpine

postgres:
	docker exec -it ps-psql psql -U postgres -d online-store-db

createdb:
	docker exec -it ps-psql createdb --username=postgres --owner=postgres online-store-db

dropdb:
	docker exec -it ps-psql dropdb --username=postgres online-store-db

migrateup:
	migrate -path ./schema -database 'postgres://postgres:Ertyu55555@ps-psql:5432/online-store-db?sslmode=disable' up

migratedown:
	migrate -path ./schema -database 'postgres://postgres:Ertyu55555@localhost:5432/online-store-db?sslmode=disable' down

build:
	docker build -t online-store:local .

run:
	docker-compose up app

build-compose:
	docker-compose build app

.PHONY: postgresinit