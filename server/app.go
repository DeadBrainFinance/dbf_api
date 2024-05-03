package server

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"dbf_api/database"
	"dbf_api/repositories"
	"dbf_api/routes"
	"dbf_api/services"
	"dbf_api/utils"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type App struct {
    httpServer *http.Server
}

func NewApp() *App {
    return &App{}
}

func InitDB() *database.Postgres {
    var envFile utils.EnvConfigs
    envFile.LoadEnvVariables()

    db, err := database.NewPostgres(envFile.DB, "localhost", envFile.DB_PORT, envFile.DB_USER, envFile.DB_PASSWORD)
	if err != nil {
		log.Fatal(err.Error())
	}
    if err := db.DB.Ping(); err != nil {
        log.Println("Postgres connected successfully...")
    }
    return db
}

func Version(db *sql.DB) chi.Router {
    r := chi.NewRouter()
    r.Mount("/v1", APIV1(db))

    return r
}

func APIV1(db *sql.DB) chi.Router {
    r := chi.NewRouter()

    transactionRepository := repositories.NewTransactionRepository(db)
    transactionService := services.NewService(transactionRepository)
    transactionRoute := routes.RegisterHTTPEndpoints(r, transactionService)

    r.Mount("/transactions", transactionRoute)

    return r
}

func (a *App) Run(port string) error {
    db := InitDB()
    router := chi.NewRouter()
    router.Use(middleware.Logger)

    router.Get("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("OK"))
    })


    router.Mount("/api", Version(db.DB))

    a.httpServer = &http.Server{
        Addr: ":" + port,
        Handler: router,
    }

    go func() {
        if err := a.httpServer.ListenAndServe(); err != nil {
            log.Fatalf("Failed to listen and server: %+v", err)
        }
    }()

    log.Println("API up and running")

    quit := make(chan os.Signal, 1)
    signal.Notify(quit, os.Interrupt, os.Interrupt)

    <-quit

    ctx, shutDown := context.WithTimeout(context.Background(), 5 * time.Second)
    defer shutDown()

    return a.httpServer.Shutdown(ctx)
}
