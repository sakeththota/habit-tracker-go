# build binaries
FROM golang:1.23.4-alpine AS builder
RUN apk add --no-cache git bash
WORKDIR /app
RUN git clone --depth 1 https://github.com/sakeththota/habit-tracker-go.git . && ls -la
RUN go mod download

# build main server
RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd

# starts new image
FROM alpine:latest
RUN apk add --no-cache ca-certificates

# copy necessary binaries / migrations
COPY --from=builder /app/server /server

EXPOSE 8080

CMD ["/server"]
