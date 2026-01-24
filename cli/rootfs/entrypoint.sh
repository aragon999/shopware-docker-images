#!/usr/bin/env bash

# Source fnm environment
eval "$(fnm --fnm-dir=/nvm/.debian-fnm env --shell=bash)"

if [ -z "$1" ]; then
    exec bash
else
    exec "$@"
fi
