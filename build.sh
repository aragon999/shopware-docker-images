#!/bin/bash

# Build script for Shopware Docker Images
# Usage: ./build.sh [PHP_VERSION]

PHP_VERSION=${1:-8.4}
REGISTRY=${REGISTRY:-shopware-docker}

echo "Building images with PHP ${PHP_VERSION}..."

# Build app image
echo "Building app image..."
docker build \
  --build-arg PHP_VERSION=${PHP_VERSION} \
  -t ${REGISTRY}/app:${PHP_VERSION} \
  -t ${REGISTRY}/app:latest \
  ./app

# Build cli image
echo "Building cli image..."
docker build \
  --build-arg PHP_VERSION=${PHP_VERSION} \
  -t ${REGISTRY}/cli:${PHP_VERSION} \
  -t ${REGISTRY}/cli:latest \
  ./cli

# Build worker image
echo "Building worker image..."
docker build \
  --build-arg PHP_VERSION=${PHP_VERSION} \
  -t ${REGISTRY}/worker:${PHP_VERSION} \
  -t ${REGISTRY}/worker:latest \
  ./worker

echo "Build complete!"
echo ""
echo "Images built:"
echo "  ${REGISTRY}/app:${PHP_VERSION}"
echo "  ${REGISTRY}/cli:${PHP_VERSION}"
echo "  ${REGISTRY}/worker:${PHP_VERSION}"
