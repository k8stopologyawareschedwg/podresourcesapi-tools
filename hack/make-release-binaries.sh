#!/bin/bash

set -eux

VERSION="${1}"

cp _out/client client-${VERSION}-linux-amd64
gzip client-${VERSION}-linux-amd64
cp _out/fakeserver fakeserver-${VERSION}-linux-amd64
gzip fakeserver-${VERSION}-linux-amd64

cp _out/client.exe client-${VERSION}-win-amd64
gzip client-${VERSION}-win-amd64
cp _out/fakeserver.exe fakeserver-${VERSION}-win-amd64
gzip fakeserver-${VERSION}-win-amd64
