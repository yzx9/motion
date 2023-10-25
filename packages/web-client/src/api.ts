import { API } from "@motion/api"

const baseURL = import.meta.env.VITE_API_BASE ?? "/"
export const api = new API(baseURL)
export * from "@motion/api"
