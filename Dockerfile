#
# 1. Build Container
#
FROM golang:1.16.3-buster AS build

ENV GO111MODULE=on \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /src

RUN mkdir -p /src && \
    mkdir -p /app

# First add modules list to better utilize caching
COPY go.sum go.mod /src/

# Download dependencies
RUN go mod download

# Adding sourc eof project
COPY . /src

# Build components.
# Put built binaries and runtime resources in /app dir ready to be copied over or used.
RUN go install -installsuffix cgo -ldflags="-w -s" && \
    cp -r $GOPATH/bin/service /app/

#
# 2. Runtime Container
#
FROM alpine:3.15

LABEL maintainer="Mahdi Imani <imani.mahdi@gmail.com>"

ENV TZ=Asia/Tehran

RUN apk update && \
    apk add --update --no-cache wget tzdata && \
    cp --remove-destination /usr/share/zoneinfo/${TZ} /etc/localtime && \
    echo "${TZ}" > /etc/timezone

# Only for advanced/unstructrual json parsing
RUN wget https://github.com/stedolan/jq/releases/download/jq-1.6/jq-linux64 && \
    mv jq-linux64 /usr/local/bin/jq && \
    chmod +x /usr/local/bin/jq

# See http://stackoverflow.com/questions/34729748/installed-go-binary-not-found-in-path-on-alpine-linux-docker
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2

WORKDIR /app

COPY --from=build /app /app/

EXPOSE 3000

CMD ["./service"]
