#!/bin/bash

# Set up port forwarding for your Kafka server
kubectl port-forward --namespace=kafka svc/example-cluster-kafka-brokers 9092:9092
