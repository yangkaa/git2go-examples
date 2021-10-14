FROM golang:1.16-alpine as builder
RUN apk update && apk upgrade && apk add --no-cache alpine-sdk openssh libgit2-dev=1.1.0-r2
WORKDIR /root
COPY . .
RUN go mod vendor && go build -o app .

FROM alpine:3.14
RUN apk update && apk upgrade && apk add --no-cache openssh libgit2-dev=1.1.0-r2
WORKDIR /root
COPY --from=0 /root/app .
CMD ["./app"]
