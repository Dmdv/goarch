# Service generation with goctl from go-zero library

## Generate API

### Generate API declaration
```bash
goctl api -o sample.api
```

### Generate API implementation
```bash
goctl api go -api sample.api -dir sample -style go-zero
```

## Generate service

It will generate API declaration with implementation

```bash
goctl api new sample
```

## Docker generation
```bash
goctl docker -go hello.go
```

It will generate docker image

```dockerfile
FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOOS linux

WORKDIR /build/zero

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .
COPY ./etc /app/etc
RUN go build -ldflags="-s -w" -o /app/hello ./hello.go


FROM alpine

RUN apk update --no-cache && apk add --no-cache ca-certificates tzdata
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/hello /app/hello
COPY --from=builder /app/etc /app/etc

CMD ["./hello", "-f", "etc/hello-api.yaml"]
```