#!/bin/bash

# generate-docs.sh - Generate OpenAPI/Swagger documentation for Voiceflow CLI API

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

echo -e "${BLUE}🚀 Generating OpenAPI/Swagger documentation...${NC}"

# Check if swag is installed
if ! command -v swag &> /dev/null; then
    echo -e "${YELLOW}⚠️  swag command not found. Installing...${NC}"
    go install github.com/swaggo/swag/cmd/swag@latest
    if ! command -v swag &> /dev/null; then
        echo -e "${RED}❌ Failed to install swag. Please ensure Go is installed and GOPATH/bin is in your PATH${NC}"
        exit 1
    fi
    echo -e "${GREEN}✅ swag installed successfully${NC}"
fi

# Get the script directory (where this script is located)
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
# Get the project root (one level up from scripts)
PROJECT_ROOT="$(cd "${SCRIPT_DIR}/.." && pwd)"

echo -e "${BLUE}📁 Project root: ${PROJECT_ROOT}${NC}"

# Change to project root
cd "${PROJECT_ROOT}"

# Generate the docs
echo -e "${BLUE}📝 Running swag init...${NC}"
swag init \
    --generalInfo server/server.go \
    --output server/docs \
    --outputTypes go,json,yaml \
    --parseInternal \
    --parseDependency \
    --markdownFiles ./server/docs \
    --codeExampleFiles ./examples

# Check if generation was successful
if [ $? -eq 0 ]; then
    echo -e "${GREEN}✅ OpenAPI documentation generated successfully!${NC}"
    
    # Show generated files
    echo -e "${BLUE}📄 Generated files:${NC}"
    ls -la server/docs/
    
    echo -e ""
    echo -e "${GREEN}🎉 Documentation is now available at:${NC}"
    echo -e "${BLUE}   • Swagger UI: http://localhost:8080/swagger/index.html${NC}"
    echo -e "${BLUE}   • OpenAPI JSON: server/docs/swagger.json${NC}"
    echo -e "${BLUE}   • OpenAPI YAML: server/docs/swagger.yaml${NC}"
    echo -e "${BLUE}   • Go docs: server/docs/docs.go${NC}"
else
    echo -e "${RED}❌ Failed to generate documentation${NC}"
    exit 1
fi
