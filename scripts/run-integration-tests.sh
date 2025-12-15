#!/bin/bash

# Integration Test Suite Runner
# This script sets up the test environment and runs all integration tests

set -e  # Exit on error

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Configuration
TEST_DB_NAME="cadastro_pessoas_test"
TEST_DB_URI="mongodb://localhost:27017"
COVERAGE_FILE="coverage-integration.out"

echo -e "${YELLOW}Starting Integration Test Suite...${NC}\n"

# Check if MongoDB is running
echo "Checking MongoDB connection..."
if ! mongosh --eval "db.runCommand({ ping: 1 })" > /dev/null 2>&1; then
    echo -e "${RED}MongoDB is not running. Please start MongoDB first.${NC}"
    echo "Run: docker-compose up -d"
    exit 1
fi
echo -e "${GREEN}✓ MongoDB is running${NC}\n"

# Clean test database
echo "Cleaning test database..."
mongosh ${TEST_DB_URI}/${TEST_DB_NAME} --eval "db.dropDatabase()" > /dev/null 2>&1
echo -e "${GREEN}✓ Test database cleaned${NC}\n"

# Set test environment variables
export MONGODB_URI="${TEST_DB_URI}"
export MONGODB_DATABASE="${TEST_DB_NAME}"
export SERVER_PORT="8081"
export LOG_LEVEL="error"

echo "Running integration tests..."
echo "Database: ${TEST_DB_NAME}"
echo "Port: 8081"
echo ""

# Run tests with coverage
if go test -v -tags=integration -coverprofile=${COVERAGE_FILE} ./tests/integration/...; then
    echo -e "\n${GREEN}✓ All integration tests passed!${NC}\n"
    
    # Generate coverage report
    echo "Generating coverage report..."
    go tool cover -func=${COVERAGE_FILE}
    
    # Show summary
    COVERAGE=$(go tool cover -func=${COVERAGE_FILE} | grep total | awk '{print $3}')
    echo -e "\n${GREEN}Integration Test Coverage: ${COVERAGE}${NC}"
    
    # Clean up
    rm -f ${COVERAGE_FILE}
    
    exit 0
else
    echo -e "\n${RED}✗ Integration tests failed!${NC}\n"
    
    # Clean up
    rm -f ${COVERAGE_FILE}
    
    exit 1
fi
