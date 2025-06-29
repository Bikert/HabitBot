#!/usr/bin/env bash
cd "$(dirname "$0")"
kill $(ps -aux | grep '[H]abitMuse' | awk '{print $2}')
