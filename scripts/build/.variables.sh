#!/usr/bin/env bash
set -eu

VERSION=${VERSION:-"unknown"}
GITCOMMIT=${GITCOMMIT:-$(git rev-parse --short HEAD 2> /dev/null || true)}
BUILDTIME=${BUILDTIME:-$(date --utc --rfc-3339 ns 2> /dev/null | sed -e 's/ /T/')}

export LDFLAGS="\
    -w \
    -X \"github.com/florinutz/go-donut-challenge/pkg.Commit=${GITCOMMIT}\" \
    -X \"github.com/florinutz/go-donut-challenge/pkg.BuildTime=${BUILDTIME}\" \
    -X \"github.com/florinutz/go-donut-challenge/pkg.Version=${VERSION}\" \
    ${LDFLAGS:-} \
"

GOOS="${GOOS:-$(go env GOHOSTOS)}"
GOARCH="${GOARCH:-$(go env GOHOSTARCH)}"

export SOURCE="main.go"
export TARGET="build/go-donut-challenge-$GOOS-$GOARCH"
