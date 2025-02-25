# Etapa de compilación
FROM golang:1.23 AS builder
WORKDIR /app
COPY . . 

RUN go build -o main .

# Etapa de ejecución
FROM debian:bookworm-slim
WORKDIR /data
COPY --from=builder /app/main /usr/local/bin/main
COPY .env /data/.env
EXPOSE 8080
CMD ["sh", "-c", "export $(grep -v '^#' /data/.env | xargs) && /usr/local/bin/main"]