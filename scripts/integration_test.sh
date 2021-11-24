#!/usr/bin/env bash

echo "===> Start Integration Test"
docker-compose build #--no-cache
RESULT=$(docker run libreoffice-integration-test_libre-server:latest   go test -count=1  -tags=integration -v ./integration_tests/...)
# RESULT=$(docker run github_action_libre-server   go test -count=1  -tags=integration -v ./integration_tests/...)

echo "${result}"

docker-compose down -v --remove-orphans --rmi all 

SUB='FAIL'
if [[ "$RESULT" == *"$SUB"* ]]; then
    exit 1
fi



