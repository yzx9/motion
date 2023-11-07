<template>
  <div class="container">
    <div class="top">
      <div class="avatar">
        <img
          src="http://npjy.oss-cn-beijing.aliyuncs.com/images/file-1575449277018pF3XL.jpg"
        />
      </div>
      <div class="detail">
        <div class="name">张波123</div>
        <div class="des">成分复杂，无法简介</div>
      </div>
      <div class="wrap">
        <div class="fan">粉丝：<span class="number">222</span></div>
        <div class="follow">关注：<span class="number">102</span></div>
      </div>
    </div>
    <div class="personal">
      <div class="title">个人作品</div>
      <div class="video">
        <Grid :column-num="3" :gutter="5">
          <GridItem v-for="(item, index) in user_info.video_info" :key="index">
            <img
              src="https://fastly.jsdelivr.net/npm/@vant/assets/apple-1.jpeg"
            />
          </GridItem>
        </Grid>
      </div>
    </div>
  </div>
</template>
<script setup>
import { Grid, GridItem } from "vant";
import { ref, reactive } from "vue";
const user_info = reactive({
  video_info: [
    {
      a: "1",
    },
    {
      a: "2",
    },
    {
      a: "3",
    },
    {
      a: "4",
    },
    {
      a: "4",
    },
    {
      a: "4",
    },
    {
      a: "4",
    },
    {
      a: "4",
    },
    {
      a: "4",
    },
    {
      a: "4",
    },
  ],
});
</script>
<style scoped>
.container {
  width: 100vw;
  height: 100vh;
  /* background-color: #262626; */
  background: linear-gradient(to top, #575656, #262626);
  color: #fff;
}
.top {
  height: 40vh;
  display: flex;
  position: relative;
}
.avatar img {
  margin: 1rem 0.5rem;
  width: 1.5rem;
  height: 1.5rem;
  border-radius: 50%;
}
.detail {
  margin: 1rem 0;
}
.name {
  font-size: 0.8rem;
  font-weight: 700;
}
.des {
  position: absolute;
  margin-top: 0.6rem;
  font-size: 0.4rem;
  left: 0.4rem;
  width: 6.5rem;
  height: 2rem;
}
.wrap {
  position: absolute;
  display: flex;
  justify-content: space-around;
  align-items: center;
  left: 0.5rem;
  font-size: 0.5rem;
  margin-top: 5rem;
}
.fan {
  font-weight: 600;
  margin-right: 1rem;
}
.follow {
  font-weight: 600;
}
.number {
  font-weight: 500;
}
.personal {
  height: 60vh;
  border-top: solid #fff;
  border-radius: 10%;
}
.title {
  text-align: center;
  font-size: 0.7rem;
  font-weight: 650;
  border-bottom: solid #474646;
}
:deep(.van-grid) {
  height: 8rem;
  overflow: scroll;
}
.video {
  height: 100%;
  background-color: #474646 !important;
}
</style>
