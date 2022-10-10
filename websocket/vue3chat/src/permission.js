import router from "./router"
import { setToken, getToken } from "@/utils/cookie"
import { judgeCookie } from "@/utils/filter"

router.beforeEach((to, from, next) => {
  if (judgeCookie()) {
    if (to.path === "/login") {
      next({ path: "/" })
    } else {
      next()
    }
  } else {
    if (to.path === "/login") {
      next()
    } else {
      next({ path: "/login" })
    }
  }
})

router.afterEach((to, from) => {})
