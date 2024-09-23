FROM golang:1.23.1-alpine3.19

# Install git
RUN apk add --no-cache git

WORKDIR /src/app

# Copy your application code
COPY . .

# Install the new air package
RUN go install github.com/air-verse/air@latest

# Additional commands if necessary
