package api

import (
	"context"
	"database/sql"
    stdlog "log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"dbf_api/database"
	"dbf_api/repositories"
	"dbf_api/services"
	"dbf_api/utils"

	"github.com/go-chi/chi/v5"
	// "github.com/go-chi/chi/v5/middleware"
	log "github.com/go-kit/log"
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

    db, err := database.NewPostgres(envFile.DB, "host.docker.internal", envFile.DB_PORT, envFile.DB_USER, envFile.DB_PASSWORD)
	if err != nil {
		stdlog.Fatal(err.Error())
	}
    if err := db.DB.Ping(); err != nil {
        stdlog.Println("Postgres connected successfully...")
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
    transactionService := services.NewTransactionService(transactionRepository)

    categoryRepository := repositories.NewCategoryRepository(db)
    categoryService := services.NewCategoryService(categoryRepository)

    balancesheetRepository := repositories.NewBalanceSheetRepository(db)
    balancesheetService := services.NewBalanceSheetService(balancesheetRepository)

    methodRepository := repositories.NewMethodRepository(db)
    methodService := services.NewMethodService(methodRepository)

    accountRepository := repositories.NewAccountRepository(db)
    accountService := services.NewAccountService(accountRepository)

    installmentRepository := repositories.NewInstallmentRepository(db)
    installmentService := services.NewInstallmentService(installmentRepository)

    debtRepository := repositories.NewDebtRepository(db)
    debtService := services.NewDebtService(debtRepository)

    transactionService.RegisterHTTPEndpoints(r)
    categoryService.RegisterHTTPEndpoints(r)
    balancesheetService.RegisterHTTPEndpoints(r)
    methodService.RegisterHTTPEndpoints(r)
    accountService.RegisterHTTPEndpoints(r)
    installmentService.RegisterHTTPEndpoints(r)
    debtService.RegisterHTTPEndpoints(r)

    return r
}

func (a *App) Run(port string) error {
    var logger log.Logger
    logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
    stdlog.SetOutput(log.NewStdlibAdapter(logger))
    logger = log.With(logger, "ts", log.DefaultTimestamp, "loc", log.DefaultCaller)
    loggingMiddleware := LoggingMiddleware(logger)

    db := InitDB()

    router := chi.NewRouter()
    router.Use(loggingMiddleware)

    router.Get("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("OK"))
    })



    router.Mount("/api", Version(db.DB))

    chi.Walk(router, func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		stdlog.Printf("[%s]: '%s' has %d middlewares\n", method, route, len(middlewares))
		return nil
	})

    a.httpServer = &http.Server{
        Addr: ":" + port,
        Handler: router,
    }

    go func() {
        if err := a.httpServer.ListenAndServe(); err != nil {
            stdlog.Fatalf("Failed to listen and server: %+v", err)
        }
    }()

    stdlog.Println("API up and running")

    quit := make(chan os.Signal, 1)
    signal.Notify(quit, os.Interrupt, os.Interrupt)

    <-quit

    ctx, shutDown := context.WithTimeout(context.Background(), 5 * time.Second)
    defer shutDown()

    return a.httpServer.Shutdown(ctx)
}
