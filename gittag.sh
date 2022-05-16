#!/bin/sh

go mod tidy

git tag v$1 -m "Release $1"
git push origin v$1

echo "v$1" > .last_release

GOPROXY=proxy.golang.org go list -m github.com/HasinduLanka/diviyago@v$1
