#!/bin/bash
echo "build image -> start"
docker build --no-cache -t kevin24ec/walmart-api:1.0.0 --file build/Dockerfile .
echo "build image -> complete"
echo "------------------------------------------------------------------------------"
echo "push image -> start"
docker push kevin24ec/walmart-api:1.0.0
echo "push image -> complete"