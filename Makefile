COMPOSE=docker-compose
DOCKER=docker
BACK=geekcamp_go
DB=geekcamp_mysql

run/build:
	$(COMPOSE) up -d --build

run/backend:
	$(COMPOSE) up -d --build $(BACK)

run/database:
	$(COMPOSE) up -d --build $(DB)

run:
	$(COMPOSE) up -d

sh/database:
	$(DOCKER) exec -it $(DB) bash

sh/backend:
	$(DOCKER) exec -it $(DB) ash

down:
	$(COMPOSE) down

down/v:
	$(COMPOSE) down -vol5