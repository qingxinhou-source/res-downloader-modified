<template>
  <div class="h-full relative p-5 overflow-y-auto [&::-webkit-scrollbar]:hidden" :key="renderKey">
    <div class="max-w-[1020px] min-h-full">
      <div class="mb-4 soft-panel p-5">
        <div class="flex flex-wrap items-end justify-between gap-3">
          <div>
            <div class="text-xs text-[var(--md-text-muted)] uppercase">Scenery Along The Way</div>
            <div class="mt-1 text-2xl font-bold text-[var(--md-text)]">{{ t('setting.basic_setting') }}</div>
          </div>
          <div class="text-sm text-[var(--md-text-muted)]">Scenery Along The Way</div>
        </div>
      </div>
    <div class="soft-panel p-4">
    <NTabs type="line" animated>
      <NTabPane name="basic" :tab="t('setting.basic_setting')">
        <div class="setting-form-panel">
          <div class="theme-setting-row">
            <div class="min-w-0">
              <div class="text-sm font-semibold text-[var(--md-text)]">盲盒主题</div>
              <div class="text-xs text-[var(--md-text-muted)] mt-1">每张抠图对应一套完整界面配色。</div>
            </div>
            <div class="theme-setting-grid">
              <button
                  v-for="(item, index) in themeOptions"
                  :key="item.key"
                  class="theme-setting-card"
                  :class="{ 'theme-setting-card-active': formValue.Theme === item.key }"
                  :style="themeCardStyle(item)"
                  @click="setTheme(item.key)"
              >
                <span class="theme-card-number">No. {{ String(index + 1).padStart(2, '0') }}</span>
                <span class="theme-card-check" aria-hidden="true"></span>
                <span class="theme-card-art" aria-hidden="true">
                  <span class="theme-card-stage"></span>
                  <img
                      class="theme-card-image"
                      :src="themeImageUrl(item)"
                      :alt="item.name"
                      draggable="false"
                  />
                </span>
                <span class="theme-card-body">
                  <span class="theme-card-series">Scenery Along The Way</span>
                  <span class="theme-card-name">{{ item.name }}</span>
                  <span class="theme-card-desc">{{ item.desc }}</span>
                  <span class="theme-card-spec">
                    <span>MOLLY</span>
                    <span>{{ item.mode.toUpperCase() }}</span>
                  </span>
                </span>
                <span class="theme-card-swatch" :style="{ background: item.swatch }"></span>
              </button>
            </div>
          </div>
          <div class="custom-theme-lab">
            <div class="custom-theme-preview" :style="themeCardStyle(customThemePreview)">
              <span class="theme-card-number">CUSTOM</span>
              <span class="theme-card-art" aria-hidden="true">
                <span class="theme-card-stage"></span>
                <img
                    class="theme-card-image"
                    :src="themeImageUrl(customThemePreview)"
                    :alt="customThemeForm.name"
                    draggable="false"
                />
              </span>
              <span class="theme-card-body">
                <span class="theme-card-series">Custom Theme Lab</span>
                <span class="theme-card-name">{{ customThemeForm.name }}</span>
                <span class="theme-card-desc">{{ customThemeForm.desc }}</span>
                <span class="theme-card-spec">
                  <span>MOLLY</span>
                  <span>{{ store.globalConfig.ThemeMode.toUpperCase() }}</span>
                </span>
              </span>
              <span class="theme-card-swatch" :style="{ background: customThemePreview.swatch }"></span>
            </div>
            <div class="custom-theme-editor">
              <div class="custom-theme-editor-head">
                <div>
                  <div class="text-sm font-semibold text-[var(--md-text)]">自定义主题工坊</div>
                  <div class="text-xs text-[var(--md-text-muted)] mt-1">保存后会作为角色卡加入上方主题列表，可继续编辑或删除。</div>
                </div>
                <NSelect
                    v-model:value="editingCustomThemeKey"
                    clearable
                    class="custom-theme-select"
                    placeholder="选择已保存主题"
                    :options="customThemeSelectOptions"
                    @update:value="loadCustomTheme"
                />
              </div>
              <div class="custom-theme-fields">
                <label class="custom-field wide">
                  <span>主题名称</span>
                  <NInput v-model:value="customThemeForm.name" placeholder="我的主题"/>
                </label>
                <label class="custom-field wide">
                  <span>角色描述</span>
                  <NInput v-model:value="customThemeForm.desc" placeholder="写一句角色卡描述"/>
                </label>
                <label class="custom-field wide">
                  <span>角色图片</span>
                  <span class="custom-image-picker">
                    <NSelect v-model:value="customThemeForm.image" :options="themeImageOptions"/>
                    <NButton secondary @click="selectCustomThemeImage">本地图片</NButton>
                  </span>
                </label>
                <label
                    v-for="field in customColorFields"
                    :key="field.key"
                    class="custom-color-field"
                >
                  <span>{{ field.label }}</span>
                  <span class="custom-color-control">
                    <input v-model="customThemeForm[field.key]" type="color" />
                    <NInput v-model:value="customThemeForm[field.key]" />
                  </span>
                </label>
              </div>
              <div class="custom-theme-actions">
                <NButton secondary @click="resetCustomThemeForm">以当前主题为模板</NButton>
                <NButton type="primary" strong @click="saveCustomTheme">保存并应用</NButton>
                <NButton
                    tertiary
                    type="error"
                    :disabled="!editingCustomThemeKey"
                    @click="deleteCustomTheme"
                >
                  删除自定义主题
                </NButton>
              </div>
            </div>
          </div>
        <NForm
          :model="formValue"
          size="medium"
          label-placement="left"
          label-width="auto"
          require-mark-placement="right-hanging"
          style="--wails-draggable:no-drag"
          class="w-[700px] max-w-full"
        >
          <NFormItem :label="t('setting.save_dir')" path="SaveDirectory">
            <NInput :value="formValue.SaveDirectory" :placeholder="t('setting.save_dir')"/>
            <NButton strong secondary type="primary" @click="selectDir" class="ml-1">{{ t('common.select') }}</NButton>
          </NFormItem>

          <NFormItem :label="t('setting.filename_rules')" path="FilenameLen">
            <NInputNumber v-model:value="formValue.FilenameLen" :min="0" :max="9999" placeholder="0"/>
            <NSwitch v-model:value="formValue.FilenameTime" class="ml-1"></NSwitch>
            <NTooltip trigger="hover">
              <template #trigger>
                <NIcon size="18" class="ml-1 text-gray-500">
                  <HelpCircleOutline/>
                </NIcon>
              </template>
              {{ t("setting.filename_rules_tip") }}
            </NTooltip>
          </NFormItem>

          <NFormItem :label="t('setting.quality')" path="Quality">
            <NSelect v-model:value="formValue.Quality" :options="options"/>
            <NTooltip trigger="hover">
              <template #trigger>
                <NIcon size="18" class="ml-1 text-gray-500">
                  <HelpCircleOutline/>
                </NIcon>
              </template>
              {{ t("setting.quality_tip") }}
            </NTooltip>
          </NFormItem>

          <NFormItem :label="t('setting.auto_proxy')" path="AutoProxy">
            <NSwitch v-model:value="formValue.AutoProxy"/>
            <NTooltip trigger="hover">
              <template #trigger>
                <NIcon size="18" class="ml-1 text-gray-500">
                  <HelpCircleOutline/>
                </NIcon>
              </template>
              {{ t("setting.auto_proxy_tip") }}
            </NTooltip>
          </NFormItem>

          <NFormItem :label="t('setting.full_intercept')" path="WxAction">
            <NSwitch v-model:value="formValue.WxAction"/>
            <NTooltip trigger="hover">
              <template #trigger>
                <NIcon size="18" class="ml-1 text-gray-500">
                  <HelpCircleOutline/>
                </NIcon>
              </template>
              {{ t("setting.full_intercept_tip") }}
            </NTooltip>
          </NFormItem>

          <NFormItem :label="t('setting.insert_tail')" path="InsertTail">
            <NSwitch v-model:value="formValue.InsertTail"/>
            <NTooltip trigger="hover">
              <template #trigger>
                <NIcon size="18" class="ml-1 text-gray-500">
                  <HelpCircleOutline/>
                </NIcon>
              </template>
              {{ t("setting.insert_tail_tip") }}
            </NTooltip>
          </NFormItem>

          <NFormItem >
            <n-popconfirm @positive-click="resetHandle">
              <template #trigger>
                <NButton tertiary type="error" style="--wails-draggable:no-drag">
                  {{ t("index.start_err_positiveText") }}
                </NButton>
              </template>
              {{t("index.reset_app_tip")}}
            </n-popconfirm>
          </NFormItem>
        </NForm>
        </div>
      </NTabPane>

      <NTabPane name="advanced" :tab="t('setting.advanced_setting')">
        <div class="setting-form-panel">
        <NForm
          :model="formValue"
          size="medium"
          label-placement="left"
          label-width="auto"
          require-mark-placement="right-hanging"
          style="--wails-draggable:no-drag"
          class="w-[700px] max-w-full"
        >
          <NFormItem label="Host" path="Host" :validation-status="hostValidationFeedback==='' ? undefined : 'error'" :feedback="hostValidationFeedback">
            <NInput v-model:value="formValue.Host" placeholder="127.0.0.1"/>
            <NTooltip trigger="hover">
              <template #trigger>
                <NIcon size="18" class="ml-1 text-gray-500">
                  <HelpCircleOutline/>
                </NIcon>
              </template>
              {{ t("setting.restart_tip") }}
            </NTooltip>
          </NFormItem>

          <NFormItem label="Port" path="Port" :validation-status="portValidationFeedback==='' ? undefined : 'error'" :feedback="portValidationFeedback">
            <NInput v-model:value="formValue.Port" placeholder="8899"/>
            <NTooltip trigger="hover">
              <template #trigger>
                <NIcon size="18" class="ml-1 text-gray-500">
                  <HelpCircleOutline/>
                </NIcon>
              </template>
              {{ t("setting.restart_tip") }}
            </NTooltip>
          </NFormItem>

          <NFormItem :label="t('setting.upstream_proxy')" path="UpstreamProxy">
            <NInput v-model:value="formValue.UpstreamProxy" placeholder="http://127.0.0.1:7890"/>
            <NSwitch v-model:value="formValue.OpenProxy" class="ml-1"/>
            <NTooltip trigger="hover">
              <template #trigger>
                <NIcon size="18" class="ml-1 text-gray-500">
                  <HelpCircleOutline/>
                </NIcon>
              </template>
              {{ t("setting.upstream_proxy_tip") }}
            </NTooltip>
          </NFormItem>

          <NFormItem :label="t('setting.download_proxy')" path="DownloadProxy">
            <NSwitch v-model:value="formValue.DownloadProxy"/>
            <NTooltip trigger="hover">
              <template #trigger>
                <NIcon size="18" class="ml-1 text-gray-500">
                  <HelpCircleOutline/>
                </NIcon>
              </template>
              {{ t("setting.download_proxy_tip") }}
            </NTooltip>
          </NFormItem>

          <NFormItem :label="t('setting.connections')" path="TaskNumber">
            <NInputNumber v-model:value="formValue.TaskNumber" :min="2" :max="64"/>
            <NTooltip trigger="hover">
              <template #trigger>
                <NIcon size="18" class="ml-1 text-gray-500">
                  <HelpCircleOutline/>
                </NIcon>
              </template>
              {{ t("setting.connections_tip") }}
            </NTooltip>
          </NFormItem>

          <NFormItem :label="t('setting.down_number')" path="DownNumber">
            <NInputNumber v-model:value="formValue.DownNumber" :min="1" :max="10"/>
            <NTooltip trigger="hover">
              <template #trigger>
                <NIcon size="18" class="ml-1 text-gray-500">
                  <HelpCircleOutline/>
                </NIcon>
              </template>
              {{ t("setting.down_number_tip") }}
            </NTooltip>
          </NFormItem>

          <NFormItem label="UserAgent" path="UserAgent">
            <NInput v-model:value="formValue.UserAgent" placeholder="UserAgent"/>
            <NTooltip trigger="hover">
              <template #trigger>
                <NIcon size="18" class="ml-1 text-gray-500">
                  <HelpCircleOutline/>
                </NIcon>
              </template>
              {{ t("setting.user_agent_tip") }}
            </NTooltip>
          </NFormItem>

          <NFormItem label="Headers" path="Headers">
            <NInput v-model:value="formValue.UseHeaders" placeholder="User-Agent,Referer,Authorization,Cookie"/>
            <NTooltip trigger="hover">
              <template #trigger>
                <NIcon size="18" class="ml-1 text-gray-500">
                  <HelpCircleOutline/>
                </NIcon>
              </template>
              {{ t("setting.use_headers_tip") }}
            </NTooltip>
          </NFormItem>

          <NFormItem :label="t('setting.domain_rule')" path="DomainRule">
            <NInput
                v-model:value="formValue.Rule"
                type="textarea"
                rows="5"
                :placeholder="t('setting.domain_rule_tip')"
            />
            <NTooltip trigger="hover">
              <template #trigger>
                <NIcon size="18" class="ml-1 text-gray-500">
                  <HelpCircleOutline/>
                </NIcon>
              </template>
              {{ t("setting.domain_rule_tip") }}
            </NTooltip>
          </NFormItem>
          <NFormItem
              :label="t('setting.mime_map')"
              path="MimeMap"
              :validation-status="mimeMapValidationFeedback === '' ? undefined : 'error'"
              :feedback="mimeMapValidationFeedback"
          >
            <NInput
                v-model:value="MimeMap"
                type="textarea"
                rows="11"
                placeholder='{"video/mp4": { "Type": "video","Suffix": ".mp4"}}'
            />
            <NTooltip trigger="hover">
              <template #trigger>
                <NIcon size="18" class="ml-1 text-gray-500">
                  <HelpCircleOutline/>
                </NIcon>
              </template>
              {{ t("setting.mime_map_tip") }}
            </NTooltip>
          </NFormItem>
        </NForm>
        </div>
      </NTabPane>
    </NTabs>
    </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import {HelpCircleOutline} from "@vicons/ionicons5"
import {ref, watch} from "vue"
import {useIndexStore} from "@/stores"
import type {appType} from "@/types/app"
import appApi from "@/api/app"
import {computed} from "vue"
import {useI18n} from 'vue-i18n'
import {isValidHost, isValidPort} from '@/func'
import {NButton, NIcon} from "naive-ui"
import * as bind from "../../wailsjs/go/core/Bind"
import {defaultThemeKey, resolveThemeSeed, resolveThemeOptions, themeImageUrl as resolveThemeImageUrl, themeSeeds} from "@/theme/presets"
import type {ThemePreset, ThemeSeed} from "@/theme/presets"

const {t} = useI18n()
const store = useIndexStore()

const options = computed(() =>
    t("setting.quality_value")
        .split(",")
        .map((value, index) => ({ value: index, label: value }))
)

const formValue = ref<appType.Config>(Object.assign({}, store.globalConfig))
if (!formValue.value.CustomThemes) {
  formValue.value.CustomThemes = {}
}

const MimeMap = ref(formValue.value.MimeMap ? JSON.stringify(formValue.value.MimeMap, null, 2) : "")
const renderKey = ref(999)
const themeOptions = computed(() => {
  store.globalConfig.CustomThemes
  return resolveThemeOptions(store.globalConfig.ThemeMode)
})
const themeImageOptions = computed(() =>
  Object.values(themeSeeds).map(item => ({
    label: item.name,
    value: item.image,
  }))
)
const customThemeSelectOptions = computed(() =>
  Object.values(formValue.value.CustomThemes ?? {}).map(item => ({
    label: item.name,
    value: item.key,
  }))
)

const hostValidationFeedback = ref("")
const portValidationFeedback = ref("")
const mimeMapValidationFeedback = ref("")
let configSaveTimer: ReturnType<typeof window.setTimeout> | undefined

type CustomColorKey = "primary" | "secondary" | "accent" | "warning" | "mollyHair" | "mollyMouth"

const customColorFields: Array<{ key: CustomColorKey; label: string }> = [
  {key: "primary", label: "主色"},
  {key: "secondary", label: "辅助色"},
  {key: "accent", label: "强调色"},
  {key: "warning", label: "提示色"},
  {key: "mollyHair", label: "Molly 发色"},
  {key: "mollyMouth", label: "Molly 嘴唇"},
]

const makeCustomTheme = (seed?: ThemeSeed): appType.CustomTheme => {
  const base = seed ?? themeSeeds[defaultThemeKey]
  return {
    key: `customTheme_${Date.now()}`,
    name: `我的${base.name}`,
    desc: base.desc,
    family: "custom",
    primary: base.primary,
    secondary: base.secondary,
    accent: base.accent,
    warning: base.warning,
    mollyHair: base.mollyHair,
    mollyMouth: base.mollyMouth,
    image: base.image,
  }
}

const editingCustomThemeKey = ref("")
const customThemeForm = ref<appType.CustomTheme>(makeCustomTheme())

const isHexColor = (value: string) => /^#[0-9a-fA-F]{6}$/.test(value)

const customThemePreview = computed(() => {
  return resolveThemeSeed(customThemeForm.value, store.globalConfig.ThemeMode)
})

const persistFormConfig = () => {
  formValue.value.Port = formValue.value.Port.trim()
  formValue.value.Host = formValue.value.Host.trim()

  if (!isValidHost(formValue.value.Host)) {
    hostValidationFeedback.value = t("setting.host_format_error")
    return
  } else {
    hostValidationFeedback.value = ''
  }

  if (!isValidPort(parseInt(formValue.value.Port))) {
    portValidationFeedback.value = t("setting.port_format_error")
    return
  } else {
    portValidationFeedback.value = ''
  }
  store.setConfig({...formValue.value})
}

watch(formValue, () => {
  if (configSaveTimer) {
    window.clearTimeout(configSaveTimer)
  }
  configSaveTimer = window.setTimeout(persistFormConfig, 120)
}, {deep: true})

watch(MimeMap, () => {
  const value = MimeMap.value.trim()

  if (!value) {
    mimeMapValidationFeedback.value = ""
    formValue.value.MimeMap = {}
    store.setConfig({MimeMap: {}})
    return
  }

  try {
    const parsed = JSON.parse(value)
    if (!parsed || Array.isArray(parsed) || typeof parsed !== "object") {
      mimeMapValidationFeedback.value = "MimeMap 必须是 JSON 对象"
      return
    }
    mimeMapValidationFeedback.value = ""
    formValue.value.MimeMap = parsed
    store.setConfig({MimeMap: parsed})
  } catch {
    mimeMapValidationFeedback.value = "JSON 格式不正确"
  }
})

watch(() => {
  return store.globalConfig.Theme
}, () => {
  formValue.value.Theme = store.globalConfig.Theme
  const customTheme = store.globalConfig.CustomThemes?.[store.globalConfig.Theme]
  if (customTheme) {
    editingCustomThemeKey.value = customTheme.key
    customThemeForm.value = {...customTheme}
  }
})

watch(() => store.globalConfig.ThemeMode, () => {
  formValue.value.ThemeMode = store.globalConfig.ThemeMode
})

watch(() => store.globalConfig.Locale, () => {
  formValue.value.Locale = store.globalConfig.Locale
  renderKey.value++
})

const selectDir = () => {
  appApi.openDirectoryDialog().then((res: any) => {
    if (res.code === 1) {
      formValue.value.SaveDirectory = res.data.folder
    }
  }).catch((err: any) => {
    window?.$message?.error(err)
  })
}

const resetHandle = ()=>{
  localStorage.clear()
  bind.ResetApp()
}

const setTheme = (themeKey: string) => {
  formValue.value.Theme = themeKey
  store.setConfig({Theme: themeKey})
}

const getCurrentSeed = () => {
  return Object.values(themeSeeds).find(item => item.key === formValue.value.Theme) ?? themeSeeds[defaultThemeKey]
}

const resetCustomThemeForm = () => {
  const activeCustom = formValue.value.CustomThemes?.[formValue.value.Theme]
  customThemeForm.value = activeCustom ? {...activeCustom} : makeCustomTheme(getCurrentSeed())
  editingCustomThemeKey.value = activeCustom?.key ?? ""
}

const loadCustomTheme = (themeKey: string | null) => {
  if (!themeKey) {
    resetCustomThemeForm()
    return
  }

  const theme = formValue.value.CustomThemes?.[themeKey]
  if (theme) {
    editingCustomThemeKey.value = theme.key
    customThemeForm.value = {...theme}
  }
}

const saveCustomTheme = () => {
  const invalid = customColorFields.some(field => !isHexColor(customThemeForm.value[field.key]))
  if (invalid) {
    window?.$message?.error("颜色必须是 #RRGGBB 格式")
    return
  }

  const key = editingCustomThemeKey.value || customThemeForm.value.key || `customTheme_${Date.now()}`
  const theme: appType.CustomTheme = {
    ...customThemeForm.value,
    key,
    family: "custom",
    name: customThemeForm.value.name.trim() || "自定义主题",
    desc: customThemeForm.value.desc.trim() || "高度客制化主题",
  }
  const customThemes = {
    ...(formValue.value.CustomThemes ?? {}),
    [key]: theme,
  }

  editingCustomThemeKey.value = key
  customThemeForm.value = {...theme}
  formValue.value.CustomThemes = customThemes
  formValue.value.Theme = key
  store.setConfig({CustomThemes: customThemes, Theme: key})
}

const deleteCustomTheme = () => {
  const key = editingCustomThemeKey.value
  if (!key) return

  const customThemes = {...(formValue.value.CustomThemes ?? {})}
  delete customThemes[key]
  const nextTheme = formValue.value.Theme === key ? defaultThemeKey : formValue.value.Theme

  formValue.value.CustomThemes = customThemes
  formValue.value.Theme = nextTheme
  store.setConfig({CustomThemes: customThemes, Theme: nextTheme})
  editingCustomThemeKey.value = ""
  customThemeForm.value = makeCustomTheme(themeSeeds[defaultThemeKey])
}

const themeImageUrl = (item: ThemePreset) =>
  resolveThemeImageUrl(item.image, store.baseUrl)

const selectCustomThemeImage = () => {
  appApi.openImageFileDialog().then((res: appType.Res) => {
    if (res.code === 0) {
      window?.$message?.error(res.message)
      return
    }
    if (res.data?.file) {
      customThemeForm.value.image = res.data.file
    }
  })
}

const themeCardStyle = (item: ThemePreset) => ({
  "--theme-primary": item.primary,
  "--theme-accent": item.accent,
  "--theme-surface": item.surface,
  "--theme-outline": item.outline,
  "--theme-swatch": item.swatch,
})
</script>
<style lang="scss">
.n-tabs-nav--top{
  @apply sticky top-0 z-10;
  background-color: var(--md-surface);
}

.setting-form-panel {
  padding: 18px 4px 4px;
}

.setting-form-panel .n-form-item {
  padding: 10px 12px;
  margin-bottom: 8px;
  border-radius: 8px;
  background: var(--md-surface-container);
  border: 1px solid transparent;
}

.setting-form-panel .n-form-item:hover {
  border-color: color-mix(in srgb, var(--md-primary) 20%, var(--md-outline));
}

.theme-setting-row {
  display: grid;
  grid-template-columns: minmax(190px, 0.55fr) minmax(520px, 1.45fr);
  gap: 16px;
  align-items: center;
  padding: 14px 12px;
  margin-bottom: 12px;
  border-radius: 8px;
  background:
    radial-gradient(circle at 92% 18%, color-mix(in srgb, var(--molly-hair) 28%, transparent) 0 36px, transparent 38px),
    color-mix(in srgb, var(--md-surface-container) 78%, var(--molly) 10%);
  border: 1px solid color-mix(in srgb, var(--md-outline) 78%, var(--molly-mouth) 22%);
}

.theme-setting-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(154px, 1fr));
  gap: 14px;
}

.theme-setting-card {
  min-height: 306px;
  aspect-ratio: 0.68;
  border-radius: 10px;
  border: 1px solid color-mix(in srgb, var(--theme-outline) 70%, var(--theme-primary) 30%);
  background:
    linear-gradient(90deg, color-mix(in srgb, var(--theme-primary) 42%, transparent) 0 1px, transparent 1px 12px),
    linear-gradient(0deg, color-mix(in srgb, var(--theme-outline) 28%, transparent) 0 1px, transparent 1px 12px),
    radial-gradient(circle at 48% 36%, color-mix(in srgb, var(--theme-primary) 22%, transparent) 0 62px, transparent 64px),
    linear-gradient(155deg, color-mix(in srgb, var(--theme-surface) 95%, var(--theme-primary) 5%) 0%, color-mix(in srgb, var(--theme-surface) 76%, var(--theme-accent) 16%) 100%);
  background-size: auto, auto, auto, 100% 100%;
  color: var(--md-text);
  display: grid;
  grid-template-rows: 22px minmax(132px, 1fr) minmax(108px, auto) 7px;
  gap: 7px;
  padding: 10px;
  overflow: hidden;
  position: relative;
  text-align: left;
  isolation: isolate;
  box-shadow:
    inset 0 0 0 1px color-mix(in srgb, white 38%, transparent),
    0 8px 18px color-mix(in srgb, var(--theme-primary) 12%, transparent);
  transition: transform 180ms ease, border-color 180ms ease, box-shadow 200ms ease;
}

.theme-setting-card::before {
  content: "";
  position: absolute;
  inset: 7px;
  border-radius: 8px;
  border: 1px solid color-mix(in srgb, var(--theme-outline) 58%, transparent);
  pointer-events: none;
  z-index: 1;
}

.theme-setting-card::after {
  content: "";
  position: absolute;
  inset: 0;
  background:
    linear-gradient(120deg, transparent 0 34%, color-mix(in srgb, white 24%, transparent) 38%, transparent 44%),
    radial-gradient(circle at 18% 12%, color-mix(in srgb, var(--theme-accent) 24%, transparent) 0 28px, transparent 30px);
  opacity: 0.75;
  pointer-events: none;
  z-index: 1;
}

.theme-setting-card:hover,
.theme-setting-card-active {
  transform: translateY(-4px) rotate(-0.5deg);
  border-color: color-mix(in srgb, var(--theme-primary) 62%, var(--theme-outline));
  box-shadow:
    inset 0 0 0 1px color-mix(in srgb, white 46%, transparent),
    0 18px 38px color-mix(in srgb, var(--theme-primary) 24%, transparent);
}

.theme-card-number {
  align-self: center;
  justify-self: start;
  padding: 2px 7px;
  border-radius: 999px;
  background: color-mix(in srgb, var(--theme-primary) 18%, var(--theme-surface));
  border: 1px solid color-mix(in srgb, var(--theme-primary) 32%, var(--theme-outline));
  color: color-mix(in srgb, var(--md-text) 88%, var(--theme-primary) 12%);
  font-size: 10px;
  font-weight: 800;
  letter-spacing: 0;
  line-height: 1.2;
  position: relative;
  z-index: 2;
}

.theme-card-check {
  position: absolute;
  right: 13px;
  top: 13px;
  width: 19px;
  height: 19px;
  border-radius: 999px;
  background: color-mix(in srgb, var(--theme-surface) 82%, transparent);
  border: 1px solid color-mix(in srgb, var(--theme-outline) 70%, transparent);
  z-index: 3;
}

.theme-setting-card-active .theme-card-check {
  background: var(--theme-primary);
  border-color: var(--theme-primary);
  box-shadow: 0 0 0 4px color-mix(in srgb, var(--theme-primary) 18%, transparent);
}

.theme-setting-card-active .theme-card-check::after {
  content: "";
  position: absolute;
  left: 5px;
  top: 3px;
  width: 6px;
  height: 9px;
  border: solid white;
  border-width: 0 2px 2px 0;
  transform: rotate(45deg);
}

.theme-card-art {
  min-height: 0;
  position: relative;
  display: grid;
  place-items: center;
  pointer-events: none;
  z-index: 2;
}

.theme-card-stage {
  position: absolute;
  width: 78%;
  aspect-ratio: 1;
  border-radius: 999px;
  background:
    radial-gradient(circle at 50% 42%, color-mix(in srgb, var(--theme-surface) 70%, transparent) 0 34%, transparent 36%),
    conic-gradient(from 220deg, color-mix(in srgb, var(--theme-primary) 74%, white 10%), color-mix(in srgb, var(--theme-accent) 76%, white 8%), color-mix(in srgb, var(--theme-primary) 74%, white 10%));
  opacity: 0.92;
  box-shadow:
    inset 0 0 0 1px color-mix(in srgb, white 42%, transparent),
    0 12px 24px color-mix(in srgb, var(--theme-primary) 18%, transparent);
}

.theme-card-stage::after {
  content: "";
  position: absolute;
  left: 18%;
  right: 18%;
  bottom: 9%;
  height: 10px;
  border-radius: 999px;
  background: color-mix(in srgb, var(--theme-primary) 28%, transparent);
  filter: blur(6px);
}

.theme-card-image {
  position: relative;
  z-index: 2;
  width: min(96%, 142px);
  height: min(100%, 140px);
  object-fit: contain;
  object-position: center bottom;
  filter:
    drop-shadow(0 11px 13px color-mix(in srgb, var(--theme-primary) 20%, transparent))
    drop-shadow(0 1px 0 color-mix(in srgb, var(--theme-surface) 80%, transparent));
  transform: translateY(4px);
}

.theme-card-body {
  min-width: 0;
  min-height: 106px;
  display: flex;
  flex-direction: column;
  justify-content: flex-end;
  position: relative;
  z-index: 2;
  padding: 9px;
  border-radius: 8px;
  background:
    linear-gradient(180deg, color-mix(in srgb, var(--theme-surface) 86%, transparent), color-mix(in srgb, var(--theme-surface) 72%, var(--theme-primary) 8%));
  border: 1px solid color-mix(in srgb, var(--theme-outline) 52%, transparent);
}

.theme-card-series {
  display: block;
  margin-bottom: 4px;
  color: color-mix(in srgb, var(--md-text-muted) 82%, var(--theme-primary) 18%);
  font-size: 9px;
  line-height: 1.1;
  text-transform: uppercase;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.theme-card-swatch {
  display: block;
  width: 100%;
  height: 7px;
  border-radius: 999px;
  opacity: 0.9;
  flex-shrink: 0;
  box-shadow: 0 0 0 1px color-mix(in srgb, var(--theme-outline) 50%, transparent);
  position: relative;
  z-index: 2;
}

.theme-card-name {
  display: block;
  font-size: 14px;
  font-weight: 800;
  line-height: 1.15;
  max-width: 100%;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.theme-card-desc {
  display: -webkit-box;
  margin-top: 4px;
  color: var(--md-text-muted);
  font-size: 11px;
  line-height: 1.28;
  overflow: hidden;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
}

.theme-card-spec {
  display: flex;
  justify-content: space-between;
  gap: 6px;
  margin-top: 6px;
  color: color-mix(in srgb, var(--md-text-muted) 76%, var(--theme-primary) 24%);
  font-size: 9px;
  font-weight: 800;
  line-height: 1;
}

.custom-theme-lab {
  display: grid;
  grid-template-columns: minmax(170px, 0.42fr) minmax(0, 1fr);
  gap: 16px;
  align-items: stretch;
  padding: 14px 12px;
  margin-bottom: 12px;
  border-radius: 8px;
  background:
    radial-gradient(circle at 8% 18%, color-mix(in srgb, var(--md-primary) 18%, transparent) 0 52px, transparent 54px),
    linear-gradient(145deg, color-mix(in srgb, var(--md-surface-container) 86%, var(--molly) 8%), var(--md-surface-container));
  border: 1px solid color-mix(in srgb, var(--md-outline) 78%, var(--md-primary) 22%);
}

.custom-theme-preview {
  width: 100%;
  min-height: 318px;
  aspect-ratio: 0.68;
  border-radius: 10px;
  border: 1px solid color-mix(in srgb, var(--theme-outline) 70%, var(--theme-primary) 30%);
  background:
    linear-gradient(90deg, color-mix(in srgb, var(--theme-primary) 42%, transparent) 0 1px, transparent 1px 12px),
    linear-gradient(0deg, color-mix(in srgb, var(--theme-outline) 28%, transparent) 0 1px, transparent 1px 12px),
    radial-gradient(circle at 48% 36%, color-mix(in srgb, var(--theme-primary) 22%, transparent) 0 62px, transparent 64px),
    linear-gradient(155deg, color-mix(in srgb, var(--theme-surface) 95%, var(--theme-primary) 5%) 0%, color-mix(in srgb, var(--theme-surface) 76%, var(--theme-accent) 16%) 100%);
  display: grid;
  grid-template-rows: 22px minmax(132px, 1fr) minmax(108px, auto) 7px;
  gap: 7px;
  padding: 10px;
  overflow: hidden;
  position: relative;
  isolation: isolate;
  box-shadow:
    inset 0 0 0 1px color-mix(in srgb, white 38%, transparent),
    0 14px 28px color-mix(in srgb, var(--theme-primary) 18%, transparent);
}

.custom-theme-preview::before {
  content: "";
  position: absolute;
  inset: 7px;
  border-radius: 8px;
  border: 1px solid color-mix(in srgb, var(--theme-outline) 58%, transparent);
  pointer-events: none;
  z-index: 1;
}

.custom-theme-editor {
  min-width: 0;
  display: flex;
  flex-direction: column;
  gap: 14px;
}

.custom-theme-editor-head {
  display: flex;
  gap: 12px;
  align-items: flex-start;
  justify-content: space-between;
}

.custom-theme-select {
  width: min(260px, 42%);
  flex-shrink: 0;
}

.custom-theme-fields {
  display: grid;
  grid-template-columns: repeat(2, minmax(0, 1fr));
  gap: 10px;
}

.custom-field,
.custom-color-field {
  min-width: 0;
  display: grid;
  gap: 6px;
  color: var(--md-text-muted);
  font-size: 12px;
  line-height: 1.2;
}

.custom-field.wide {
  grid-column: span 2;
}

.custom-color-control {
  display: grid;
  grid-template-columns: 42px minmax(0, 1fr);
  gap: 8px;
  align-items: center;
}

.custom-image-picker {
  display: grid;
  grid-template-columns: minmax(0, 1fr) auto;
  gap: 8px;
  align-items: center;
}

.custom-color-control input[type="color"] {
  width: 42px;
  height: 34px;
  padding: 0;
  border: 1px solid var(--md-outline);
  border-radius: 8px;
  background: var(--md-surface);
  cursor: pointer;
}

.custom-theme-actions {
  display: flex;
  flex-wrap: wrap;
  gap: 8px;
  justify-content: flex-end;
}

@media (max-width: 900px) {
  .theme-setting-row {
    grid-template-columns: 1fr;
  }

  .theme-setting-grid {
    grid-template-columns: repeat(2, minmax(0, 1fr));
  }

  .custom-theme-lab {
    grid-template-columns: 1fr;
  }

  .custom-theme-preview {
    width: min(220px, 100%);
    justify-self: center;
  }

  .custom-theme-editor-head,
  .custom-theme-fields {
    grid-template-columns: 1fr;
  }

  .custom-theme-select,
  .custom-field.wide {
    width: 100%;
    grid-column: auto;
  }
}
</style>
