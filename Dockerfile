FROM golang:latest

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o SportNotes ./cmd/main.go

EXPOSE 8080

CMD ["./SportNotes"]