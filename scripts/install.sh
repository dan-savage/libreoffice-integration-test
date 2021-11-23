echo "==> Installing..."
go install \
    -v \
    -ldflags "${LD_FLAGS}" \
    ./...