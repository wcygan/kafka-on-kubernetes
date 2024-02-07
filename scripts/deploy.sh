#!/bin/bash

# Get the name of the current directory
current_directory=$(basename "$(pwd)")

# Check if the current directory is not "kafka-on-kubernetes"
if [ "$current_directory" != "kafka-on-kubernetes" ]; then
    echo "Please run this from the root directory of the project (kafka-on-kubernetes)."
    exit 1
fi

echo "Deploying the application..."
kubectl apply -f kafka/deployment.yaml -n kafka

echo "Waiting for Kafka to be ready..."
kubectl wait kafka/example-cluster --for=condition=Ready --timeout=300s -n kafka
echo "Kafka is ready!"

echo "Creating kafka topics..."
kubectl apply -f kafka/topics/*.yaml -n kafka

echo "Deploying the producer..."
kubectl apply -f producer/deployment.yaml