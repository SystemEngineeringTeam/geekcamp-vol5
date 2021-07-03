COMPOSE=docker-compose
DOCKER=docker
BACK=geekcamp_go
DB=geekcamp_mysql
BACK_SERVICE = backend
DB_SERVICE = mysql

run/build:
	$(COMPOSE) up -d --build

run/backend:
	$(COMPOSE) up -d --build $(BACK_SERVICE)

run/database:
	$(COMPOSE) up -d --build $(DB_SERVICE)

run:
	$(COMPOSE) up -d

sh/database:
	$(DOCKER) exec -it $(DB) bash

sh/backend:
	$(DOCKER) exec -it $(DB) ash

down:
	$(COMPOSE) down

down/v:
	$(COMPOSE) down -v

logs:
	$(COMPOSE) logs

ps:
	$(COMPOSE) ps