package models

type Email struct {
	MessageID string   `json:"message_id"`
	Date      string   `json:"date"`
	From      string   `json:"from"`
	XFrom     string   `json:"x_from"`
	To        string   `json:"to"`
	XTo       string   `json:"x_to"`
	Subject   string   `json:"subject"`
	Cc        string   `json:"cc"`
	XCc       string   `json:"x_cc"`
	Bcc       string   `json:"bcc"`
	XBcc      string   `json:"x_bcc"`
	Content   string   `json:"content"`
	Highlight []string `json:"highlight"`
}

type SubEmail struct {
	From    string `json:"from"`
	Sent    string `json:"sent"`
	To      string `json:"to"`
	Subject string `json:"subject"`
	Content string `json:"content"`
}
