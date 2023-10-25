import { Module } from "./base"
import { Tokens } from "./session"

export type User = {
  id: string
  name: string
  avatar: string
  avatarUrl: string
  mobile: string
}

export type UserSMS = {
  sent: boolean
  exists: boolean
  wait: number
  code?: string // DEBUG mode only
}

export type FormUser = {
  name?: string
  mobile?: string
}

export class UserAPI extends Module {
  /*
   * user
   */

  async create(
    mobile: string,
    code: string
  ): Promise<User & { session: Tokens }> {
    return this.api.post(`/users`, { mobile, code })
  }

  async update(user: User): Promise<User> {
    return this.api.put(`/users/${user.id}`, user)
  }

  /*
   * current user
   */

  async getCurrent(): Promise<User> {
    return await this.api.get("/users/current")
  }
  async sendVerificationCode(
    mobile: string,
    exists: boolean = true // send even not exists if false
  ): Promise<UserSMS> {
    return this.api.post(
      "/users/sms",
      { mobile, exists },
      { withoutAuth: true }
    )
  }
}

export const UserDefault: User = {
  id: "",
  name: "",
  avatar: "",
  avatarUrl: "",
  mobile: "",
}
