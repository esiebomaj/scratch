package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/esiebomaj/scratch/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	_ "github.com/lib/pq"
)

type ApiConfig struct {
	DB *database.Queries
}

func main() {

	AppEnvs, err := ConfigureEnvs()
	if err != nil {
		log.Fatalf("could not configre envs %v", err)
	}

	db, err := sql.Open("postgres", AppEnvs.DB_URL)
	if err != nil {
		log.Fatalf("could not open db %v", err)
	}

	dbQueries := database.New(db)

	Config := ApiConfig{
		DB: dbQueries,
	}

	go ScrapeFeeds(dbQueries, AppEnvs.SCRAPE_INTERVAL, 10)

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

	v1Router.Get("/posts", Config.AuthMiddleWare(Config.GetRecentPosts))
	router.Mount("/v1", v1Router)

	server := &http.Server{
		Addr:    ":" + AppEnvs.PORT,
		Handler: router,
	}

	fmt.Printf("server starting on PORT %v \n", AppEnvs.PORT)
	err = server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}
