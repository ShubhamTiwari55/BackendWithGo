package main

import (
	"fmt"
	"log"
	"os"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

func main() {
	fmt.Println("Hello, World!")

	godotenv.Load()

	portString := os.Getenv("PORT")
	if portString == ""{
		log.Fatal("PORT is not found in the environment")
	}
	
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
	v1Router.Get("/ready", handlerReadiness)
	v1Router.Get("/err", handlerErr)
	router.Mount("/v1", v1Router)


	//creating the server
	srv := &http.Server{
		Handler: router,
		Addr: ":"+portString,
	}

	log.Printf("Server starting on port %v", portString)

	err := srv.ListenAndServe()
	if err != nil{
		log.Fatal(err)
	}

	fmt.Println("Port:",portString)

}