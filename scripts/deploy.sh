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
kubectl apply -f kafka/topics/packet.yaml -n kafka