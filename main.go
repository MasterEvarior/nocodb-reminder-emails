package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	baseURL := getEnvVar("NDBRE_BASE_URL")
	apiToken := getEnvVar("NDBRE_API_TOKEN")
	tableId := getEnvVar("NDBRE_TABLE")

	getRecords(baseURL, apiToken, tableId)
}

func getEnvVar(name string) string {
	value, ok := os.LookupEnv(name)
	if !ok {
		log.Fatalf("Environment variable '%s' was not defined", name)
	}
	return value
}

func getRecords(baseURL string, apiToken string, tableId string) error {
	url := baseURL + "/" + tableId + "/records"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	req.Header.Set("xc-token", apiToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(body))

	return nil
}
