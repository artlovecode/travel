package main

import (
	"artlovecode/travel/apiclients"
	"artlovecode/travel/formatters"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)


type Country struct {
	ISOShortCode string `json:"iso"`
	Name         string `json:"name"`
}

/*
Select a country at random, and display some fun travel facts about it!
*/
func main() {
	// Contains { iso: "XX", name: "XXXX" } mappings of country names and country codes
	countriesFile, fileIoErr := os.Open("mappings.json")

	if fileIoErr != nil {
		fmt.Print(fileIoErr)
		os.Exit(1)
	}

	defer countriesFile.Close()

	countryBytes, fileIoErr := ioutil.ReadAll(countriesFile)
	var countries [238]Country
	json.Unmarshal(countryBytes, &countries)

	rand.Seed(time.Now().Unix())
	country := countries[rand.Intn(238)]

  metadataClient := apiclients.ComposeApiClient("https://www.distance24.org/")
  advisoryClient := apiclients.ComposeApiClient("https://www.travel-advisory.info/api")
  
	advisoryResponse, advisoryErr := advisoryClient("/?countrycode=" + country.ISOShortCode)
	metadataResponse, metadataErr := metadataClient("/route.json?stops=Norway|" + country.Name)

	if advisoryErr != nil {
		fmt.Print(advisoryErr)
		os.Exit(1)
	}

	if metadataErr != nil {
		fmt.Print(metadataErr)
		os.Exit(1)
	}

	fmt.Println("Your next destination is...", strings.ToUpper(country.Name)+"\n\n")
	fmt.Println(formatters.FormatAdvisory(advisoryResponse, country.ISOShortCode) + "\n\n")
	fmt.Println(formatters.FormatMetaData(metadataResponse) + "\n\n")
}
