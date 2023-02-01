FROM golang:1.19.5-alpine3.16 as builder

WORKDIR /go/src/github.com/vorobeyme/chatgpt-telegram-bot
COPY . .

RUN apk add --no-cache ca-certificates curl git
RUN CGO_ENABLED=0 GOOS=`go env GOHOSTOS` GOARCH=`go env GOHOSTARCH` go build -v -o gptbot

FROM alpine:3.17.1
WORKDIR /
RUN apk add --no-cache ca-certificates
COPY --from=builder /go/src/github.com/vorobeyme/chatgpt-telegram-bot .
RUN chmod +x gptbot

CMD ["./gptbot"]
