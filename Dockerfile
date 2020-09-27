
#build stage
FROM golang:alpine AS builder
WORKDIR /go/src/app
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...

#final stage
FROM alpine:latest
COPY --from=builder /go/bin/cashcalc-backend /cashcalc-backend
COPY --from=builder /go/src/app/app.properties /app.properties
COPY --from=builder /go/src/app/static /static
LABEL Name=cashcalc-backend Version=0.0.1
EXPOSE 8080
CMD ["/cashcalc-backend"]
