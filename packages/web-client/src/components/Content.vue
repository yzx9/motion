<template>
  <div class="van_swipe">
    <van-swipe
      :show-indicators="false"
      @change="onChange"
      vertical
      :loop="false"
    >
      <van-swipe-item
        v-for="(item, index) in videoList"
        :key="index"
        class="product_swiper"
      >
        <div class="video_container">
          <!--video属性
                    webkit-playsinline ios 小窗播放，使视频不脱离文本流，安卓则无效
                    微信内置x5内核，
                    x5-video-player-type="h5-page" 启用H5播放器,是wechat安卓版特性，添加此属性，微信浏览器会自动将视频置为全屏
                    x5-video-player-fullscreen="true" 全屏设置，设置为 true 是防止横屏
                    x5-video-orientation 控制横竖屏 landscape 横屏，portrain竖屏； 默认portrain
                    poster：封面video_container
                    src：播放地址
                    -->
          <video
            class="video_box"
            loop
            webkit-playsinline="true"
            x5-video-player-type="h5-page"
            x5-video-player-fullscreen="true"
            playsinline
            preload="auto"
            :poster="item.cover"
            :src="item.url"
            :playOrPause="playOrPause"
            @click="pauseVideo"
            @ended="onPlayerEnded($event)"
          ></video>
          <!-- 封面 -->
          <img
            v-show="isVideoShow"
            class="play"
            @click="playvideo"
            :src="item.cover"
          />
          <!-- 播放暂停按钮 -->
          <img
            v-show="iconPlayShow"
            class="icon_play"
            @click="playvideo"
            src="http://npjy.oss-cn-beijing.aliyuncs.com/images/file-1575340653940esdHR.png"
          />
        </div>
        <!-- 右侧头像、点赞、评论、分享功能 -->
        <div class="tools_right">
          <div class="tools_r_li" style="margin: 0.4rem 0.2rem">
            <img :src="item.tag_image" class="tag_image" />
            <!-- <i
              class="iconfont icon-jiahao tag_add"
              v-show="!item.tagFollow"
              @click="checkSubscribe(item, index)"
            ></i>
            <i
              class="iconfont icon-duihao tag_dui"
              v-show="item.tagFollow"
              :class="item.tagFollow ? 'tag_dui_active' : ''"
            ></i> -->
            <svg
              class="iconfont tag_add"
              aria-hidden="true"
              v-show="!item.tagFollow"
              @click="checkSubscribe(item, index)"
            >
              <use xlink:href="#icon-jiahao-copy"></use>
            </svg>
            <svg
              class="iconfont tag_dui"
              aria-hidden="true"
              v-show="item.tagFollow"
              :class="item.tagFollow ? 'tag_dui_active' : ''"
              @click="cancelSubscribe(item, index)"
            >
              <use xlink:href="#icon-duihao"></use>
            </svg>
          </div>
          <div class="tools_r_li" @click="changeLike(item, index)">
            <!-- <i
              class="iconfont icon-shoucang icon_right"
              :class="item.fabulous ? 'fabulous_active' : ''"
            ></i> -->
            <svg
              class="iconfont icon_right"
              v-show="!item.fabulous"
              aria-hidden="true"
              :class="item.fabulous ? 'fabulous_active' : ''"
            >
              <use xlink:href="#icon-icon-test"></use>
            </svg>
            <svg
              class="iconfont icon_right"
              v-show="item.fabulous"
              aria-hidden="true"
              :class="item.fabulous ? 'fabulous_active' : ''"
            >
              <use xlink:href="#icon-aixin3"></use>
            </svg>
            <div class="tools_r_num">52.1w</div>
          </div>
          <div class="tools_r_li" @click="changeComment">
            <!-- <i class="iconfont icon-liuyan icon_right icon_right_change"></i> -->
            <svg
              class="iconfont icon_right icon_right_change"
              aria-hidden="true"
            >
              <use xlink:href="#icon-liuyan"></use>
            </svg>
            <div class="tools_r_num">12.5w</div>
          </div>
          <div class="tools_r_li" @click="changeShare">
            <!-- <i class="iconfont icon-iconfontforward icon_right"></i> -->
            <svg class="iconfont icon_right" aria-hidden="true">
              <use xlink:href="#icon-zhuanfa"></use>
            </svg>
            <div class="tools_r_num">22.2w</div>
          </div>
        </div>
        <!-- 底部作品描述 -->
        <div class="production_box">
          <div class="production_name">@{{ item.author }}</div>
          <div class="production_des">
            {{ item.des }}
          </div>
        </div>
      </van-swipe-item>
    </van-swipe>
    <!-- 评论弹窗 -->
    <van-popup
      v-model:show="commentPop"
      closeable
      :overlay="true"
      class="comment_container"
      position="bottom"
    >
      <div class="comment_box">
        <div class="comment_top">
          12.5w条评论
          <i class="iconfont icon-guanbi1 guanbi3" @click="closeComment"></i>
        </div>
        <ul class="comment_ul">
          <div v-if="videoComment.length != 0">
            <transition-group appear>
              <li
                class="comment_li"
                v-for="(item, index) in videoComment"
                :key="index"
                @click="replayUser(item, index, -1)"
              >
                <div class="comment_author_left">
                  <img :src="item.avatar" />
                </div>
                <div class="comment_author_right">
                  <div class="comment_author_top">
                    <div class="comment_author_name">@{{ item.nickname }}</div>
                    <div
                      class="icon-shoucang1_box"
                      @click.stop="commentLove(item, index, -1)"
                    >
                      <div
                        class="icon_right_change"
                        :class="item.love_comment ? 'love_active' : ''"
                      >
                        <i class="iconfont icon-shoucang icon-shoucang1"></i>
                      </div>
                    </div>
                    <div class="shoucang1_num">{{ item.love_count }}</div>
                  </div>
                  <div class="comment_author_text">
                    {{ item.comment_content
                    }}<span>{{ item.create_time }}</span>
                  </div>
                </div>
                <div class="clear"></div>
                <div class="comment_replay_box">
                  <transition-group appear>
                    <div
                      class="comment_replay_li"
                      v-for="(item2, index2) in item.child_comment"
                      :key="index2"
                      @click.stop="replayUser(item2, index, index2)"
                    >
                      <div class="comment_replay_left">
                        <img :src="item2.avatar" />
                      </div>
                      <div class="comment_replay_right">
                        <div class="comment_replay_top">
                          <div class="comment_replay_name">
                            @{{ item2.nickname }}
                          </div>
                          <div
                            class="icon-shoucang1_box"
                            @click.stop="commentLove(item2, index, index2)"
                          >
                            <div
                              class="icon_right_change"
                              :class="item2.love_comment ? 'love_active' : ''"
                            >
                              <!-- <i
                                class="iconfont icon-shoucang icon-shoucang1"
                              ></i> -->
                              <svg
                                style="transform: scale(0.15)"
                                aria-hidden="true"
                              >
                                <use xlink:href="#icon-aixin3"></use>
                              </svg>
                            </div>
                          </div>
                        </div>
                        <div class="shoucang1_num">
                          {{ item2.love_count }}
                        </div>
                        <div class="comment_replay_text">
                          <span
                            v-if="
                              item.user_id != item2.be_commented_user_id &&
                              item.user_id != item2.user_id
                            "
                            >回复 {{ item2.be_commented_nickname }}：</span
                          >
                          {{ item2.comment_content }}
                          <span>{{ item2.create_time }}</span>
                        </div>
                      </div>
                      <div class="clear"></div>
                    </div>
                  </transition-group>
                </div>
              </li>
            </transition-group>
          </div>
          <div class="no_message" v-if="videoComment.length == 0">
            <i class="iconfont iconfont_style icon-zanwupinglun"></i>
            <div class="no_message_tips">暂无评论</div>
          </div>
        </ul>
      </div>
    </van-popup>
    <!-- 评论输入 -->
    <div class="comment_input_box_hover"></div>
    <div class="comment_input_box" v-show="commentPop">
      <!--<form action="#" class="comment_form">-->
      <input
        :placeholder="commentPlaceholder"
        class="comment_input"
        v-model="comment_text"
        ref="content"
        @keyup.enter="checkComment"
      />
      <!--</form>-->
      <div class="comment_input_right" @click="checkComment">
        <i
          class="iconfont icon-fasong comment_i"
          :class="canSend ? 'comment_i_active' : ''"
        ></i>
      </div>
    </div>
  </div>
</template>

<script>
import { Swipe, SwipeItem, Popup } from "vant";
let videoProcessInterval;

export default {
  components: {
    [Swipe.name]: Swipe,
    [SwipeItem.name]: SwipeItem,
    [Popup.name]: Popup,
  },
  data() {
    let u = navigator.userAgent;
    return {
      current: 0,
      videoList: [
        {
          url: "http://s3jbnm16b.hd-bkt.clouddn.com/videos/The%20first%2020%20hours%20--%20how%20to%20learn%20anything%20-%20Josh%20Kaufman%20-%20TEDxCSU.mp4", //视频源
          cover: "http://oss.jishiyoo.com/images/file-1575341210559.png", //封面
          tag_image:
            "http://npjy.oss-cn-beijing.aliyuncs.com/images/file-1575449277018pF3XL.jpg", //作者头像
          fabulous: false, //是否赞过
          tagFollow: false, //是否关注过该作者
          author_id: 1, //作者ID
          author: "superKM1",
          des: "上海加油",
        },
        {
          url: "http://video.jishiyoo.com/1eedc49bba7b4eaebe000e3721149807/d5ab221b92c74af8976bd3c1473bfbe2-4518fe288016ee98c8783733da0e2da4-ld.mp4",
          cover: "http://oss.jishiyoo.com/images/file-1575343195934.jpg",
          tag_image:
            "http://npjy.oss-cn-beijing.aliyuncs.com/images/file-1575449298299M3V50.jpg",
          fabulous: true, //是否赞过
          tagFollow: false, //是否关注过该作者
          author_id: 2, //作者ID
          author: "superKM2",
          des: "北京加油",
        },
        {
          url: "http://video.jishiyoo.com/161b9562c780479c95bbdec1a9fbebcc/8d63913b46634b069e13188b03073c09-d25c062412ee3c4a0758b1c48fc8c642-ld.mp4",
          cover: "http://oss.jishiyoo.com/images/file-1575343262684.jpg",
          tag_image:
            "http://npjy.oss-cn-beijing.aliyuncs.com/images/file-1575449277018pF3XL.jpg",
          fabulous: false, //是否赞过
          tagFollow: false, //是否关注过该作者
          author_id: 1, //作者ID
          author: "superKM3",
          des: "武汉加油",
        },
        {
          url: "http://video.jishiyoo.com/549ed372c9d14b029bfb0512ba879055/8e2dc540573d496cb0942273c4a4c78c-15844fe70971f715c01d57c0c6595f45-ld.mp4",
          cover: "http://oss.jishiyoo.com/images/file-1575343508574.jpg",
          tag_image:
            "http://npjy.oss-cn-beijing.aliyuncs.com/images/file-1575449277018pF3XL.jpg",
          fabulous: false, //是否赞过
          tagFollow: false, //是否关注过该作者
          author_id: 1, //作者ID
          author: "superKM4",
          des: "广州加油",
        },
      ],
      isVideoShow: true,
      playOrPause: true,
      video: null,
      iconPlayShow: true,
      isAndroid: u.indexOf("Android") > -1 || u.indexOf("Adr") > -1, // android终端
      isiOS: !!u.match(/\(i[^;]+;( U;)? CPU.+Mac OS X/), // ios终端
      showShareBox: false, //展示分享框
      commentPop: false, //是否展示评论弹框
      // 评论相关
      comment_text: "", //评论输入内容
      canSend: false, //是否可以发送
      videoComment: [],
      commentPlaceholder: "留下你精彩的评论吧", //评论Placeholder
      replayUserData: "",
      to_comment_id: "",
      videoProcess: 0, //视频播放进度
    };
  },
  watch: {
    //监听输入变化
    comment_text(newV, oldV) {
      newV == "" ? (this.canSend = false) : (this.canSend = true);
    },
  },
  mounted() {
    //获取到视频资源后默认自动播放
    setTimeout(() => {
      this.playvideo();
    }, 500);
  },
  methods: {
    //获取评论
    getComment() {
      //setTimeout模拟Ajax请求
      setTimeout(() => {
        let data = [
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
        ]; //获取评论数据
        this.videoComment = [...this.videoComment, ...data];
        let to_comment_id = this.to_comment_id;
        if (to_comment_id != 0) {
          //从评论列表进来回复
          this.getCommentDetail(to_comment_id);
        }
      }, 500);
    },
    //获取单个评论
    getCommentDetail(to_comment_id) {
      let obj = {
        action: "show_comment_info",
        comment_id: to_comment_id,
      };
      setTimeout(() => {
        let index;
        let item = res.data;
        this.videoComment.map((v, i) => {
          v.child_comment.map((v2, i2) => {
            if (v2.comment_id == to_comment_id) {
              index = i;
            }
          });
        });
        setTimeout(() => {
          //对评论重新排序，存在to_comment_id的时候排至第一个
          let firstComment = this.videoComment.splice(index, 1);
          this.videoComment.unshift(firstComment[0]);
          //重设回复人
          item.index = 0;
          item.index2 = 0;
          this.replayUserData = item;
          this.commentPlaceholder = `回复 @${item.nickname}：`;
          this.$refs.content.focus();
        }, 10);
      }, 500);
    },
    //检测评论内容
    checkComment() {
      if (this.comment_text == "") {
        Toast("评论内容不能为空");
      } else {
        let comment_id = 0,
          p_id = "",
          p_user_id = "",
          content = this.comment_text;
        if (this.replayUserData != "") {
          comment_id = this.replayUserData.comment_id;
          p_id = this.replayUserData.p_id;
          p_user_id = this.replayUserData.user_id;
        }
        this.sendComment(comment_id, p_id, p_user_id, content);
      }
    },
    //发送评论
    sendComment(comment_id, p_id, p_user_id, content) {
      this.comment_text = "";
      this.isSending = true;
      setTimeout(() => {
        let newData = {
          comment_id: comment_id,
          p_id: p_id,
          comment_content: content,
          love_count: 0,
          create_time: "刚刚",
          user_id: p_user_id,
          nickname: "游客", //当前用户
          avatar: "https://profile.csdnimg.cn/B/1/E/3_ridingfish", //当前用户头像
          be_commented_user_id: this.replayUserData.user_id,
          be_commented_nickname: this.replayUserData.nickname,
        };
        if (this.replayUserData == "") {
          //回复作品
          this.videoComment.splice(0, 0, newData);
        } else {
          let index = this.replayUserData.index;
          let index2 = this.replayUserData.index2;
          if (this.replayUserData.index2 == -1) {
            //回复一级人
            this.videoComment[index].child_comment.splice(0, 0, newData);
          } else {
            //回复二级人
            this.videoComment[index].child_comment.splice(index2, 0, newData);
          }
        }
        this.replayUserData = "";
        this.isSending = false;
      }, 500);
    },
    //评论点赞
    commentLove(item, index, index2) {
      let comment_id = item.comment_id,
        user_id = item.user_id,
        love_comment = item.love_comment,
        love_flag = 0; //0:取消点赞；1：点赞
      if (love_comment) {
        //取消点赞
        love_flag = 0;
      } else {
        //添加点赞
        love_flag = 1;
      }
      //setTimeout模拟Ajax请求
      setTimeout(() => {
        if (index2 == -1) {
          //点赞一级评论
          this.videoComment[index].love_comment =
            !this.videoComment[index].love_comment;
          if (love_flag == 1) {
            this.videoComment[index].love_count++;
          } else {
            this.videoComment[index].love_count--;
          }
        } else {
          //点赞二级评论
          this.videoComment[index].child_comment[index2].love_comment =
            !this.videoComment[index].child_comment[index2].love_comment;
          if (love_flag == 1) {
            this.videoComment[index].child_comment[index2].love_count++;
          } else {
            this.videoComment[index].child_comment[index2].love_count--;
          }
        }
      }, 500);
    },
    //点击回复
    replayUser(item, index, index2) {
      item.index = index;
      item.index2 = index2;
      this.replayUserData = item;
      this.commentPlaceholder = `回复 @${item.nickname}：`;
      this.$refs.content.focus();
    },
    //弹出评论窗口
    changeComment() {
      this.commentPop = true;
      console.log("1111", this.commentPop);
      this.videoComment = [];
      this.getComment();
    },
    //关闭评论弹窗
    closeComment() {
      this.commentPop = false;
      this.commentPlaceholder = "留下你精彩的评论吧";
      this.replayUserData = "";
    },
    //关注该作者
    checkSubscribe(item, index) {
      this.videoList.map((v) => {
        v.author_id == item.author_id ? (v.tagFollow = true) : "";
      });
    },
    // 取消关注
    cancelSubscribe(item, index) {
      this.videoList.map((v) => {
        v.author_id == item.author_id ? (v.tagFollow = false) : "";
      });
    },
    //改变菜单

    //改变收藏状态
    changeLike(item, index) {
      this.videoList[index].fabulous = !this.videoList[index].fabulous;
    },
    //展示分享弹窗
    changeShare() {
      this.showShareBox = true;
    },
    //取消分享
    cancelShare() {
      this.showShareBox = false;
    },
    //滑动改变播放的视频
    onChange(index) {
      //改变的时候 暂停当前播放的视频
      clearInterval(videoProcessInterval);
      this.videoProcess = 0;
      let video = document.querySelectorAll("video")[this.current];
      video.pause();
      this.playOrPause = false;
      this.showShareBox = false;
      this.current = index;
      //非ios切换直接自动播放下一个
      if (!this.isiOS) {
        this.isVideoShow = false;
        setTimeout(() => {
          this.pauseVideo();
        }, 100);
      } else {
        //ios官方禁止video自动播放，未找到合适的方法，如果您发现了，麻烦告诉我一下谢谢啦
        this.playOrPause = true;
        this.iconPlayShow = true;
      }
    },
    // 开始播放
    playvideo(event) {
      let video = document.querySelectorAll("video")[this.current];
      console.log("playvideo：" + this.current);
      this.isVideoShow = false;
      this.iconPlayShow = false;
      this.showShareBox = false;
      video.play();

      if (this.isiOS) {
        setTimeout(() => {
          //处理ios宽视频
          let documentW =
            document.documentElement.clientWidth || document.body.clientWidth;
          let docB = parseFloat(video.videoWidth / documentW);
          console.log("获取视频宽和屏幕比：" + docB);
          // 计算视频最适高度
          let realH = parseInt(video.videoHeight / docB);
          this.realH = realH + "px";
          console.log("视频最适高度：" + this.realH);
          this.$forceUpdate();
        }, 200);
      }

      videoProcessInterval = setInterval(() => {
        this.changeProcess(video);
      }, 100);
    },
    pauseVideo() {
      //暂停\播放
      try {
        let video = document.querySelectorAll("video")[this.current];
        if (this.playOrPause) {
          video.pause();
          this.iconPlayShow = true;
          clearInterval(videoProcessInterval);
        } else {
          // wx.ready(() => {
          //     // 在微信的ready中进行播放 不管成功配置与否 都会执行ready
          //     video.play();
          // })
          video.play();
          video.pause();
          setTimeout(() => {
            video.play();
            this.iconPlayShow = false;
            videoProcessInterval = setInterval(() => {
              this.changeProcess(video);
            }, 100);
          }, 100);
        }
        this.playOrPause = !this.playOrPause;
        this.showShareBox = false;
      } catch (e) {
        alert(e);
      }
    },
    //记录播放进度
    changeProcess() {
      let video = document.querySelectorAll("video")[this.current];
      let currentTime = video.currentTime.toFixed(1);
      let duration = video.duration.toFixed(1);
      this.videoProcess = parseInt((currentTime / duration).toFixed(2) * 100);
    },
    onPlayerEnded(player) {
      //视频结束
      this.isVideoShow = true;
      this.current += this.current;
    },
    //复制当前链接
    copyUrl() {
      let httpUrl = window.location.href;
      var oInput = document.createElement("input");
      oInput.value = httpUrl;
      document.body.appendChild(oInput);
      oInput.select(); // 选择对象
      document.execCommand("Copy"); // 执行浏览器复制命令
      oInput.className = "oInput";
      oInput.style.display = "none";
      alert("链接复制成功");
    },
  },
};
</script>
<style scoped>
.icon {
  width: 1rem;
  height: 0.5rem;
  vertical-align: -0.15em;
  fill: currentColor;
  overflow: hidden;
}

.video_container {
  position: relative;
}
.video-player {
  height: 100vh;
  width: 100vw;
}

.product_swiper {
  width: 100vw;
  height: 100vh;
}

.van_swipe {
  width: 100vw;
  height: 100vh;
  max-width: 550px;
  margin: 0 auto;
  position: relative;
}

.van-swipe {
  width: 100%;
  height: 100%;
}

.video_box {
  position: relative;
  top: 5rem;
  background-color: aqua;
  object-fit: fill !important;
  z-index: 999;
  width: 100%;
}

video {
  object-position: 0 0;
}

.icon_play {
  position: absolute;
  top: 6rem;
  right: 0;
  left: 0;
  bottom: auto;
  margin: auto;
  z-index: 999;
  height: 1.2rem;
  background: rgba(0, 0, 0, 0.5);
  border-radius: 50%;
}

.play,
.platStart {
  background-color: aqua;
  position: absolute;
  margin: auto;
  top: 0;
  left: 0;
  z-index: 999;
  width: 100%;
  height: 100%;
}

/*头像， 点赞，转发 */
.tag_image {
  width: 50px;
  height: 50px;
  border: 1px solid #fff;
  box-sizing: border-box;
  border-radius: 50%;

  position: relative;
  box-shadow: 0px 0px 10px #9d9d9d;
}

.tag_add {
  position: absolute;
  top: 0.85rem;
  left: 0;
  right: 0;
  border-radius: 67px;
  display: inline-block;
  width: 18px;
  height: 18px;
  line-height: 18px;
  margin: 0 auto;
  z-index: 1;
  font-size: 20px;
  color: #fff;
  background: #f44;
}

.tag_dui {
  position: absolute;
  top: 0.85rem;
  left: 0;
  right: 0;
  border-radius: 67px;
  display: inline-block;
  width: 20px;
  height: 20px;
  line-height: 20px;
  margin: 0 auto;
  z-index: 1;
  font-size: 11px;
  color: #f44;
  background: #fff;
}

.tag_dui_active {
  opacity: 0;
  animation: showFollow 1.4s ease-in-out 0s;
}

.tag_add_num {
  position: absolute;
  bottom: 0px;
  left: 0;
  right: 0;
}

@keyframes showFollow {
  0% {
    opacity: 1;
    color: #fff;
    background: #f44;
    transform: rotate(-180deg) scale(1);
  }
  35% {
    opacity: 1;
    color: #f44;
    background: #fff;
    transform: rotate(0deg) scale(1.2);
  }
  80% {
    opacity: 1;
    color: #f44;
    transform: scale(1.2);
  }
  100% {
    opacity: 0;
    display: none;
    color: #f44;
    transform: scale(0);
  }
}

.tools_right {
  z-index: 1001;
  position: relative;
  float: right;
  right: 0.2rem;
  top: 3rem;
}

.tools_r_li {
  margin: 0.3rem 0.2rem;
  color: #fff;
  font-size: 14px;
  position: relative;
}

.tools_r_li:last-child {
  margin-bottom: 0px;
}

.icon_right {
  margin-bottom: 0.2rem;
  width: 1rem;
  height: 0.8rem;
  font-size: 42px;
  display: block;
  text-shadow: 0px 0px 10px #9d9d9d;
  /*transition: .5s;*/
}
.fabulous_active {
  color: #f44;
  animation: showHeart 0.5s ease-in-out 0s;
}
@keyframes showHeart {
  0% {
    color: #f44;
    transform: scale(1);
  }

  25% {
    color: #fff;
    transform: scale(0);
  }

  80% {
    color: #f44;
    transform: scale(1.2);
  }

  100% {
    color: #f44;
    transform: scale(1);
  }
}
.tools_r_num {
  position: relative;
  right: -0.15rem;
}

.production_box {
  z-index: 1001;
  position: relative;
  /* right: 16px; */
  top: 9.5rem;
  text-align: left;
  padding: 0 0.3rem 0.3rem 0.3rem;
  color: #fff;
  width: 100%;
  left: 0;
  box-sizing: border-box;
  background: linear-gradient(
    to top,
    rgba(255, 255, 255, 0.2),
    rgba(0, 0, 0, 0)
  );
}

.production_name {
  font-weight: 700;
  font-size: 18px;
  margin-bottom: 10px;
}

.production_des {
  overflow: hidden;
  text-overflow: ellipsis;
  display: -webkit-box;
  -webkit-line-clamp: 3;
  -webkit-box-orient: vertical;
  min-height: 62px;
  font-size: 13px;
}
/* 评论的样式 */
::-webkit-input-placeholder {
  color: rgba(0, 0, 0, 0.2);
}

:-moz-placeholder {
  color: rgba(0, 0, 0, 0.2);
}

::-moz-placeholder {
  color: rgba(0, 0, 0, 0.2);
}

:-ms-input-placeholder {
  color: rgba(0, 0, 0, 0.2);
}

.comment_container {
  height: 50%;
  background-color: #fff;
  bottom: 0rem;
  width: 100%;
  border-top-left-radius: 10px;
  border-top-right-radius: 10px;
  position: absolute;
}

.comment_box {
  padding: 0 15px 52px 15px;
}

.comment_top {
  text-align: center;
  font-size: 12px;
  color: #000;
  line-height: 40px;
}

.guanbi3 {
  float: right;
  font-size: 12px;
  padding: 0 10px;
  position: absolute;
  right: 6px;
}

.comment_li {
  margin-bottom: 20px;
  font-size: 14px;
  text-align: left;
}

.comment_author_left {
  float: left;
}

.comment_author_left img {
  width: 35px;
  height: 35px;
  border-radius: 50%;
}

.comment_author_right {
  margin-left: 46px;
  padding-top: 4px;
}

.comment_author_top {
  position: relative;
}

.comment_author_name {
  margin-bottom: 6px;
  color: #777;
}

.icon-shoucang1_box {
  position: absolute;
  width: 0.4rem;
  height: 0.4rem;
  left: 1.9rem;
  top: -1.3rem;
  text-align: center;
  color: #777;
}

.comment_author_text {
  color: #555;
  margin-bottom: 10px;
  padding-right: 35px;
}

.comment_replay_box {
  padding-left: 46px;
  box-sizing: border-box;
}

.comment_replay_li {
  margin-bottom: 10px;
}

.comment_replay_left {
  float: left;
}

.comment_replay_left img {
  width: 25px;
  height: 25px;
  border-radius: 50%;
}

.comment_replay_right {
  margin-left: 35px;
  padding-top: 2px;
}

.comment_replay_top {
  position: relative;
  margin-bottom: 6px;
}

.comment_replay_text {
  padding-right: 35px;
  margin-bottom: 10px;
  color: #555;
}

.comment_author_text span,
.comment_replay_text span {
  color: #999;
  font-size: 13px;
}

.shoucang1_num {
  text-align: center;
  width: 30px;
  font-size: 12px;
  /* right: -4px; */
  position: relative;
  float: right;
  bottom: 0.2rem;
  right: 0.3rem;
}

.comment_ul {
  height: 400px;
  overflow-y: auto;
}

.comment_input_box {
  position: fixed;
  bottom: 0;
  z-index: 2999;
  width: 100%;
  border-top: 1px solid #e8e8e8;
  background: #fff;
  padding: 10px 15px;
  box-sizing: border-box;
}

/*.comment_form {*/
/**/
/*}*/

.comment_input {
  border: none;
  resize: none;
  width: 80%;
  float: left;
  color: #555;
  caret-color: #f44;
  line-height: 0.44rem;
}

.comment_input_right {
  float: right;
}

.comment_i {
  font-size: 22px;
  color: #999;
  transition: 0.3s;
}

.comment_i_active {
  color: #f44;
}

.icon-zanwupinglun {
  font-size: 100px;
  color: #777;
}

.v-enter,
.v-leave-to {
  opacity: 0;
  transform: translateY(80px);
}

.v-enter-active,
.v-leave-active {
  transition: all 0.5s ease;
}

/*添加进场效果*/
.v-move {
  transition: all 1s ease;
}

.v-leave-active {
  position: absolute;
}

.list-complete-item {
  transition: all 1s;
  display: inline-block;
  margin-right: 10px;
}

.list-complete-enter,
.list-complete-leave-to {
  opacity: 0;
  transform: translateY(30px);
}

.list-complete-leave-active {
  position: absolute;
}

.love_active {
  color: #f44;
}
/* 评论的样式 */
</style>
