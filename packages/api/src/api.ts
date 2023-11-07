import { RequestProvider } from "./base"
import { SessionAPI } from "./session"
import { UserAPI } from "./user"
import { VideoAPI } from "./video"

export class API {
  provider: RequestProvider

  session: SessionAPI
  user: UserAPI
  video: VideoAPI

  constructor(
    baseURL: string,
    apis: {
      session?: typeof SessionAPI
      user?: typeof UserAPI
      video?: typeof VideoAPI
    } = {}
  ) {
    const { session = SessionAPI, user = UserAPI, video = VideoAPI } = apis

    this.provider = new RequestProvider(baseURL)

    this.session = new session(this.provider)
    this.user = new user(this.provider)
    this.video = new video(this.provider)
  }

  setToken(token: string | null) {
    this.provider.token = token
  }
}
