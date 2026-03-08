.PHONY: help build run clean test docker-build docker-run docker-compose-up docker-compose-down

# Variáveis
BINARY_NAME=api
DOCKER_IMAGE=inventory-service:latest
DOCKER_REGISTRY=
SERVICE_PORT=8081

help:
	@echo "Inventory Service - Available commands:"
	@echo ""
	@echo "Development:"
	@echo "  make build           - Build local binary"
	@echo "  make run             - Run application locally"
	@echo "  make clean           - Remove build artifacts"
	@echo "  make test            - Run tests"
	@echo ""
	@echo "Docker:"
	@echo "  make docker-build    - Build Docker image"
	@echo "  make docker-run      - Run Docker container"
	@echo "  make docker-push     - Push to registry"
	@echo ""
	@echo "Compose:"
	@echo "  make compose-up      - Start services with docker-compose"
	@echo "  make compose-down    - Stop services with docker-compose"
	@echo "  make compose-logs    - View compose logs"
	@echo ""

build:
	@echo "Building binary..."
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
		-ldflags="-w -s" \
		-o bin/$(BINARY_NAME) \
		./cmd/api
	@echo "✓ Binary built: bin/$(BINARY_NAME)"

run: build
	@echo "Running application..."
	./bin/$(BINARY_NAME)

clean:
	@echo "Cleaning..."
	rm -rf bin/
	go clean
	@echo "✓ Clean complete"

test:
	@echo "Running tests..."
	go test -v -cover ./...

docker-build:
	@echo "Building Docker image: $(DOCKER_IMAGE)"
	docker build -t $(DOCKER_IMAGE) .
	@echo "✓ Image built: $(DOCKER_IMAGE)"

docker-run: docker-build
	@echo "Running Docker container..."
	docker run -p $(SERVICE_PORT):$(SERVICE_PORT) \
		--name inventory-service \
		--rm \
		$(DOCKER_IMAGE)

docker-push: docker-build
	@echo "Pushing to registry: $(DOCKER_REGISTRY)$(DOCKER_IMAGE)"
	docker push $(DOCKER_REGISTRY)$(DOCKER_IMAGE)
	@echo "✓ Image pushed"

compose-up:
	@echo "Starting services with docker-compose..."
	docker-compose up -d
	@echo "✓ Services started"

compose-down:
	@echo "Stopping services..."
	docker-compose down
	@echo "✓ Services stopped"

compose-logs:
	docker-compose logs -f inventory-service

fmt:
	@echo "Formatting code..."
	go fmt ./...
	@echo "✓ Format complete"

lint:
	@echo "Linting code..."
	golangci-lint run ./...
	@echo "✓ Lint complete"
