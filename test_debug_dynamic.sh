#!/bin/bash
echo "Starting server with logs..."
go run main.go > /tmp/server_debug.log 2>&1 &
SRV=$!
sleep 6

echo "[1] Current config:"
curl -s http://localhost:8080/api/admin/config/current | jq -r '.data.base_url'

echo ""
echo "[2] Update to fake-test-domain.xyz:"
curl -s -X PUT http://localhost:8080/api/admin/config/base_url \
  -H "Content-Type: application/json" \
  -d '{"value":"https://fake-test-domain.xyz"}' | jq -r '.message'

echo ""
echo "[3] Testing anime endpoint (watch the logs)..."
curl -s http://localhost:8080/api/v1/anime-terbaru?page=1 > /tmp/result.json

echo ""
echo "=== SERVER LOGS (last 20 lines) ==="
tail -20 /tmp/server_debug.log

echo ""
echo "=== API RESPONSE ==="
cat /tmp/result.json | jq '{error, message, source, count: (.data | length)}' 2>/dev/null || cat /tmp/result.json | head -c 200

kill $SRV 2>/dev/null
