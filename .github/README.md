# GitHub Workflows

This directory contains GitHub Actions workflows for the Shopware Docker Images project.

## Workflows

### build.yml
Main workflow that builds and pushes Docker images to the GitHub Container Registry (GHCR).

**Triggers:**
- Push to main branch
- Manual dispatch

**Features:**
- Builds for PHP versions 8.1, 8.2, 8.3, and 8.4
- Creates PHP version-specific tags:
  - `8.1`, `8.2`, `8.3`, `8.4`
- Uses GitHub cache for faster builds
- Generates SBOM and security scans with Trivy
- Attestations for provenance

## Image Tags

Images are pushed to: `ghcr.io/aragon999/shopware-docker-images`

### Available Tags

For each image type (app, cli, worker):
- `8.1`, `8.2`, `8.3`, `8.4` - PHP version specific

## Usage

### Using the Images

```yaml
services:
  app:
    image: ghcr.io/aragon999/shopware-docker-images/app:8.4
  cli:
    image: ghcr.io/aragon999/shopware-docker-images/cli:8.4
  worker:
    image: ghcr.io/aragon999/shopware-docker-images/worker:8.4
```

### Pulling Images

```bash
docker pull ghcr.io/aragon999/shopware-docker-images/app:8.4
docker pull ghcr.io/aragon999/shopware-docker-images/cli:8.4
docker pull ghcr.io/aragon999/shopware-docker-images/worker:8.4
```

## Security

- Images are scanned with Trivy for vulnerabilities
- Results are uploaded to GitHub Security tab
- SBOM (Software Bill of Materials) is generated for each image
