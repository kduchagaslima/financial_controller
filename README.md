# FinOps Controller

This project is a minimal example of a financial control web application. It uses a FastAPI backend and a small frontend. Authentication integration should be implemented using Authentik or social OAuth providers (GitHub/Facebook) – currently a placeholder endpoint is provided.

The backend now runs on PostgreSQL and can be started together with the database (and an optional Angular frontend) using Docker Compose.

## Running with Docker

Ensure you have Docker and Docker Compose installed. To start the stack (backend, PostgreSQL and Angular frontend) simply run:

```bash
docker-compose up --build
```

The API will be available at `http://localhost:8000` and the Angular frontend at `http://localhost:4200`.

For development without Docker you can still run the backend manually:

```bash
cd backend
pip install -r requirements.txt
DATABASE_URL=postgresql://postgres:postgres@localhost:5432/finops uvicorn app:app --reload
```

The old static HTML frontend is kept for reference in `frontend-static`.

## Features

- Summary of expenses by category
- Add new financial entries when expanding categories
- On-demand reports filtered by category and period
- Create custom expense categories and assign them when recording expenses

This is a basic starting point and does not include full authentication or a production-ready setup.

