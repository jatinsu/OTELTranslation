#!/bin/bash
set -x
rm -rf vector_data
rm logs/new.json
mkdir vector_data
cat logs/viaq.json | vector -c vector.toml
go run go_code/format.go
