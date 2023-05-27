SCHEMA_NAME := ""
# Run your server
run: 
	go run ./cmd/main.go 
# Generate Ent boilerplate
generate:
	go generate ./...
	
# Generates a new ent schema in target directory 
generate-schema:
	@read -p "Enter the schema name: " SCHEMA_NAME; \
    echo "creating ent schema: $$SCHEMA_NAME"; \
	eval "go run -mod=mod entgo.io/ent/cmd/ent new --target ./app/database/schema/ $$SCHEMA_NAME"