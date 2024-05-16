package main

import (
	"fmt"
	"log"
	"os"

	_ "github.com/go-chi/chi"
	_ "github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("hello world")

	godotenv.Load()

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("Port is not found in env")
	}

 
}