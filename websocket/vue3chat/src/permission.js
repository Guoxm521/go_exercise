import router from "./router"
import { judgeCookie } from "@/utils/filter"
import { useUserStore } from "@/store/index"

let whiteList = ["/room"]
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
      if (whiteList.indexOf(to.path) !== -1) {
        next()
      } else {
        next({ path: "/login" })
      }
    }
  }
})

router.afterEach((to, from) => {})
