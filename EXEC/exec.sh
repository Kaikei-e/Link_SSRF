#!/bin/bash

# Set common options
URL="http://localhost:8080/api/v1/register"
HEADERS="-H 'Connection: keep-alive' -H 'Keep-Alive: timeout=5, max=100'"

# pattern1
curl -vvv -XPOST $HEADERS -H "Content-Type: application/json" -d  '{"username":"aaaa", "profile_link":"http://localhost:8080/api/v1/users"}' $URL

# pattern2
# curl -vvv http://localhost:8080?http://localhost:9000/admin