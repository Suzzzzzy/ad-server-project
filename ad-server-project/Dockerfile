FROM golang:1.18-alpine

RUN apk --no-cache add bash

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GIN_MODE=release

WORKDIR /build

COPY . .

RUN go mod download
RUN go build -o ad-server-project ./src

WORKDIR /app

RUN cp /build/ad-server-project .
RUN cp /build/wait-for-it.sh .

EXPOSE 8080

CMD ["/app/ad-server-project"]

