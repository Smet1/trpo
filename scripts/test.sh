#!/usr/bin/env bash

check_cover()
{
    cd ..
    sleep 15
    go test -coverpkg=./internal/... -coverprofile=cover.out.tmp ./internal/...
    cat cover.out.tmp | grep -v -E "_easyjson.go|.pb.go|chat|cmd|HelloWorld" > cover.out
    go tool cover -func cover.out
}

check_cover