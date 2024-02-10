#!/bin/bash

# Get the name of the current directory
current_directory=$(basename "$(pwd)")

# Check if the current directory is not "kafka-on-kubernetes"
if [ "$current_directory" != "kafka-on-kubernetes" ]; then
    echo "Please run this from the root directory of the project (kafka-on-kubernetes)."
    exit 1
fi

helm install example-kafka oci://registry-1.docker.io/bitnamicharts/kafka --version 26.8.4 -f kafka/values.yaml
kubectl apply -f producer/deployment.yaml
kubectl apply -f admin-dashboard.yaml
./scripts/token.sh