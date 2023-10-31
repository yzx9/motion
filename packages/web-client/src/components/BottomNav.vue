<template>
  <div class="container_bottom">
    <div class="border_progress" :style="'width:' + videoProcess + '%'"></div>
    <div
      class="bottom_tab"
      :class="tabIndex == 0 ? 'tab_active' : ''"
      @click="changeTab(0)"
    >
      <span class="bottom_tab_span">首页</span>
    </div>
    <div
      class="bottom_tab align-middle content-center"
      :class="tabIndex == 1 ? 'tab_active' : ''"
      @click="changeTab(1)"
    >
      <svg class="icon" aria-hidden="true">
        <use xlink:href="#icon-vuesax-linear-add-square"></use>
      </svg>
    </div>
    <div
      class="bottom_tab"
      :class="tabIndex == 2 ? 'tab_active' : ''"
      @click="changeTab(2)"
    >
      <span class="bottom_tab_span">我的</span>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from "vue";

let videoProcess = ref<number>(0);
let tabIndex = ref<number>(0);
const changeTab = (index: number) => {
  tabIndex.value = index;
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
      fill: #ffffff;
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
      transition: 0.5s all;
    }
  }
}
</style>
