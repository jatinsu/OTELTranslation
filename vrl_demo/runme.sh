#!/bin/bash
# Local Vector/VRL demo
set -x
rm -rf vector_data
mkdir vector_data
vector -c vector.toml
