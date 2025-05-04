FROM golang:1.24

WORKDIR /app
COPY src/ ./
RUN go mod download

CMD ["go", "run", "main.go"]