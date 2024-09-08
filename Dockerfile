FROM golang:1.23.1-alpine
WORKDIR /app
COPY go.mod ./
COPY . .
RUN go build -o /goapp main.go
CMD ["/goapp"]