# Habit Tracker REST API in Go

- [x] Register and authenticate users
- [x] Create, read, and delete habits
- [x] Create and read habit progress

### Prereqs

- [Go](https://golang.org/dl/) 1.23 or higher
- [Docker](https://www.docker.com/get-started) (for running the PostgreSQL database)

### Installation

1. Clone Repository

```bash
git clone https://github.com/sakeththota/habit-tracker-go.git
cd habit-tracker-go
```

2. Configure env variables

Create a .env file in the project root and specify the required environment variables. You can use .env.example as a template:

```bash
cp .env.example .env
```

3. Spin up database

```bash
make docker-run
```

4. Run migrations

```bash
make migrate-up
```

5. Install dependencies and run

The server will start on <http://localhost:8080>

```bash
go mod tidy
make build
make run
```
