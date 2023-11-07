import { Module } from "./base"
import { Tokens } from "./session"

export type User = {
  id: string
  nickname: string
  des: string
  avatarURL: string
  fan: number
  follow: number
  videos: { img: string }[]
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
  /**
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

  /**
   * current user
   */

  async getCurrent(): Promise<User> {
    // return await this.api.get("/users/current")

    return {
      id: "fake",
      nickname: "张波123",
      des: "成分复杂，无法简介",
      avatarURL:
        "http://npjy.oss-cn-beijing.aliyuncs.com/images/file-1575449277018pF3XL.jpg",
      fan: 222,
      follow: 102,
      videos: new Array(10).fill(0).map((_) => ({
        img: "https://fastly.jsdelivr.net/npm/@vant/assets/apple-1.jpeg",
      })),
    }
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
  nickname: "unauthed",
  des: "",
  avatarURL: "",
  fan: 0,
  follow: 0,
  videos: [],
}
