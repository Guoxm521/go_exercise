FROM golang:1.16-alpine
MAINTAINER GxmProjecrt
WORKDIR  /app
COPY . .
RUN GOPROXY="https://goproxy.io" go get -d -v ./...
RUN GOPROXY="https://goproxy.io" go build -o go_docker_project
ENV TZ=Asia/Shanghai
CMD ["/app/go_docker_project"]