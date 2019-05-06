FROM golang:latest AS build

WORKDIR /go/Weather/server
COPY . .

RUN go get -d -v ./...

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main

FROM alpine:latest

RUN apk --update add ca-certificates

WORKDIR /root/
COPY --from=build /go/weather/server/main ./

ENV AWS_ACCESS_KEY_ID=AKIA34XNLPJYL72JMBX4
ENV AWS_SECRET_ACCESS_KEY=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

CMD ["./main"]
