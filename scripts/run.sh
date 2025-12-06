set -e

GREEN='\033[0;32m'
NC='\033[0m'

mkdir -p bin

echo -e "${GREEN}▶ Building service...${NC}"
go build -o bin/service ./cmd/service

echo -e "${GREEN}▶ Running with example files...${NC}"
./bin/service examples/permissions.txt examples/queries.txt
