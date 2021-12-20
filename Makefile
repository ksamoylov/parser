include .env

init:
	docker run -d --name postgresql -e POSTGRES_DB=${DB_NAME} \
      -e POSTGRES_PASSWORD=${DB_PASS} -p ${DB_PORT}:${DB_PORT} postgres:11

m_up:
	migrate -path db/migration -database "${DB_URL}&sslmode=disable" -verbose up

m_down:
	migrate -path db/migration -database "${DB_URL}&sslmode=disable" -verbose down

parse:
	go run ./cmd/main.go