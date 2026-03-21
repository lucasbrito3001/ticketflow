# Stage 1: Build
FROM golang:1.24.5-alpine AS builder

# Instala build essentials
RUN apk add --no-cache git ca-certificates tzdata

WORKDIR /build

# Copia go.mod e go.sum para aproveitar cache do Docker
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copia código-fonte
COPY . .

# Compila o binário com otimizações para compactação
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags="-w -s" \
    -o /build/api \
    ./cmd/api

# Stage 2: Runtime
FROM alpine:3.19

# Instala apenas certificados e timezone (para logs com data correta)
RUN apk add --no-cache ca-certificates tzdata

# Cria usuário non-root por segurança
RUN addgroup -g 1000 -S appgroup && \
    adduser -u 1000 -S appuser -G appgroup

WORKDIR /app

# Copia binário do stage de build
COPY --from=builder /build/api ./

# Muda propriedário do arquivo
RUN chown -R appuser:appgroup /app

# Muda para usuário non-root
USER appuser

# Expõe porta
EXPOSE 8081

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD wget --no-verbose --tries=1 --spider http://localhost:8081/ping || exit 1

# Executa aplicação
CMD ["./api"]
