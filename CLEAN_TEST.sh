#!/bin/bash
echo "=== CLEAN DYNAMIC CONFIG TEST ==="

# Ensure clean start
pkill -9 -f "go run" 2>/dev/null
rm -f winbu.db

# Start server
go run main.go &
SRV=$!
sleep 8

echo "[1] Test with winbu.net:"
curl -s http://localhost:8080/api/v1/anime-terbaru?page=1 | jq '{source, count: (.data|length)}'

echo ""
echo "[2] Change to fake.domain:"
curl -s -X PUT http://localhost:8080/api/admin/config/base_url \
  -H "Content-Type: application/json" \
  -d '{"value":"https://fake.domain"}' | jq '{error, message}'

echo ""  
echo "[3] Test with fake.domain:"
curl -s http://localhost:8080/api/v1/anime-terbaru?page=1 | jq '{source, error, message}' | head -5

kill $SRV
