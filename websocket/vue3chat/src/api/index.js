import request from "@/utils/requeset"

export const ping = (param) =>
  request({
    url: "/ping",
    method: "get",
    data: param,
  })

export const login = (param) => {
  return request({
    url: "/account/login",
    method: "post",
    data: param,
  })
}

export const getSocketInfo = (param) => {
  return request({
    url: "/socket/info",
    method: "get",
    data: param,
  })
}

export const getAccountInfo = (param) => {
  return request({
    url: "/account/info",
    method: "get",
    data: param,
  })
}


export const getMessageList = (param) => {
  return request({
    url: "/message/list",
    method: "post",
    data: param,
  })
}
