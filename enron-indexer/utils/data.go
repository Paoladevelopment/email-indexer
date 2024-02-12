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
	headerName := ""
	emailParts := make(map[string][]string)
	//To track the byte offset in the file, until it reaches the content of the email
	var offset int64

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return email, fmt.Errorf("Could not open file: %v\n", err)
	}
	defer file.Close()

	//Read file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := scanner.Text()
		offset += int64(len(row) + 1)
		if strings.TrimSpace(row) == "" {
			break
		}
		ParseHeader(row, &headerName, &emailParts)
	}
	content := parseContent(file, offset)

	email.MessageID = strings.Join(emailParts["Message-ID"], "")
	email.Date = strings.Join(emailParts["Date"], "")
	email.From = strings.Join(emailParts["From"], "")
	email.To = strings.Join(emailParts["To"], "")
	email.Subject = strings.Join(emailParts["Subject"], "")
	email.XFrom = strings.Join(emailParts["X-From"], "")
	email.XTo = strings.Join(emailParts["X-To"], "")
	email.XCc = strings.Join(emailParts["X-cc"], "")
	email.XBcc = strings.Join(emailParts["X-bcc"], "")
	email.Content = content

	return email, nil
}

// Parse the headers of the email to a map that contains the headers of email.
func ParseHeader(text string, headerName *string, emailParts *map[string][]string) {

	partsHeader := strings.SplitN(text, ":", 2)
	if !isAHeader(text) {
		partsHeader = strings.SplitN(text, ":", 1) //This is for the case when there is a continuation of an already existing header but its new values has a ":" somewhere
	}
	//Appends to an already saved header new values that it has in a new line.
	if len(partsHeader) == 1 {
		(*emailParts)[*headerName] = append((*emailParts)[*headerName], partsHeader[0])
		return
	}

	*headerName = partsHeader[0]
	switch *headerName {
	case "Message-ID":
		(*emailParts)["Message-ID"] = []string{partsHeader[1]}
	case "Date":
		(*emailParts)["Date"] = []string{partsHeader[1]}
	case "From":
		(*emailParts)["From"] = []string{partsHeader[1]}
	case "To":
		(*emailParts)["To"] = []string{partsHeader[1]}
	case "Subject":
		(*emailParts)["Subject"] = []string{partsHeader[1]}
	case "X-From":
		(*emailParts)["X-From"] = []string{partsHeader[1]}
	case "X-To":
		(*emailParts)["X-To"] = []string{partsHeader[1]}
	case "X-cc":
		(*emailParts)["X-cc"] = []string{partsHeader[1]}
	case "X-bcc":
		(*emailParts)["X-bcc"] = []string{partsHeader[1]}
	}
}

func isAHeader(text string) bool {
	return strings.Contains(text, "Message-ID") || strings.Contains(text, "Date") || strings.Contains(text, "From") ||
		strings.Contains(text, "To") || strings.Contains(text, "Subject") || strings.Contains(text, "Mime-Version") ||
		strings.Contains(text, "Content-Type") || strings.Contains(text, "Content-Transfer-Encoding") || strings.Contains(text, "X-From") ||
		strings.Contains(text, "X-To") || strings.Contains(text, "X-cc") || strings.Contains(text, "X-bcc") || strings.Contains(text, "X-Folder") ||
		strings.Contains(text, "X-Origin") || strings.Contains(text, "X-FileName")
}

func parseContent(file *os.File, offset int64) string {
	_, seekErr := file.Seek(offset, 0)
	if seekErr != nil {
		fmt.Printf("Error seeking file: %v\n", seekErr)
	}

	const maxTokenSize = 10 * 1024 * 1024

	scanner := bufio.NewScanner(file)
	scanner.Buffer(make([]byte, maxTokenSize), maxTokenSize)
	var contentBuilder strings.Builder

	foundEmptyLine := false
	for scanner.Scan() {
		row := scanner.Text()

		if !foundEmptyLine {
			if strings.TrimSpace(row) == "" {
				foundEmptyLine = true
			}
			continue
		}

		contentBuilder.WriteString(row)
		contentBuilder.WriteString("\n")
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading content: %v\n", err)
	}

	return contentBuilder.String()
}
