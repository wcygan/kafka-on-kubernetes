#!/bin/bash

echo "Creating the 'kafka' namespace..."
kubectl get namespace kafka &> /dev/null || kubectl create namespace kafka
echo "Applying the Strimzi installation..."
kubectl create -f 'https://strimzi.io/install/latest?namespace=kafka' -n kafka