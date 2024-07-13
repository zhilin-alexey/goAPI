package main

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5"
	pgxPool "github.com/jackc/pgx/v5/pgxpool"
	"github.com/lmittmann/tint"
	sloggin "github.com/samber/slog-gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	pgxUUID "github.com/vgarvardt/pgx-google-uuid/v5"
	"goAPI/controllers"
	databaseInterface "goAPI/database/sqlc"
	_ "goAPI/docs"
	"goAPI/routes"
	"goAPI/seeders"
	"log"
	"log/slog"
	"os"
	"strings"
	"time"
)

var (
	config   Config
	router   *gin.Engine
	database *databaseInterface.Queries

	PeopleController controllers.PeopleController
	TasksController  controllers.TasksController

	InfoRoute   routes.InfoRoute
	PeopleRoute routes.PeopleRoute
	TasksRoute  routes.TasksRoute
)

func init() {
	log.Println("Loading config from environment or .env file...")
	var err error

	config, err = LoadConfig()
	if err != nil {
		log.Fatalf("Cannot load config: %v", err)
	}

	log.Println("Initializing logger...")
	slogLevel, err := ParseSlogLevel(config.LogLevel)
	if err != nil {
		slogLevel = slog.LevelInfo
	}

	logger := slog.New(
		tint.NewHandler(os.Stdout, &tint.Options{Level: slogLevel, TimeFormat: time.DateTime}),
	)
	slog.SetDefault(logger)

	ctx := context.Background()

	slog.Debug("Connecting to database...")
	pgxConfig, err := pgxPool.ParseConfig(config.PostgresURL)
	if err != nil {
		log.Fatalf("Cannot parse database URL: %v", err)
	}

	pgxConfig.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		pgxUUID.Register(conn.TypeMap())
		return nil
	}

	conn, err := pgxPool.NewWithConfig(ctx, pgxConfig)
	if err != nil {
		log.Fatalf("Cannot connect to database: %v", err)
	}
	database = databaseInterface.New(conn)

	slog.Debug("Running migrations...")
	migrations, err := migrate.New("file://database/migrations", config.PostgresURL)
	if err != nil {
		log.Fatalf("Cannot create database migrations client: %v", err)
	}
	err = migrations.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatalf("Cannot run database migrations: %v", err)
	} else if errors.Is(err, migrate.ErrNoChange) {
		slog.Debug("No migrations to run")
	} else {
		slog.Info("Migrations ran successfully")
	}

	slog.Debug("Checking if seeds executed...")
	databaseSeeder := seeders.NewDatabaseSeeder(database)
	err = databaseSeeder.Start()
	if err != nil {
		slog.Error("Cannot run seeds: ", err)
	}

	slog.Debug("Creating routes and controllers...")
	PeopleController = *controllers.NewPeopleController(database)
	TasksController = *controllers.NewTasksController(database)

	InfoRoute = routes.NewInfoRoute(PeopleController)
	PeopleRoute = routes.NewPeopleRoute(PeopleController)
	TasksRoute = routes.NewTasksRoute(TasksController)

	slog.Debug("Configuring router...")

	router = gin.Default()
	router.Use(sloggin.New(logger))

	slog.Debug("Registering custom validation function for passport number...")
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err := v.RegisterValidation("passportNumber", PassportNumber)
		if err != nil {
			slog.Error("Failed to register custom validation function for passport number ")
		}
	}

	router.ForwardedByClientIP = true
	addresses := strings.Split(config.TrustedProxies, " ")
	err = router.SetTrustedProxies(addresses)

	if err != nil {
		log.Printf("Cannot set router trusted proxies: %v\n", err)
	}
}

// @title			goAPI
// @version		1.0
// @description	go API with task tracking
// @host			localhost:8080
// @BasePath		/
// @schemes		http
func main() {
	InfoRoute.Register(router)
	PeopleRoute.Register(router)
	TasksRoute.Register(router)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	slog.Info("Starting server...")
	err := router.Run(":" + config.Port)
	if err != nil {
		log.Fatalf("Cannot run server from router: %v", err)
	}
}
