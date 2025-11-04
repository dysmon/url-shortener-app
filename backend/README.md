# URL Shortener

This project is a simple backend service implemented in **Go**, containerized with **Docker**, and deployed to Kubernetes using **Helm**.  
It accepts HTTP requests to shorten URLs and redirect short codes to their original URLs.  

---

## Features

- Shorten a URL (`POST /shorten`)
- Redirect short code to original URL (`GET /{shortCode}`)
- Health check endpoint (`GET /healthcheck`)
- Logs requests and handles CORS
- Configurable via environment variables

---

## Technologies Used

- **Go** — Backend server
- **Docker** — Containerization
- **Helm** — Kubernetes deployment

---

## Prerequisites

- **Go** installed for local development (`go1.15` used in base image)
- **Docker** installed and configured
- Kubernetes cluster & Helm installed for deployment

---
## Endpoints

| Method | Path               | Description                           |
|--------|--------------------|---------------------------------------|
| `GET`  | `/healthcheck`     | Returns a simple health check status. |
| `POST` | `/shorten`         | Accepts a URL, returns short code.    |
| `GET`  | `/{shortCode}`     | Redirects to original URL.            |

---

## Configuration

The app reads configuration from environment variables:

- `PORT` — Port to run the server (default: `8080`)
- `ENV` — Runtime environment (`dev`, `prod`, etc.)

---
