# Ticketflow - Microservices

Plataforma de gerenciamento de vendas e reservas de tickets com arquitetura em microserviços.

## 🏗️ Arquitetura

### Serviços
```
ticketflow/
├── reservation-service/    (Porta 8080)
│   └─ Gerencia reservas de tickets
│       
└── inventory-service/      (Porta 8081)
    └─ Gerencia disponibilidade de tickets
```

## 🔄 Fluxo de Requisição

```
Cliente
    │
    └─→ POST /events/:id/reserve [Reservation]
        │
        ├─ Valida quantidade
        │
        ├─→ POST /events/:id/consume [Inventory]  ← Chamada síncrona
        │   ├─ Valida disponibilidade
        │   ├─ Decrementa quantidade
        │   └─ Retorna sucesso/erro
        │
        ├─ Se sucesso: Cria reserva no BD
        │
        └─ Retorna ticket_ids
```

## 🚀 Como Executar

### Pré-requisitos
- Go 1.24+
- Docker & Docker Compose
- MySQL 8.0+

### Start dos Serviços

**Option 1: Via Docker Compose**
```bash
cd reservation-service
docker-compose up -d

cd inventory-service
docker-compose up -d
```

**Option 2: Via Go CLI**
```bash
# Terminal 1 - Inventory
cd inventory-service
go run cmd/api/main.go

# Terminal 2 - Reservation
cd reservation-service
go run cmd/api/main.go
```

## 📝 Exemplos de API

### Consumir Tickets (Inventory Service)
```bash
curl -X POST http://localhost:8081/events/1/consume \
  -H "Content-Type: application/json" \
  -d '{
    "event_id": 1,
    "tickets": {
      "GENERAL": 2,
      "VIP": 1
    }
  }'
```

**Resposta:**
```json
{
  "message": "Tickets consumed successfully"
}
```

### Reservar Tickets (Reservation Service)
```bash
curl -X POST http://localhost:8080/events/1/reserve \
  -H "Content-Type: application/json" \
  -d '{
    "event_id": 1,
    "tickets": {
      "GENERAL": 2,
      "VIP": 1
    }
  }'
```

**Resposta:**
```json
{
  "message": "Tickets reserved successfully",
  "data": {
    "ticket_ids": ["uuid1", "uuid2", "uuid3"]
  }
}
```

## 📊 Estrutura de Banco de Dados

Ambos os serviços compartilham a mesma instância MySQL:

```
Database: app
├── events
│   ├── id (PK)
│   ├── name
│   ├── starts_at
│   ├── ends_at
│   ├── venue_name
│   ├── address_id (FK)
│   ├── open_sales_at
│   └── close_sales_at
│
├── ticket_catalog_items
│   ├── id (PK)
│   ├── event_id (FK)
│   ├── type
│   ├── price
│   ├── total_quantity
│   └── sold_quantity
│
├── tickets
│   ├── id (PK)
│   ├── event_id (FK)
│   ├── type
│   └── price
│
└── addresses
    ├── id (PK)
    ├── street
    ├── number
    ├── city
    ├── state
    └── zip_code
```

## 🔒 Segurança e Transações

### Unit of Work Pattern
- Ambos os serviços implementam `UnitOfWork` para transações ACID
- Garante consistência de dados

### Inventory Control
- Inventory Service: Decrementa com validação
  ```sql
  sold_quantity = sold_quantity + ?
  WHERE (sold_quantity + ?) <= total_quantity
  ```
- Previne overselling

### Correlation Tracking
- Header `X-Request-ID` propagado entre serviços
- Facilita rastreamento de requisições em logs

## 📚 Tipos de Tickets

```go
const (
    TicketTypeGeneral TicketType = "GENERAL"
    TicketTypeVIP     TicketType = "VIP"
    TicketTypeBoxSeat TicketType = "BOX_SEAT"
)
```

## 🛠️ Dependências

### Principais
- `github.com/gin-gonic/gin` - Framework web
- `github.com/go-sql-driver/mysql` - Driver MySQL
- `github.com/google/uuid` - Geração de UUIDs
- `github.com/lucasbrito3001/go-kit` - Shared utilities (logging, observability)

## 📖 Documentação Adicional

- [ARCHITECTURE.md](../ARCHITECTURE.md) - Visão geral da arquitetura
- [REFACTORING_SUMMARY.md](../REFACTORING_SUMMARY.md) - Detalhes da refatoração
- [CHECKLIST.md](../CHECKLIST.md) - Checklist de implementação

## 🌐 Variáveis de Ambiente

```env
# MySQL
MYSQL_USER=app
MYSQL_PASSWORD=app
MYSQL_DATABASE=app
MYSQL_HOST=localhost
MYSQL_PORT=3306

# Logging
LOG_LEVEL=info  # debug, info, warn, error

# Services
INVENTORY_SERVICE_URL=http://localhost:8081
```

## 🧪 Testes

### Inventory Service
```bash
cd inventory-service
go test ./...
```

### Reservation Service
```bash
cd reservation-service
go test ./...
```

## 📈 Monitoramento

### Health Checks
```bash
curl http://localhost:8080/ping
curl http://localhost:8081/ping
```

### Logs
- Ambos usam `slog` (structured logging)
- Correlação de requisições via `X-Request-ID`
- Propagação automática entre serviços

## 🔄 Integração

### Sincronização
- Reservation Service → Inventory Service (HTTP)
- Síncrono para garantir consistência imediata
- Falha de Inventory causa falha de Reservation

### Futuro: Assíncrono
- Considerar Event Bus (Kafka/RabbitMQ)
- Maior desacoplamento
- Melhor resiliência

## 📞 Suporte

Para dúvidas sobre a arquitetura, consulte:
1. [REFACTORING_SUMMARY.md](../REFACTORING_SUMMARY.md)
2. Código comentado em `internal/app/usecases/`
3. Testes de integração (quando implementados)

---

**Última Atualização**: Março 2026
