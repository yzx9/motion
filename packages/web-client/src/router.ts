import { useSessionStore } from "@/store"
import NProgress from "nprogress"
import { RouteRecordRaw, createRouter, createWebHistory } from "vue-router"

export const QUERY_TO = "to"
export const QUERY_FROM = "from"

const routes: RouteRecordRaw[] = [
  { path: "/", component: () => import("./views/Home.vue") },
  {
    path: "/signin",
    meta: { public: true },
    component: () => import("./views/SignIn.vue"),
  },
  {
    path: "/signup",
    meta: { public: true },
    component: () => import("./views/SignUp.vue"),
  },
  {
    path: "/:pathMatch(.*)*",
    name: "NotFound",
    meta: { public: true },
    component: () => import("./views/NotFound.vue"),
  },
]

export const router = createRouter({
  routes,
  history: createWebHistory(import.meta.env.BASE_URL),
  scrollBehavior(to, from, savedPosition) {
    // always scroll to top
    return { top: 0 }
  },
})

// Auth
router.beforeEach((to, from) => {
  if (to.meta.public) return true

  const store = useSessionStore()
  if (store.authed) return true

  return {
    path: "/signin",
    query: {
      [QUERY_FROM]: from.fullPath,
      [QUERY_TO]: to.fullPath,
    },
  }
})

// NProgress
router.beforeEach((to, from, next) => {
  NProgress.start()
  next()
})

router.afterEach(() => NProgress.done())
