package main

import (
	"log"
	"os"
)

func main() {
	BASE_URL := getEnvVar("NDBRE_BASE_URL")
	API_TOKEN := getEnvVar("NDBRE_API_TOKEN")
	FROM := getEnvVar("NDBRE_EMAIL_FROM")
	SMTP := getEnvVar("NDBRE_SMTP_SERVER")
	TO := getEnvVar("NDBRE_EMAIL_TO")

	log.Printf("Collecting data from NocoDB")
	results, err := getRecords(BASE_URL, API_TOKEN)
	if err != nil {
		log.Fatalf("Could not get all results from NocoDB because of an error: %v", err)
	}

	if len(results.Rows) < 1 {
		log.Println("No applicable rows found, quitting...")
		os.Exit(0)
	}

	log.Printf("Found %d applicable rows", len(results.Rows))
	err = SendEmail(SMTP, FROM, TO, results.Rows)
	if err != nil {
		log.Fatalf("Could not send an email because of an error: %v", err)
	}
}

func getEnvVar(name string) string {
	value, ok := os.LookupEnv(name)
	if !ok {
		log.Fatalf("Environment variable '%s' was not defined but is mandatory", name)
	}
	return value
}
