#!/bin/bash
# Local Vector/VRL demo
set -x
rm -rf vector_data
rm new.json
mkdir vector_data
vector -c vector.toml
