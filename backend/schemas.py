from datetime import date
from pydantic import BaseModel

class TransactionBase(BaseModel):
    description: str
    amount: float
    date: date
    category_id: int

class TransactionCreate(TransactionBase):
    pass

class Transaction(TransactionBase):
    id: int

    class Config:
        orm_mode = True

class Category(BaseModel):
    id: int
    name: str

    class Config:
        orm_mode = True

class CategorySummary(BaseModel):
    category: str
    total: float

class CategoryCreate(BaseModel):
    name: str

