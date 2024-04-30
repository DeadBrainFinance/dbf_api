package main

import (
	"log"
	"net/http"

	// "database/sql"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"dbf_api/database"
	"dbf_api/models"
	"dbf_api/services"

	_ "github.com/lib/pq"
)

func main() {
    r := chi.NewRouter()
    r.Use(middleware.Logger)

    r.Get("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("OK"))
    })

    conn, err := database.NewPostgres("dbf_db", "localhost:5499", "postgres", "postgres")
	if err != nil {
		log.Fatal(err.Error())
	}
    db := models.New(conn.DB)
    if db != nil {
        log.Println("PostgreSql connected successfully...")
    }
    transServices := services.NewService(db)

    r.Mount("/transactions", transServices.RegisterHandlers())

    http.ListenAndServe(":4000", r)
}

// func TransactionRoutes() chi.Router {
//     r := chi.NewRouter()
//     handler := services.TransactionHandler{}
//
//     r.Get("/", handler.ListAll)
//     r.Post("/", handler.CreateRecord)
//     r.Get("/{id}", handler.GetById)
//     r.Put("/{id}", handler.UpdateRecord)
//     r.Delete("/{id}", handler.DeleteRecord)
//
//     return r
// }
