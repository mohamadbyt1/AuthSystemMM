postgresinit:
	docker run --name pg  -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres
postgres:
	docker exec -it pg psql
createdb:
		docker exec -it pg createdb --username=root --owner=root ghodsp
dropdb:
	docker exec -it pg dropdb chatapp

migrationup:
	migrate -path db/migration -database postgresql://root:password@localhost:5432/chatapp?sslmode=disable -verbose up

migrationdown:
	migrate -path db/migration -database postgresql://root:password@localhost:5432/chatapp?sslmode=disable -verbose down


.PHONY:  postgresinit postgres createdb dropdb migrationup migrationdown 