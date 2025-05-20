#!/bin/bash

echo "Running database migrations..."
./migrate up

echo "Starting Go server..."
./server
