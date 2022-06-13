#!/bin/bash

golangci-lint run ./...

go build -o invoice-api github.com/samverrall/invoice-api-service/cmd/invoice