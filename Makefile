build:
	@go build -o bin/habit-tracker-go cmd/main.go

test:
	@go test -v ./...

run:
	@./bin/habit-tracker-go
