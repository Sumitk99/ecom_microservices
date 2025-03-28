# Define variables for each service
ACCOUNT_CMD = account/cmd/account/main.go
CATALOG_CMD = catalog/cmd/catalog/main.go
CART_CMD = cart/cmd/cart/main.go
ORDER_CMD = order/cmd/order/main.go
GATEWAY_CMD = gateway/cmd/gateway/main.go

# Targets to run each service individually
run-account:
	@echo "Starting account service..."
	go run $(ACCOUNT_CMD)

run-catalog:
	@echo "Starting catalog service..."
	go run $(CATALOG_CMD)

run-cart:
	@echo "Starting cart service..."
	go run $(CART_CMD)

run-order:
	@echo "Starting order service..."
	go run $(ORDER_CMD)

# Target to run the API gateway
run-gateway:
	@echo "Starting API gateway..."
	go run $(GATEWAY_CMD)

# Target to run all services (including gateway) concurrently
run-all:
	@echo "Starting all services..."
	@make -s run-gateway & \
	make -s run-account & \
	make -s run-catalog & \
	make -s run-cart & \
	make -s run-order

# Target to stop all services
stop-all:
	@echo "Stopping all services"
	@pkill -f $(ACCOUNT_CMD) || true
	@pkill -f $(CATALOG_CMD) || true
	@pkill -f $(CART_CMD) || true
	@pkill -f $(ORDER_CMD) || true
	@pkill -f $(GATEWAY_CMD) || true

# Phony targets
.PHONY: run-account run-catalog run-cart run-order run-gateway run-all stop-all