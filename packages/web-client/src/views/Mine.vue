<script setup lang="ts">
import { api, UserDefault } from "@/api"
import { Grid, GridItem } from "vant"
import { ref } from "vue"

const user = ref(UserDefault)

updateData()

async function updateData() {
  user.value = await api.user.getCurrent()
}
</script>

<template>
  <div class="container w-screen h-screen text-white">
    <div class="flex relative h-[40vh]">
      <div>
        <img class="my-4 mx-2 w-6 h-6 rounded-[50%]" :src="user.avatarURL" />
      </div>

      <div class="my-4">
        <div class="text-[0.8rem] font-bold">{{ user.nickname }}</div>

        <div
          class="w-[6.5rem] h-8 absolute left-[0.4rem] mt-[0.6rem] text-[0.4rem]"
        >
          {{ user.des }}
        </div>
      </div>

      <div
        class="absolute left-2 flex justify-around items-center mt-20 text-[0.5rem]"
      >
        <div class="ml-4 font-semibold">
          粉丝：<span class="font-normal">{{ user.fan }}</span>
        </div>

        <div class="font-semibold">
          关注：<span class="font-normal">{{ user.follow }}</span>
        </div>
      </div>
    </div>

    <div class="personal h-[60vh] rounded-[10%]">
      <div class="title text-[0.7rem] text-center font-bold">个人作品</div>

      <div class="video h-full">
        <Grid :column-num="3" :gutter="5">
          <GridItem v-for="(item, index) in user.videos" :key="index">
            <img :src="item.img" />
          </GridItem>
        </Grid>
      </div>
    </div>
  </div>
</template>

<style scoped>
.container {
  /* background-color: #262626; */
  background: linear-gradient(to top, #575656, #262626);
}

.personal {
  border-top: solid #fff;
}

.title {
  border-bottom: solid #474646;
}

:deep(.van-grid) {
  height: 8rem;
  overflow: scroll;
}

.video {
  background-color: #474646 !important;
}
</style>
