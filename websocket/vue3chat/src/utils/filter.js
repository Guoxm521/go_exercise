import { setToken, getToken } from "@/utils/cookie"
const judgeCookie = () => {
  if (getToken()) {
    return true
  } else {
    return false
  }
}

export { judgeCookie }
