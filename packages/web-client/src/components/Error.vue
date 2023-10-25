<script setup lang="ts">
import EmptyPage from "@/assets/imgs/illustrations/empty-page.svg"
import { useRouter } from "vue-router"
import MyButton from "./Button.vue"

const props = withDefaults(
  defineProps<{
    info: string
    image: string
    buttonTitle?: string
  }>(),
  {
    image: EmptyPage,
  }
)

const emit = defineEmits<{
  (event: "click"): void
  (event: "click:minor"): void
}>()

const router = useRouter()

function handleClickMajor() {
  emit("click")
}

function handleClickMinor() {
  router.push("/")
}
</script>

<template>
  <section class="flex flex-col items-center gap-4">
    <img :src="props.image" />

    <h1 class="text-gray-700">{{ props.info }}</h1>

    <MyButton
      v-if="props.buttonTitle"
      class="w-[70vw]"
      type="primary"
      @click="handleClickMajor"
    >
      {{ props.buttonTitle }}
    </MyButton>

    <MyButton
      class="w-[70vw]"
      :type="props.buttonTitle ? 'info' : 'primary'"
      @click="handleClickMinor"
    >
      返回首页
    </MyButton>
  </section>
</template>
