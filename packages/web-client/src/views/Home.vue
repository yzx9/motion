<template>
  <div @mousewheel.prevent>
    <Topnav></Topnav>
    <div class="container_box">
      <Content></Content>
    </div>
    <Bottomnav></Bottomnav>
  </div>
</template>
<script setup lang="ts">
import Topnav from "../components/TopNav.vue";

import Bottomnav from "../components/BottomNav.vue";
import Content from "../components/Content.vue";
</script>
<style scoped>
.container_box {
  width: 100vw;
  height: 100vh;
  background: #000;
}
</style>
