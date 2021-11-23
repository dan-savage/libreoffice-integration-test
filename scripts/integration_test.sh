#!/usr/bin/env bash

echo "===> Start Integration Test"
docker-compose build #--no-cache
docker run github_action_libre-server   go test -count=1  -tags=integration -v ./integration_tests/...
docker-compose down -v --rmi all --remove-oprhans
docker-compose down -v --rmi all --remove-orphans