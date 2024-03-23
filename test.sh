echo "Building images"
docker compose -f tests/docker-compose.yml up -d

echo "Running tests"
go test ./tests -v