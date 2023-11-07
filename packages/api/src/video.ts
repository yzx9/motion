import { Module } from "./base"

export type Video = {
  url: string //视频源
  cover: string //封面
  tag_image: string //作者头像
  fabulous: boolean //是否赞过
  tagFollow: boolean //是否关注过该作者
  author_id: string //作者ID
  author: string
  des: string
}

export type VideoComment = {
  comment_id: number
  p_id: number
  comment_content: string
  love_count: number
  create_time: string
  user_id: number
  nickname: string
  avatar: string
  be_commented_user_id: number
  be_commented_nickname: string
  be_commented_avatar: string
  love_comment: boolean
  child_comment: VideoComment[]
}

export class VideoAPI extends Module {
  async list(): Promise<Video[]> {
    return [
      {
        url: "http://s3jbnm16b.hd-bkt.clouddn.com/videos/The%20first%2020%20hours%20--%20how%20to%20learn%20anything%20-%20Josh%20Kaufman%20-%20TEDxCSU.mp4", //视频源
        cover: "http://oss.jishiyoo.com/images/file-1575341210559.png", //封面
        tag_image:
          "http://npjy.oss-cn-beijing.aliyuncs.com/images/file-1575449277018pF3XL.jpg", //作者头像
        fabulous: false, //是否赞过
        tagFollow: false, //是否关注过该作者
        author_id: "1", //作者ID
        author: "superKM1",
        des: "上海加油",
      },
      {
        url: "http://s3jbnm16b.hd-bkt.clouddn.com/videos/%E6%8A%96%E9%9F%B3-%E8%AE%B0%E5%BD%95%E7%BE%8E%E5%A5%BD%E7%94%9F%E6%B4%BB.mp4",
        cover: "http://oss.jishiyoo.com/images/file-1575343195934.jpg",
        tag_image:
          "http://npjy.oss-cn-beijing.aliyuncs.com/images/file-1575449298299M3V50.jpg",
        fabulous: true, //是否赞过
        tagFollow: false, //是否关注过该作者
        author_id: "1", //作者ID
        author: "superKM2",
        des: "北京加油",
      },
      {
        url: "http://s3jbnm16b.hd-bkt.clouddn.com/videos/%E6%8A%96%E9%9F%B3-%E8%AE%B0%E5%BD%95%E7%BE%8E%E5%A5%BD%E7%94%9F%E6%B4%BB_2.mp4",
        cover: "http://oss.jishiyoo.com/images/file-1575343262684.jpg",
        tag_image:
          "http://npjy.oss-cn-beijing.aliyuncs.com/images/file-1575449277018pF3XL.jpg",
        fabulous: false, //是否赞过
        tagFollow: false, //是否关注过该作者
        author_id: "1", //作者ID
        author: "superKM3",
        des: "武汉加油",
      },
      {
        url: "http://s3jbnm16b.hd-bkt.clouddn.com/videos/%E6%8A%96%E9%9F%B3-%E8%AE%B0%E5%BD%95%E7%BE%8E%E5%A5%BD%E7%94%9F%E6%B4%BB_3.mp4",
        cover: "http://oss.jishiyoo.com/images/file-1575343508574.jpg",
        tag_image:
          "http://npjy.oss-cn-beijing.aliyuncs.com/images/file-1575449277018pF3XL.jpg",
        fabulous: false, //是否赞过
        tagFollow: false, //是否关注过该作者
        author_id: "1", //作者ID
        author: "superKM4",
        des: "广州加油",
      },
    ]
  }

  async getComment(id: string): Promise<VideoComment[]> {
    return new Promise<VideoComment[]>((res) => {
      setTimeout(() => {
        res([
          {
            comment_id: 297,
            p_id: 0,
            comment_content: "你好，我叫蓝湛",
            love_count: 0,
            create_time: "1月前",
            user_id: 78634,
            nickname: "蓝忘机\uD83C\uDF1F",
            avatar:
              "http://npjy.oss-cn-beijing.aliyuncs.com/images/file-1575449277018pF3XL.jpg",
            be_commented_user_id: 0,
            be_commented_nickname: "",
            be_commented_avatar: "",
            child_comment: [
              {
                comment_id: 298,
                p_id: 296,
                comment_content: "蓝二公子，今天天气不错",
                love_count: 1,
                create_time: "7天前",
                user_id: 55163,
                nickname: "魏婴",
                avatar:
                  "http://thirdwx.qlogo.cn/mmopen/vi_32/Q0j4TwGTfTKPJb1k8zia02PjVibdaDJ83JIDGm0hIsY34kAlXyZMT6FMBibdw6rhdPPjpxtp6d8B75x5Kpicxp2gqw/132",
                be_commented_user_id: 78480,
                be_commented_nickname: "chenchen",
                be_commented_avatar:
                  "http://thirdwx.qlogo.cn/mmopen/vi_32/icxHc0Ym1p4hQAFAUnjpxDPMkEUyojnibBj9wUSS2OmibiazdBAicSLpoicricVYP6QG6XicjTzQPx9koMPqcGfhTOy5qA/132",
                love_comment: true,
                child_comment: [],
              },
            ],
            love_comment: false,
          },
          {
            comment_id: 281,
            p_id: 0,
            comment_content: "楼主好帅，我要嫁给你！！",
            love_count: 0,
            create_time: "1月前",
            user_id: 74164,
            nickname: "冰雪奇缘2",
            avatar:
              "http://npjy.oss-cn-beijing.aliyuncs.com/images/file-1575449298299M3V50.jpg",
            be_commented_user_id: 0,
            be_commented_nickname: "",
            be_commented_avatar: "",
            child_comment: [],
            love_comment: false,
          },
        ])
      }, 500)
    })
  }
}
