#!/usr/bin/env bash
go test -v -coverprofile=build/coverage.out ./...
go tool cover -html=build/coverage.out