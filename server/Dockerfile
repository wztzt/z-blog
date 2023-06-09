FROM golang:1.20.5-alpine3.18
WORKDIR /data/server
COPY . ./
RUN go build -tags netgo -ldflags "-s -w" -o app &&\
    mkdir /data/app && \
    mv ./app /data/app/ && \
    mv ./conf /data/app/ && \
    cd /data/app && \
    rm -rf /data/server

EXPOSE 8081


CMD [ "./app" ]
