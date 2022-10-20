import { createRouter, createWebHashHistory } from "vue-router"
import layout from "@/layout/index.vue"

const routes = [
  {
    path: "/login",
    name: "login",
    component: () => import("./../views/login.vue"),
  },
  {
    path: "/",
    component: layout,
    children: [
      {
        path: "/",
        name: "home",
        component: () => import("@/views/HomeView.vue"),
      },
      {
        path: "/room",
        name: "room",
        component: () => import("@/views/Room.vue"),
      },
    ],
  },
]

const router = createRouter({
  history: createWebHashHistory("/"),
  routes,
})

export default router
