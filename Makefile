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
app:
	docker-compose exec -it app sh
db:
	docker-compose exec -it db bash
sql:
	docker compose exec db bash -c 'mysql -u $$MYSQL_USER -p$$MYSQL_PASSWORD $$MYSQL_DATABASE'
