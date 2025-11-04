# URL Shortener Frontend

This is the frontend of the **URL Shortener** application, built with **React**.  
It allows users to enter a long URL, which is then shortened through the backend API.  
The app is containerized with Docker, deployable via Helm charts.

---

## 🚀 Features
- Shorten long URLs using the backend service.
- Display and copy the shortened URL to clipboard.
- Simple and responsive user interface.
- Configurable backend API endpoint via environment variables.

---

## 📂 Project Structure
```
frontend/
│── src/
│   ├── components/
│   │   ├── UrlShortenerForm.js     # Form to submit URLs
│   │   ├── ShortUrlDisplay.js      # Displays shortened URL with copy feature
│   │   └── Header.js               # Application header
│   ├── services/
│   │   └── api.js                  # API integration
│── public/                         # Static assets
│── Dockerfile                      # Docker image definition
│── charts/                         # Helm chart for Kubernetes deployment
```

---

## ⚙️ Configuration

### Environment Variables
- `REACT_APP_API_BASE` → The base URL of the backend API.  
  Example:
  ```bash
  export REACT_APP_API_BASE=http://localhost:8080
  ```

---

## 🐳 Docker

### Build Image
```bash
docker build -t url-shortener-frontend .
```

### Run Container
```bash
docker run -p 3000:3000 -e REACT_APP_API_BASE=http://backend:8080 url-shortener-frontend
```

---

## ⎈ Helm Deployment

Helm chart is included for Kubernetes deployment.

---

## 🖼️ Example UI
1. Enter a long URL in the input field.
2. Click **Shorten**.
3. Get your **shortened URL** with a **copy button**.

---
