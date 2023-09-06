FROM golang:latest

WORKDIR /go/src/app

COPY go.* ./

RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 3000

CMD ["./main"]
