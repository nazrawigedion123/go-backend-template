MIGRATION_DIR := internal/constant/db/schemas
SWAG_CMD := swag
SWAG_INIT_FLAGS := -g cmd/main.go
.PHONY: create-migration-dir

DB_URL := postgresql://postgres:secret@172.28.22.99:5436/audit_db
SCHEMA_FILE := audit_schema.sql
MIGRATION_PATH := platform/auditor/query/schema

create-migration-dir:
	@mkdir -p $(MIGRATION_DIR)/up
	@mkdir -p $(MIGRATION_DIR)/down
	@echo "Created directory: $(MIGRATION_DIR)"
# It depends on the 'create-migration-dir' target.
create-migration-file: create-migration-dir

	@touch $(MIGRATION_DIR)/up/$(shell date +'%Y%m%d%H%M%S')_$(name).up.sql
	@touch $(MIGRATION_DIR)/$(shell date +'%Y%m%d%H%M%S')_$(name).down.sql
	@echo "Created new migration file in $(MIGRATION_DIR)/up and $(MIGRATION_DIR)/up"

.PHONY: clean


clean:
	@rm -rf $(MIGRATION_DIR)
	@echo "Cleaned up migrations directory."

migrate-down:
	- migrate -database postgresql://postgres:secret@172.28.22.99:5436/mobile?sslmode=disable -path internal/constant/db/schemas -verbose down $(N)
migrate-up:
	- migrate -database postgresql://postgres:secret@172.28.22.99:5436/mobile?sslmode=disable -path internal/constant/db/schemas -verbose up
migrate-down-test:
	- migrate -postgresql://postgres:secret@172.28.22.99:5436/mobile?sslmode=disable -path internal/constant/db/schemas -verbose down
migrate-up-test:
	- migrate -postgresql://postgres:secret@172.28.22.99:5436/mobile?sslmode=disable -path internal/constant/db/schemas -verbose up
migrate-create:
	- migrate create -ext sql -dir internal/constant/db/schemas -tz "UTC" $(name)
swagger:
	-swag fmt && swag init -g cmd/main.go
run:
	go run cmd/main.go
sqlc:
	cd ./config && sqlc generate
air:
	@echo "Running air..."
	air -c .air.toml
up:
	@echo "Starting Docker images..."
	docker-compose -f docker-compose.yaml up --build -d
	@echo "Docker images started!"
down:
	@echo "Stopping docker compose..."
	docker-compose -f docker-compose.yaml down
	@echo "Done!"
test:
	go test -v $(path) | grep -v '"level"' | grep -v 'Error #'

# Mock generation
.PHONY: generate-mocks
generate-mocks: install-mockgen
	@echo "Generating mocks..."
	@mkdir -p internal/storage/mocks platform/auditor/mocks
	@echo "Generating storage mocks..."
	@mockgen -source=internal/storage/storage.go -destination=internal/storage/mocks/storage_mock.go -package=mocks
	@echo "Generating module mocks..."
	@mockgen -source=internal/module/module.go -destination=internal/module/mocks/module_mock.go -package=mocks
	@echo "Generating handler mocks..."
	@mockgen -source=internal/handler/handler.go -destination=internal/handler/mocks/handler_mock.go -package=mocks
	@echo "Generating auditor mocks..."
	@mockgen -source=platform/auditor/auditor.go -destination=platform/auditor/mocks/auditor_mock.go -package=mocks
	@echo "Mocks generated successfully!"

# Install mockgen if not already installed (tries both versions)
.PHONY: install-mockgen
install-mockgen:
	@echo "Checking for mockgen..."
	@which mockgen > /dev/null 2>&1 || (echo "Installing mockgen from go.uber.org/mock..." && go install go.uber.org/mock/mockgen@latest)
	@echo "mockgen ready!"


# ----------------------------------------------------------------------
# 1. VERSIONED MIGRATION (Requires 'migrate' CLI tool like golang-migrate)
# ----------------------------------------------------------------------

# Target to run UP migrations using the 'migrate' tool
migrate-audit-up:
	@echo "--- Starting Versioned Database Migration (UP) ---"
	migrate -database "$(DB_URL)?sslmode=disable" -path "$(MIGRATION_PATH)" -verbose up
	@echo "--- Migration Complete! (Versioned) ---"
