FROM golang:1.19.3-alpine
# Tag as driscoll-gcp-golang-template:latest
WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o app .
CMD ["./app"]
