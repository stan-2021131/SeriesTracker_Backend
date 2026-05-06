FROM golang:1.26.2-alpine AS builder

WORKDIR /app

# Copiar go.mod y descargar dependencias
COPY go.mod go.sum ./
RUN go mod download

# Copiar el código
COPY . .

RUN go install github.com/swaggo/swag/cmd/swag@latest

ENV GO111MODULE=on

RUN swag init --parseDependency --parseInternal

# Compilar
RUN go build -o main .

# Etapa final
FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/main .

# carpeta de uploads
RUN mkdir -p /app/uploads

# Puerto
EXPOSE 3000

# Ejecutar app
CMD ["./main"]