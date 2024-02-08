package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	serverAddress := os.Getenv("SERVER_ADDRESS")

	// Endpoints
	homePage := "/"
	infoPage := "/info"

	// Request the Homepage
	response, err := http.Get(serverAddress + homePage)
	if err != nil {
		log.Fatal("Error fetching homepage:", err)
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal("Error reading homepage response:", err)
	}
	fmt.Println("Homepage Response:", string(body))

	// Request the Info Page
	response, err = http.Get(serverAddress + infoPage)
	if err != nil {
		log.Fatal("Error fetching info page:", err)
	}
	defer response.Body.Close()
	body, err = io.ReadAll(response.Body)
	if err != nil {
		log.Fatal("Error reading info page response:", err)
	}
	fmt.Println("Info Page Response:", string(body))
}
