#!/bin/bash
set -x
rm -rf vector_data
rm new.json
mkdir vector_data
cat viaq.json | vector -c vector.toml
go run format.go
