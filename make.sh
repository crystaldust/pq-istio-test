#!/bin/bash
go build -o test-client
docker build -t test-client ./

distribute-image.sh '43.7' test-client
rm test-client
