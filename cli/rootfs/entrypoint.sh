#!/usr/bin/env sh

if [ -z "$1" ]; then
    exec bash
else
    exec "$@"
fi
