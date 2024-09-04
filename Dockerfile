FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

ADD cmd internal ui ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /snippets

EXPOSE 8080

CMD ["/snippets"]

