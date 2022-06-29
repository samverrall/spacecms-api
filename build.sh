#!/bin/bash

golangci-lint run ./...

go build -o invoice-api github.com/samverrall/spacecms-api/cmd/invoice
