package zincutilities

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Paoladevelopment/search-api/config"
	"github.com/Paoladevelopment/search-api/models"
)

func SearchZinc(reqBody models.SearchZincRequest, indexName string) (models.SearchZincResponse, error) {
	fmt.Println(config.EnvVars.ZincURL)
	requestUrl := fmt.Sprintf("%s/%s/_search", config.EnvVars.ZincURL, indexName)

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return models.SearchZincResponse{}, fmt.Errorf("Error marshaling JSON: %v\n", err)
	}

	req, err := makeZincRequest(http.MethodPost, requestUrl, &jsonBody)
	if err != nil {
		return models.SearchZincResponse{}, fmt.Errorf("Error creating the request: %v\n", err)
	}

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return models.SearchZincResponse{}, fmt.Errorf("Error making http request: %v\n", err)
	}

	defer res.Body.Close()

	var response models.SearchZincResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return models.SearchZincResponse{}, fmt.Errorf("Error decoding JSON response: %v", err)
	}

	return response, nil
}
