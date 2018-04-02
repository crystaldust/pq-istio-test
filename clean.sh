#!/bin/bash
kubectl delete -f ./mysql.test.yaml
kubectl delete -f ./postgresql.test.yaml
kubectl delete -f ./test-client.yaml

