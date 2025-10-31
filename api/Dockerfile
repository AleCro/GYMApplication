# Backend Dockerfile
FROM golang:1.25.3

RUN go install github.com/air-verse/air@latest

WORKDIR /app
COPY . .

RUN go mod tidy

EXPOSE 7284
CMD ["air"]
