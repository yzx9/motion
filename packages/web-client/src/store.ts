import { api, Tokens, User, UserDefault } from "@/api"
import { acceptHMRUpdate, createPinia, defineStore } from "pinia"
import { ref, watch } from "vue"

const KEY_REFRESH_TOKEN = "MOTION_REFRESH_TOKEN"
const PREFETCH_TOKEN = 60

export const pinia = createPinia()

export const useSessionStore = defineStore("user", () => {
  const authed = ref(true)
  const refreshToken = ref(localStorage.getItem(KEY_REFRESH_TOKEN) ?? "")
  watch(refreshToken, (v) => localStorage.setItem(KEY_REFRESH_TOKEN, v))

  const user = ref<User>(UserDefault)

  async function trySignIn(): Promise<boolean> {
    return await trySignInByRefreshToken()
  }

  async function signIn(tokens: Tokens) {
    authed.value = true
    setupTokens(tokens)

    try {
      user.value = await api.user.getCurrent()
    } catch {
      showError("获取当前用户失败")
    }
  }

  function signOut() {
    api.setToken(null)

    authed.value = false
    refreshToken.value = ""
    user.value = { ...UserDefault }
  }

  async function trySignInByRefreshToken(): Promise<boolean> {
    try {
      const tokens = await api.session.signInByRefreshToken(refreshToken.value)
      await signIn(tokens)
      return true
    } catch {
      return false
    }
  }

  function setupTokens(tokens: Tokens) {
    refreshToken.value = tokens.refreshToken
    api.setToken(tokens.accessToken)

    const t =
      tokens.expiresIn > PREFETCH_TOKEN
        ? tokens.expiresIn - PREFETCH_TOKEN
        : tokens.expiresIn
    setTimeout(async () => {
      try {
        const tokens = await api.session.signInByRefreshToken(
          refreshToken.value
        )
        setupTokens(tokens)
      } catch (e) {
        console.error(e)
      }
    }, 1000 * t)
  }

  return {
    authed,
    refreshToken,
    user,

    trySignIn,
    signIn,
    signOut,
  }
})

function showError(err: string) {
  // TODO
}

if (import.meta.hot) {
  // @ts-ignore
  import.meta.hot.accept(acceptHMRUpdate(useSessionStore, import.meta.hot))
}
