FROM golang:latest

LABEL maintainer="Bhojpur Consulting <product@bhojpur-consulting.com>"

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go test
RUN go build -o sugactl ./cmd/client/main.go
RUN go build -o sugasvr ./cmd/server/main.go
EXPOSE $PORT

CMD ["./sugasvr", "-port=$PORT"]