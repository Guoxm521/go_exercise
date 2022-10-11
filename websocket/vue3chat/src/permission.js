import router from "./router"
import { setToken, getToken } from "@/utils/cookie"
import { judgeCookie } from "@/utils/filter"
import { useUserStore } from "@/store/index"

router.beforeEach((to, from, next) => {
  if (judgeCookie()) {
    const userStore = useUserStore()
    if (!userStore.user_name) {
      userStore.getAccountInfo()
    }
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
