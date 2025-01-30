package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/ShubhamTiwari55/helloGo/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct{
	DB *database.Queries
}

func main() {
	fmt.Println("Hello, World!")
	godotenv.Load()

	//connecting to the database
	portString := os.Getenv("PORT")
	if portString == ""{
		log.Fatal("PORT is not found in the environment")
	}
	
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL is not found in the environment")
	}

	conn, err := sql.Open("postgres", dbURL)
	if err != nil{
		log.Fatal("Error connecting to the database: ", err)
	}


	apiCfg := apiConfig{
		DB: database.New(conn),
	}

	go startScraping(database.New(conn), 10, time.Minute)

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},	//domain names
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, 	//allowed http methods
		AllowedHeaders: []string{"*"},		//allowed headers that client is allowed to send
		ExposedHeaders: []string{"Link"},	//headers that client is allowed to access in response
		AllowCredentials: false,	//credentials like cookies are allowed
		MaxAge: 300,		//maximum duration in seconds that the result of preflight req can be cached by client
	}))
	
	v1Router := chi.NewRouter()
	router.Mount("/v1", v1Router)
	v1Router.Get("/ready", handlerReadiness)
	v1Router.Get("/err", handlerErr)
	v1Router.Post("/users", apiCfg.handleCreateUser)
	v1Router.Get("/users", apiCfg.middlewareAuth(apiCfg.handleGetUser))
	v1Router.Post("/feeds", apiCfg.middlewareAuth(apiCfg.handleCreateFeed))
	v1Router.Get("/feeds", apiCfg.handleGetFeeds)
	v1Router.Post("/feeds_follows", apiCfg.middlewareAuth(apiCfg.handlerFeedFollowCreate))
	v1Router.Get("/feeds_follows", apiCfg.middlewareAuth(apiCfg.handlerGetFeedFollow))
	v1Router.Delete("/feeds_follows/{feedFollowID}", apiCfg.middlewareAuth(apiCfg.handlerFeedFollowDelete))
	v1Router.Get("/posts", apiCfg.middlewareAuth(apiCfg.handleGetPostForUsers))

	//creating the server
	srv := &http.Server{
		Handler: router,
		Addr: ":"+portString,
	}

	log.Printf("Server starting on port %v", portString)

	err = srv.ListenAndServe()
	if err != nil{
		log.Fatal(err)
	}

	fmt.Println("Port:",portString)

}