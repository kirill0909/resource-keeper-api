#!/bin/bash

export $(xargs < .env)
migrate -path ./schema -database "postgres://postgres:${DB_POSTGRES_PASSWORD}@localhost:5432/postgres?sslmode=disable" up
