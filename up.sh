#!/bin/bash
kubectl apply -f ./mysql.test.yaml
kubectl apply -f ./postgresql.test.yaml
kubectl apply -f ./test-client.yaml

