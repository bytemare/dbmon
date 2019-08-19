# 1. Build exec
FROM golang:alpine AS builder
RUN apk update && apk add --no-cache git
RUN adduser -D -g '' dbmon

WORKDIR $GOPATH/src/github.com/bytemare/dbmon/
COPY *.go ./
COPY dbmon ./dbmon/
COPY connectors ./connetors/
COPY app ./app/
RUN go get -d -v ./...
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-w -s" -a -installsuffix cgo app/dbmon.go -o $GOPATH/bin/dbmon

# 2. Build image
FROM scratch
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder $GOPATH/bin/dbmon $GOPATH/bin/dbmon
USER dbmon
EXPOSE 4000
ENTRYPOINT ["$GOPATH/bin/dbmon"]


