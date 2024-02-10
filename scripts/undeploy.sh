#!/bin/bash

# Get the name of the current directory
current_directory=$(basename "$(pwd)")

# Check if the current directory is not "kafka-on-kubernetes"
if [ "$current_directory" != "kafka-on-kubernetes" ]; then
    echo "Please run this from the root directory of the project (kafka-on-kubernetes)."
    exit 1
fi

echo "Undeploying the application..."
kubectl delete -f producer/deployment.yaml
helm delete example-kafka