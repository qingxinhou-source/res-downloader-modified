<template>
  <div
    class="sider-shell flex flex-col h-full border-r border-[var(--md-outline)]"
    :class="collapsed ? 'w-[104px] min-w-[104px]' : 'w-[204px] min-w-[204px]'"
  >
    <Screen v-if="envInfo.platform !== 'darwin'" />

    <div class="px-3" :class="envInfo.platform === 'darwin' ? 'pt-8' : 'pt-3'">
      <button class="brand-card" @click="handleFooterUpdate('about')">
        <MollyAvatar class="brand-avatar" :preset="activeTheme" :size="44" />
        <span v-if="!collapsed" class="brand-copy">
          <span class="brand-title">Res Downloader</span>
          <span class="brand-subtitle">{{ activeTheme.name }}</span>
        </span>
        <span v-if="showUpdate" class="update-dot">New</span>
      </button>
    </div>

    <MollyWidget class="mx-3 mt-3" :compact="collapsed" />

    <main v-if="is" class="flex-1 min-h-0 pt-3">
      <NScrollbar :size="1">
        <NLayout has-sider>
          <NLayoutSider
            :bordered="false"
            show-trigger
            collapse-mode="width"
            :collapsed-width="84"
            :collapsed="collapsed"
            :width="184"
            :native-scrollbar="false"
            :inverted="inverted"
            :on-update:collapsed="collapsedChange"
            class="bg-inherit material-sider"
          >
            <NMenu
              :inverted="inverted"
              :collapsed-width="84"
              :collapsed-icon-size="22"
              :options="menuOptions"
              :value="menuValue"
              @update:value="handleUpdateValue"
            />
          </NLayoutSider>
        </NLayout>
        <NLayoutFooter position="absolute" :inverted="inverted" class="bg-inherit">
          <NMenu
            :inverted="inverted"
            :collapsed-width="84"
            :collapsed-icon-size="22"
            :options="footerOptions"
            :value="menuValue"
            @update:value="handleFooterUpdate"
          />
        </NLayoutFooter>
      </NScrollbar>
    </main>

    <div class="px-3 pb-3 space-y-2" style="--wails-draggable:no-drag">
      <NPopover placement="right" trigger="click">
        <template #trigger>
          <button class="theme-trigger" :title="activeTheme.name">
            <span class="theme-dot"></span>
            <span v-if="!collapsed">{{ activeTheme.name }}</span>
          </button>
        </template>
        <div class="theme-popover">
          <div class="mb-3">
            <div class="text-sm font-semibold text-[var(--md-text)]">沿途的风景主题</div>
            <div class="text-xs text-[var(--md-text-muted)] mt-1">每款盲盒一套配色。</div>
          </div>
          <div class="theme-list">
            <button
              v-for="item in displayThemeOptions"
              :key="item.key"
              class="theme-choice"
              :class="{ 'theme-choice-active': item.key === globalConfig.Theme }"
              :style="themeChoiceStyle(item)"
              @click="setTheme(item.key)"
            >
              <!-- 主题角色头像（小尺寸） -->
              <span class="theme-choice-avatar">
                <MollyAvatar :preset="item" :size="38" />
              </span>
              <span class="min-w-0">
                <span class="block text-sm font-medium">{{ item.name }}</span>
                <span class="block text-xs text-[var(--md-text-muted)]">{{ item.desc }}</span>
              </span>
              <!-- 色板条 -->
              <span class="theme-choice-swatch-bar" :style="{ background: item.swatch }"></span>
            </button>
          </div>
        </div>
      </NPopover>
    </div>
  </div>
  <Footer v-model:showModal="showAppInfo" />
</template>

<script lang="ts" setup>
import { MenuOption, NIcon } from "naive-ui"
import { computed, h, onMounted, ref, watch } from "vue"
import { useRoute, useRouter } from "vue-router"
import {
  CloudOutline,
  SettingsOutline,
  HelpCircleOutline,
  ColorPaletteOutline,
  LanguageSharp,
} from "@vicons/ionicons5"
import { useIndexStore } from "@/stores"
import Footer from "@/components/Footer.vue"
import MollyWidget from "@/components/MollyWidget.vue"
import MollyAvatar from "@/components/MollyAvatar.vue"
import Screen from "@/components/Screen.vue"
import { useI18n } from "vue-i18n"
import request from "@/api/request"
import { compareVersions } from "@/func"
import { getThemeOptions, resolveTheme, resolveThemeOptions } from "@/theme/presets"
import type { ThemePreset } from "@/theme/presets"

const { t } = useI18n()
const route = useRoute()
const router = useRouter()
const inverted = ref(false)
const collapsed = ref(false)
const showAppInfo = ref(false)
const menuValue = ref(route.fullPath.substring(1))
const store = useIndexStore()
const is = ref(false)
const showUpdate = ref(false)

const envInfo = store.envInfo
const globalConfig = computed(() => store.globalConfig)
const activeTheme = computed(() => resolveTheme(store.globalConfig.Theme, store.globalConfig.ThemeMode))
const displayThemeOptions = computed(() => {
  store.globalConfig.CustomThemes
  return resolveThemeOptions(store.globalConfig.ThemeMode)
})

watch(() => route.path, () => {
  menuValue.value = route.fullPath.substring(1)
})

onMounted(() => {
  const collapsedCache = localStorage.getItem("collapsed")
  if (collapsedCache) {
    collapsed.value = JSON.parse(collapsedCache).collapsed
  }
  is.value = true

  request({
    url: "https://res.putyy.com/version.json?v=" + Date.now(),
    method: "get",
  }).then((res) => {
    showUpdate.value = compareVersions(res.version, store.appInfo.Version) === 1
  })
})

const renderIcon = (icon: any) => {
  return () => h(NIcon, null, { default: () => h(icon) })
}

const menuOptions = ref([
  {
    label: computed(() => t("menu.index")),
    key: "index",
    icon: renderIcon(CloudOutline),
  },
  {
    label: computed(() => t("menu.setting")),
    key: "setting",
    icon: renderIcon(SettingsOutline),
  },
])

const footerOptions = ref([
  {
    label: computed(() => t("menu.locale")),
    key: "locale",
    icon: renderIcon(LanguageSharp),
  },
  {
    label: computed(() => t("menu.theme")),
    key: "theme",
    icon: renderIcon(ColorPaletteOutline),
  },
  {
    label: computed(() => t("menu.about")),
    key: "about",
    icon: renderIcon(HelpCircleOutline),
  },
])

const handleUpdateValue = (key: string, item?: MenuOption) => {
  menuValue.value = key
  return router.push({ path: "/" + key })
}

const handleFooterUpdate = (key: string, item?: MenuOption) => {
  if (key === "about") {
    showAppInfo.value = true
    return
  }

  if (key === "theme") {
    const options = getThemeOptions()
    const index = options.findIndex(item => item.key === globalConfig.value.Theme)
    const next = options[(index + 1) % options.length]
    setTheme(next.key)
    return
  }

  if (key === "locale") {
    store.setConfig({ Locale: globalConfig.value.Locale === "zh" ? "en" : "zh" })
    return
  }

  menuValue.value = key
  return router.push({ path: "/" + key })
}

const collapsedChange = (value: boolean) => {
  collapsed.value = value
  localStorage.setItem("collapsed", JSON.stringify({ collapsed: value }))
}

const setTheme = (themeKey: string) => {
  store.setConfig({ Theme: themeKey })
}

const themeChoiceStyle = (item: ThemePreset) => ({
  "--choice-primary": item.primary,
  "--choice-accent": item.accent,
  "--choice-surface": item.surface,
  "--choice-outline": item.outline,
})
</script>

<style scoped>
.sider-shell {
  background:
    radial-gradient(circle at 18% 12%, color-mix(in srgb, var(--molly-hair) 22%, transparent) 0 70px, transparent 72px),
    radial-gradient(circle at 96% 42%, color-mix(in srgb, var(--md-accent) 14%, transparent) 0 86px, transparent 88px),
    linear-gradient(180deg, color-mix(in srgb, var(--md-surface-container) 82%, var(--molly) 10%), var(--md-surface-container));
  transition: width 180ms ease, min-width 180ms ease;
}

.brand-card {
  position: relative;
  width: 100%;
  min-height: 68px;
  display: flex;
  align-items: center;
  gap: 10px;
  padding: 10px;
  border-radius: 8px;
  color: var(--md-text);
  background:
    radial-gradient(circle at 86% 18%, color-mix(in srgb, var(--md-accent) 18%, transparent) 0 28px, transparent 30px),
    linear-gradient(145deg, color-mix(in srgb, var(--md-surface) 90%, var(--molly) 10%), var(--md-surface));
  border: 1px solid color-mix(in srgb, var(--md-outline) 70%, var(--md-primary) 24%);
  box-shadow: var(--md-shadow-1);
  text-align: left;
  overflow: hidden;
  transition: transform 160ms ease, box-shadow 160ms ease, border-color 160ms ease;
}

.brand-card::after {
  content: "";
  position: absolute;
  inset: 0;
  background: linear-gradient(110deg, transparent 0 38%, color-mix(in srgb, white 20%, transparent) 45%, transparent 52%);
  transform: translateX(-110%);
  transition: transform 620ms ease;
  pointer-events: none;
}

.brand-card:hover {
  transform: translateY(-2px);
  border-color: color-mix(in srgb, var(--md-primary) 44%, var(--md-outline));
  box-shadow: var(--md-shadow-2);
}

.brand-card:hover::after {
  transform: translateX(110%);
}

.brand-card:active,
.theme-trigger:active,
.author-chip:active,
.theme-choice:active {
  transform: translateY(0) scale(0.98);
}

.brand-avatar {
  flex: 0 0 auto;
}

.brand-copy {
  min-width: 0;
  display: flex;
  flex-direction: column;
}

.brand-title {
  font-weight: 800;
  line-height: 1.15;
}

.brand-subtitle {
  color: var(--md-text-muted);
  font-size: 12px;
  line-height: 1.25;
  margin-top: 3px;
}

.update-dot {
  position: absolute;
  right: 8px;
  top: 8px;
  font-size: 10px;
  line-height: 1;
  padding: 3px 5px;
  border-radius: 999px;
  color: white;
  background: var(--md-error);
}

.material-sider :deep(.n-layout-toggle-button) {
  border: 1px solid var(--md-outline);
  box-shadow: var(--md-shadow-1);
}

.theme-trigger,
.author-chip {
  width: 100%;
  min-height: 40px;
  border-radius: 8px;
  border: 1px solid var(--md-outline);
  background: var(--md-surface);
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 8px;
  color: var(--md-primary);
  box-shadow: var(--md-shadow-1);
  transition: transform 160ms ease, box-shadow 160ms ease, border-color 160ms ease;
}

.theme-trigger:hover,
.author-chip:hover {
  transform: translateY(-1px);
  border-color: color-mix(in srgb, var(--md-primary) 42%, var(--md-outline));
  box-shadow: var(--md-shadow-2);
}

.author-chip span {
  font-weight: 800;
  letter-spacing: 0;
}

.theme-popover {
  width: 300px;
  padding: 6px;
}

.theme-list {
  display: grid;
  gap: 6px;
}

.theme-choice {
  width: 100%;
  display: flex;
  align-items: center;
  gap: 10px;
  padding: 8px;
  border-radius: 8px;
  border: 1px solid color-mix(in srgb, var(--choice-outline) 48%, transparent);
  background:
    radial-gradient(circle at 12% 22%, color-mix(in srgb, var(--choice-primary) 12%, transparent) 0 30px, transparent 32px),
    linear-gradient(135deg, color-mix(in srgb, var(--choice-surface) 92%, var(--choice-primary) 8%), var(--choice-surface));
  color: var(--md-text);
  text-align: left;
  transition: background-color 160ms ease, border-color 160ms ease, transform 160ms ease, box-shadow 160ms ease;
  overflow: hidden;
}

.theme-choice:hover,
.theme-choice-active {
  transform: translateX(2px);
  background:
    radial-gradient(circle at 92% 18%, color-mix(in srgb, var(--choice-accent) 22%, transparent) 0 22px, transparent 24px),
    linear-gradient(135deg, color-mix(in srgb, var(--choice-surface) 84%, var(--choice-primary) 16%), var(--choice-surface));
  border-color: color-mix(in srgb, var(--choice-primary) 42%, var(--choice-outline));
  box-shadow: 0 10px 22px color-mix(in srgb, var(--choice-primary) 14%, transparent);
}

/* 角色头像容器 */
.theme-choice-avatar {
  flex-shrink: 0;
  width: 46px;
  height: 46px;
  display: flex;
  align-items: flex-end;
  justify-content: center;
  overflow: visible;
  border-radius: 999px;
  background: transparent;
  position: relative;
}

/* 色板条（紧凑横条） */
.theme-choice-swatch-bar {
  flex-shrink: 0;
  width: 5px;
  align-self: stretch;
  border-radius: 999px;
  opacity: 0.6;
}
</style>
