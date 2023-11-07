import { Module } from "./base"

export type Tokens = {
  accessToken: string
  tokenKind: string
  expiresIn: number
  refreshToken: string
}

export class SessionAPI extends Module {
  async signInByRefreshToken(refreshToken: string): Promise<Tokens> {
    return await this.api.post(
      "/sessions/refresh-token",
      { refreshToken },
      { withoutAuth: true }
    )
  }

  async signInBySMS(mobile: string, code: string): Promise<Tokens> {
    return await this.api.post(
      "/sessions/sms",
      { mobile, code },
      { withoutAuth: true }
    )
  }

  async signOut(): Promise<void> {
    await this.api.delete("/sessions")
  }
  async signIn(mobile: string, password: string): Promise<Tokens> {
    return await this.api.post(
      "/v1/user/vail/login",
      { mobile, password },
      { withoutAuth: true }
    )
  }
  async getVideo() {
    return await this.api.get(
      "/v1/video/recommend/videos"
    )
  }
}

