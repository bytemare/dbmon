#!/usr/bin/env bash

docker run --security-opt="seccomp=.\dbmon.seccomp" bytemare/dbmon:dbmon.server.v0