package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"os"
)

func main() {
	baseURL := getEnvVar("NDBRE_BASE_URL")
	apiToken := getEnvVar("NDBRE_API_TOKEN")
	tableId := getEnvVar("NDBRE_TABLE")
	from := getEnvVar("NDBRE_EMAIL_FROM")
	smtp := getEnvVar("NDBRE_SMTP_SERVER")
	to := getEnvVar("NDBRE_EMAIL_TO")

	results, err := getRecords(baseURL, apiToken, tableId)
	if err != nil {
		log.Fatalf("Could not get all results from NocoDB because of an error: %v", err)
	}

	if len(results.Rows) < 1 {
		log.Println("No applicable rows found, quitting...")
		os.Exit(0)
	}

	sendEmail(smtp, from, to)
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

func sendEmail(smtpAddress string, sender string, receiver string) error {
	receivers := []string{receiver}
	msg := fmt.Sprintf("To: %s\r\n"+
		"Subject: NocoDB Reminder\r\n"+
		"\r\n"+
		"This is something written in the body...\r\n", receiver)

	err := smtp.SendMail(smtpAddress, nil, sender, receivers, []byte(msg))
	if err != nil {
		log.Printf("Could not send email because of the following error: %v", err)
		return err
	}

	return nil
}
