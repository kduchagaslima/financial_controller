import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';

interface Category { id: number; name: string; }
interface CategorySummary { category: string; total: number; }
interface Transaction { description: string; amount: number; date: string; category_id: number; }

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit {
  categories: Category[] = [];
  summary: CategorySummary[] = [];
  newCategory = '';
  newTransaction: Partial<Transaction> = {};

  constructor(private http: HttpClient) {}

  ngOnInit(): void {
    this.loadCategories();
    this.loadSummary();
  }

  loadCategories() {
    this.http.get<Category[]>('/categories').subscribe(data => this.categories = data);
  }

  loadSummary() {
    this.http.get<CategorySummary[]>('/categories/summary').subscribe(data => this.summary = data);
  }

  addCategory() {
    if (!this.newCategory) return;
    this.http.post('/categories', { name: this.newCategory }).subscribe(() => {
      this.newCategory = '';
      this.loadCategories();
      this.loadSummary();
    });
  }

  addTransaction() {
    if (!this.newTransaction.description || !this.newTransaction.amount || !this.newTransaction.date || !this.newTransaction.category_id) return;
    this.http.post('/transactions', this.newTransaction).subscribe(() => {
      this.newTransaction = {};
      this.loadSummary();
    });
  }
}
