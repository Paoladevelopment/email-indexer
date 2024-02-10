package handlers

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"

	"github.com/Paoladevelopment/search-api/models"
	zincutilities "github.com/Paoladevelopment/search-api/zincUtilities"
)

func SearchByPhrase(w http.ResponseWriter, r *http.Request) {
	searchType := "matchphrase"
	page := 0
	resultsPerPage := 100
	query := r.URL.Query()
	term := query.Get("term")

	err := ParseQueryParameterToInt(query, "page", &page)
	if err != nil {
		ResponseWithError(w, http.StatusBadRequest, `"page" parameter is invalid.`)
	}

	err = ParseQueryParameterToInt(query, "results_per_page", &resultsPerPage)
	if err != nil {
		ResponseWithError(w, http.StatusBadRequest, `"results_per_page" parameter is invalid.`)
	}
	if term == "" {
		searchType = "alldocuments"
	}

	searchBody := models.SearchZincRequest{
		SearchType: searchType,
		Query: struct {
			Term  string "json:\"term\""
			Field string "json:\"field\""
		}{
			Term:  term,
			Field: "_all",
		},
		SortFields: []string{},
		From:       uint(page),
		MaxResults: uint(resultsPerPage),
		Source:     []string{},
		Highlight: struct {
			Fields map[string]interface{} "json:\"fields\""
		}{
			Fields: map[string]interface{}{
				"content": struct{}{},
				"subject": struct{}{},
				"from":    struct{}{},
				"to":      struct{}{},
			},
		},
	}

	response, err := zincutilities.SearchZinc(searchBody, "emails")
	if err != nil {
		ResponseWithError(w, http.StatusInternalServerError, err.Error())
	}

	emails := MapHitsToMails(response)
	searchResponse := models.SearchApiResponse{
		TotalEmails: response.Hits.Total.Value,
		Emails:      emails,
	}

	ResponseJSON(w, http.StatusOK, searchResponse)
}

func MapHitsToMails(zincResponse models.SearchZincResponse) []models.Email {
	emails := []models.Email{}
	for _, email := range zincResponse.Hits.Hits {
		emails = append(emails, models.Email{
			MessageID: email.Source.MessageID,
			Date:      email.Source.Date,
			From:      email.Source.From,
			XFrom:     email.Source.XFrom,
			To:        email.Source.To,
			XTo:       email.Source.XTo,
			Subject:   email.Source.Subject,
			Cc:        email.Source.Cc,
			XCc:       email.Source.XCc,
			Bcc:       email.Source.Bcc,
			XBcc:      email.Source.XBcc,
			Content:   email.Source.Content,
			Subemails: email.Source.Subemails,
		})
	}

	return emails
}

// Parses an integer query parameter from the given URL values and stores the result in the provided integer pointer.
// If the parameter is missing or cannot be parsed as an integer, it returns an error.
func ParseQueryParameterToInt(query url.Values, parameterName string, varInt *int) error {
	var err error
	if query.Has(parameterName) {
		*varInt, err = strconv.Atoi(query.Get(parameterName))
		if err != nil {
			return err
		}
	}
	return nil
}

func ResponseWithError(w http.ResponseWriter, statusCode int, err string) {
	ResponseJSON(w, statusCode, map[string]string{"error": err})
}

func ResponseJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(response)
}
