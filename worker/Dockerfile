FROM golang:latest AS build

WORKDIR /go/server
COPY . .

RUN go get -d -v ./...

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main

FROM alpine:latest

RUN apk --update add ca-certificates

WORKDIR /root/
COPY --from=build /go/server/main ./

ENV AWS_ACCESS_KEY_ID=AKIAJGYAAV2LWEN2RS6Q
ENV AWS_SECRET_ACCESS_KEY=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

CMD ["./main"]
