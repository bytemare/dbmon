# 1. Build exec
FROM golang:alpine AS builder
RUN apk update && apk add --no-cache git
RUN adduser -D -g '' dbmon

WORKDIR $GOPATH/src/github.com/bytemare/dbmon/
COPY *.go ./
COPY .git ./
COPY connectors/ ./connectors/
COPY app/ ./app/
RUN go get -d ./...
#RUN GOOS=linux GOARCH=amd64 go build -gccgoflags="-w -s" -a -installsuffix cgo -o $GOPATH/bin/dbmon app/dbmon.go
#RUN GOOS=linux GOARCH=amd64 go build -o $GOPATH/bin/dbmon ./app/dbmon.go
RUN GOOS=linux GOARCH=amd64 go build -o /bin/dbmon ./app/dbmon.go

# 2. Build image
FROM scratch
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /bin/dbmon $GOPATH/bin/dbmon
USER dbmon
EXPOSE 4000
ENTRYPOINT ["$GOPATH/bin/dbmon"]