import { createPinia, defineStore } from "pinia"
import { getAccountInfo } from "@/api"
const useUserStore = defineStore("user", {
  state: () => ({
    user_name: undefined,
    avatar: undefined,
  }),
  getters: {},
  actions: {
    setUserName(name) {
      this.user_name = name
    },
    setAvatar(avatar) {
      this.avatar = avatar
    },
    async getAccountInfo() {
      let res = await getAccountInfo({})
      if (res.code === 200) {
        this.setAvatar(res.data.avatar)
        this.setUserName(res.data.account)
      }
    },
  },
})
const pinia = createPinia()
export { useUserStore }
export default pinia
