COMPOSE=docker-compose

build: 
	$(COMPOSE) build

up: 
	$(COMPOSE) up

down:
	$(COMPOSE) down

stop:
	$(COMPOSE) stop

api:
	docker exec -it api bash

db:
	docker exec -it db bash

.PHONY: db api
