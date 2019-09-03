#!/usr/bin/env bash

docker run --security-opt="no-new-privileges" --security-opt="seccomp=.\dbmon.seccomp" bytemare/dbmon:dbmon.server.v0