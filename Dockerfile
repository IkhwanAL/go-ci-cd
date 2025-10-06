# Build Stage
FROM golang:1.24.1-alpine AS build

WORKDIR /app
COPY go.mod ./

RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o run ./cmd/server

# Runtime
FROM alpine:3.18
COPY --from=build /app/run /app/run

EXPOSE 8080

ENTRYPOINT ["/app/run"]
