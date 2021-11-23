#!/usr/bin/env bash
#
# This script runs tests for the attention measurement platform services.
set -e

# Test!
echo "==> Testing..."
go test \
    -ldflags "${LD_FLAGS}" \
    ./...
