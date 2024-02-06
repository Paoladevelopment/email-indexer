package zincutilities

import (
	"bytes"
	"io"
	"net/http"

	"github.com/Paoladevelopment/enron-indexer/env"
)

func makeZincRequest(method, requestURL string, jsonBody *[]byte) (*http.Request, error) {

	var bodyReader io.Reader
	if jsonBody != nil {
		bodyReader = bytes.NewReader(*jsonBody)
	}

	req, err := http.NewRequest(method, requestURL, bodyReader)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(env.EnvVars.ZincUser, env.EnvVars.ZincPass)
	req.Header.Add("Content-Type", "application/json")
	return req, nil
}
