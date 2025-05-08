package main

import (
	"encoding/json"
	"log"
	"net/http"
)

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

func getRecords(baseURL string, apiToken string) (*NocoDBResult, error) {
	url := baseURL + "/api/v2/tables/m681e9ecsz4b669/records?fields=Title%2CStatus%2CSubject%2CReminder&where=%28Status%2Cneq%2CClosed%29~and%28Reminder%2Ceq%2Ctoday%29&limit=1000&shuffle=0&offset=0"

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
