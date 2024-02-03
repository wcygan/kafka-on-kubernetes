#!/bin/bash

# Set up port forwarding for your Kafka server
kubectl port-forward --namespace=kafka svc/example-cluster-kafka-bootstrap 9092:9092
