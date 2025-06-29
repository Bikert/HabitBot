#!/usr/bin/env bash
cd "$(dirname "$0")"

echo "Building BE"
go build

cd ./webapp
echo "Installing node modules"
npm i
echo "Building FE"
npm run build
