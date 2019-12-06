#! /bin/bash

# only exit with zero if all commands of the pipeline exit successfully
set -o pipefail
# error on unset variables
set -u

TRAVIS_OS_NAME="osx"
TRAVIS_CPU_ARCH="amd64"
TRAVIS_TAG="0.2.3"

TIME=`date +%FT%T%z`
NAME="${TRAVIS_OS_NAME}_${TRAVIS_CPU_ARCH}"

mkdir -p "dist/${NAME}"

go build -ldflags "-s -w -X github.com/ajdnik/decrypo/build.version=${TRAVIS_TAG} -X github.com/ajdnik/decrypo/build.datetime=${TIME}" -o "dist/${NAME}/decrypo"

cd "dist/${NAME}" && tar -cvzf "../${NAME}.tar.gz" *

cd ../ && shasum -a 256 "${NAME}.tar.gz" > "${NAME}.sha256"

rm -rf "${NAME}"
