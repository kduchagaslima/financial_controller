package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Transaction struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Amount      float64   `json:"amount"`
	Date        time.Time `json:"date"`
	CategoryID  int       `json:"category_id"`
}

var categories = []Category{}
var transactions = []Transaction{}
var categorySeq, transSeq int

func main() {
	http.HandleFunc("/categories", categoriesHandler)
	http.HandleFunc("/categories/summary", summaryHandler)
	http.HandleFunc("/transactions", transactionsHandler)
	http.HandleFunc("/reports", reportsHandler)
	http.HandleFunc("/auth/login", loginHandler)

	http.ListenAndServe(":8000", nil)
}

func categoriesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		json.NewEncoder(w).Encode(categories)
	case http.MethodPost:
		var c Category
		if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		categorySeq++
		c.ID = categorySeq
		categories = append(categories, c)
		json.NewEncoder(w).Encode(c)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func summaryHandler(w http.ResponseWriter, r *http.Request) {
	type Summary struct {
		Category string  `json:"category"`
		Total    float64 `json:"total"`
	}
	summaryMap := map[int]float64{}
	for _, t := range transactions {
		summaryMap[t.CategoryID] += t.Amount
	}
	var result []Summary
	for _, c := range categories {
		total := summaryMap[c.ID]
		result = append(result, Summary{Category: c.Name, Total: total})
	}
	json.NewEncoder(w).Encode(result)
}

func transactionsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var t Transaction
		if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		transSeq++
		t.ID = transSeq
		transactions = append(transactions, t)
		json.NewEncoder(w).Encode(t)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func reportsHandler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	catParam := q.Get("category_id")
	startParam := q.Get("start")
	endParam := q.Get("end")

	var catID int
	var err error
	if catParam != "" {
		catID, err = strconv.Atoi(catParam)
		if err != nil {
			http.Error(w, "invalid category_id", http.StatusBadRequest)
			return
		}
	}
	var start, end time.Time
	if startParam != "" {
		start, err = time.Parse("2006-01-02", startParam)
		if err != nil {
			http.Error(w, "invalid start", http.StatusBadRequest)
			return
		}
	}
	if endParam != "" {
		end, err = time.Parse("2006-01-02", endParam)
		if err != nil {
			http.Error(w, "invalid end", http.StatusBadRequest)
			return
		}
	}
	var result []Transaction
	for _, t := range transactions {
		if catID != 0 && t.CategoryID != catID {
			continue
		}
		if !start.IsZero() && t.Date.Before(start) {
			continue
		}
		if !end.IsZero() && t.Date.After(end) {
			continue
		}
		result = append(result, t)
	}
	json.NewEncoder(w).Encode(result)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Auth integration not implemented", http.StatusNotImplemented)
}
