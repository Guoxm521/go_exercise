FROM golang:1.16-alpine
MAINTAINER Gxm
WORKDIR  /app
COPY . .
RUN GOPROXY="https://goproxy.io" go get -d -v ./...
RUN GOPROXY="https://goproxy.io" go build -o go_docker
CMD ["/app/go_docker"]