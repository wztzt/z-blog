FROM golang:1.20.5-alpine3.18
WORKDIR /data/server
COPY . ./
go build -tags netgo -o app
EXPOSE 8081
CMD [ "./app" ]
