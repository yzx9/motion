<template>
  <div class="topNav">
    <div class="topLeft">
      <!-- <svg class="icon" aria-hidden="true">
        <use xlink:href="#icon-31liebiao"></use>
      </svg> -->
    </div>

    <div class="topContent">
      <div @click="changeid(0)">
        <span :class="currentid == 0 ? 'tab_active' : ''">直播</span>
      </div>
      <div @click="changeid(1)">
        <span :class="currentid == 1 ? 'tab_active' : ''">热点</span>
      </div>
      <div @click="changeid(2)">
        <span :class="currentid == 2 ? 'tab_active' : ''">同城</span>
      </div>
      <div @click="changeid(3)">
        <span :class="currentid == 3 ? 'tab_active' : ''">推荐</span>
      </div>
    </div>

    <div class="topRight">
      <svg class="icon" aria-hidden="true">
        <use xlink:href="#icon-search"></use>
      </svg>
    </div>
  </div>
</template>

<script>
import { api } from "../api";
export default {
  data() {
    return {
      currentid: 3,
    };
  },
  methods: {
    changeid(id) {
      this.currentid = id;
      this.$emit("childByValue", this.currentid);
      const token = api.session.signIn(213, 22312);
      token.then((res, rej) => console.log("214"));

      // TODO
    },
  },
};
</script>

<style scoped>
.topNav {
  z-index: 9999;
  position: fixed;
  background-color: #161823;
  color: #ffffff;
  width: 100%;
  height: 1rem;
  padding: 0.2rem;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.icon {
  fill: #ffffff;
  width: 0.5rem;
  height: 0.5rem;
}
.topContent {
  width: 65%;
  height: 100%;
  display: flex;
  justify-content: space-around;
  align-items: center;
  font-size: 0.36rem;
}
.tab_active {
  border-bottom: 2px solid #fff;
  font-weight: 700;
}
</style>
