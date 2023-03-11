#!/bin/bash

DBSTRING="host=$DATABASE_HOST port=$DB_PORT user=$POSTGRES_USER password=$POSTGRES_PASSWORD dbname=$POSTGRES_DB sslmode=disable"
echo "$(date): run migrations $DBSTRING"
goose postgres "$DBSTRING" up
