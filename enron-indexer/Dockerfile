FROM golang:1.21.5-alpine

# Env vars
ENV PROJECT_DIR=/app \
    GO111MODULE=on \
    CGO_ENABLED=0 \
    GO_ENV=production

## Project setup

# Create app directory inside container
RUN mkdir app

# cd /app
WORKDIR /app

COPY go.mod .

# Download all dependencies
RUN go mod download

# Declare custom args
ARG ZINC_FIRST_ADMIN_USER \
    ZINC_FIRST_ADMIN_PASSWORD \
    ZINC_SERVER_URL \
    ENRON_URL

ENV ZINC_FIRST_ADMIN_USER=${ZINC_FIRST_ADMIN_USER} \
    ZINC_FIRST_ADMIN_PASSWORD=${ZINC_FIRST_ADMIN_PASSWORD} \
    ZINC_SERVER_URL=${ZINC_SERVER_URL} \
    ENRON_URL=${ENRON_URL}

# Copies all files from local to container
COPY . .

# Build
RUN go build -o zincindexer .

# Run
CMD ["/app/zincindexer"]