#!/usr/bin/env bash

echo "===> Start Integration Test"
docker-compose build #--no-cache
docker run libreoffice-integration-test_libre-server:latest   go test -count=1  -tags=integration -v ./integration_tests/...
docker-compose down -v --rmi all --remove-oprhans
docker-compose down -v --rmi all --remove-orphans