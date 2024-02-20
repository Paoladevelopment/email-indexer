package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/Paoladevelopment/enron-indexer/models"
)

// Reads a file and parses its content in an email structure.
func GenerateEmail(filePath string) (models.Email, error) {
	var email models.Email

	// Read the file
	file, err := os.ReadFile(filepath.Clean(filePath))
	if err != nil {
		return email, fmt.Errorf("could not open file: %v", err)
	}

	//Separate headers emails and content
	emailDetails := strings.SplitN(string(file), "\r\n\r\n", 2)

	if len(emailDetails) != 2 {
		return email, fmt.Errorf("email in bad format found at: %v", filePath)
	}

	header, content := emailDetails[0], emailDetails[1]

	headersPart := strings.Split(header, "\r\n")

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		ParseHeader(headersPart, &email)
	}()

	wg.Wait()

	email.Content = content
	return email, nil
}

func ParseHeader(headersPart []string, email *models.Email) {
	headerName := ""
	parsedHeader := make(map[string][]string)

	for _, part := range headersPart {
		MapToHeaderEmail(part, &headerName, &parsedHeader)
	}

	email.MessageID = strings.Join(parsedHeader["Message-ID"], "")
	email.Date = strings.Join(parsedHeader["Date"], "")
	email.From = strings.Join(parsedHeader["From"], "")
	email.To = strings.Join(parsedHeader["To"], "")
	email.Subject = strings.Join(parsedHeader["Subject"], "")
	email.XFrom = strings.Join(parsedHeader["X-From"], "")
	email.XTo = strings.Join(parsedHeader["X-To"], "")
	email.XCc = strings.Join(parsedHeader["X-cc"], "")
	email.XBcc = strings.Join(parsedHeader["X-bcc"], "")
}

func MapToHeaderEmail(text string, headerName *string, parsedHeader *map[string][]string) {

	partsHeader := strings.SplitN(text, ":", 2)
	if !isAHeader(text) {
		partsHeader = strings.SplitN(text, ":", 1) //This is for the case when there is a continuation of an already existing header but its new values has a ":" somewhere
	}
	//Appends to an already saved header new values that it has in a new line.
	if len(partsHeader) == 1 {
		(*parsedHeader)[*headerName] = append((*parsedHeader)[*headerName], partsHeader[0])
		return
	}

	*headerName = partsHeader[0]
	switch *headerName {
	case "Message-ID":
		(*parsedHeader)["Message-ID"] = []string{partsHeader[1]}
	case "Date":
		(*parsedHeader)["Date"] = []string{partsHeader[1]}
	case "From":
		(*parsedHeader)["From"] = []string{partsHeader[1]}
	case "To":
		(*parsedHeader)["To"] = []string{partsHeader[1]}
	case "Subject":
		(*parsedHeader)["Subject"] = []string{partsHeader[1]}
	case "X-From":
		(*parsedHeader)["X-From"] = []string{partsHeader[1]}
	case "X-To":
		(*parsedHeader)["X-To"] = []string{partsHeader[1]}
	case "X-cc":
		(*parsedHeader)["X-cc"] = []string{partsHeader[1]}
	case "X-bcc":
		(*parsedHeader)["X-bcc"] = []string{partsHeader[1]}
	}
}

func isAHeader(text string) bool {
	return strings.Contains(text, "Message-ID") || strings.Contains(text, "Date") || strings.Contains(text, "From") ||
		strings.Contains(text, "To") || strings.Contains(text, "Subject") || strings.Contains(text, "Mime-Version") ||
		strings.Contains(text, "Content-Type") || strings.Contains(text, "Content-Transfer-Encoding") || strings.Contains(text, "X-From") ||
		strings.Contains(text, "X-To") || strings.Contains(text, "X-cc") || strings.Contains(text, "X-bcc") || strings.Contains(text, "X-Folder") ||
		strings.Contains(text, "X-Origin") || strings.Contains(text, "X-FileName")
}
