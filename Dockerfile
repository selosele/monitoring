FROM golang:1.24

WORKDIR /app

COPY src/go.mod src/go.sum src/main.go ./

RUN go mod download

CMD ["go", "run", "main.go"]