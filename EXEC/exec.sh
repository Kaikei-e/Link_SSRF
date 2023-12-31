#!/bin/bash

# Set common options
URL="http://localhost:8080/api/v1/register"
HEADERS="-H 'Connection: keep-alive' -H 'Keep-Alive: timeout=5, max=100'"

# pattern1
curl -vvv -X POST $HEADERS -H "Content-Type: application/json" -d  '{"username":"aaaa", "profile_link":"http://localhost:9000/api/v1/users"}' $URL
