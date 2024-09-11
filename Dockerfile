FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY cmd ./cmd/
COPY internal ./internal/
COPY ui ./ui/

RUN CGO_ENABLED=1 GOOS=linux go build -o /snippets ./cmd/web

EXPOSE 4000

CMD ["/snippets"]

