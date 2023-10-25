<script setup lang="ts">
import Illustration404 from "@/assets/imgs/illustrations/404.svg"
import { Error } from "@/components"
import { useRouter } from "vue-router"

const router = useRouter()
function handleClick() {
  router.go(-1)
}
</script>

<template>
  <Error
    info="页面不见了"
    :image="Illustration404"
    button-title="返回上级"
    @click="handleClick"
  />
</template>
