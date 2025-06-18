#!/bin/bash

# validate-openapi.sh - Validate the generated OpenAPI specification

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

echo -e "${BLUE}🔍 Validating OpenAPI specification...${NC}"

# Get the script directory and project root
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(cd "${SCRIPT_DIR}/.." && pwd)"

cd "${PROJECT_ROOT}"

# Check if swagger files exist
SWAGGER_JSON="server/docs/swagger.json"
SWAGGER_YAML="server/docs/swagger.yaml"

if [ ! -f "${SWAGGER_JSON}" ]; then
    echo -e "${RED}❌ swagger.json not found. Run 'make docs' first.${NC}"
    exit 1
fi

if [ ! -f "${SWAGGER_YAML}" ]; then
    echo -e "${RED}❌ swagger.yaml not found. Run 'make docs' first.${NC}"
    exit 1
fi

echo -e "${GREEN}✅ Swagger files found${NC}"

# Basic JSON validation
echo -e "${BLUE}📄 Validating JSON syntax...${NC}"
if jq empty "${SWAGGER_JSON}" 2>/dev/null; then
    echo -e "${GREEN}✅ swagger.json is valid JSON${NC}"
else
    echo -e "${RED}❌ swagger.json contains invalid JSON${NC}"
    exit 1
fi

# Basic YAML validation
echo -e "${BLUE}📄 Validating YAML syntax...${NC}"
if python3 -c "import yaml; yaml.safe_load(open('${SWAGGER_YAML}'))" 2>/dev/null; then
    echo -e "${GREEN}✅ swagger.yaml is valid YAML${NC}"
elif python -c "import yaml; yaml.safe_load(open('${SWAGGER_YAML}'))" 2>/dev/null; then
    echo -e "${GREEN}✅ swagger.yaml is valid YAML${NC}"
else
    echo -e "${YELLOW}⚠️  Could not validate YAML (Python with PyYAML not available)${NC}"
fi

# Validate OpenAPI structure
echo -e "${BLUE}📋 Checking OpenAPI structure...${NC}"

# Check required fields
REQUIRED_FIELDS=("swagger" "info" "paths")
for field in "${REQUIRED_FIELDS[@]}"; do
    if jq -e ".${field}" "${SWAGGER_JSON}" >/dev/null 2>&1; then
        echo -e "${GREEN}✅ Field '${field}' present${NC}"
    else
        echo -e "${RED}❌ Required field '${field}' missing${NC}"
        exit 1
    fi
done

# Check info object required fields
INFO_FIELDS=("title" "version")
for field in "${INFO_FIELDS[@]}"; do
    if jq -e ".info.${field}" "${SWAGGER_JSON}" >/dev/null 2>&1; then
        echo -e "${GREEN}✅ Info field '${field}' present${NC}"
    else
        echo -e "${RED}❌ Required info field '${field}' missing${NC}"
        exit 1
    fi
done

# Count paths and definitions
PATHS_COUNT=$(jq '.paths | length' "${SWAGGER_JSON}")
DEFINITIONS_COUNT=$(jq '.definitions | length' "${SWAGGER_JSON}")

echo -e "${BLUE}📊 OpenAPI Statistics:${NC}"
echo -e "   • Paths: ${PATHS_COUNT}"
echo -e "   • Definitions: ${DEFINITIONS_COUNT}"
echo -e "   • API Version: $(jq -r '.info.version' "${SWAGGER_JSON}")"
echo -e "   • API Title: $(jq -r '.info.title' "${SWAGGER_JSON}")"

# List all endpoints
echo -e "${BLUE}🔗 Available Endpoints:${NC}"
jq -r '.paths | keys[]' "${SWAGGER_JSON}" | while read -r path; do
    METHODS=$(jq -r ".paths[\"${path}\"] | keys | join(\", \")" "${SWAGGER_JSON}")
    echo -e "   • ${path} [${METHODS}]"
done

# Advanced validation with swagger CLI if available
if command -v swagger >/dev/null 2>&1; then
    echo -e "${BLUE}🔍 Running advanced validation with swagger CLI...${NC}"
    if swagger validate "${SWAGGER_YAML}"; then
        echo -e "${GREEN}✅ Advanced validation passed${NC}"
    else
        echo -e "${RED}❌ Advanced validation failed${NC}"
        exit 1
    fi
else
    echo -e "${YELLOW}⚠️  swagger CLI not available for advanced validation${NC}"
    echo -e "${YELLOW}   Install with: go install github.com/go-swagger/go-swagger/cmd/swagger@latest${NC}"
fi

# Check for common issues
echo -e "${BLUE}🔍 Checking for common issues...${NC}"

# Check if all paths have at least one method
INVALID_PATHS=$(jq -r '.paths | to_entries[] | select(.value | length == 0) | .key' "${SWAGGER_JSON}")
if [ -n "${INVALID_PATHS}" ]; then
    echo -e "${RED}❌ Found paths without methods:${NC}"
    echo "${INVALID_PATHS}"
    exit 1
else
    echo -e "${GREEN}✅ All paths have valid methods${NC}"
fi

# Check if all definitions are referenced
UNUSED_DEFINITIONS=$(jq -r '
    .definitions | keys[] as $def |
    select(
        [.. | strings | test("#/definitions/" + $def)] | length == 0
    ) | $def
' "${SWAGGER_JSON}")

if [ -n "${UNUSED_DEFINITIONS}" ]; then
    echo -e "${YELLOW}⚠️  Found potentially unused definitions:${NC}"
    echo "${UNUSED_DEFINITIONS}"
else
    echo -e "${GREEN}✅ All definitions appear to be referenced${NC}"
fi

echo -e ""
echo -e "${GREEN}🎉 OpenAPI specification validation completed successfully!${NC}"
echo -e ""
echo -e "${BLUE}📖 Next steps:${NC}"
echo -e "${BLUE}   • Start server: make dev${NC}"
echo -e "${BLUE}   • View Swagger UI: http://localhost:8080/swagger/index.html${NC}"
echo -e "${BLUE}   • Download spec: http://localhost:8080/swagger/doc.json${NC}"
