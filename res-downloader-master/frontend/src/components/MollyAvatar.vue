<script setup lang="ts">
import { computed } from "vue"
import { themeImageUrl } from "@/theme/presets"
import type { ThemePreset } from "@/theme/presets"
import { useIndexStore } from "@/stores"

const props = defineProps<{
  preset: ThemePreset
  size?: number
}>()

const containerSize = computed(() => props.size ?? 64)
const containerHeight = computed(() => containerSize.value)
const store = useIndexStore()

const imageUrl = computed(
  () => themeImageUrl(props.preset.image, store.baseUrl),
)

const style = computed(() => ({
  width: `${containerSize.value}px`,
  height: `${containerHeight.value}px`,
}))
</script>

<template>
  <span class="molly-avatar" :style="style">
    <img
      class="molly-avatar__img"
      :src="imageUrl"
      :alt="preset.name"
      draggable="false"
    />
  </span>
</template>

<style scoped>
.molly-avatar {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  overflow: visible;
  border-radius: 999px;
  flex-shrink: 0;
  position: relative;
  background:
    radial-gradient(circle at 48% 58%, color-mix(in srgb, var(--md-primary) 18%, transparent) 0 52%, transparent 54%),
    linear-gradient(145deg, color-mix(in srgb, var(--md-surface) 84%, var(--molly) 16%), var(--md-surface));
  box-shadow:
    inset 0 0 0 1px color-mix(in srgb, var(--md-outline) 58%, transparent),
    0 10px 24px color-mix(in srgb, var(--md-primary) 12%, transparent);
}

.molly-avatar__img {
  width: 100%;
  height: 100%;
  object-fit: contain;
  object-position: center;
  border-radius: inherit;
  opacity: 0.98;
  transition:
    transform 220ms cubic-bezier(0.34, 1.56, 0.64, 1),
    opacity 180ms ease;
  filter: drop-shadow(0 6px 10px color-mix(in srgb, var(--md-primary) 18%, transparent));
}

.molly-avatar:hover .molly-avatar__img {
  transform: scale(1.06);
  opacity: 1;
}
</style>
