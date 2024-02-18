# Email indexer from the Enron database
This repository contains three aplications: `enron-indexer`,`search-api` and `search-client`. The purpose of the three is to allow searching information from enron database and visualize it. 

## First steps
- Download the latest version of ZincSearch from its [repository](https://github.com/zincsearch/zincsearch/releases) or run it with docker.
- Download the Enron emails database [here](http://www.cs.cmu.edu/~enron/enron_mail_20110402.tgz).
- Use a tool to extract the folder from the archive and place it here:
  ```bash
  email-indexer
  │
  ├── enron-indexer
  │   ├── Place enron emails folder db here
  │   │   
  │   └── ... 
  ├── search-api
  │   └── ... 
  ├── .gitignore
  ├── docker-compose.yml
  └── README.md
  ```

## Enron-indexer
This app parses the raw files from the Enron Database and upload emails to a ZincZearch index.

### Build
- Get dependencies
```bash
go mod download
```
- Build app
```bash
go build -o zincindexer
```
### Environment variables
Don't forget to include the following environment variables before running the `./zincindexer` command:

```bash
export ZINC_FIRST_ADMIN_USER="admin"
export ZINC_FIRST_ADMIN_PASSWORD="Complexpass#123"
export ZINC_SERVER_URL="http://localhost:4080/api"
export ENRON_URL="./enron_mail_20110402" #where you place your enron database folder
./zincindexer
```
_Remember that your environment credentials variables must match with your ZincSearch server credentials._

## Search-api
Small API to facilitate searching for words within the Enron database of emails.

### Build
- Get dependencies
```bash
go mod download
```
- Build app
```bash
go build -o api-search
```
### Environment variables
Don't forget to include the following environment variables before running the `./api-search` command:

```bash
export ZINC_FIRST_ADMIN_USER="admin"
export ZINC_FIRST_ADMIN_PASSWORD="Complexpass#123"
export ZINC_SERVER_URL="http://localhost:4080/api"
./api-search
```
_Remember that your environment credentials variables must match with your ZincSearch server credentials._

_Server will be listening on port 8080_

### Documentation api

```http
  GET /api/search_emails
```
| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `term` | `string` | The word to be searched in emails.|
| `from_email` | `integer` | The starting index from which to retrieve emails. Defaults to `0`.|
| `max_emails` | `integer` | The maximum number of emails to return per page. Defaults to `100`.|

####Examples

-Retrieve the first 100 emails:
```http
  GET /api/search_emails
```

-Search emails containing the word "ice cream" starting from the 100th email and return 20 emails per page:
```http
GET /api/search_emails?term="word"&from_email=100&max_emails=20
```
Ensure to replace `http://localhost:8080` with your actual base URL when utilizing this documentation.

## Search client
A visualizer for the emails and to search emails based on user's input. Build with Vue 3 and Tailwind.
### Environment variables
Create `.env` file and inclued:
`VITE_API_URL=http://localhost`
`VITE_API_PORT=8080`

### Install
```bash
  npm install
```

## Run locally, in development environment.
```bash
  npm run dev
```
