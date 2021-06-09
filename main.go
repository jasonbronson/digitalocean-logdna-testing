package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/logdna/logdna-go/logger"
)

func main() {
	log.Println("Testing logdna")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	key := os.Getenv("KEY")

	// Configure your options with your desired level, hostname, app, ip address, mac address and environment.
	// Hostname is the only required field in your options- the rest are optional.
	options := logger.Options{}
	options.Level = "Debug"
	options.Hostname = "gotest"
	options.App = "logdnatestapp"
	options.IPAddress = "192.168.1.1"
	options.MacAddress = "C0:FF:EE:C0:FF:EE"
	options.Env = "local"
	options.Tags = "logging,golang"

	myLogger, err := logger.NewLogger(options, key)
	if err != nil {
		fmt.Println(err)
	}
	myLogger.Log("this is a test")
}
