<template>
  <div>
    <Topnav v-on:childByValue="childByValue"></Topnav>
    <div class="container_box">
      <Content :videoList="videoList"></Content>
    </div>
  </div>
</template>
<script>
import { api } from "../api";
import Topnav from "../components/TopNav.vue";

import Content from "../components/Content.vue";
export default {
  components: {
    Topnav,
    Content,
  },
  data() {
    return {
      id: 3,
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
          url: "http://s3jbnm16b.hd-bkt.clouddn.com/videos/%E6%8A%96%E9%9F%B3-%E8%AE%B0%E5%BD%95%E7%BE%8E%E5%A5%BD%E7%94%9F%E6%B4%BB.mp4",
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
          url: "http://s3jbnm16b.hd-bkt.clouddn.com/videos/%E6%8A%96%E9%9F%B3-%E8%AE%B0%E5%BD%95%E7%BE%8E%E5%A5%BD%E7%94%9F%E6%B4%BB_2.mp4",
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
          url: "http://s3jbnm16b.hd-bkt.clouddn.com/videos/%E6%8A%96%E9%9F%B3-%E8%AE%B0%E5%BD%95%E7%BE%8E%E5%A5%BD%E7%94%9F%E6%B4%BB_3.mp4",
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
    };
  },

  mounted() {
    this.getVideo();
  },
  methods: {
    getVideo() {
      const token = api.session.getVideo();
      token.then((res) => console.log(res.data));
    },
    childByValue(childValue) {
      // childValue就是子组件传过来的值
      this.id = childValue;
    },
  },
};
</script>
<style scoped>
.container_box {
  width: 100vw;
  height: 100vh;
  background: #000;
}
</style>
