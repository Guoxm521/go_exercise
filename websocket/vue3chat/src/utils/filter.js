import { setToken, getToken } from "@/utils/cookie"
const judgeCookie = () => {
  if (getToken()) {
    return true
  } else {
    return false
  }
}

const parseTime = (time, cFormat) => {
  if (arguments.length === 0) {
    return null
  }
  const format = cFormat || "{y}-{M}-{d} {h}:{m}:{s}"
  let date
  if (typeof time === "object") {
    date = time
  } else {
    if (("" + time).length === 10) time = time * 1000
    date = new Date(time)
  }
  const formatObj = {
    y: date.getFullYear(),
    M: date.getMonth() + 1,
    d: date.getDate(),
    h: date.getHours(),
    m: date.getMinutes(),
    s: date.getSeconds(),
    w: date.getDay(),
  }
  const time_str = format.replace(/{(y|M|d|h|m|s|w)+}/g, (result, key) => {
    let value = formatObj[key]
    if (key === "w") {
      return ["一", "二", "三", "四", "五", "六", "日"][value - 1]
    }
    if (result.length > 0 && value < 10) {
      value = "0" + value
    }
    return value || 0
  })
  return time_str
}

const getUserAvatar = (num) => {
  console.log(import.meta)
  return new URL(`./../assets/user/${num}.png`, import.meta.url).href
}

export { judgeCookie, parseTime, getUserAvatar }
