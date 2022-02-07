FROM golang:1.17-alpine as builder
LABEL maintainer="Hetao<hetao@hetao.name>"
LABEL version="1.0"

WORKDIR /data/scws/

ENV GOPROXY=https://goproxy.cn,direct

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories \ 
	&& apk update && apk add tree build-base git

COPY . .

RUN	--mount=type=cache,target=/root/.cache/go-build \
	--mount=type=cache,target=/go/pkg/mod \
	go build -v -ldflags "-X main.version=1.0.0 -X main.build=`date -u +%Y-%m-%d%H:%M:%S`" -o bin/scws

FROM alpine:3.15 as prod

RUN apk --no-cache add ca-certificates

WORKDIR /data/scws/

RUN mkdir bin/
RUN mkdir -p etc/dict/
COPY etc/config.json  /data/scws/etc/
COPY etc/dict/*  /data/scws/etc/dict/
COPY --from=0 /data/scws/bin/scws bin/

HEALTHCHECK --interval=5s --timeout=5s --retries=3 \
    CMD ps aux | grep "scws" | grep -v "grep" > /dev/null; if [ 0 != $? ]; then exit 1; fi

CMD ["/data/scws/bin/scws", "-b" , "0.0.0.0:80"]
