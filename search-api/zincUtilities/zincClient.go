package zincutilities

import (
	"bytes"
	"net/http"

	"github.com/Paoladevelopment/search-api/config"
)

func makeZincRequest(method, requestURL string, jsonBody *[]byte) (*http.Request, error) {

	bodyReader := bytes.NewReader(*jsonBody)
	req, err := http.NewRequest(method, requestURL, bodyReader)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(config.EnvVars.ZincUser, config.EnvVars.ZincPass)
	req.Header.Add("Content-Type", "application/json")
	return req, nil
}
