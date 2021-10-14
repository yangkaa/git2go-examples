FROM golang:1.16-alpine
RUN apk update && apk upgrade && apk add --no-cache bash git openssh alpine-sdk libgit2-dev=1.1.0-r2

WORKDIR /root
COPY . .
RUN go mod vendor && go build -o app . && ls -al
CMD ["./app"]
