import axios from "axios"
import { getToken, removeToken } from "./cookie"
import router from "@/router/index"
import { ElMessage } from "element-plus"

const service = axios.create({
  baseURL: process.env.VUE_APP_BASE_API,
  timeout: 10000,
})

service.interceptors.request.use(
  (config) => {
    if (getToken()) {
      config.headers["Authorization"] = getToken()
    }
    return config
  },
  (error) => {
    Promise.reject(error)
  }
)

service.interceptors.response.use((response) => {
  const { status, data } = response
  console.log(status)
  if (data.code === 40101) {
    removeToken()
    ElMessage({
      message: "token已失效，请重新登录",
      type: "warning",
    })
    router.push("/login")
  }
  return data
})

export default service
