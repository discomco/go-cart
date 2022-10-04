#! /bin/bash

git config --global --add url."git@github.com:".insteadOf "https://github.com/"

GOPROXY="" GONOSUMDB=* go mod tidy

GOPROXY="" GONOSUMDB=* go mod vendor