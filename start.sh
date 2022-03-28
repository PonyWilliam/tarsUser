#!/bin/bash
set -ex
make
./Users --config=config.conf
