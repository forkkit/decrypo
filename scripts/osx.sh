#! /bin/bash

# only exit with zero if all commands of the pipeline exit successfully
set -o pipefail
# error on unset variables
set -u

TIME=`date +%FT%T%z`
NAME="${TRAVIS_OS_NAME}_${TRAVIS_CPU_ARCH}"

mkdir -p "dist/${NAME}"

CGO_ENABLED=1 go build -ldflags "-s -w -X github.com/ajdnik/decrypo/build.version=${TRAVIS_TAG} -X github.com/ajdnik/decrypo/build.datetime=${TIME}" -o "dist/${NAME}/decrypo"

cd "dist/${NAME}" && tar -cvzf "../${NAME}.tar.gz" *

cd ../ && shasum -a 256 "${NAME}.tar.gz" > "${NAME}.sha256"

rm -rf "${NAME}"

go get -u github.com/git-chglog/git-chglog/cmd/git-chglog

cd ../ && git-chglog -c .chglog/release/config.yml -o RELEASE.md ${TRAVIS_TAG}
