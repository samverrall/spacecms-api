#!/bin/bash

set -e

echo "Linting..."

golangci-lint --enable gosec,misspell run ./...

echo "Running tests..."

go test --cover --race --count=1 ./...

echo "Building..."

go build -o spacecms github.com/samverrall/spacecms-api/cmd/spacecms

echo "Built"
