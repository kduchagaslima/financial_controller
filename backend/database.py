import os
from sqlalchemy import create_engine
from sqlalchemy.orm import sessionmaker, declarative_base

# Database URL can be provided via the DATABASE_URL environment variable. It
# defaults to a local SQLite file for development purposes. When running in
# Docker the variable will point to the Postgres container.
SQLALCHEMY_DATABASE_URL = os.getenv("DATABASE_URL", "sqlite:///./finops.db")

connect_args = {
    "check_same_thread": False
} if SQLALCHEMY_DATABASE_URL.startswith("sqlite") else {}

engine = create_engine(SQLALCHEMY_DATABASE_URL, **connect_args)

SessionLocal = sessionmaker(autocommit=False, autoflush=False, bind=engine)

Base = declarative_base()

