# FinOps Controller

This project is a minimal example of a financial control web application. It uses a FastAPI backend with SQLite and a small HTML/JS frontend. Authentication integration should be implemented using Authentik or social OAuth providers (GitHub/Facebook) – currently a placeholder endpoint is provided.

## Running Backend

```bash
cd backend
pip install -r requirements.txt
uvicorn app:app --reload
```

## Frontend

Open `frontend/index.html` in your browser. The frontend expects the backend to be running at the same origin (e.g., localhost:8000).

## Features

- Summary of expenses by category
- Add new financial entries when expanding categories
- On-demand reports filtered by category and period

This is a basic starting point and does not include full authentication or a production-ready setup.

