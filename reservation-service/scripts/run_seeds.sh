#!/bin/bash

# Seed runner for inventory-service database
# Usage: ./run_seeds.sh [environment]
# Environments: local, dev, staging, prod

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Configuration
ENV=${1:-local}
SEEDS_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)/../db/seeds"

# Load environment configuration
case $ENV in
    local)
        DB_HOST="127.0.0.1"
        DB_PORT="3306"
        DB_USER="app"
        DB_PASSWORD="app"
        DB_NAME="ticketflow_reservation"
        echo -e "${YELLOW}Running seeds for LOCAL environment${NC}"
        ;;
    dev)
        DB_HOST="${DB_HOST_DEV}"
        DB_PORT="${DB_PORT_DEV:-3306}"
        DB_USER="${DB_USER_DEV}"
        DB_PASSWORD="${DB_PASSWORD_DEV}"
        DB_NAME="${DB_NAME_DEV}"
        echo -e "${YELLOW}Running seeds for DEV environment${NC}"
        ;;
    staging)
        DB_HOST="${DB_HOST_STAGING}"
        DB_PORT="${DB_PORT_STAGING:-3306}"
        DB_USER="${DB_USER_STAGING}"
        DB_PASSWORD="${DB_PASSWORD_STAGING}"
        DB_NAME="${DB_NAME_STAGING}"
        echo -e "${YELLOW}Running seeds for STAGING environment${NC}"
        ;;
    prod)
        echo -e "${RED}ERROR: Seeds in production require explicit confirmation${NC}"
        echo "Run with: CONFIRM_PROD=true ./run_seeds.sh prod"
        exit 1
        ;;
    *)
        echo -e "${RED}Unknown environment: $ENV${NC}"
        echo "Supported: local, dev, staging, prod"
        exit 1
        ;;
esac

# Validate seeds directory exists
if [ ! -d "$SEEDS_DIR" ]; then
    echo -e "${YELLOW}⚠ Seeds directory not found: $SEEDS_DIR${NC}"
    echo "Creating seeds directory..."
    mkdir -p "$SEEDS_DIR"
    echo -e "${GREEN}✓ Seeds directory created${NC}"
    exit 0
fi

# Check if there are any seed files
if ! ls "$SEEDS_DIR"/*.sql 1>/dev/null 2>&1; then
    echo -e "${YELLOW}⚠ No seed files found in $SEEDS_DIR${NC}"
    exit 0
fi

# Validate database connection
echo "Validating database connection to $DB_HOST:$DB_PORT..."
if ! docker exec -i ticketflow-mysql mysql -h "$DB_HOST" -P "$DB_PORT" -u "$DB_USER" -p"$DB_PASSWORD" -e "SELECT 1" &>/dev/null; then
    echo -e "${RED}✗ Failed to connect to database${NC}"
    exit 1
fi
echo -e "${GREEN}✓ Database connection successful${NC}"

# Run seeds
echo ""
echo -e "${BLUE}Running seeds from: $SEEDS_DIR${NC}"
seed_count=0
for seed_file in "$SEEDS_DIR"/*.sql; do
    filename=$(basename "$seed_file")
    
    echo "Running $filename..."
    docker exec -i ticketflow-mysql mysql -h "$DB_HOST" -P "$DB_PORT" -u "$DB_USER" -p"$DB_PASSWORD" "$DB_NAME" < "$seed_file" 2>&1
    
    if [ $? -eq 0 ]; then
        echo -e "${GREEN}✓ $filename completed${NC}"
        ((seed_count++))
    else
        echo -e "${RED}✗ $filename failed${NC}"
        exit 1
    fi
done

echo ""
echo -e "${GREEN}✓ All $seed_count seeds completed successfully${NC}"
