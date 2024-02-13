postgresinit:
	docker run --name postgres15 -p 5432:5432 -e POSTGRES_PASSWORD=password -d postgres:15-alpine

postgres:
	docker exec -it postgres15 psql -U postgres -d online-store

createdb:
	docker exec -it postgres15 createdb --username=postgres --owner=postgres online-store

dropdb:
	docker exec -it postgres15 dropdb online-store

.PHONY: postgresinit