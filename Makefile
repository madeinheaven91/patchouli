include .env
export $(shell sed 's/=.*//' .env)

.PHONY: compose logs psql psql-test psql-dev migration test

compose:
	@docker compose down
	docker-compose -f 'compose.yml' up -d --build

psql:
	@docker exec -it patch-db psql -U ${DB_USER} -d ${DB_NAME}

migration:
	@docker exec -it patch-db psql -U ${DB_USER} -d ${DB_NAME} -f $(m)

test:
	docker exec -it patch-catalog go test ./...
