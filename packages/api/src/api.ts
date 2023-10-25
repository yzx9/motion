import { RequestProvider } from "./base"
import { SessionAPI } from "./session"
import { UserAPI } from "./user"

export class API {
  provider: RequestProvider

  session: SessionAPI
  user: UserAPI

  constructor(
    baseURL: string,
    apis: {
      session?: typeof SessionAPI
      user?: typeof UserAPI
    } = {}
  ) {
    const { session = SessionAPI, user = UserAPI } = apis

    this.provider = new RequestProvider(baseURL)

    this.session = new session(this.provider)
    this.user = new user(this.provider)
  }

  setToken(token: string | null) {
    this.provider.token = token
  }
}
