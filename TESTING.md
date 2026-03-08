# 🧪 Guia de Testes - Ticketflow

## Preparação

### 1. Iniciar Banco de Dados
```bash
cd ticketflow/reservation-service
docker-compose up -d

# Verificar se MySQL está rodando
docker ps | grep mysql
```

### 2. Iniciar Serviços
```bash
cd ticketflow

# Opção 1: Via script
chmod +x start.sh
./start.sh

# Opção 2: Manual
# Terminal 1
cd inventory-service && go run cmd/api/main.go

# Terminal 2
cd reservation-service && go run cmd/api/main.go
```

### 3. Verificar Serviços
```bash
# Ambos devem responder com "pong"
curl http://localhost:8080/ping
curl http://localhost:8081/ping
```

---

## 📋 Testes Básicos

### Test 1: Consumir Ticket (Inventory Service)
```bash
curl -X POST http://localhost:8081/events/1/consume \
  -H "Content-Type: application/json" \
  -H "X-Request-ID: test-001" \
  -d '{
    "event_id": 1,
    "tickets": {
      "GENERAL": 2,
      "VIP": 1
    }
  }'
```

**Esperado:**
```json
{
  "message": "Tickets consumed successfully"
}
```

**Status**: 200 OK

---

### Test 2: Reservar Ticket (Reservation Service)
```bash
curl -X POST http://localhost:8080/events/1/reserve \
  -H "Content-Type: application/json" \
  -H "X-Request-ID: test-002" \
  -d '{
    "event_id": 1,
    "tickets": {
      "GENERAL": 2,
      "VIP": 1
    }
  }'
```

**Esperado:**
```json
{
  "message": "Tickets reserved successfully",
  "data": {
    "ticket_ids": [
      "uuid-1",
      "uuid-2",
      "uuid-3"
    ]
  }
}
```

**Status**: 200 OK

---

## 🔍 Testes de Cenários

### Scenario 1: Fluxo Completo (Sucesso)

**Pré-requisito**: Event ID 1 deve existir no BD com tickets disponíveis

```bash
# Passo 1: Reservar tickets
RESERVATION=$(curl -s -X POST http://localhost:8080/events/1/reserve \
  -H "Content-Type: application/json" \
  -d '{
    "event_id": 1,
    "tickets": {"GENERAL": 1}
  }')

echo "Reservation Response:"
echo $RESERVATION | jq .

# Passo 2: Extrair IDs de ticket
TICKET_IDS=$(echo $RESERVATION | jq -r '.data.ticket_ids[]')
echo "Ticket IDs: $TICKET_IDS"

# Passo 3: Verificar em logs que inventory foi chamado
tail -f logs/inventory.log | grep consumed
```

---

### Scenario 2: Quantidade Insuficiente

**Pré-requisito**: Tickets já consumidos além do total

```bash
curl -X POST http://localhost:8080/events/1/reserve \
  -H "Content-Type: application/json" \
  -d '{
    "event_id": 1,
    "tickets": {
      "GENERAL": 10000
    }
  }'
```

**Esperado**: 500 com mensagem de erro
```json
{
  "error": "There are not enough tickets available for the requested ticket type"
}
```

---

### Scenario 3: Evento Não Encontrado

```bash
curl -X POST http://localhost:8080/events/9999/reserve \
  -H "Content-Type: application/json" \
  -d '{
    "event_id": 9999,
    "tickets": {"GENERAL": 1}
  }'
```

**Esperado**: 500 com erro de evento não encontrado

---

### Scenario 4: Tipo de Ticket Inválido

```bash
curl -X POST http://localhost:8080/events/1/reserve \
  -H "Content-Type: application/json" \
  -d '{
    "event_id": 1,
    "tickets": {
      "PLATINUM": 1
    }
  }'
```

**Esperado**: 500 com mensagem de tipo inválido

---

## 🔗 Testes de Integração

### Test: Communicação Entre Serviços

```bash
#!/bin/bash

echo "=== Testing Service Communication ==="
echo ""

# 1. Check inventory service is up
echo "1. Checking Inventory Service..."
INV_HEALTH=$(curl -s -o /dev/null -w "%{http_code}" http://localhost:8081/ping)
echo "   Status: $INV_HEALTH"

# 2. Check reservation service is up
echo "2. Checking Reservation Service..."
RES_HEALTH=$(curl -s -o /dev/null -w "%{http_code}" http://localhost:8080/ping)
echo "   Status: $RES_HEALTH"

# 3. Test end-to-end flow
echo ""
echo "3. Testing End-to-End Flow..."
RESPONSE=$(curl -s -X POST http://localhost:8080/events/1/reserve \
  -H "Content-Type: application/json" \
  -d '{
    "event_id": 1,
    "tickets": {"GENERAL": 1}
  }')

echo "Response: $RESPONSE"

# 4. Verify the response
if echo "$RESPONSE" | grep -q "ticket_ids"; then
    echo "✓ Integration test PASSED"
else
    echo "✗ Integration test FAILED"
fi
```

---

## 📊 Verificações no Banco de Dados

### Conectar ao MySQL
```bash
docker exec -it mysql mysql -u app -p app

# Verificar events
SELECT * FROM events;

# Verificar tickets
SELECT * FROM tickets;

# Verificar inventário
SELECT * FROM ticket_catalog_items;

# Verificar histórico de mudanças
SELECT * FROM tickets ORDER BY id DESC LIMIT 5;
```

---

## 🐛 Debugging

### Ver logs da Reservation Service
```bash
# Com correlação de IDs
curl -X POST http://localhost:8080/events/1/reserve \
  -H "Content-Type: application/json" \
  -H "X-Request-ID: debug-001" \
  -d '{"event_id": 1, "tickets": {"GENERAL": 1}}'

# Procurar nos logs pela correlação
echo "Procure por [debug-001] nos logs"
```

### Habilitar DEBUG Logging
```bash
# Na variável de ambiente
export LOG_LEVEL=debug

# Reiniciar serviços
go run cmd/api/main.go
```

### Monitorar Requisições HTTP
```bash
# Terminal separado com tcpdump (se disponível)
tcpdump -i lo -A 'tcp port 8080 or tcp port 8081'

# Ou usar curl com verbose
curl -v -X POST http://localhost:8080/events/1/reserve \
  -H "Content-Type: application/json" \
  -d '{"event_id": 1, "tickets": {"GENERAL": 1}}'
```

---

## ✅ Checklist de Testes

- [ ] Serviço Inventory responde em porta 8081
- [ ] Serviço Reservation responde em porta 8080
- [ ] MySQL está rodando
- [ ] Consumo de ticket funciona
- [ ] Reserva de ticket funciona
- [ ] Integração entre serviços funciona
- [ ] Headers de correlação são propagados
- [ ] Erros são tratados corretamente
- [ ] Inventário é verificado antes de reservar
- [ ] Transações são ACID

---

## 🚨 Troubleshooting

### Porta Já Está em Uso
```bash
# Encontrar processo usando a porta
lsof -i :8080
lsof -i :8081

# Matar processo
kill -9 <PID>
```

### Falha de Conexão com Banco
```bash
# Verificar se MySQL está rodando
docker ps

# Reiniciar MySQL
docker-compose restart mysql

# Verificar logs
docker logs mysql
```

### Erro de Módulo Go
```bash
# Atualizar dependências
go mod tidy

# Baixar módulos novamente
go mod download
```

### Serviço não inicia
```bash
# Verificar erros de compilação
go build ./cmd/api

# Ver logs completos
go run cmd/api/main.go 2>&1 | head -50
```

---

## 📞 Recursos Úteis

- [Documentação: ARCHITECTURE.md](../ARCHITECTURE.md)
- [Documentação: REFACTORING_SUMMARY.md](../REFACTORING_SUMMARY.md)
- [API Endpoints](#)
- [Schema BD](#)

---

**Última Atualização**: Março 2026
