init:
	docker network create go_network
	@make build
	@make up
	docker-compose exec front bash -c 'npm install'
	@make sql-run
up:
	docker-compose up -d
down:
	docker-compose down
stop:
	docker-compose stop
build:
	docker-compose build --no-cache --force-rm
ps:
	docker-compose ps
back:
	docker-compose exec -it back sh
front:
	docker-compose exec -it front bash
db:
	docker-compose exec -it db bash
sql:
	docker-compose exec db bash -c 'mysql -u $$MYSQL_USER -p$$MYSQL_PASSWORD $$MYSQL_DATABASE'
sql-run:
	docker-compose exec db bash -c 'mysql -p$$MYSQL_PASSWORD < docker-entrypoint-initdb.d/*.sql'
front-log:
	docker logs --follow front
db-log:
	docker logs db
back-log:
	docker logs --follow back
back-serve:
	docker-compose exec back sh -c 'go run cmd/main.go'
front-serve:
	docker-compose exec front bash -c 'npm run dev'