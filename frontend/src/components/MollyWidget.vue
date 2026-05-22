<template>
  <button
    class="molly-widget"
    :class="{ compact }"
    :title="activeTheme.name"
    style="--wails-draggable:no-drag"
    @click="tap"
  >
    <span class="widget-art">
      <MollyAvatar :preset="activeTheme" :size="compact ? 48 : 70" />
      <span class="widget-ping ping-a"></span>
      <span class="widget-ping ping-b"></span>
    </span>

    <span v-if="!compact" class="widget-copy">
      <span class="widget-kicker">Scenery Along The Way</span>
      <span class="widget-name">{{ activeTheme.name }}</span>
      <span class="widget-desc">{{ activeTheme.desc }}</span>
    </span>

    <span v-if="compact" class="compact-copy">
      <span class="compact-name">{{ activeTheme.name }}</span>
      <span class="compact-swatch" :style="{ background: activeTheme.swatch }"></span>
    </span>

    <span class="widget-swatch" :style="{ background: activeTheme.swatch }"></span>
  </button>
</template>

<script setup lang="ts">
import { computed } from "vue"
import { useIndexStore } from "@/stores"
import { getThemeOptions, resolveTheme } from "@/theme/presets"
import MollyAvatar from "@/components/MollyAvatar.vue"

defineProps({
  compact: {
    type: Boolean,
    default: false,
  },
})

const store = useIndexStore()

const activeTheme = computed(() => resolveTheme(store.globalConfig.Theme, store.globalConfig.ThemeMode))

const tap = () => {
  const options = getThemeOptions()
  const index = options.findIndex(item => item.key === store.globalConfig.Theme)
  const next = options[(index + 1) % options.length]
  store.setConfig({ Theme: next.key })
}
</script>

<style scoped>
.molly-widget {
  position: relative;
  min-height: 108px;
  width: calc(100% - 1.5rem);
  padding: 12px;
  display: grid;
  grid-template-columns: 78px minmax(0, 1fr);
  align-items: center;
  gap: 12px;
  border-radius: 8px;
  border: 1px solid color-mix(in srgb, var(--md-outline) 62%, var(--md-primary) 18%);
  color: var(--md-text);
  background:
    radial-gradient(circle at 20% 18%, color-mix(in srgb, var(--molly-hair) 32%, transparent) 0 34px, transparent 36px),
    radial-gradient(circle at 94% 12%, color-mix(in srgb, var(--md-accent) 18%, transparent) 0 40px, transparent 42px),
    linear-gradient(145deg, color-mix(in srgb, var(--md-surface) 88%, var(--molly) 12%), var(--md-surface));
  box-shadow: var(--md-shadow-1);
  overflow: hidden;
  text-align: left;
  transition: transform 180ms ease, box-shadow 200ms ease, border-color 180ms ease;
}

.molly-widget::before {
  content: "";
  position: absolute;
  inset: 0;
  background: linear-gradient(110deg, transparent 0 36%, color-mix(in srgb, white 20%, transparent) 43%, transparent 50%);
  transform: translateX(-110%);
  transition: transform 620ms ease;
  pointer-events: none;
  z-index: 1;
}

.molly-widget:hover {
  transform: translateY(-2px);
  border-color: color-mix(in srgb, var(--md-primary) 46%, var(--md-outline));
  box-shadow: var(--md-shadow-2);
}

.molly-widget:hover::before {
  transform: translateX(110%);
}

.molly-widget:active {
  transform: translateY(0) scale(0.99);
}

.molly-widget.compact {
  min-height: 82px;
  grid-template-columns: 52px minmax(0, 1fr);
  padding: 9px;
  gap: 8px;
}

.widget-art {
  position: relative;
  z-index: 2;
  display: flex;
  align-items: center;
  justify-content: center;
}

.widget-ping {
  position: absolute;
  border-radius: 999px;
  background: var(--md-primary);
  box-shadow: 0 0 0 5px color-mix(in srgb, var(--md-primary) 18%, transparent);
  opacity: 0;
  transform: scale(0.4);
}

.ping-a {
  width: 7px;
  height: 7px;
  right: 8px;
  top: 9px;
  background: var(--md-accent);
}

.ping-b {
  width: 5px;
  height: 5px;
  left: 10px;
  bottom: 10px;
  background: var(--molly-hair);
}

.molly-widget:hover .widget-ping {
  animation: widget-ping 900ms ease forwards;
}

.molly-widget:hover .ping-b {
  animation-delay: 120ms;
}

.widget-copy,
.compact-copy {
  position: relative;
  z-index: 2;
  min-width: 0;
  display: flex;
  flex-direction: column;
}

.widget-kicker {
  color: var(--md-primary);
  font-size: 9px;
  font-weight: 800;
  line-height: 1.1;
  text-transform: uppercase;
}

.widget-name {
  margin-top: 5px;
  color: var(--md-text);
  font-size: 15px;
  font-weight: 850;
  line-height: 1.1;
}

.widget-desc {
  margin-top: 5px;
  color: var(--md-text-muted);
  font-size: 11px;
  line-height: 1.35;
}

.compact-name {
  color: var(--md-text);
  font-size: 11px;
  font-weight: 800;
  line-height: 1.15;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.compact-swatch {
  width: 44px;
  height: 5px;
  margin-top: 7px;
  border-radius: 999px;
}

.widget-swatch {
  position: absolute;
  left: 0;
  right: 0;
  bottom: 0;
  height: 5px;
}

@keyframes widget-ping {
  0% {
    opacity: 0;
    transform: scale(0.4);
  }
  35% {
    opacity: 0.9;
  }
  100% {
    opacity: 0;
    transform: scale(1.6);
  }
}
</style>
