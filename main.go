package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func main() {
	baseURL := getEnvVar("NDBRE_BASE_URL")
	apiToken := getEnvVar("NDBRE_API_TOKEN")
	tableId := getEnvVar("NDBRE_TABLE")
	receiver := getEnvVar("NDBRE_RECEIVER")

	results, err := getRecords(baseURL, apiToken, tableId)
	if err != nil {
		log.Fatalf("Could not get all results from NocoDB because of an error: %v", err)
	}

	if len(results.Rows) < 1 {
		log.Println("No applicable rows found, quitting...")
		os.Exit(0)
	}

	sendEmail(receiver)
}

func getEnvVar(name string) string {
	value, ok := os.LookupEnv(name)
	if !ok {
		log.Fatalf("Environment variable '%s' was not defined", name)
	}
	return value
}

type NocoDBResult struct {
	Rows []Row    `json:"list"`
	Info PageInfo `json:"pageInfo"`
}

type Row struct {
	Title   string `json:"Title"`
	Status  string `json:"Status"`
	Subject string `json:"Subject"`
}

type PageInfo struct {
	TotalRows   int  `json:"totalRows"`
	Page        int  `json:"page"`
	PageSize    int  `json:"pageSize"`
	IsFirstPage bool `json:"isFirstPage"`
	IsLastPage  bool `json:"isLastPage"`
}

func getRecords(baseURL string, apiToken string, tableId string) (*NocoDBResult, error) {
	url := baseURL + "/api/v2/tables/" + tableId + "/records?fields=Title%2CStatus%2CSubject%2CReminder&where=%28Status%2Cneq%2CClosed%29~and%28Reminder%2Ceq%2Ctoday%29&limit=1000&shuffle=0&offset=0"

	log.Println(url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("xc-token", apiToken)
	req.Header.Set("accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var body NocoDBResult
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		log.Printf("Could not unmarshal of the NocoDB answer: %v", err)
		return nil, err
	}

	return &body, nil
}

func sendEmail(receiver string) {

}
