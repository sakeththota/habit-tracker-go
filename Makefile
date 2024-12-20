build:
	@go build -o bin/habit-tracker-go cmd/main.go

test:
	@go test -v ./...

run:
	@./bin/habit-tracker-go

migration:
	@migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	@go run cmd/migrate/main.go up

migrate-down:
	@go run cmd/migrate/main.go down
