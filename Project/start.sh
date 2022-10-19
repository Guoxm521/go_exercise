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