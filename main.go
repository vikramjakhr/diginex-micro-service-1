package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

// Command line flags
var port = flag.String("port", "8080", "port on which application will run")
var diginex_micro_sercvice_2_base_url = flag.String("diginex_micro_sercvice_2_base_url", "",
	"Endpoint of diginex micro sercvice 2 for accessing the reverse string API")

const usage = `Diginex Micro Service 1.

Usage:

  diginex-micro-service-1 [commands|flags]

The commands & flags are:

  help              prints help

  --port              port on which application will run (default: 8080)
  -- diginex_micro_sercvice_2_base_url		Base URL of diginex micro sercvice 2 for accessing the reverse string API

Examples:

  # prints help:
  diginex-micro-service-1 help

  # sample usage
  diginex-micro-service-1 --diginex_micro_sercvice_2_base_url=http://localhost:8081
  diginex-micro-service-1 --port=8080 --diginex_micro_sercvice_2_base_url=http://localhost:8081
`

// Show usage message and exit
func usageExit(rc int) {
	fmt.Println(usage)
	os.Exit(rc)
}

// Validate the command line flags
func validate() {
	if _, err := strconv.Atoi(*port); err != nil {
		log.Printf("ERROR: Please enter a valid port number.\n\n")
		usageExit(1)
	}

	if *diginex_micro_sercvice_2_base_url == "" {
		log.Printf("ERROR: Please specify the Diginex micro service 2 API endpoint.\n\n")
		usageExit(1)
	}
}

// Entrypoint
func main() {

	// Parse flags
	flag.Usage = func() { usageExit(0) }
	flag.Parse()
	args := flag.Args()

	if len(args) > 0 {
		switch args[0] {
		case "help":
			usageExit(0)
			return
		}
	}

	// Validate flags
	validate()

	// Register API endpoints
	http.HandleFunc("/health", health)
	http.HandleFunc("/api", handler)

	// Listen and serve API's
	if err := http.ListenAndServe(":"+*port, nil); err != nil {
		log.Printf("ERROR: %q", err)
	}
}
