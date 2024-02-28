# Run this command from root project directory
# sh scripts/swag_init.sh

swag init --dir ./ -g ./cmd/main.go --output ./docs/apidocs --parseDependency --parseInternal --parseDepth 8
