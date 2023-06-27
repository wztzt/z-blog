FROM golang:1.20.5-alpine3.18 as builder
WORKDIR /data/server
COPY . ./
RUN set -x &&   \
    go env -w  GOPROXY=https://goproxy.cn,direct && \
    go build -tags netgo -ldflags "-s -w" -o app 

FROM alpine:3.18 as runner

WORKDIR /data/app

COPY --from=builder /data/server/app  ./
COPY --from=builder /data/server/conf ./conf

CMD ["./app"]
