# devops-pet-project

A small end-to-end DevOps pet project built around a simple Go application.

The application runs an HTTP server, checks connectivity to PostgreSQL and exposes the result through a web page. I use the project to practice containerization, CI/CD and Kubernetes deployment.

## Stack

- Go
- PostgreSQL
- Docker
- GitHub Actions
- Docker Hub
- Kubernetes

## Application

The Go service runs on port `8080` and uses the `DATABASE_URL` environment variable to connect to PostgreSQL.

It reports whether the application is running and whether the database connection is available.

## CI

A GitHub Actions workflow runs on pushes to `main`.

The pipeline:

1. checks out the repository
2. configures Docker Buildx
3. authenticates to Docker Hub using GitHub Secrets
4. builds the application image
5. pushes `latest` and commit-specific image tags

```text
coodiallen/my-app:latest
coodiallen/my-app:<commit-sha>
```

## Structure

```text
.
├── .github/workflows/   # CI pipeline
├── app/                 # Go application and Dockerfile
├── infrastructure/      # infrastructure files
└── kubernetes/          # Kubernetes configuration
```

This project is mainly a hands-on environment where I practice connecting application development with containers, CI/CD and infrastructure.

---

# AWS Builder Challenge Website

A simple static website I built as part of the AWS Builder Challenge #2: **Build a website on the cloud**.

The goal was to get hands-on experience with hosting and delivering a static website using AWS.

## Stack

- HTML
- CSS
- Amazon S3
- Amazon CloudFront

## Project

The website contains a small personal profile and links to my AWS Builder Center profile.

```text
.
├── index.html
├── styles.css
└── profile.jpg
```

This was a small learning project focused on understanding the basic workflow of building a static website and publishing it through AWS.
