FROM golang:1.8

WORKDIR /go/src/euler
COPY . .

RUN go-wrapper install

CMD ["go-wrapper", "run"]
