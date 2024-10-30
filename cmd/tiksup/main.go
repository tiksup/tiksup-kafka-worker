/*
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⡀⡀⢀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣿⣿⣿⣿⣿⣿
⠀⠀⢘⢸⠀⡇⠁⠰⠀⡀⠆⢀⢰⠁⡇⡆⢸⢀⣁⣗⣺⢘⡀⡇⣱⠸⠀⠀⠆⠸⠐⠀⡇⣎⢸⠐⡃⣗⢸⣈⡂⡇⢰⢸⠀⡆⠂⢀⠰⠀⡄⠘⣿⣿⣿⣿⣿⣿
⠀⠀⠜⢸⠀⡇⠃⢠⠀⡄⡄⢸⢠⠀⡇⡄⢸⠤⠇⡧⢼⠸⠄⡇⢣⢰⠀⡀⡄⢸⢠⠀⡇⡼⢸⠤⠇⡧⢼⠸⠤⡇⢢⢰⠀⡅⡄⠈⢠⠀⡃⠸⣿⣿⣿⣿⣿⣿
⠀⠀⢱⢸⠀⡇⠀⢸⠀⣩⣇⢸⢨⡀⡏⡁⢸⠉⡆⣏⣹⢸⠁⡇⣎⢸⠀⠁⡁⠘⢈⠀⠃⣱⢸⣉⣇⣏⣽⣰⠁⡇⢈⢸⠀⠁⠀⠈⠈⠀⠃⢰⣿⣿⣿⣿⣿⣿
⠀⠀⢠⢰⠀⡁⡄⢈⠀⢻⣿⣷⣬⣤⠇⠣⢸⠐⡤⡇⢸⢰⠀⡇⡜⢈⠀⠀⠁⢠⠈⠀⡇⢣⢸⢐⣾⣿⣿⣦⠀⡇⠘⢘⠀⠅⠁⢠⠈⠀⡄⢠⣿⣿⣿⣿⣿⣿
⠀⠀⢈⢸⠀⡇⠁⠘⠀⡇⣿⡏⠙⢿⣷⣭⣹⢬⡉⡇⣼⣈⠄⡇⡱⢸⠀⡄⡇⢸⢸⡁⡇⣮⣽⣿⡿⠃⣿⣯⠄⡇⢹⢸⠀⠇⠂⠀⠰⠀⡀⢈⣿⣿⣿⣿⣿⣿
⠀⠀⠘⠘⠀⡇⠀⠸⠀⡄⣿⡇⠀⠀⠉⠿⣿⣷⣅⡇⢾⢘⣖⣇⣣⣸⣠⣄⣇⣸⣤⣤⣿⣿⠏⠉⢀⠀⢹⣿⠀⡇⢠⢸⠀⠀⠄⠀⢠⠀⠀⠘⣿⣿⣿⣿⣿⣿
⠀⠀⠸⢸⠀⡇⠇⢰⠀⠁⣻⣿⠀⠀⠀⠀⠈⠙⠻⣿⣿⠿⠛⠿⠛⠛⠛⠛⠛⠿⠿⠿⣿⡄⠀⠀⠀⠀⠘⣿⣇⡇⢍⢸⠠⡃⡤⢸⢀⠀⡃⠸⣿⣿⣿⣿⣿⣿
⠀⠀⢲⢸⠒⡇⠆⢨⠀⠃⢸⣿⠀⠀⠀⠀⠀⠀⠀⠀⠙⠃⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢻⣿⣷⡚⢺⠐⡇⠅⠸⠘⠀⠇⢲⣿⣿⣿⣿⣿⣿
⠀⠀⢥⢸⠤⡇⡄⢸⢠⠆⢼⣿⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢠⣼⡷⠾⢿⣬⡻⣿⣯⡤⡧⡇⢰⠼⠀⡦⢤⣿⣿⣿⣿⣿⣿
⠀⠀⣚⢸⡂⡇⡃⢸⠘⣆⡾⣿⡆⠀⠀⠀⠀⢀⣤⣾⠿⠿⠶⣶⣤⣀⠀⠀⠀⠀⠀⠀⠀⠀⢠⣿⠁⠀⠀⠀⠈⠻⣿⣿⣿⣖⠏⢸⢰⠃⡇⢚⣿⣿⣿⣿⣿⣿
⠀⠀⠼⢸⠠⡇⠁⢸⠀⣤⣼⡿⡃⠀⠀⠀⠠⣾⠏⠀⠀⠀⠀⠀⠙⢿⣷⣄⠀⠀⠀⠀⢐⠀⢿⣇⠀⠀⠀⠀⠀⠀⠈⢿⣞⣿⣧⢘⢠⠀⠁⠸⣿⣿⣿⣿⣿⣿
⠀⠀⣱⢸⠁⡇⣆⢸⣲⣿⡟⠀⠀⠀⠀⠀⢰⣿⠀⠀⠀⠀⠀⠀⠀⠀⠹⣿⡀⠀⠀⠀⠀⠀⠈⣿⡀⠀⠀⠀⠀⠀⠀⠀⢿⣾⣿⣏⢈⠀⡇⢰⣿⣿⣿⣿⣿⣿
⠀⠀⢠⢸⠀⡇⠄⣸⣿⡟⠀⠀⠀⠀⠀⠁⠈⣿⠀⠀⠀⠀⠀⠀⡀⠀⠀⢹⡧⠀⠀⠀⠀⠀⠀⠺⣷⡄⠀⠀⢰⣾⣶⠀⢸⣿⢻⣿⡜⠀⡇⢠⣿⣿⣿⣿⣿⣿
⠰⢶⣭⣸⢈⡇⡈⣿⡿⠀⠀⠀⠀⠀⠀⠀⠀⠻⣧⡀⠀⠀⠀⣿⣿⠀⠀⠀⣿⠀⠀⠀⠀⠀⠀⠀⠉⢿⣦⡀⠀⠙⠁⢀⣸⡿⣆⣿⣿⣀⠆⣈⣿⣿⣿⣿⣿⣿
⠀⠀⠘⢹⠻⠶⣵⣿⠃⠀⠀⠀⠀⠀⠀⠀⢀⠈⠹⣷⣄⠀⠀⠈⠁⠀⣀⣼⠿⣂⡀⢠⣦⣦⣶⣶⣶⣶⡟⠛⠷⠦⢤⡾⠋⠀⠈⠙⠛⠿⣷⣾⣿⣿⣿⣿⣿⣿
⠀⠀⡹⢼⣀⡇⢸⣿⠷⠆⠀⠀⠀⠀⠀⠀⠀⠀⢀⣬⣽⠿⢶⣤⣤⣴⠿⠋⠀⠀⠉⠉⠻⣿⣿⣿⣿⣿⠃⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣨⣿⣿⣿⣿⣿⣿⣿
⢀⣀⣰⣸⣀⣇⣸⣿⣄⡀⠀⠀⠀⠀⠀⠀⢰⠞⠉⠀⠀⠀⠀⠀⠀⢀⡀⠀⠀⠀⠀⠀⠀⠀⠉⠛⠛⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠉⠁⣿⣿⣿⣿⣿⣿
⠈⠉⢉⢹⠛⡏⡽⣿⣿⠀⠀⠀⠀⠀⠄⠆⡇⠀⠀⠀⠀⠀⠀⠀⠀⣿⠛⠛⠶⣦⣤⣀⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣀⠀⣿⣿⣿⣿⣿⣿
⠀⠀⡚⢺⣒⡗⢓⠻⣿⣷⡀⠀⠀⠀⠂⠀⡇⠀⠀⠀⠀⠀⠀⠀⠀⡟⠀⠀⠀⠀⠀⠉⠙⠛⠶⠶⣤⣤⣀⣀⣀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠉⠛⣿⣿⣿⣿⣿⣿
⠀⠀⠼⢸⠀⡇⠧⢸⠙⣿⣿⣄⠀⠀⢀⠀⠸⡀⠀⠀⠀⠀⠀⠀⠀⣿⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠉⠙⣛⣿⠟⠀⠀⠀⠀⠀⠀⢀⣴⣿⣿⣿⣿⣿⣿
⠀⠀⣰⢸⠀⡇⡆⢸⡀⡉⠻⣿⣵⣄⡀⠀⠀⠙⢄⠀⠀⠀⠀⠀⠀⠘⢷⣄⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣠⡾⠛⠁⠀⠀⠀⠀⠀⠀⣠⣿⡿⣿⣿⣿⣿⣿⣿
⠀⠀⢢⢸⠀⡇⠄⢈⠀⠇⠧⢨⠻⣿⣿⣢⢀⠀⠈⢑⢦⡀⠀⠀⠀⠀⠀⠉⠛⠶⣦⣤⣀⣀⣀⣤⣶⠾⠛⠉⠀⠀⠀⠀⠀⢀⣠⣶⣿⡿⡏⢦⣿⣿⣿⣿⣿⣿
⠀⠀⢈⢈⠀⠇⡀⠐⠀⡀⠇⢸⢰⣀⡟⠻⣿⣿⣶⣠⣄⢚⣷⠦⣄⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣀⣤⣶⣿⣿⢛⢹⡈⡇⠈⣿⣿⣿⣿⣿⣿
⠀⠀⠸⢸⠀⡇⠃⠠⠀⠀⡄⠈⢰⠂⡇⡔⢸⠉⠛⡿⢿⢿⣿⣿⣿⣭⣷⣶⣤⣄⣀⣀⣤⣤⣤⣤⣤⣴⣶⣾⡿⡿⢿⢻⠙⡆⡗⢺⢠⠀⠁⠘⣿⣿⣿⣿⣿⣿
⠀⠀⠹⢸⠀⡇⠁⢀⠀⠁⡁⠈⢈⠀⡏⡉⢸⠉⠇⣏⣹⢸⡀⡏⢏⠉⠉⠙⡛⢛⢛⠛⡉⡻⢹⢉⡍⣏⣸⠹⣉⡇⢈⢹⠈⠃⠁⠈⢈⠀⠃⠸⣿⣿⣿⣿⣿⣿
⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿
⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿
⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿
*/
package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/joho/godotenv"
	"github.com/jsusmachaca/go-router/pkg/router"
	"github.com/tiksup/tiksup-kafka-worker/api/handler"
	"github.com/tiksup/tiksup-kafka-worker/api/middleware"
	"github.com/tiksup/tiksup-kafka-worker/internal/config"
	"github.com/tiksup/tiksup-kafka-worker/internal/database"
	"github.com/tiksup/tiksup-kafka-worker/internal/service"
	"github.com/tiksup/tiksup-kafka-worker/pkg/movie"

	"go.mongodb.org/mongo-driver/mongo"
)

var (
	collection *mongo.Collection
	configMap  kafka.ConfigMap
	ctx        = context.TODO()
	db         *sql.DB
	mongoConn  movie.MongoConnection
	client     *http.Client
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("\033[31mNot .env file found. Using system variables\033[0m")
	}

	client = &http.Client{}

	configMap = config.KafkaConfig()

	var err error
	db, err = database.PGConnection()
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}

	if err = database.PGMigrate(db); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	collection, err = database.MongoConnection(ctx)
	if err != nil {
		log.Fatalf("Error trying to connect to mongo: %v", err)
	}
	mongoConn = movie.MongoConnection{Collection: collection, CTX: ctx}
}

func main() {
	go service.KafkaWorker(client, &configMap, db, mongoConn)

	router := router.NewRouter()

	userInfo := &handler.GetUserInfo{DB: db}
	randomMovies := &handler.GetRandomMovies{DB: db, MongoConn: mongoConn}
	login := &handler.Login{DB: db}
	register := &handler.Register{DB: db}

	router.Post("/api/login", nil, login)
	router.Post("/api/register", nil, register)
	router.Get("/user-info", middleware.AuthMiddleware, userInfo)
	router.Get("/random/movies", middleware.AuthMiddleware, randomMovies)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", os.Getenv("PORT")),
		Handler: router.ServeMux,
	}
	fmt.Printf("Server listen on http://localhost%s\n", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Error to initialize server %v", err)
	}
}
