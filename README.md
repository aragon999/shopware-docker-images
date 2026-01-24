# Shopware 6 Docker Images

A set of lightweight, flexible Docker images for Shopware 6 development environments. These images are designed for development use only, not for production.

## Overview

This project provides three main container types:

- **app**: Nginx + PHP-FPM server for serving Shopware 6 applications
- **cli**: PHP CLI container with Shopware CLI tools and development utilities
- **worker**: Dedicated container for running Shopware messenger consumers and scheduled tasks

## Features

- PHP version configurable (see Available PHP Versions)
- Separated worker processes for better control during setup
- Includes all necessary PHP extensions for Shopware 6
- Optimized for development with proper volume mounts
- Built on Alpine Linux for smaller image sizes (app/worker)
- Includes Node.js via fnm for frontend build tools

## Quick Start

### Using Pre-built Images from GitHub Container Registry

```bash
# Pull images for PHP 8.4
docker pull ghcr.io/aragon999/shopware-docker-images/app:8.4
docker pull ghcr.io/aragon999/shopware-docker-images/cli:8.4
docker pull ghcr.io/aragon999/shopware-docker-images/worker:8.4

# Pull images for PHP 8.3
docker pull ghcr.io/aragon999/shopware-docker-images/app:8.3
docker pull ghcr.io/aragon999/shopware-docker-images/cli:8.3
docker pull ghcr.io/aragon999/shopware-docker-images/worker:8.3
```

### Building Images Locally

```bash
# Build with default PHP 8.4
docker build -t shopware-app ./app
docker build -t shopware-cli ./cli
docker build -t shopware-worker ./worker

# Build with specific PHP version
docker build --build-arg PHP_VERSION=8.3 -t shopware-app:php8.3 ./app
docker build --build-arg PHP_VERSION=8.3 -t shopware-cli:php8.3 ./cli
docker build --build-arg PHP_VERSION=8.3 -t shopware-worker:php8.3 ./worker
```

### Docker Compose Example

#### Option 1: Using Pre-built Images

```yaml
services:
  app:
    image: ghcr.io/aragon999/shopware-docker-images/app:8.4
    # ... (rest of configuration)

  worker:
    image: ghcr.io/aragon999/shopware-docker-images/worker:8.4
    # ... (rest of configuration)

  cli:
    image: ghcr.io/aragon999/shopware-docker-images/cli:8.4
    # ... (rest of configuration)
```

#### Option 2: Building Locally

```yaml
services:
  app:
    build: ./shopware-docker-images/app
    build-args:
      PHP_VERSION: 8.4  # Adjust as needed
    labels:
      - traefik.enable=true
      - traefik.http.routers.myproject_insecure.entrypoints=web
      - traefik.http.routers.myproject_insecure.rule=Host(`myproject.dev.localhost`)
      - traefik.http.routers.myproject_insecure.middlewares=redirect@file
      - traefik.http.routers.myproject.entrypoints=web-secure
      - traefik.http.routers.myproject.rule=Host(`myproject.dev.localhost`)
      - traefik.http.routers.myproject.tls.certresolver=letsencrypt
      - traefik.docker.network=traefik-proxy
    env_file:
      - ./docker.env
    extra_hosts:
      .dev.localhost: 127.0.0.1
    environment:
      VIRTUAL_HOST: .dev.localhost
      CERT_NAME: shared
      HTTPS_METHOD: noredirect
      PHP_MEMORY_LIMIT: 4G
      FPM_PM_MAX_CHILDREN: 128
    networks:
      - default
      - traefik-proxy
      - myproject
      - mailcatcher
    links:
      - mysql
      - elastic
      - redis
    volumes:
      - ./code/:/var/www/html
      - ./htpasswd:/etc/nginx/.htpasswd

  worker:
    build: ./shopware-docker-images/worker
    build-args:
      PHP_VERSION: 8.4  # Should match app container
    env_file:
      - ./docker.env
    networks:
      - myproject
    links:
      - mysql
      - elastic
      - redis
    volumes:
      - ./code/:/var/www/html
    profiles:
      - workers  # Use profile to control when workers start

  cli:
    build: ./shopware-docker-images/cli
    build-args:
      PHP_VERSION: 8.4  # Adjust as needed
    env_file:
      - ./docker.env
    tty: true
    volumes:
      - ./code/:/var/www/html
    networks:
      - myproject
      - mailcatcher
    links:
      - mysql

  mysql:
    image: mariadb:10.11
    env_file: ./docker.env
    volumes:
      - ./mysql-data:/var/lib/mysql:delegated
      - ./sql-mode.cnf:/etc/mysql/mariadb.conf.d/sql-mode.cnf
    networks:
      - myproject

  elastic:
    image: blacktop/elasticsearch:7
    environment:
      VIRTUAL_HOST: elastic
      VIRTUAL_PORT: 9200
      discovery.type: single-node
    networks:
      - myproject

  redis:
    image: redis:7-alpine
    command: ["redis-server", "--maxmemory", "2gb", "--maxmemory-policy", "allkeys-lru"]
    volumes:
      - ./redis:/data
    networks:
      - myproject

networks:
  myproject:
  mailcatcher:
    external: true
  traefik-proxy:
    external: true
```

## Usage Workflow

1. **Initial Setup (without workers):**
   ```bash
   docker-compose up app cli mysql elastic redis
   ```

2. **Install Shopware:**
   ```bash
   docker-compose exec cli bash
   composer create-project shopware/core .
   ```

3. **Start Workers:**
   ```bash
   docker-compose --profile workers up -d worker
   ```

## Available PHP Versions

- 8.1
- 8.2
- 8.3
- 8.4

## Environment Variables

### App Container
- `PHP_MEMORY_LIMIT`: PHP memory limit (default: 512m)
- `FPM_PM_MAX_CHILDREN`: Maximum PHP-FPM children (default: 5)
- `APP_DOCUMENT_ROOT`: Document root (default: /var/www/html/public)

### All Containers
- `TZ`: Timezone (default: Europe/Berlin)

## Development Tools

The CLI container includes:
- Composer
- Shopware CLI
- Node.js via fnm
- yq for YAML processing
- MariaDB client
- Git, wget, zip, unzip

## Image Registry and Tags

Images are published to: `ghcr.io/aragon999/shopware-docker-images`

For each image type (`app`, `cli`, `worker`), tags correspond to the PHP versions listed in the "Available PHP Versions" section.

## Security

- Images are scanned with Trivy for vulnerabilities
- Results are uploaded to the GitHub Security tab
- An SBOM (Software Bill of Materials) is generated for each image

## Continuation

This is a continuation of the images from shyim: https://github.com/shyim/shopware-docker-images
