FROM golang:1.17-alpine as builder
LABEL maintainer="Hetao<hetao@hetao.name>"
LABEL version="1.0"

WORKDIR /data/scws/

COPY . .
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories \
	&& apk update && apk add git tree build-base \
	&& tree -L 3
RUN export GOPROXY=https://goproxy.cn && go build -o bin/scws && rm -rf /var/lib/apk/*

FROM alpine:3.14 as prod

RUN apk --no-cache add ca-certificates

WORKDIR /data/scws/

RUN mkdir bin/
RUN mkdir dict/
COPY dict/*  /data/scws/dict/
COPY --from=0 /data/scws/bin/scws bin/

HEALTHCHECK --interval=5s --timeout=5s --retries=3 \
    CMD ps aux | grep "scws" | grep -v "grep" > /dev/null; if [ 0 != $? ]; then exit 1; fi

CMD ["/data/scws/bin/scws", "-b" , "0.0.0.0:80"]
