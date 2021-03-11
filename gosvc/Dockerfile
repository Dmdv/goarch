FROM golang:1.16 AS builder

RUN mkdir /app
ADD . /app
WORKDIR /app

ENV CGO_ENABLED 0
ENV GOOS linux

RUN go build -ldflags="-s -w" -o app cmd/server/main.go

FROM alpine:latest as production

RUN apk update --no-cache && apk add --no-cache ca-certificates tzdata
ENV TZ Asia/Shanghai

COPY --from=builder /app .

CMD ["./app"]