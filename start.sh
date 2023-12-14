#!/bin/sh

# if the file does not exist, make .env file
if [ ! -f .env ]; then
    touch .env
fi

# start docker compose
docker compose up -d