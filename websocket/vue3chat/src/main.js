import { createApp } from "vue"
import App from "./App.vue"
import router from "./router"
import "./assets/css/idnex.css"
import "element-plus/es/components/message/style/css"

import * as ElementPlusIconsVue from "@element-plus/icons-vue"

const app = createApp(App)
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
  app.component(key, component)
}
app.use(router).mount("#app")
