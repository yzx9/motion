<template>
  <div>
    <Topnav></Topnav>
    <div class="container_box">
      <Content></Content>
    </div>
  </div>
</template>
<script setup lang="ts">
import Topnav from "../components/TopNav.vue";

import Content from "../components/Content.vue";
</script>
<style scoped>
.container_box {
  width: 100vw;
  height: 100vh;
  background: #000;
}
</style>
