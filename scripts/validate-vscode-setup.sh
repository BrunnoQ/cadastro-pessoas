#!/bin/bash

# Colors
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

echo -e "${BLUE}================================================${NC}"
echo -e "${BLUE}VS Code Setup Validation${NC}"
echo -e "${BLUE}================================================${NC}\n"

ERRORS=0
WARNINGS=0

# Function to check file existence
check_file() {
    local file=$1
    local description=$2
    
    if [ -f "$file" ]; then
        echo -e "${GREEN}✓${NC} $description: $file"
    else
        echo -e "${RED}✗${NC} $description: $file (NOT FOUND)"
        ERRORS=$((ERRORS + 1))
    fi
}

# Function to check JSON file validity
check_json() {
    local file=$1
    local description=$2
    
    if [ -f "$file" ]; then
        if python3 -m json.tool "$file" > /dev/null 2>&1; then
            echo -e "${GREEN}✓${NC} $description: Valid JSON"
        else
            echo -e "${RED}✗${NC} $description: Invalid JSON"
            ERRORS=$((ERRORS + 1))
        fi
    fi
}

echo -e "${YELLOW}Checking VS Code configuration files...${NC}\n"

# Check .vscode directory
if [ -d ".vscode" ]; then
    echo -e "${GREEN}✓${NC} .vscode directory exists\n"
else
    echo -e "${RED}✗${NC} .vscode directory NOT FOUND\n"
    ERRORS=$((ERRORS + 1))
    exit 1
fi

# Check configuration files
check_file ".vscode/launch.json" "Debug configurations"
check_file ".vscode/tasks.json" "Task configurations"
check_file ".vscode/settings.json" "Workspace settings"

echo ""

# Validate JSON files
check_json ".vscode/launch.json" "launch.json"
check_json ".vscode/tasks.json" "tasks.json"
check_json ".vscode/settings.json" "settings.json"

echo -e "\n${YELLOW}Checking launch configurations...${NC}\n"

# Count launch configurations
if [ -f ".vscode/launch.json" ]; then
    CONFIG_COUNT=$(grep -c '"name":' .vscode/launch.json || echo "0")
    echo -e "${BLUE}Found $CONFIG_COUNT debug configurations:${NC}"
    grep '"name":' .vscode/launch.json | sed 's/.*"name": "\(.*\)".*/  - \1/'
fi

echo -e "\n${YELLOW}Checking tasks...${NC}\n"

# Count tasks
if [ -f ".vscode/tasks.json" ]; then
    TASK_COUNT=$(grep -c '"label":' .vscode/tasks.json || echo "0")
    echo -e "${BLUE}Found $TASK_COUNT tasks:${NC}"
    grep '"label":' .vscode/tasks.json | sed 's/.*"label": "\(.*\)".*/  - \1/'
fi

echo -e "\n${YELLOW}Checking Go settings...${NC}\n"

# Check Go-specific settings
if [ -f ".vscode/settings.json" ]; then
    if grep -q "go.useLanguageServer" .vscode/settings.json; then
        echo -e "${GREEN}✓${NC} Go language server enabled"
    else
        echo -e "${YELLOW}⚠${NC} Go language server not configured"
        WARNINGS=$((WARNINGS + 1))
    fi
    
    if grep -q "go.formatOnSave" .vscode/settings.json; then
        echo -e "${GREEN}✓${NC} Format on save enabled"
    else
        echo -e "${YELLOW}⚠${NC} Format on save not configured"
        WARNINGS=$((WARNINGS + 1))
    fi
    
    if grep -q "go.lintTool" .vscode/settings.json; then
        LINT_TOOL=$(grep "go.lintTool" .vscode/settings.json | sed 's/.*: "\(.*\)".*/\1/')
        echo -e "${GREEN}✓${NC} Linter configured: $LINT_TOOL"
    else
        echo -e "${YELLOW}⚠${NC} Linter not configured"
        WARNINGS=$((WARNINGS + 1))
    fi
fi

echo -e "\n${YELLOW}Checking required tools...${NC}\n"

# Check Go installation
if command -v go &> /dev/null; then
    GO_VERSION=$(go version | awk '{print $3}')
    echo -e "${GREEN}✓${NC} Go installed: $GO_VERSION"
else
    echo -e "${RED}✗${NC} Go not installed"
    ERRORS=$((ERRORS + 1))
fi

# Check golangci-lint
if command -v golangci-lint &> /dev/null; then
    LINT_VERSION=$(golangci-lint --version | head -1)
    echo -e "${GREEN}✓${NC} golangci-lint installed: $LINT_VERSION"
else
    echo -e "${YELLOW}⚠${NC} golangci-lint not installed (recommended)"
    WARNINGS=$((WARNINGS + 1))
fi

# Check goimports
if command -v goimports &> /dev/null; then
    echo -e "${GREEN}✓${NC} goimports installed"
else
    echo -e "${YELLOW}⚠${NC} goimports not installed (recommended)"
    echo -e "    Install with: go install golang.org/x/tools/cmd/goimports@latest"
    WARNINGS=$((WARNINGS + 1))
fi

# Check air (hot reload)
if command -v air &> /dev/null; then
    echo -e "${GREEN}✓${NC} air installed (hot reload)"
else
    echo -e "${YELLOW}⚠${NC} air not installed (optional)"
    echo -e "    Install with: go install github.com/cosmtrek/air@latest"
    WARNINGS=$((WARNINGS + 1))
fi

# Check Docker
if command -v docker &> /dev/null; then
    echo -e "${GREEN}✓${NC} Docker installed"
else
    echo -e "${RED}✗${NC} Docker not installed (required for tasks)"
    ERRORS=$((ERRORS + 1))
fi

# Check docker-compose (v1 or v2)
if command -v docker-compose &> /dev/null; then
    echo -e "${GREEN}✓${NC} docker-compose installed"
elif docker compose version &> /dev/null; then
    echo -e "${GREEN}✓${NC} docker compose installed (v2)"
else
    echo -e "${RED}✗${NC} docker-compose not installed (required for tasks)"
    ERRORS=$((ERRORS + 1))
fi

echo -e "\n${YELLOW}Checking project files...${NC}\n"

# Check important project files
check_file "go.mod" "Go module file"
check_file "docker-compose.yml" "Docker Compose file"
check_file ".air.toml" "Air configuration (optional)"
check_file "cmd/api/main.go" "Main application file"

echo -e "\n${BLUE}================================================${NC}"
echo -e "${BLUE}Summary${NC}"
echo -e "${BLUE}================================================${NC}\n"

if [ $ERRORS -eq 0 ] && [ $WARNINGS -eq 0 ]; then
    echo -e "${GREEN}✓ VS Code setup is complete and ready!${NC}\n"
    echo -e "${BLUE}Next steps:${NC}"
    echo -e "  1. Open this project in VS Code: ${YELLOW}code .${NC}"
    echo -e "  2. Press ${YELLOW}F5${NC} to start debugging"
    echo -e "  3. Or use ${YELLOW}Ctrl+Shift+P${NC} → 'Tasks: Run Task' to see all available tasks"
    echo ""
    exit 0
elif [ $ERRORS -eq 0 ]; then
    echo -e "${YELLOW}⚠ Setup complete with $WARNINGS warning(s)${NC}"
    echo -e "  The warnings above are for optional tools that improve the development experience.\n"
    exit 0
else
    echo -e "${RED}✗ Setup incomplete: $ERRORS error(s), $WARNINGS warning(s)${NC}"
    echo -e "  Please fix the errors above before proceeding.\n"
    exit 1
fi
