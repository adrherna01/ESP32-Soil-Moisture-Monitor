COMPOSE=docker-compose

build: 
	$(COMPOSE) build

up: 
	$(COMPOSE) up

down:
	$(COMPOSE) down

stop:
	$(COMPOSE) stop

logs:
	$(COMPOSE) logs -f

re:
	$(COMPOSE) up -d --build backend

prune:
	docker system prune -a

backend:
	docker exec -it backend bash

frontend:
	docker exec -it frontend bash

db:
	docker exec -it db bash

front:
	cd src/frontend && docker build -t frontend . && docker run -p 80:80 frontend

.PHONY: db api
