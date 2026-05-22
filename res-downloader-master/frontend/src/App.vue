<template>
  <NConfigProvider
    class="h-full"
    :class="{ 'visual-feedback-off': !store.globalConfig.Personalization?.VisualFeedback }"
    :theme="theme"
    :theme-overrides="themeOverrides"
    :locale="uiLocale"
  >
    <!-- MOLLY 全局背景装饰层 -->
    <MollyBg v-if="store.globalConfig.Personalization?.ShowMollyBackground" :preset="activeTheme" />
    <MollyCompanion v-if="store.globalConfig.Personalization?.ShowMollyCompanion" :preset="activeTheme" />

    <NaiveProvider>
      <RouterView/>
    </NaiveProvider>
    <NGlobalStyle/>
    <NModalProvider/>
  </NConfigProvider>
</template>

<script setup lang="ts">
import NaiveProvider from '@/components/NaiveProvider.vue'
import MollyBg from '@/components/MollyBg.vue'
import MollyCompanion from '@/components/MollyCompanion.vue'
import {darkTheme, lightTheme, zhCN, enUS} from 'naive-ui'
import type {GlobalThemeOverrides} from 'naive-ui'
import {useIndexStore} from "@/stores"
import {computed, onMounted} from "vue"
import {useEventStore} from "@/stores/event"
import type {appType} from "@/types/app"
import {useI18n} from 'vue-i18n'
import {resolveTheme} from "@/theme/presets"

const store = useIndexStore()
const eventStore = useEventStore()
const {locale} = useI18n()

const activeTheme = computed(() => {
  return resolveTheme(store.globalConfig.Theme, store.globalConfig.ThemeMode)
})

const theme = computed(() => {
  const preset = activeTheme.value
  document.documentElement.classList.toggle('dark', preset.mode === "dark")
  Object.entries({
    "--md-primary": preset.primary,
    "--md-primary-hover": preset.primaryHover,
    "--md-on-primary": preset.mode === "dark" ? "#08111f" : "#ffffff",
    "--md-secondary": preset.secondary,
    "--md-error": preset.error,
    "--md-warning": preset.warning,
    "--md-background": preset.background,
    "--md-surface": preset.surface,
    "--md-surface-container": preset.surfaceContainer,
    "--md-surface-container-high": preset.surfaceContainerHigh,
    "--md-outline": preset.outline,
    "--md-text": preset.text,
    "--md-text-muted": preset.textMuted,
    "--md-shadow-1": preset.shadow1,
    "--md-shadow-2": preset.shadow2,
    "--md-theme-glow": preset.glow,
    "--md-accent": preset.accent,
    "--molly": preset.molly,
    "--molly-hair": preset.mollyHair,
    "--molly-eye": preset.mollyEye,
    "--molly-mouth": preset.mollyMouth,
  }).forEach(([key, value]) => document.documentElement.style.setProperty(key, value))
  return preset.mode === "dark" ? darkTheme : lightTheme
})

const uiLocale = computed(() => {
  locale.value = store.globalConfig.Locale
  if (store.globalConfig.Locale === "zh") {
    return zhCN
  }
  return enUS
})

const themeOverrides = computed<GlobalThemeOverrides>(() => {
  const preset = activeTheme.value
  const isDark = preset.mode === "dark"
  const primaryColor = preset.primary
  const primaryHover = preset.primaryHover
  const primaryPressed = preset.primaryPressed
  const surface = preset.surface
  const surfaceVariant = preset.surfaceContainer
  const textBase = preset.text
  const textMuted = preset.textMuted
  const border = preset.outline

  return {
    common: {
      fontFamily: "Inter, 'Segoe UI', 'Microsoft YaHei UI', Arial, sans-serif",
      fontFamilyMono: "'Roboto Mono', 'SFMono-Regular', Consolas, monospace",
      primaryColor,
      primaryColorHover: primaryHover,
      primaryColorPressed: primaryPressed,
      primaryColorSuppl: primaryColor,
      infoColor: primaryColor,
      infoColorHover: primaryHover,
      infoColorPressed: primaryPressed,
      successColor: preset.secondary,
      warningColor: preset.warning,
      errorColor: preset.error,
      textColorBase: textBase,
      textColor1: textBase,
      textColor2: textMuted,
      textColor3: textMuted,
      borderColor: border,
      dividerColor: border,
      bodyColor: preset.background,
      cardColor: surface,
      modalColor: surface,
      popoverColor: surface,
      tableColor: surface,
      tableHeaderColor: surfaceVariant,
      borderRadius: "10px",
      borderRadiusSmall: "8px",
      heightMedium: "40px",
      heightSmall: "32px",
      heightLarge: "48px",
      cubicBezierEaseInOut: "cubic-bezier(0.2, 0, 0, 1)",
    },
    Button: {
      borderRadiusMedium: "10px",
      borderRadiusSmall: "8px",
      fontWeight: "500",
      paddingMedium: "0 18px",
      paddingSmall: "0 14px",
    },
    Card: {
      borderRadius: "12px",
      paddingMedium: "24px",
    },
    DataTable: {
      borderRadius: "12px",
      thColor: surfaceVariant,
      thColorHover: surfaceVariant,
      tdColorHover: isDark ? "#242b3b" : surfaceVariant,
      thTextColor: textMuted,
      tdTextColor: textBase,
      borderColor: border,
    },
    Input: {
      borderRadius: "12px",
    },
    InputNumber: {
      borderRadius: "12px",
    },
    Select: {
      peers: {
        InternalSelection: {
          borderRadius: "12px",
        },
      },
    },
    Menu: {
      itemBorderRadius: "10px",
      itemTextColorActive: primaryColor,
      itemIconColorActive: primaryColor,
      itemColorActive: isDark
        ? `color-mix(in srgb, ${primaryColor} 20%, transparent)`
        : `color-mix(in srgb, ${primaryColor} 14%, ${surface})`,
      itemColorActiveHover: isDark
        ? `color-mix(in srgb, ${primaryColor} 26%, transparent)`
        : `color-mix(in srgb, ${primaryColor} 18%, ${surface})`,
      itemColorHover: isDark
        ? "rgba(255, 255, 255, 0.08)"
        : `color-mix(in srgb, ${surfaceVariant} 84%, ${primaryColor} 8%)`,
    },
    Modal: {
      borderRadius: "20px",
    },
    Popover: {
      borderRadius: "16px",
    },
    Tabs: {
      tabTextColorActiveLine: primaryColor,
      barColor: primaryColor,
    },
  }
})

onMounted(async () => {
  await store.init()

  eventStore.init()
  eventStore.addHandle({
    type: "message",
    event: (res: appType.Message) => {
      switch (res?.code) {
        case 0:
          window.$message?.error(res.message)
          break
        case 1:
          window.$message?.success(res.message)
          break
      }
    }
  })
})
</script>
