#!/bin/bash
# Script para iniciar ambos os serviços do Ticketflow

set -e

# Cores para output
GREEN='\033[0;32m'
BLUE='\033[0;34m'
RED='\033[0;31m'
NC='\033[0m' # No Color

echo -e "${BLUE}========================================${NC}"
echo -e "${BLUE}Ticketflow - Microservices Startup${NC}"
echo -e "${BLUE}========================================${NC}"
echo ""

# Função para iniciar um serviço
start_service() {
    local service_name=$1
    local service_path=$2
    local port=$3
    
    echo -e "${GREEN}▶ Iniciando $service_name (porta $port)...${NC}"
    
    cd "$script_dir/$service_path"
    
    # Verificar se go.mod existe
    if [ ! -f "go.mod" ]; then
        echo -e "${RED}✗ Erro: go.mod não encontrado em $service_path${NC}"
        exit 1
    fi
    
    # Iniciar em background
    go run cmd/api/main.go &
    local pid=$!
    echo -e "${GREEN}✓ $service_name iniciado (PID: $pid)${NC}"
    
    # Aguardar até que o serviço esteja respondendo
    echo "  Aguardando $service_name estar pronto..."
    sleep 2
    
    for i in {1..10}; do
        if curl -s http://localhost:$port/ping > /dev/null 2>&1; then
            echo -e "${GREEN}✓ $service_name está pronto!${NC}"
            return 0
        fi
        echo "  Tentativa $i/10..."
        sleep 1
    done
    
    echo -e "${RED}✗ Timeout aguardando $service_name${NC}"
    kill $pid 2>/dev/null
    return 1
}

# Obter diretório do script
script_dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

# Verificar se estamos na pasta correta
if [ ! -d "reservation-service" ] || [ ! -d "inventory-service" ]; then
    echo -e "${RED}✗ Erro: Não encontradas pastas de serviços em $script_dir${NC}"
    echo "  Este script deve ser executado a partir de ticketflow/"
    exit 1
fi

echo -e "${BLUE}📁 Diretório: $script_dir${NC}"
echo ""

# Iniciar Inventory Service
start_service "Inventory Service" "inventory-service" 8081
if [ $? -ne 0 ]; then
    echo -e "${RED}✗ Falha ao iniciar Inventory Service${NC}"
    exit 1
fi

echo ""

# Iniciar Reservation Service
start_service "Reservation Service" "reservation-service" 8080
if [ $? -ne 0 ]; then
    echo -e "${RED}✗ Falha ao iniciar Reservation Service${NC}"
    exit 1
fi

echo ""
echo -e "${GREEN}========================================${NC}"
echo -e "${GREEN}✓ Todos os serviços iniciados com sucesso!${NC}"
echo -e "${GREEN}========================================${NC}"
echo ""
echo "Endpoints disponíveis:"
echo -e "  ${BLUE}Reservation Service${NC}:   http://localhost:8080"
echo -e "  ${BLUE}Inventory Service${NC}:      http://localhost:8081"
echo ""
echo "Exemplo de requisição:"
echo "  curl -X POST http://localhost:8080/events/1/reserve \\"
echo "    -H 'Content-Type: application/json' \\"
echo "    -d '{\"event_id\": 1, \"tickets\": {\"GENERAL\": 2}}'"
echo ""
echo "Pressione Ctrl+C para parar os serviços"
echo ""

# Aguardar infinitamente
wait
