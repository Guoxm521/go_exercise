#!/usr/bin/env bash

IMAGENAME=node_demo
CONTAINERNAME=node_demo_container
ARG1=$(docker ps -aqf "name=${CONTAINERNAME}")
ARG2=$(docker images -q --filter reference=${IMAGENAME})

build-node-demo() {
    docker build -t ${IMAGENAME} .  
}
run-node-demo() {
    docker run --privileged=true --name ${CONTAINERNAME} -p 5050:3000 -itd  ${IMAGENAME} 
}
restart-node-demo() {
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

# build-node-demo
# run-node-demo