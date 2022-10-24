### 一、简介

> 一个练手的小项目，为熟悉运用go、gin框架、websocket的使用，前端运用的是vue3
>
> [在线demo](http://demo.gxming.top/socket/#/)
>
> [Github仓库](https://github.com/Guoxm521/go_exercise)

### 二、技术使用

- go
- gin框架
- websocket
- mysql
- docker编译部署
- 前端nginx部署
- shell脚本部署

三、项目结构

```js

Project
├─ common
│  └─ response.go
├─ config
│  └─ app.yml
├─ config.pro
│  └─ app.yml
├─ controller
│  ├─ account
│  │  ├─ ctrAccountAdd.go
│  │  ├─ ctrAccountInfo.go
│  │  └─ ctrAccountLogin.go
│  ├─ github
│  │  └─ ctrTrendingList.go
│  ├─ message
│  │  └─ ctrMessageList.go
│  ├─ ping
│  │  └─ ctrPing.go
│  └─ ws
│     └─ ctrWs.go
├─ database
│  ├─ config.go
│  └─ dbCon.go
├─ Dockerfile
├─ go.mod
├─ go.sum
├─ go_docker_project
├─ init.go
├─ main.go
├─ make_image.sh
├─ middleware
│  ├─ token.go
│  └─ utils.go
├─ model
│  ├─ Account.go
│  ├─ db
│  │  ├─ init.go
│  │  ├─ tb.go
│  │  ├─ tbAccount.go
│  │  ├─ tbGithubLanguage.go
│  │  ├─ tbGithubSince.go
│  │  ├─ tbGithubTrending.go
│  │  └─ tbMessage.go
│  ├─ githubLanguage.go
│  ├─ githubTrending.go
│  ├─ init.go
│  ├─ Message.go
│  └─ sync.go
├─ README.md
├─ route
│  └─ route.go
├─ run_image.sh
├─ spider
│  └─ githubTrending.go
├─ start.sh
├─ vue3chat.sh
└─ websocket
   ├─ main.go
   ├─ manager.go
   └─ ws.go
```

### 三、nginx部署

```nginx
server {
        listen       80;
        server_name  api.gxming.top;
        location / {
            proxy_pass http://172.17.0.1:3000;
        }

        location ^~ /go/socket/ {
            proxy_pass http://172.17.0.1:8010;
            proxy_redirect off;
            proxy_http_version 1.1;
            proxy_set_header Upgrade $http_upgrade;
            proxy_set_header Connection "upgrade";
            proxy_set_header Host $host:$server_port;
            proxy_set_header X-Real-IP $remote_addr;
            proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;    
            proxy_read_timeout 6000s;
        }
        location ^~ /go {
            proxy_pass http://172.17.0.1:8010;
        }
    } 
```

### 四、docker部署

```yaml
FROM golang:1.16-alpine
MAINTAINER GxmProjecrt
WORKDIR  /app
COPY . .
RUN GOPROXY="https://goproxy.io" go get -d -v ./...
RUN GOPROXY="https://goproxy.io" go build -o go_docker_project -ldflags "-X database/config.c=config.pro"
ENV TZ=Asia/Shanghai
CMD ["/app/go_docker_project"]
```

### 五、shell脚本

```sh
#!/usr/bin/env bash
IMAGENAME=go_exercise_project
CONTAINERNAME=go_exercise_project_container
ARG1=$(docker ps -aqf "name=${CONTAINERNAME}")
ARG2=$(docker images -q --filter reference=${IMAGENAME})

build() {
    docker build -t ${IMAGENAME} .
}
run() {
    docker run --privileged=true --name ${CONTAINERNAME} -p 8010:8010 -itd  ${IMAGENAME}
}
restart() {
    docker restart ${CONTAINERNAME}
}

# 删除容器
if [  -n "$ARG1" ]; then
 docker rm -f $(docker stop $ARG1)
 echo "$CONTAINERNAME容器停止删除成功.....！！！"
fi

#删除镜像
if [  -n "$ARG2" ]; then
 docker rmi -f $ARG2
 echo "$IMAGENAME镜像删除成功.....！！！"
fi

build
run

echo '执行成功'
```

### 六、License

MIT

















