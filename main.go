package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main()  {
	
	godotenv.Load()
	port := os.Getenv("PORT")
	if port == ""{
		log.Fatal("PORT not configured")
	}

	
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"HEAD", "GET", "POST", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: false,
		MaxAge: 300,
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/health", HealthHandler)
	v1Router.Get("/error", ErrorHandler)
	router.Mount("/v1", v1Router)
	
	server := &http.Server{
		Addr: ":"+port,
		Handler: router,
	}
	
	fmt.Printf("server starting on PORT %v", port)
	err := server.ListenAndServe()

	if err != nil{
		log.Fatal(err)
	}
}
