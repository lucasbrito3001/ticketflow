# Database Migrations

Schema migrations para o inventory-service.

## 📋 Migrations

| # | Migration | Descrição |
|---|-----------|-----------|
| 001 | `create_table_addresses` | Tabela de endereços para eventos |
| 002 | `create_table_events` | Tabela de eventos com informações de venda |
| 003 | `create_table_ticket_catalog_items` | Tabela de tipos de ticket com inventário |
| 004 | `seed_data` | Dados de teste para desenvolvimento |

## 🚀 Como usar

### Script automatizado (recomendado)

```bash
# Local (default)
./db/migrations/run_migrations.sh

# Desenvolvimento
./db/migrations/run_migrations.sh dev

# Staging (com dados de seed)
./db/migrations/run_migrations.sh staging

# Produção (requer confirmação explícita)
export CONFIRM_PROD=true
./db/migrations/run_migrations.sh prod
```

### Manual com MySQL CLI

```bash
# Executar todas as migrations
mysql -h localhost -u app -p app app < db/migrations/001_create_table_addresses.sql
mysql -h localhost -u app -p app app < db/migrations/002_create_table_events.sql
mysql -h localhost -u app -p app app < db/migrations/003_create_table_ticket_catalog_items.sql
mysql -h localhost -u app -p app app < db/migrations/004_seed_data.sql
```

### Com Docker Compose

```bash
# Migrations rodam automaticamente ao iniciar
docker-compose up -d

# Ou manualmente
docker exec mysql-db mysql -u app -p app app < db/migrations/001_create_table_addresses.sql
```

## 📊 Schema

### addresses
```sql
- id: BIGINT (PK)
- street: VARCHAR(255)
- number: INT
- city: VARCHAR(100)
- state: VARCHAR(2)      -- UF (Estado)
- zip_code: VARCHAR(10)
- created_at, updated_at: TIMESTAMP
```

### events
```sql
- id: BIGINT (PK)
- name: VARCHAR(255)
- starts_at, ends_at: DATETIME
- venue: VARCHAR(255)
- address_id: BIGINT (FK → addresses)
- open_sales_at, close_sales_at: DATETIME
- created_at, updated_at: TIMESTAMP
- deleted_at: TIMESTAMP (soft delete)
```

### ticket_catalog_items
```sql
- id: BIGINT (PK)
- event_id: BIGINT (FK → events, cascade delete)
- type: VARCHAR(50)          -- INTEIRA, MEIA, VIP, etc
- price: BIGINT              -- em centavos (R$ 100.00 = 10000)
- total_quantity: INT
- sold_quantity: INT
- created_at, updated_at: TIMESTAMP
```

## 🔍 Características

✅ **Soft Delete**: Events suportam soft delete com `deleted_at`  
✅ **Foreign Keys**: Relacionamentos com cascata de delete  
✅ **Índices**: Otimizados para queries frequentes  
✅ **Timestamps**: Auditoria automática com `created_at`/`updated_at`  
✅ **COLLATION**: UTF8MB4 para suportar emojis e caracteres especiais  
✅ **Seed Data**: Dados de teste com 3 eventos e 9 tipos de ticket  

## 💰 Preços (formato de centavos)

| Tipo | Preço Exemplo | Valor BD |
|------|---------------|----------|
| INTEIRA | R$ 500.00 | 50000 |
| MEIA | R$ 250.00 | 25000 |
| VIP | R$ 1000.00 | 100000 |

## ⚠️ Notas de segurança

- Seed data (`004_seed_data.sql`) é skipped em staging/prod
- Production migrations requerem confirmação explícita (`CONFIRM_PROD=true`)
- Soft delete em events mantém histórico sem apagar registros
- Foreign key constraints protegem integridade referencial

## 📝 Versioning

Migrations seguem padrão numérico: `NNN_description.sql`
- Cada número representa a ordem de execução
- Nunca modificar migrations já executadas
- Sempre criar uma nova migration para alterações
