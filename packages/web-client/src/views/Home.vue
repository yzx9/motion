<template>
  <div @mousewheel.prevent>
    <router-view v-slot="{ Component }">
      <!-- Component: VNodes, 传递给 <component>的is prop
        VNode的全称是VirtualNode,也就是虚拟节点;Vnode是一个类,
        可以生成不同类型的vnode实例,而不同类型的vnode表示不同类型的真实DOM元素。 -->

      <component :is="Component"></component>
    </router-view>
    <Bottomnav></Bottomnav>
  </div>
</template>

<script>
import { api } from "../api";
import Bottomnav from "../components/BottomNav.vue";
import "vant/lib/index.css";
export default {
  components: {
    Bottomnav,
  },
  mounted() {
    this.login();
  },
  methods: {
    login() {
      const token = api.session.signIn("17017376601", "123456qwer");
      token.then((res) => {
        console.log(res);
      });
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
