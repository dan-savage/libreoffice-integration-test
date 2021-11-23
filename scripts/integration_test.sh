#!/usr/bin/env bash

echo "===> Start Integration Test"
docker-compose up -d 
go test -count=1  -tags=integration -v ./integration_tests/...
docker-compose down -v --rmi all --remove-oprhans
