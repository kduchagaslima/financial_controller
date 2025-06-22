from fastapi import FastAPI, Depends, HTTPException
from sqlalchemy.orm import Session
from datetime import date
from . import models, schemas, database

models.Base.metadata.create_all(bind=database.engine)

app = FastAPI(title="FinOps Controller")

# Dependency

def get_db():
    db = database.SessionLocal()
    try:
        yield db
    finally:
        db.close()

@app.get("/categories/summary", response_model=list[schemas.CategorySummary])
def read_summary(db: Session = Depends(get_db)):
    results = db.query(models.Category.name, models.Transaction)
    summaries = {}
    for cat_name, transaction in db.query(models.Category.name, models.Transaction.amount).join(models.Transaction).all():
        summaries.setdefault(cat_name, 0)
        summaries[cat_name] += transaction
    return [schemas.CategorySummary(category=k, total=v) for k, v in summaries.items()]


@app.get("/categories", response_model=list[schemas.Category])
def read_categories(db: Session = Depends(get_db)):
    return db.query(models.Category).all()


@app.post("/categories", response_model=schemas.Category)
def create_category(category: schemas.CategoryCreate, db: Session = Depends(get_db)):
    db_cat = models.Category(name=category.name)
    db.add(db_cat)
    db.commit()
    db.refresh(db_cat)
    return db_cat

@app.post("/transactions", response_model=schemas.Transaction)
def create_transaction(transaction: schemas.TransactionCreate, db: Session = Depends(get_db)):
    db_trans = models.Transaction(**transaction.dict())
    db.add(db_trans)
    db.commit()
    db.refresh(db_trans)
    return db_trans

@app.get("/reports", response_model=list[schemas.Transaction])
def read_reports(category_id: int | None = None, start: date | None = None, end: date | None = None, db: Session = Depends(get_db)):
    query = db.query(models.Transaction)
    if category_id:
        query = query.filter(models.Transaction.category_id == category_id)
    if start:
        query = query.filter(models.Transaction.date >= start)
    if end:
        query = query.filter(models.Transaction.date <= end)
    return query.all()

# Placeholder authentication endpoint
@app.post("/auth/login")
def login():
    """This endpoint should integrate with authentik or social auth"""
    raise HTTPException(status_code=501, detail="Auth integration not implemented")

