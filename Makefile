COMPOSE_FILE=docker-compose.yml
DOCKER_COMPOSE=docker-compose

build:
	$(DOCKER_COMPOSE) -f $(COMPOSE_FILE) build

up:
	$(DOCKER_COMPOSE) -f $(COMPOSE_FILE) up -d

down:
	$(DOCKER_COMPOSE) -f $(COMPOSE_FILE) down

stop:
	$(DOCKER_COMPOSE) -f $(COMPOSE_FILE) stop

logs-backend:
	$(DOCKER_COMPOSE) -f $(COMPOSE_FILE) logs -f backend

restart:
	$(DOCKER_COMPOSE) -f $(COMPOSE_FILE) down
	$(DOCKER_COMPOSE) -f $(COMPOSE_FILE) up -d

clean:
	$(DOCKER_COMPOSE) -f $(COMPOSE_FILE) down -v