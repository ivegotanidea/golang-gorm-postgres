#!/usr/bin/env bash

docker compose down
docker volume rm golang-gorm-postgres_postgres -f

docker compose up --build --force-recreate -d
