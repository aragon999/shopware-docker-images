#!/bin/bash

# Build all PHP versions for Shopware Docker Images
# Usage: ./build-all.sh

REGISTRY=${REGISTRY:-shopware-docker}
PHP_VERSIONS=("8.1" "8.2" "8.3" "8.4")

echo "Building all PHP versions..."

for PHP_VERSION in "${PHP_VERSIONS[@]}"; do
    echo ""
    echo "========================================"
    echo "Building PHP ${PHP_VERSION} images..."
    echo "========================================"

    # Build app image
    echo "Building app image..."
    docker build \
      --build-arg PHP_VERSION=${PHP_VERSION} \
      -t ${REGISTRY}/app:${PHP_VERSION} \
      ./app

    # Build cli image
    echo "Building cli image..."
    docker build \
      --build-arg PHP_VERSION=${PHP_VERSION} \
      -t ${REGISTRY}/cli:${PHP_VERSION} \
      ./cli

    # Build worker image
    echo "Building worker image..."
    docker build \
      --build-arg PHP_VERSION=${PHP_VERSION} \
      -t ${REGISTRY}/worker:${PHP_VERSION} \
      ./worker
done

# Tag latest as 8.4
echo ""
echo "Tagging latest versions..."
docker tag ${REGISTRY}/app:8.4 ${REGISTRY}/app:latest
docker tag ${REGISTRY}/cli:8.4 ${REGISTRY}/cli:latest
docker tag ${REGISTRY}/worker:8.4 ${REGISTRY}/worker:latest

echo ""
echo "Build complete!"
echo ""
echo "Available images:"
for PHP_VERSION in "${PHP_VERSIONS[@]}"; do
    echo "  ${REGISTRY}/app:${PHP_VERSION}"
    echo "  ${REGISTRY}/cli:${PHP_VERSION}"
    echo "  ${REGISTRY}/worker:${PHP_VERSION}"
done
echo ""
echo "Latest tags:"
echo "  ${REGISTRY}/app:latest"
echo "  ${REGISTRY}/cli:latest"
echo "  ${REGISTRY}/worker:latest"
