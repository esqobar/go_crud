include .env

stop_containers:
	@echo "Stopping other docker containers"
	@containers=$$(docker ps -q); \
	if [ -n "$$containers" ]; then \
		echo "Found and stopping containers..."; \
		docker stop $$containers; \
	else \
		echo "No containers running..."; \
	fi

create_container:
	docker run --name ${DB_DOCKER_CONTAINER} -p 5432:5432 -e POSTGRES_USER=${DB_USER} -e POSTGRES_PASSWORD=${DB_PASSWORD} -d postgres:16-alpine

create_db:
	@docker exec -i ${DB_DOCKER_CONTAINER} psql -U ${DB_USER} \
		-tAc "SELECT 1 FROM pg_database WHERE datname='${DB_NAME}'" \
		| grep -q 1 && \
		echo "Database ${DB_NAME} already exists" || \
		(docker exec -it ${DB_DOCKER_CONTAINER} createdb \
			--username=${DB_USER} \
			--owner=${DB_USER} \
			${DB_NAME} && \
		echo "Database ${DB_NAME} created")

check_db:
	docker exec -it ${DB_DOCKER_CONTAINER} psql -U ${DB_USER} -l

list_tables:
	docker exec -it ${DB_DOCKER_CONTAINER} psql -U ${DB_USER} -d ${DB_NAME} -c "\dt"

list_schemas:
	docker exec -it ${DB_DOCKER_CONTAINER} psql -U ${DB_USER} -d ${DB_NAME} -c "\dn"

start_container:
	docker start ${DB_DOCKER_CONTAINER}

create_migrations:
	sqlx migrate add -r init

migrate_up:
	sqlx migrate run --database-url "postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable"

migrate_down:
	sqlx migrate revert --database-url "postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable"

build:
	if [ -f "${BINARY}" ]; then \
		rm ${BINARY}; \
		echo "Deleted ${BINARY}"; \
	fi
	@echo "Building binary..."
	go build -o ${BINARY} cmd/api/*.go

run: build
	./${BINARY}
#	@echo "api server started..."

stop:
	@echo "api server stopping..."
	@-pkill -SIGTERM -f "./${BINARY}"
	@echo "api server stopped..."

#stop:
#	@echo "api server stopping..."
#	@pkill -SIGTERM -f "${BINARY}" || true
#	@echo "api server stopped..."