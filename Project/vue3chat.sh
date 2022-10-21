#!/usr/bin/env bash
if [ ! -d "/home/node/vue3chat" ];then
  echo "文件不存在1"
else
  rm -f /home/node/vue3chat
fi

if [ ! -d "/home/go/go_exercise/websocket" ];then
  echo "文件不存在2"
else
  cp -r /home/go/go_exercise/websocket/vue3chat /home/node/vue3chat
  docker exec -it node-demo  sh
  sleep 2s
  cd /home/node/vue3chat
  npm --version
  npm install
  sleep 2s
  npm run build
  sleep 2s
  exit
  cp -r /home/node/vue3chat/dist/* /home/nginx/html/socket/
fi