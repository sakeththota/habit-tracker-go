FROM golang:1.23.4-alpine AS builder
RUN apk add --no-cache git bash
WORKDIR /app
RUN git clone --depth 1 https://github.com/sakeththota/habit-tracker-go.git . && ls -la
RUN ls -R /app
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd
FROM alpine:latest
RUN apk add --no-cache ca-certificates
COPY --from=builder /app/server /server
EXPOSE 8080
CMD ["/server"]
