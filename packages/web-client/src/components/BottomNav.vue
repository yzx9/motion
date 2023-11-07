<template>
  <div class="container_bottom">
    <div class="border_progress" :style="'width:' + videoProcess + '%'"></div>
    <div class="bottom_tab" @click="changeTab(0)">
      <span class="bottom_tab_span" :class="tabIndex == 0 ? 'tab_active' : ''"
        >首页</span
      >
    </div>
    <div class="bottom_tab align-middle content-center" @click="changeTab(1)">
      <svg
        class="icon"
        aria-hidden="true"
        :class="tabIndex == 1 ? 'tab_active' : ''"
      >
        <use xlink:href="#icon-vuesax-linear-add-square"></use>
      </svg>
    </div>
    <div class="bottom_tab" @click="changeTab(2)">
      <span class="bottom_tab_span" :class="tabIndex == 2 ? 'tab_active' : ''"
        >我的</span
      >
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from "vue";
import { useRouter } from "vue-router";
const router = useRouter();
let videoProcess = ref<number>(0);
let tabIndex = ref<number>(0);
onMounted(() => {
  if (tabIndex.value === 0) {
    router.push({
      path: "main",
    });
  }
  if (tabIndex.value === 1) {
    router.push({
      path: "uploadvideo",
    });
  }
  if (tabIndex.value === 2) {
    router.push({
      path: "mine",
    });
  }
});
const changeTab = (index: number) => {
  tabIndex.value = index;
  if (tabIndex.value === 0) {
    router.push({
      path: "main",
    });
  }
  if (tabIndex.value === 1) {
    router.push({
      path: "uploadvideo",
    });
  }
  if (tabIndex.value === 2) {
    router.push({
      path: "mine",
    });
  }
};
</script>

<style scoped>
.container_bottom {
  position: fixed;
  bottom: 0;
  width: 100%;
  background: rgba(0, 0, 0, 0.85);
  height: 0.96rem;
  /*border-top: 1px solid rgba(255, 255, 255, 0.7);*/
  /* max-width: 550px; */
  & .border_progress {
    width: 0;
    height: 0.02rem;
    background: #bababa;
    position: absolute;
    top: 0;
    transition: 0.1s all;
  }

  & .bottom_tab {
    width: 33%;
    text-align: center;
    float: left;
    color: #c9c9c9;
    padding: 0.2rem 0;
    & .bottom_tab_icon {
      width: 0.88rem;
    }
    & .icon {
      fill: #c9c9c9;
      width: 0.5rem;
      height: 0.5rem;
      margin: 0 auto;
    }
    & .bottom_tab_span {
      line-height: 0.56rem;
      border-bottom: 2px solid transparent;
      font-weight: 600;
      font-size: 0.32rem;
      display: inline-block;
      padding: 0 0.12rem;
      transition: 1s all;
    }
  }
}
.tab_active {
  color: #fff;
  fill: #fff !important;
}
</style>
