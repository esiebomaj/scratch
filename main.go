package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/esiebomaj/rssarg/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type ApiConfig struct {
	DB *database.Queries
}

func main() {

	godotenv.Load()
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT not configured")
	}

	DB_URL := os.Getenv("DB_URL")
	if port == "" {
		log.Fatal("DB_URL not configured")
	}

	db, err := sql.Open("postgres", DB_URL)
	if err != nil {
		log.Fatalf("could not open db %v", err)
	}

	dbQueries := database.New(db)

	Config := ApiConfig{
		DB: dbQueries,
	}

	fmt.Println("DB connected successfully")

	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"HEAD", "GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/health", HealthHandler)

	v1Router.Post("/users", Config.CreateUserHandler)
	v1Router.Get("/users", Config.GetAllUser)
	v1Router.Get("/current-user", Config.AuthMiddleWare(Config.GetUser))
	
	v1Router.Get("/feeds", Config.GetAllFeeds)
	v1Router.Post("/feeds", Config.AuthMiddleWare(Config.CreateFeedHandler))

	v1Router.Post("/feed_follows", Config.AuthMiddleWare(Config.FollowFeedHandler))
	v1Router.Delete("/feed_follows/{feedFollowID}", Config.AuthMiddleWare(Config.UnfollowFeed))
	v1Router.Get("/feed_follows", Config.AuthMiddleWare(Config.GetUserFollowedFeeds))
	router.Mount("/v1", v1Router)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	fmt.Printf("server starting on PORT %v", port)
	err = server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}
