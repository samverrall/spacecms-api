#!/bin/bash

export INVOICE_SECRET="1234"
export INVOICE_SECRET_REFRESH="1234"
export DATABASE_DIR="/Users/samverrall/projects/invoice-app/.data/database.db"
export MIGRATIONS_DIR="/Users/samverrall/projects/invoice-app/database/migration"

./invoice-api
