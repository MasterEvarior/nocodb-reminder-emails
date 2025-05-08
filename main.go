package main

import (
	"log"
	"os"
)

func main() {
	baseURL := getEnvVar("NDBRE_BASE_URL")
	apiToken := getEnvVar("NDBRE_API_TOKEN")
	from := getEnvVar("NDBRE_EMAIL_FROM")
	smtp := getEnvVar("NDBRE_SMTP_SERVER")
	to := getEnvVar("NDBRE_EMAIL_TO")

	log.Printf("Collecting data from NocoDB")
	results, err := getRecords(baseURL, apiToken)
	if err != nil {
		log.Fatalf("Could not get all results from NocoDB because of an error: %v", err)
	}

	if len(results.Rows) < 1 {
		log.Println("No applicable rows found, quitting...")
		os.Exit(0)
	}

	log.Printf("Found %d applicable rows", len(results.Rows))
	err = SendEmail(smtp, from, to, results.Rows)
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
