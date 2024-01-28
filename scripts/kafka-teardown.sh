#!/bin/bash

kubectl -n kafka delete $(kubectl get strimzi -o name -n kafka)
