package models

type SearchApiResponse struct {
	TotalEmails uint    `json:"total_emails"`
	Emails      []Email `json:"emails"`
}
