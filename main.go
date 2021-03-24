package main

import (
	"fmt"
	"github.com/Mishanki/rest-api-course/handlers"
	"github.com/Mishanki/rest-api-course/helpers"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func init() {
	helpers.InitEnv()
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/grab", handlers.CreateGrab).Methods("POST")
	router.HandleFunc("/solve", handlers.GetSolve).Methods("GET")

	port := os.Getenv("PORT")
	if port == "" {
		panic("Param PORT is not found in .env")
	}

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Print(err)
	}
}
