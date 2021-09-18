package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/sachinsmc/shodan/shodan"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Usage: shodan searchterm")
	}
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apiKey := os.Getenv("API_KEY")

	shodan := shodan.New(apiKey)

	info, err := shodan.APIInfo()
	if err != nil {
		log.Panicln("🚀 ~ file: main.go ~ line 26 ~ funcmain ~ err : ", err)
	}

	fmt.Printf(
		"Query Credits: %d\nScan Credits: %d\n\n",
		info.QueryCredits,
		info.ScanCredits,
	)

	hostSearch, err := shodan.HostSearch(os.Args[1])
	if err != nil {
		log.Panicln("🚀 ~ file: main.go ~ line 26 ~ funcmain ~ err : ", err)
	}

	for _, host := range hostSearch.Matches {
		fmt.Printf("%18s%8d\n", host.IPString, host.Port)
	}

}
