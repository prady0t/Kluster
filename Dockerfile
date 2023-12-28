FROM golang:latest
LABEL authors="pradyotranjan"
WORKDIR /app
COPY . /app
RUN go build -o main main.go
CMD ["./main"]