#!/usr/bin/env bash
cd "$(dirname "$0")"
./build.sh
./stop.sh
./start.sh
