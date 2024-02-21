package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Paoladevelopment/enron-indexer/models"
)

// Reads a file and parses its content in an email structure.
func GenerateEmail(filePath string) (models.Email, error) {
	var email models.Email
	partName := ""
	emailParts := make(map[string][]string)
	emailParts["Content"] = []string{}
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return email, fmt.Errorf("could not open file: %v", err)
	}
	defer file.Close()

	const maxTokenSize = 10 * 1024 //allows to read larger files
	scanner := bufio.NewScanner(file)
	scanner.Buffer(make([]byte, maxTokenSize), maxTokenSize)

	//Read file line by line
	for scanner.Scan() {
		row := scanner.Text()

		if strings.TrimSpace(row) == "" {
			partName = "Content"
		}
		ParseEmail(row, &partName, &emailParts)
	}

	email.MessageID = strings.Join(emailParts["Message-ID"], "")
	email.Date = strings.Join(emailParts["Date"], "")
	email.From = strings.Join(emailParts["From"], "")
	email.To = strings.Join(emailParts["To"], "")
	email.Subject = strings.Join(emailParts["Subject"], "")
	email.XFrom = strings.Join(emailParts["X-From"], "")
	email.XTo = strings.Join(emailParts["X-To"], "")
	email.XCc = strings.Join(emailParts["X-cc"], "")
	email.XBcc = strings.Join(emailParts["X-bcc"], "")
	email.Content = strings.Join(emailParts["Content"], "\n")

	return email, nil
}

// Parse the email to a map that contains the headers of email.
func ParseEmail(text string, partName *string, emailParts *map[string][]string) {

	partsText := strings.SplitN(text, ":", 2)
	if !isAHeader(text) {
		partsText = strings.SplitN(text, ":", 1) //This is for the case when there is a continuation of an already existing header but its new values has a ":" somewhere
	}
	//Appends to an already saved header new values that it has in a new line.
	if len(partsText) == 1 {
		(*emailParts)[*partName] = append((*emailParts)[*partName], partsText[0])
		return
	}

	*partName = partsText[0]
	switch *partName {
	case "Message-ID":
		(*emailParts)["Message-ID"] = []string{partsText[1]}
	case "Date":
		(*emailParts)["Date"] = []string{partsText[1]}
	case "From":
		(*emailParts)["From"] = []string{partsText[1]}
	case "To":
		(*emailParts)["To"] = []string{partsText[1]}
	case "Subject":
		(*emailParts)["Subject"] = []string{partsText[1]}
	case "X-From":
		(*emailParts)["X-From"] = []string{partsText[1]}
	case "X-To":
		(*emailParts)["X-To"] = []string{partsText[1]}
	case "X-cc":
		(*emailParts)["X-cc"] = []string{partsText[1]}
	case "X-bcc":
		(*emailParts)["X-bcc"] = []string{partsText[1]}
	}
}

func isAHeader(text string) bool {
	return strings.Contains(text, "Message-ID") || strings.Contains(text, "Date") || strings.Contains(text, "From") ||
		strings.Contains(text, "To") || strings.Contains(text, "Subject") || strings.Contains(text, "Mime-Version") ||
		strings.Contains(text, "Content-Type") || strings.Contains(text, "Content-Transfer-Encoding") || strings.Contains(text, "X-From") ||
		strings.Contains(text, "X-To") || strings.Contains(text, "X-cc") || strings.Contains(text, "X-bcc") || strings.Contains(text, "X-Folder") ||
		strings.Contains(text, "X-Origin") || strings.Contains(text, "X-FileName")
}
