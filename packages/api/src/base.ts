import { Err as i18n } from "@motion/i18n"
import { joinPath, startsWithProtocol } from "@motion/shared"
import qs from "qs"

export type Params = Record<string, string | number> | null
export type Form = FormData | Record<string, any> | string | any[] | null
export type Pager<T> = {
  total: number
  data?: T[] | null
}

export class RequestProvider {
  token: string | null = null
  baseURL: string

  constructor(baseURL: string) {
    if (!baseURL.endsWith("/")) {
      baseURL += "/"
    }

    this.baseURL = baseURL
  }

  async get(
    url: string,
    params: Params = null,
    options: {
      withoutAuth?: boolean
    } = {}
  ): Promise<any> {
    return this.fetchWithOptions({
      method: "GET",
      ...options,
      url,
      params,
    })
  }

  async post(
    url: string,
    form: Form = null,
    options: {
      params?: Params
      withoutAuth?: boolean
    } = {}
  ): Promise<any> {
    return this.fetchWithOptions({
      method: "POST",
      ...options,
      url,
      form,
    })
  }

  async put(
    url: string,
    form: Form = null,
    options: {
      params?: Params
      withoutAuth?: boolean
    } = {}
  ): Promise<any> {
    return this.fetchWithOptions({
      method: "PUT",
      ...options,
      url,
      form,
    })
  }

  async delete(
    url: string,
    params: Params = null,
    options: {
      withoutAuth?: boolean
    } = {}
  ): Promise<any> {
    return this.fetchWithOptions({
      method: "DELETE",
      ...options,
      url,
      params,
    })
  }

  async fetch(url: string, params: Params, init: RequestInit): Promise<any> {
    url = this.getUrl(url, params)

    try {
      const res = await fetch(url, init)
      const { status } = res
      if (status < 200 || status >= 400) {
        throw res
      }
      if (res.headers.get("Content-Type")?.includes("application/json")) {
        return res.json()
      } else {
        return res.text()
      }
    } catch (e) {
      throw await this.getErrorMsg(e)
    }
  }

  async fetchWithOptions(options: {
    method: string
    url: string
    params?: Params
    form?: Form
    withoutAuth?: boolean
  }): Promise<any> {
    const {
      method,
      url,
      params = null,
      form = null,
      withoutAuth = false,
    } = options

    const headers: HeadersInit = {}
    const init: RequestInit = {
      method,
      headers,
    }
    if (this.token != null && !withoutAuth) {
      headers["Authorization"] = `Bearer ${this.token}`
    }

    if (form instanceof FormData) {
      init.body = form
    } else if (form != null) {
      headers["Content-Type"] = "application/json"
      init.body = JSON.stringify(form)
    }
    return this.fetch(url, params, init)
  }

  getUrl(url: string, params: Params = null) {
    if (!startsWithProtocol(url)) {
      url = joinPath(this.baseURL, url)
    }

    if (params != null && Reflect.ownKeys(params).length !== 0) {
      url += url.includes("?") ? "&" : "?"
      url += qs.stringify(params)
    }

    return url
  }

  async getErrorMsg(e: any): Promise<string> {
    if (e instanceof Response) {
      // service error
      try {
        const body = await e.json()
        const msg = body?.msg ?? body?.message ?? i18n.SystemDown
        return msg
      } catch {
        return i18n.SystemDown
      }
    } else if (!navigator.onLine) {
      // offline
      return i18n.NetworkDown
    } else {
      // request abort
      return i18n.SystemDown
    }
  }

  setToken(token: string | null) {
    this.token = token
  }
}

export class Module {
  api: RequestProvider

  constructor(provider: RequestProvider) {
    this.api = provider
  }
}
