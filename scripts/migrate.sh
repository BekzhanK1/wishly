#!/bin/bash

source .env

MIGRATION_DIR="file://migrations"
DB_URL="postgres://$DB_USERNAME:$DB_PASSWORD@$DB_HOST:$DB_PORT/$DB_DATABASE?sslmode=disable"

COMMAND=$1

case $COMMAND in
    up)
        migrate -path migrations -database "$DB_URL" up
        ;;
    down)
        migrate -path migrations -database "$DB_URL" down
        ;;
    force)
        migrate -path migrations -database "$DB_URL" force $2
        ;;
    create)
        if [ -z "$2" ]; then
        echo "Provide a name: ./migrate.sh create create_users_table"
        else
        migrate create -ext sql -dir migrations -seq "$2"
        fi
    ;;
    drop)
        read -p "Are you sure you want to drop the database? This action cannot be undone. (Y/N): " confirm
        if [[ "$confirm" =~ ^[Yy]$ ]]; then
            migrate -path migrations -database "$DB_URL" drop $2
        else
            echo "Drop operation canceled."
        fi
        ;;
    status)
        migrate -path migrations -database "$DB_URL" version
        ;;
    *)
        echo "Usage: ./migrate.sh [up|down|create|status|force|drop]"
        ;;
    esac
