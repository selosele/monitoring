FROM golang:1.24

# 디버깅용 라이브러리 설치
RUN go install github.com/go-delve/delve/cmd/dlv@latest

WORKDIR /app
COPY . .
RUN go mod download

# 디버깅 포트 2345 열기
EXPOSE 2345

# Delve를 사용한 디버깅 실행
CMD ["dlv", "debug", "--headless", "--listen=:2345", "--api-version=2", "--accept-multiclient", "--log"]