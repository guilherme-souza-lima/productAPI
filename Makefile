env:
	cp ./.env.example ./.env
env-docker:
	cp ./.env.docker ./.env
up:
	docker-compose up -d
down:
	docker-compose down
run:
	go run main.go
