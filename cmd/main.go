package main

import (
    "log"
    "net/http"
    
    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
    "github.com/lalka1231/todo-list1/database"
)

func main() {
    // Подключаемся к базе данных
    database.ConnectDB()
    
    r := chi.NewRouter()
    
    r.Use(middleware.Logger)
    r.Use(middleware.Recoverer)
    
    r.Get("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("TODO List API with PostgreSQL - Connected to DB!"))
    })
    
    r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("OK"))
    })
    
    log.Println("Server starting on :3000...")
    http.ListenAndServe(":3000", r)
}
