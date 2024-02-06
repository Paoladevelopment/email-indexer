package zincutilities

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/Paoladevelopment/enron-indexer/env"
	"github.com/Paoladevelopment/enron-indexer/models"
)

func CreateIndex(index models.Index) {
	if isIndexCreated(index.Name) {
		return
	}

	requestURL := fmt.Sprintf("%s/index", env.EnvVars.ZincURL)

	jsonBody, err := json.Marshal(index)
	if err != nil {
		log.Printf("Error marshaling JSON: %s", err)
	}

	req, err := makeZincRequest(http.MethodPost, requestURL, &jsonBody)
	if err != nil {
		log.Printf("Error creating the request: %s", err)
	}

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Printf("Error making http request: %s", err)
	}

	defer res.Body.Close()
	if res.StatusCode == http.StatusOK {
		fmt.Println("Index created correctly!")
	}
}

func isIndexCreated(index string) bool {

	requestURL := fmt.Sprintf("%s/index/%s", env.EnvVars.ZincURL, index)
	req, err := makeZincRequest(http.MethodHead, requestURL, nil)
	if err != nil {
		log.Printf("Error creating the request: %s", err)
	}

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Printf("Error making http request: %s", err)
	}

	defer res.Body.Close()
	if res.StatusCode == http.StatusOK {
		return true
	}

	return false
}

func SaveBulk(name string, emails []models.Email) {
	requestURL := fmt.Sprintf("%s/_bulkv2", env.EnvVars.ZincURL)
	bulk := models.IndexBody{
		IndexName: name,
		Records:   emails,
	}

	bulkJson, _ := json.Marshal(bulk)
	fmt.Println("Posting bulk to Zinc server")

	req, err := makeZincRequest(http.MethodPost, requestURL, &bulkJson)
	if err != nil {
		log.Printf("Error creating the request: %s", err)
	}

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Printf("Error making http request: %s", err)
	}

	defer res.Body.Close()
	fmt.Printf("Zinc server response code: %d\n", res.StatusCode)

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		log.Print(err)
		return
	}
	fmt.Printf("Zinc server response body: %s\n", string(resBody))
}
