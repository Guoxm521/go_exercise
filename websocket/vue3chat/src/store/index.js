import { createPinia, defineStore } from "pinia"
import { getAccountInfo } from "@/api"
const useUserStore = defineStore("user", {
  state: () => ({
    user_name: undefined,
    avatar: undefined,
    uid: undefined,
  }),
  getters: {},
  actions: {
    setUserName(name) {
      this.user_name = name
    },
    setAvatar(avatar) {
      this.avatar = avatar
    },
    setUid(uid) {
      this.uid = uid
    },
    async getAccountInfo() {
      let res = await getAccountInfo({})
      if (res.code === 200) {
        this.setAvatar(res.data.avatar)
        this.setUserName(res.data.account)
        this.setUid(res.data.account_id)
      }
    },
  },
})
const pinia = createPinia()
export { useUserStore }
export default pinia
