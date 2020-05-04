FROM golang:latest

ENV PORT 8080

WORKDIR /go/src/github.com/Istvan/cashcalc-backend

COPY . .

RUN go get -v ./...
RUN go install ./...

EXPOSE ${PORT}

CMD ["/go/bin/cashcalc-backend"]