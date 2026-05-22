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

const tiltX = ref("0px")
const tiltY = ref("0px")
const isWaving = ref(false)
const message = ref("")
const sparkles = ref<Array<{ id: number; x: string; y: string; delay: string }>>([])

const messages = ["收到", props.preset.name, "一起看看", "准备好了"]
let sparkleId = 0
let waveTimer: number | undefined
let messageTimer: number | undefined

const move = (event: MouseEvent) => {
  const target = event.currentTarget as HTMLElement
  const rect = target.getBoundingClientRect()
  const x = (event.clientX - rect.left) / rect.width - 0.5
  const y = (event.clientY - rect.top) / rect.height - 0.5
  tiltX.value = `${x * 14}px`
  tiltY.value = `${y * 10}px`
}

const leave = () => {
  tiltX.value = "0px"
  tiltY.value = "0px"
}

const interact = () => {
  window.clearTimeout(waveTimer)
  window.clearTimeout(messageTimer)
  isWaving.value = true
  message.value = messages[Math.floor(Math.random() * messages.length)]
  sparkles.value = Array.from({length: 8}, (_, index) => ({
    id: sparkleId++,
    x: `${Math.cos(index * 0.78) * (44 + index * 4)}px`,
    y: `${Math.sin(index * 0.78) * (28 + index * 3) - 38}px`,
    delay: `${index * 22}ms`,
  }))

  waveTimer = window.setTimeout(() => {
    isWaving.value = false
    sparkles.value = []
  }, 900)

  messageTimer = window.setTimeout(() => {
    message.value = ""
  }, 1500)
}
</script>

<template>
  <button
    class="molly-bg"
    :class="{'is-waving': isWaving}"
    :title="preset.name"
    type="button"
    style="--wails-draggable:no-drag"
    @mousemove="move"
    @mouseleave="leave"
    @click="interact"
  >
    <span v-if="message" class="molly-bg__bubble">{{ message }}</span>
    <span
      v-for="sparkle in sparkles"
      :key="sparkle.id"
      class="molly-bg__sparkle"
      :style="{
        '--sparkle-x': sparkle.x,
        '--sparkle-y': sparkle.y,
        '--sparkle-delay': sparkle.delay,
      }"
    ></span>
    <img
      class="molly-bg__img"
      :src="imageUrl"
      :alt="preset.name"
      draggable="false"
      :style="{
        '--molly-shift-x': tiltX,
        '--molly-shift-y': tiltY,
      }"
    />
  </button>
</template>

<style scoped>
.molly-bg {
  position: fixed;
  right: -20px;
  bottom: -34px;
  z-index: 1;
  width: clamp(300px, 30vw, 520px);
  height: auto;
  border: 0;
  padding: 0;
  background: transparent;
  overflow: visible;
  cursor: pointer;
  pointer-events: auto;
}

.molly-bg::before {
  content: "";
  position: absolute;
  inset: 18% 8% 4% 6%;
  border-radius: 50%;
  background:
    radial-gradient(circle at 34% 28%, color-mix(in srgb, var(--molly-hair) 34%, transparent) 0 18%, transparent 20%),
    radial-gradient(circle at 70% 36%, color-mix(in srgb, var(--md-accent) 22%, transparent) 0 22%, transparent 24%),
    color-mix(in srgb, var(--md-primary) 10%, transparent);
  filter: blur(1px);
  opacity: 0.9;
}

.molly-bg__img {
  width: 100%;
  height: auto;
  object-fit: contain;
  object-position: bottom right;
  opacity: 0.72;
  filter:
    drop-shadow(0 24px 34px color-mix(in srgb, var(--md-primary) 22%, transparent))
    drop-shadow(0 -4px 12px color-mix(in srgb, var(--md-surface) 48%, transparent));
  position: relative;
  transform: translate(var(--molly-shift-x, 0), var(--molly-shift-y, 0));
  transition: opacity 400ms ease, transform 260ms ease, filter 260ms ease;
  user-select: none;
}

:global(.dark) .molly-bg__img {
  opacity: 0.65;
}

.molly-bg:hover .molly-bg__img {
  opacity: 0.92;
  transform: translate(var(--molly-shift-x, 0), calc(var(--molly-shift-y, 0) - 3px)) scale(1.02);
}

.molly-bg.is-waving .molly-bg__img {
  animation: molly-wave 620ms cubic-bezier(0.2, 0.75, 0.2, 1);
}

.molly-bg__bubble {
  position: absolute;
  right: 50%;
  top: 22%;
  z-index: 2;
  min-width: 72px;
  max-width: 132px;
  padding: 8px 12px;
  border-radius: 14px 14px 4px;
  background: color-mix(in srgb, var(--md-surface) 92%, var(--molly) 8%);
  color: var(--md-text);
  border: 1px solid color-mix(in srgb, var(--md-outline) 70%, var(--molly-mouth) 30%);
  box-shadow: 0 10px 24px color-mix(in srgb, var(--md-primary) 18%, transparent);
  font-size: 13px;
  font-weight: 700;
  line-height: 1.2;
  white-space: nowrap;
  animation: molly-bubble 1500ms ease forwards;
}

.molly-bg__sparkle {
  position: absolute;
  left: 50%;
  top: 48%;
  z-index: 2;
  width: 10px;
  height: 10px;
  border-radius: 3px;
  background: linear-gradient(135deg, var(--molly-hair), var(--md-accent));
  opacity: 0;
  transform: translate(-50%, -50%) rotate(0deg) scale(0.4);
  animation: molly-sparkle 760ms cubic-bezier(0.16, 0.72, 0.18, 1) var(--sparkle-delay) forwards;
}

@keyframes molly-wave {
  0%,
  100% {
    transform: translate(var(--molly-shift-x, 0), var(--molly-shift-y, 0)) rotate(0deg) scale(1);
  }
  28% {
    transform: translate(calc(var(--molly-shift-x, 0) - 8px), calc(var(--molly-shift-y, 0) - 12px)) rotate(-3deg) scale(1.035);
  }
  58% {
    transform: translate(calc(var(--molly-shift-x, 0) + 7px), calc(var(--molly-shift-y, 0) - 8px)) rotate(2.5deg) scale(1.02);
  }
}

@keyframes molly-bubble {
  0% {
    opacity: 0;
    transform: translateY(8px) scale(0.94);
  }
  16%,
  78% {
    opacity: 1;
    transform: translateY(0) scale(1);
  }
  100% {
    opacity: 0;
    transform: translateY(-6px) scale(0.98);
  }
}

@keyframes molly-sparkle {
  0% {
    opacity: 0;
    transform: translate(-50%, -50%) rotate(0deg) scale(0.4);
  }
  18% {
    opacity: 1;
  }
  100% {
    opacity: 0;
    transform: translate(calc(-50% + var(--sparkle-x)), calc(-50% + var(--sparkle-y))) rotate(150deg) scale(1.1);
  }
}

@media (prefers-reduced-motion: reduce) {
  .molly-bg,
  .molly-bg__img,
  .molly-bg__bubble,
  .molly-bg__sparkle {
    animation: none !important;
    transition: none !important;
  }
}
</style>
