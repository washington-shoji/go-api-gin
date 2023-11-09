package server

import (
	"log"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	cloudstorage "github.com/washington-shoji/gin-api/cloud_storage"
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

	cld, err := cloudstorage.CloudinaryConnection()
	if err != nil {
		log.Fatal(err)
	}

	validate := validator.New()

	bookRepository := repositories.NewBookRepositoryImp(db)
	userAccRepository := repositories.NewUserAccountRePository(db)
	tableTopGameRepository := repositories.NewTableTopGameRepositoryImpl(db)
	eventRepository := repositories.NewEventRepositoryImp(db)
	dynamicDataRepository := repositories.NewDynamicDataRepositoryImpl(db)

	bookService := services.NewBookService(bookRepository, validate, cld)
	userAccService := services.NewUserAccountService(userAccRepository, validate)
	tableTopGameService := services.NewTableTopGameService(tableTopGameRepository)
	eventService := services.NewEventService(eventRepository, cld)
	dynamicDataService := services.NewDynamicDataService(dynamicDataRepository)

	bookHandler := handlers.NewBookHandler(bookService)
	userAccHandler := handlers.NewUserAccountHandler(userAccService)
	tableTopGameHandler := handlers.NewTableTopGameHandler(tableTopGameService)
	eventHandler := handlers.NewEventHandler(eventService)
	dynamicDataHandler := handlers.NewDynamicDataHandler(dynamicDataService)

	loginService := services.NewLoginService(userAccRepository, validate)

	loginHandler := handlers.NewLoginHandler(loginService)

	dashboardHandler := handlers.NewDashboardHandler()

	expHandler := handlers.NewExpHandler(db)

	router := routers.NewRouter(
		bookHandler,
		userAccHandler,
		loginHandler,
		tableTopGameHandler,
		dynamicDataHandler,
		eventHandler,
		expHandler,
		dashboardHandler,
	)

	// testMeta := exp.NewMetaDatabaseImp(db)
	// testMeta.MetaDatabase()

	server := &http.Server{
		Addr:           ":3030",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}
