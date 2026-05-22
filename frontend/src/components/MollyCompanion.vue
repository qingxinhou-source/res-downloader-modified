<template>
  <button
    class="molly-companion"
    :title="preset.name"
    style="--wails-draggable:no-drag"
    @click="pop"
  >
    <span class="companion-halo"></span>
    <img class="companion-img" :src="imageUrl" :alt="preset.name" draggable="false" />
    <span class="companion-dot dot-a"></span>
    <span class="companion-dot dot-b"></span>
    <span class="companion-dot dot-c"></span>
    <span
      v-for="burst in bursts"
      :key="burst.id"
      class="companion-burst"
      :style="{
        '--burst-x': burst.x,
        '--burst-y': burst.y,
        '--burst-delay': burst.delay,
      }"
    ></span>
  </button>
</template>

<script setup lang="ts">
import { computed, ref } from "vue"
import { themeImageUrl } from "@/theme/presets"
import type { ThemePreset } from "@/theme/presets"
import { useIndexStore } from "@/stores"

const props = defineProps<{
  preset: ThemePreset
}>()
const store = useIndexStore()

const imageUrl = computed(
  () => themeImageUrl(props.preset.image, store.baseUrl),
)

const bursts = ref<Array<{ id: number; x: string; y: string; delay: string }>>([])
let burstId = 0

const pop = () => {
  const points = [
    ["-44px", "-22px"],
    ["-18px", "-46px"],
    ["24px", "-48px"],
    ["48px", "-18px"],
    ["34px", "22px"],
    ["-36px", "24px"],
  ]

  bursts.value = points.map(([x, y], index) => ({
    id: burstId++,
    x,
    y,
    delay: `${index * 26}ms`,
  }))

  window.setTimeout(() => {
    bursts.value = []
  }, 820)
}
</script>

<style scoped>
.molly-companion {
  position: fixed;
  right: 22px;
  bottom: 18px;
  z-index: 35;
  width: 74px;
  height: 74px;
  border: 0;
  border-radius: 999px;
  padding: 0;
  background:
    radial-gradient(circle at 50% 56%, color-mix(in srgb, var(--md-surface) 78%, transparent) 0 42%, transparent 44%),
    conic-gradient(from 220deg, color-mix(in srgb, var(--md-primary) 76%, white 8%), color-mix(in srgb, var(--md-accent) 72%, white 8%), color-mix(in srgb, var(--molly-hair) 72%, white 8%), color-mix(in srgb, var(--md-primary) 76%, white 8%));
  box-shadow:
    inset 0 0 0 1px color-mix(in srgb, white 42%, transparent),
    0 14px 34px color-mix(in srgb, var(--md-primary) 24%, transparent);
  cursor: pointer;
  overflow: visible;
  animation: companion-float 4.6s ease-in-out infinite;
  transition: transform 180ms ease, box-shadow 180ms ease, filter 180ms ease;
}

.molly-companion:hover {
  transform: translateY(-5px) scale(1.04);
  filter: saturate(1.08);
  box-shadow:
    inset 0 0 0 1px color-mix(in srgb, white 52%, transparent),
    0 20px 42px color-mix(in srgb, var(--md-primary) 32%, transparent);
}

.molly-companion:active {
  transform: translateY(-1px) scale(0.97);
}

.companion-halo {
  position: absolute;
  inset: -8px;
  border-radius: inherit;
  border: 1px solid color-mix(in srgb, var(--md-primary) 34%, transparent);
  opacity: 0.7;
  animation: companion-pulse 2.8s ease-in-out infinite;
}

.companion-img {
  position: absolute;
  left: 50%;
  bottom: 6px;
  width: 82%;
  height: 82%;
  object-fit: contain;
  object-position: center bottom;
  transform: translateX(-50%);
  filter: drop-shadow(0 8px 10px color-mix(in srgb, var(--md-primary) 22%, transparent));
}

.companion-dot {
  position: absolute;
  width: 8px;
  height: 8px;
  border-radius: 999px;
  background: var(--md-primary);
  box-shadow: 0 0 0 4px color-mix(in srgb, var(--md-primary) 16%, transparent);
  opacity: 0.76;
}

.dot-a {
  left: -2px;
  top: 18px;
  background: var(--md-accent);
}

.dot-b {
  right: 2px;
  top: 10px;
  width: 6px;
  height: 6px;
  background: var(--molly-hair);
}

.dot-c {
  right: 6px;
  bottom: 10px;
  width: 5px;
  height: 5px;
}

.companion-burst {
  position: absolute;
  left: 50%;
  top: 50%;
  width: 9px;
  height: 9px;
  border-radius: 3px;
  background: linear-gradient(135deg, var(--md-primary), var(--md-accent));
  opacity: 0;
  transform: translate(-50%, -50%) rotate(0deg);
  animation: companion-burst 720ms cubic-bezier(0.2, 0.7, 0.1, 1) var(--burst-delay) forwards;
}

@keyframes companion-float {
  0%,
  100% {
    translate: 0 0;
  }
  50% {
    translate: 0 -7px;
  }
}

@keyframes companion-pulse {
  0%,
  100% {
    transform: scale(0.94);
    opacity: 0.48;
  }
  50% {
    transform: scale(1.06);
    opacity: 0.84;
  }
}

@keyframes companion-burst {
  0% {
    opacity: 0;
    transform: translate(-50%, -50%) scale(0.4) rotate(0deg);
  }
  18% {
    opacity: 1;
  }
  100% {
    opacity: 0;
    transform: translate(calc(-50% + var(--burst-x)), calc(-50% + var(--burst-y))) scale(1.05) rotate(135deg);
  }
}

@media (max-width: 900px) {
  .molly-companion {
    right: 14px;
    bottom: 14px;
    width: 62px;
    height: 62px;
  }
}
</style>
