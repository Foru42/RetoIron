# Etapa de compilación
FROM golang:1.23-alpine AS builder
WORKDIR /app

# Instalar dependencias necesarias
RUN apk add --no-cache git

COPY . . 
RUN go build -o main .

# Etapa de ejecución
FROM alpine:latest
WORKDIR /data

# Copiar el binario desde la etapa de compilación
COPY --from=builder /app/main /usr/local/bin/main
COPY .env /data/.env

EXPOSE 8080
CMD ["sh", "-c", "export $(grep -v '^#' /data/.env | xargs) && /usr/local/bin/main"]
