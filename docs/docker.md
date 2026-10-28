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
docker build --pull -t dsatutor .
docker run -p 8080:8080 dsatutor
```

`--pull` fetches the newest Oxigen release (see [Image Details](#image-details));
with podman use `podman build --pull=always`.

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

This builds the image and starts the container with port 8080 exposed. The
compose file sets `pull: true`, so Docker Compose re-pulls the base images on
every build. podman-compose ignores that setting; build with
`podman-compose build --pull always` instead.

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
- **Oxigen**: copied from [`javanhut/oxigen`](https://hub.docker.com/r/javanhut/oxigen)
- **Exposed Port**: 8080

The image includes Python 3.12 and Oxigen for the code execution sandbox that
powers the practice problems. The `oxigen` binary goes to `/usr/local/bin` and
its standard library to `/usr/local/lib/oxigen/stdlib`, where the binary looks
for it.

The Oxigen release comes from the `OXIGEN_VERSION` build argument, which
defaults to `latest`. A build only gets the newest release if it re-pulls that
tag (`docker build --pull`, `podman build --pull=always`); otherwise it reuses
whichever `latest` is cached. The publishing workflow always pulls. To pin a
release instead:

```bash
docker build --build-arg OXIGEN_VERSION=0.2.2 -t dsatutor .
```

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
