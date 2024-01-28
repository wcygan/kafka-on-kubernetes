#!/bin/bash

# Set up port forwarding for your Kafka server
kubectl port-forward --namespace=kafka example-cluster-kafka-0 9092:9092