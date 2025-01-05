.PHONY: up
up:
	docker compose -f docker/docker-compose.yml up --build -d

.PHONY: down
down:
	docker compose -f docker/docker-compose.yml down --remove-orphans

.PHONY: migrate
migrate:
	docker exec -i $$(docker ps -qf "name=postgres") psql -U user -d users -c "CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, name VARCHAR(100), email VARCHAR(100) UNIQUE, password VARCHAR(100));"

.PHONY: run
run:
	go run ./cmd/main.go

.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: test
test:
	go test ./...

.PHONY: clean
clean:
	rm -rf build/
