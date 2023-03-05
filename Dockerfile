FROM golang:1.19.6

RUN mkdir /app

ADD . /app

WORKDIR /app

COPY go.mod ./

RUN go mod download

RUN go build -o main .

CMD ["/app/main"]