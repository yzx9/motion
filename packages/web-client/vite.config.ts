/// <reference types="@types/node" />

import vue from "@vitejs/plugin-vue"
import { resolve } from "upath"
import { defineConfig } from "vite"

// https://vitejs.dev/config/
export default defineConfig({
  base: process.env.BASE_PATH ?? "/",
  plugins: [vue()],
  resolve: {
    alias: { "@": resolve(__dirname, "src") },
  },
})
