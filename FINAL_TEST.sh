#!/bin/bash
echo "=========================================="
echo "FINAL DYNAMIC CONFIG TEST - CLEAN START"
echo "=========================================="

# Start fresh
rm -f winbu.db
go run main.go > /tmp/final_test.log 2>&1 &
SRV=$!
sleep 8

echo "[1] Initial state (should be winbu.net from schema):"
INIT=$(curl -s http://localhost:8080/api/admin/config/current | jq -r '.data.base_url')
echo "   Config: $INIT"

echo ""
echo "[2] Test scraping with winbu.net (should work):"
R1=$(curl -s http://localhost:8080/api/v1/anime-terbaru?page=1)
C1=$(echo "$R1" | jq -r '.data | length')
S1=$(echo "$R1" | jq -r '.source')
echo "   Source: $S1"
echo "   Items: $C1"

echo ""
echo "[3] Change to INVALID domain (fake-xyz-test.com):"
curl -s -X PUT http://localhost:8080/api/admin/config/base_url \
  -H "Content-Type: application/json" \
  -d '{"value":"https://fake-xyz-test.com"}' | jq -r '.message'

NEW=$(curl -s http://localhost:8080/api/admin/config/current | jq -r '.data.base_url')
echo "   New config: $NEW"

echo ""
echo "[4] Test scraping with FAKE domain (should fail):"
R2=$(curl -s http://localhost:8080/api/v1/anime-terbaru?page=1)
ERR=$(echo "$R2" | jq -r '.error')
SRC=$(echo "$R2" | jq -r '.source')
echo "   Source field: $SRC"
echo "   Has error: $ERR"

if [ "$SRC" = "fake-xyz-test.com" ]; then
    echo "   ✓✓✓ Source changed to fake domain!"
    if [ "$ERR" = "true" ]; then
        echo "   ✓✓✓ Scraping FAILED as expected!"
        echo "   ✓✓✓✓✓ DYNAMIC CONFIG IS WORKING 100%!"
    else
        echo "   ✗ Scraping still works (should fail)"
    fi
else
    echo "   ✗✗✗ Source still old domain: $SRC"
    echo "   ✗✗✗ DYNAMIC CONFIG NOT WORKING!"
fi

echo ""
echo "[5] Server logs (last 10 lines):"
tail -10 /tmp/final_test.log | grep -E "Visiting|Error visiting"

kill $SRV 2>/dev/null
echo ""
echo "=========================================="
