# Etapa 1: Build da aplicação
FROM golang:1.23.1 AS builder

WORKDIR /app

# Copiar arquivos do projeto
COPY go.mod go.sum ./
RUN go mod download

COPY . ./

# Build da aplicação
RUN go build -o load-tester ./cmd/main.go

# Etapa 2: Imagem final para execução
FROM alpine:latest

WORKDIR /root/

# Instalar GLIBC na imagem final
RUN apk add --no-cache libc6-compat

# Copiar o binário da etapa de build
COPY --from=builder /app/load-tester .

ENTRYPOINT ["./load-tester"]



