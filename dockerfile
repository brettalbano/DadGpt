FROM golang:1.20-alpine

ENV GO111MODULE=on
WORKDIR /dadgpt

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY controllers/ ./controllers/
COPY helpers/ ./helpers/
COPY initializers/ ./initializers
COPY middleware/ ./middleware/
COPY models/ ./models/
COPY .env .
COPY main.go .
COPY routes.go .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /DadGpt

EXPOSE 3000

CMD ["/DadGpt"]
