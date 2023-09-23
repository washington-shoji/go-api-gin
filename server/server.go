package server

import (
	"log"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/washington-shoji/gin-api/databases"
	"github.com/washington-shoji/gin-api/handlers"
	"github.com/washington-shoji/gin-api/repositories"
	"github.com/washington-shoji/gin-api/routers"
	"github.com/washington-shoji/gin-api/services"
)

func InitServer() {

	db, err := databases.DatabaseConnection()
	if err != nil {
		log.Fatal(err)
	}

	if err := databases.CreateTables(db); err != nil {
		log.Fatal(err)
	}

	validate := validator.New()

	bookRepository := repositories.NewBookRepositoryImp(db)
	bookService := services.NewBookService(bookRepository, validate)
	bookHandler := handlers.NewBookHandler(bookService)
	bookRouter := routers.NewRouter(bookHandler)

	// testMeta := exp.NewMetaDatabaseImp(db)
	// testMeta.MetaDatabase()

	server := &http.Server{
		Addr:           ":3030",
		Handler:        bookRouter,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}
