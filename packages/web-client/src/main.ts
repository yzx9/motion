import { pinia } from "@/store"
import "nprogress/nprogress.css"
import { createApp } from "vue"
import App from "./App.vue"
import { router } from "./router"
import "./style.css"

createApp(App).use(pinia).use(router).mount("#app")
