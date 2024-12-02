#!/usr/bin/env bash

docker compose down
docker volume rm golang-gorm-postgres_data golang-gorm-postgres_postgres -f

docker compose up --no-deps --build --force-recreate -d
