BINARY_NAME=edushare
BUILD_DIR=bin/build
CMD_DIR=./cmd/edushare
GOOSE=goose

## LOAD ENVIRONMENT VARIABLES
include .env
export

build:
	@echo "Building the project ..."
	@go build -o $(BUILD_DIR)/$(BINARY_NAME) $(CMD_DIR)
	@echo "Build Complete: $(BUILD_DIR)/$(BINARY_NAME)"

run: build
	@echo "Running the project..."
	@$(BUILD_DIR)/$(BINARY_NAME)

clean:
	@echo "Cleaning build artifacts..."
	@rm -rf $(BUILD_DIR)
	@echo "Clean complete."

# Format the code
fmt:
	@echo "Formatting the code..."
	@go fmt ./...
	@echo "Format complete."

test: 

goose-status:
	@echo "Status"
	$(GOOSE) -dir ./internal/migrations postgres $(DB) status  

migrate-up:
	@echo "Migrating up..."
	$(GOOSE) -dir ./internal/migrations postgres $(DB) up  

migrate-down:
	@echo "Migrating down..."
	$(GOOSE) -dir ./internal/migrations postgres $(DB) down  
