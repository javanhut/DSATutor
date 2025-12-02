# Docker Deployment

DSATutor can be deployed as a Docker container for easy distribution and deployment.

## Quick Start

### Pull from Docker Hub

```bash
docker pull yourusername/dsatutor:latest
docker run -p 8080:8080 yourusername/dsatutor:latest
```

Then open http://localhost:8080 in your browser.

### Build Locally

```bash
docker build -t dsatutor .
docker run -p 8080:8080 dsatutor
```

## Configuration

### Port Configuration

The default port is 8080. To use a different host port:

```bash
docker run -p 3000:8080 dsatutor
```

This maps host port 3000 to container port 8080.

### Custom Container Port

To change the port inside the container, override the CMD:

```bash
docker run -p 9000:9000 dsatutor -serve -addr=:9000
```

## Development with Docker Compose

For local development and testing, use the included docker-compose.yml:

```bash
docker compose up --build
```

This builds the image and starts the container with port 8080 exposed.

To run in the background:

```bash
docker compose up -d --build
```

To stop:

```bash
docker compose down
```

## Image Details

- **Base Image**: Python 3.12 slim (Debian Bookworm)
- **Build Stage**: Go 1.22 (multi-stage build)
- **Final Size**: ~150MB
- **Exposed Port**: 8080

The image includes Python 3.12 for the code execution sandbox that powers the practice problems.

## Automated Publishing with GitHub Actions

The repository includes a GitHub Actions workflow that automatically builds and pushes Docker images when you create a version tag.

### Setup

1. Create a Docker Hub access token at https://hub.docker.com/settings/security

2. Add the following secrets to your GitHub repository (Settings > Secrets and variables > Actions):
   - `DOCKERHUB_USERNAME`: Your Docker Hub username
   - `DOCKERHUB_TOKEN`: Your Docker Hub access token

### Creating a Release

Push a version tag to trigger the automated build:

```bash
git tag v1.0.0
git push origin v1.0.0
```

The workflow will:
- Build the image for both `linux/amd64` and `linux/arm64` platforms
- Push with multiple tags: `1.0.0`, `1.0`, `1`, and `latest`

### Tag Format

Use semantic versioning tags:
- `v1.0.0` - Creates tags: `1.0.0`, `1.0`, `1`, `latest`
- `v1.1.0` - Creates tags: `1.1.0`, `1.1`, `1`, `latest`
- `v2.0.0` - Creates tags: `2.0.0`, `2.0`, `2`, `latest`

## Manual Publishing to Docker Hub

1. Build the image with your Docker Hub username:

```bash
docker build -t yourusername/dsatutor:latest .
```

2. Log in to Docker Hub:

```bash
docker login
```

3. Push the image:

```bash
docker push yourusername/dsatutor:latest
```

## Health Check

The application serves a web interface at the root path. You can verify it's running:

```bash
curl http://localhost:8080/
```

A successful response returns the HTML page for the DSATutor interface.
