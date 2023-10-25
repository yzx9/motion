<script setup lang="ts">
const props = withDefaults(
  defineProps<{
    type?: "primary" | "info"
  }>(),
  { type: "primary" }
)

const emit = defineEmits<{
  (event: "click"): void
}>()
</script>

<template>
  <div
    class="button w-fit min-w-[5.5rem] h-8 px-3 py-1 rounded-full text-base text-center"
    :class="{
      'text-white bg-theme': props.type === 'primary',
      'text-gray-500 border border-gray-300': props.type === 'info',
    }"
    @click="emit('click')"
  >
    <slot></slot>
  </div>
</template>
