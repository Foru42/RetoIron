# Etapa de compilación
FROM golang:1.21 AS builder
WORKDIR /app
COPY ./src . 

# Inicializa el módulo y descarga dependencias
RUN go mod tidy
RUN go build -o main .

# Etapa de ejecución
FROM debian:bookworm-slim
WORKDIR /data
COPY --from=builder /app/main /usr/local/bin/main
EXPOSE 8080
CMD ["/usr/local/bin/main"]
