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
}
