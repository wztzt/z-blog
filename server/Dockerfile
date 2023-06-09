FROM golang:1.20.5-alpine3.18
WORKDIR /data/server
COPY . ./
RUN go build -tags netgo -ldflags "-s -w" -o app
EXPOSE 8081
CMD [ "./app" ]
