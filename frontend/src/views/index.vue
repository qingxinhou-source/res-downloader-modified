<template>
  <div class="h-full flex flex-col overflow-hidden">
    <div class="z-40 px-5 pt-5 pb-3" id="header">
      <div class="grid grid-cols-1 xl:grid-cols-[1fr_360px] gap-3" style="--wails-draggable:no-drag">
        <div class="soft-panel p-4">
          <div class="flex flex-wrap items-center justify-between gap-3">
            <div class="min-w-0">
              <div class="flex items-center gap-2 text-xs uppercase tracking-normal text-[var(--md-text-muted)]">
                <span class="theme-dot"></span>
                Resource Monitor
              </div>
              <div class="mt-2 text-2xl font-bold text-[var(--md-text)]">{{ isProxy ? t("index.close_grab") : t("index.open_grab") }}</div>
              <div class="mt-1 text-sm text-[var(--md-text-muted)]">{{ store.envInfo.platform || 'desktop' }}</div>
            </div>
            <div class="flex flex-wrap items-center gap-2">
              <div v-if="personalization.ShowClock" class="live-clock">
                <span>{{ currentTime }}</span>
                <small>{{ currentDate }}</small>
              </div>
              <NTooltip trigger="hover">
                <template #trigger>
                  <NButton
                      circle
                      secondary
                      :type="isDarkTheme ? 'warning' : 'primary'"
                      class="theme-mode-toggle"
                      @click.stop="toggleThemeMode"
                  >
                    <template #icon>
                      <NIcon size="18">
                        <SunnyOutline v-if="isDarkTheme"/>
                        <MoonOutline v-else/>
                      </NIcon>
                    </template>
                  </NButton>
                </template>
                {{ isDarkTheme ? '切换到日间模式' : '切换到夜间模式' }}
              </NTooltip>
              <NButton secondary :type="editMode ? 'warning' : 'default'" @click.stop="toggleEditMode">
                {{ editMode ? '完成编辑' : '编辑页面' }}
              </NButton>
              <NButton secondary type="info" @click.stop="openGuide">
                新手引导
              </NButton>
              <NButton v-if="isProxy" strong type="error" @click.stop="close">
                <span class="inline-block w-2 h-2 bg-white rounded-full mr-1.5 animate-pulse"></span>
                {{ t("index.close_grab") }}
              </NButton>
              <NButton v-else strong type="primary" @click.stop="open">
                {{ t("index.open_grab") }}
              </NButton>
              <NSelect class="min-w-[190px]" :placeholder="t('index.grab_type')" v-model:value="resourcesType" multiple clearable
                       :max-tag-count="3" :options="classify"></NSelect>
            </div>
          </div>
          <div class="mt-4 flex flex-wrap gap-2 items-center">
          <NButtonGroup>

          <NButton v-if="rememberChoice" tertiary type="error" @click.stop="clear">
            <template #icon>
              <n-icon>
                <TrashOutline/>
              </n-icon>
            </template>
            {{ t("index.clear_list") }}
          </NButton>
          <n-popconfirm
              v-else
              @positive-click="()=>{rememberChoice=rememberChoiceTmp;clear()}"
              :show-icon="false"
          >
            <template #trigger>
              <NButton tertiary type="error">
                <template #icon>
                  <n-icon>
                    <TrashOutline/>
                  </n-icon>
                </template>
                {{ t("index.clear_list") }}
              </NButton>
            </template>
            <div>
              <div class="flex flex-row items-center text-red-700 my-2 text-base">
                <n-icon>
                  <TrashOutline/>
                </n-icon>
                <p class="ml-1">{{ t("index.clear_list_tip") }}</p>
              </div>
              <NCheckbox
                  v-model:checked="rememberChoiceTmp"
              >
                <span class="text-gray-400">{{ t('index.remember_clear_choice') }}</span>
              </NCheckbox>
            </div>
          </n-popconfirm>

          <NButton tertiary type="primary" @click.stop="batchDown">
            <template #icon>
              <n-icon>
                <DownloadOutline/>
              </n-icon>
            </template>
            {{ t('index.batch_download') }}
          </NButton>
          <NButton tertiary type="info">
            <NPopover placement="bottom" trigger="hover">
              <template #trigger>
                <NIcon size="18" class="">
                  <Apps/>
                </NIcon>
              </template>
              <div class="flex flex-col">
                <NButton tertiary type="error" @click.stop="batchCancel" class="my-1">
                  <template #icon>
                    <n-icon>
                      <CloseOutline/>
                    </n-icon>
                  </template>
                  {{ t('index.cancel_down') }}
                </NButton>
                <NButton tertiary type="warning" @click.stop="batchExport()" class="my-1">
                  <template #icon>
                    <n-icon>
                      <ArrowRedoCircleOutline/>
                    </n-icon>
                  </template>
                  {{ t('index.batch_export') }}
                </NButton>
                <NButton tertiary type="info" @click.stop="showImport=true" class="my-1">
                  <template #icon>
                    <n-icon>
                      <ServerOutline/>
                    </n-icon>
                  </template>
                  {{ t('index.batch_import') }}
                </NButton>
                <NButton tertiary type="primary" @click.stop="batchExport('url')" class="my-1">
                  <template #icon>
                    <n-icon>
                      <ArrowRedoCircleOutline/>
                    </n-icon>
                  </template>
                  {{ t('index.export_url') }}
                </NButton>
              </div>
            </NPopover>
          </NButton>
          </NButtonGroup>
            <span class="material-chip" v-if="checkedRowKeysValue.length > 0">{{ t("index.choice") }} {{ checkedRowKeysValue.length }}</span>
            <span class="material-chip" v-if="downloadQueue.length > 0">{{ t("index.pending") }} {{ downloadQueue.length }}</span>
          </div>
        </div>
        <div v-if="personalization.ShowMetrics" class="grid grid-cols-3 gap-3">
          <div class="metric-card">
            <div class="metric-label">资源总数</div>
            <div class="metric-value">{{ data.length }}</div>
          </div>
          <div class="metric-card">
            <div class="metric-label">{{ t("index.running") }}</div>
            <div class="metric-value">{{ runningCount }}</div>
          </div>
          <div class="metric-card">
            <div class="metric-label">{{ t("index.done") }}</div>
            <div class="metric-value">{{ doneCount }}</div>
          </div>
        </div>
      </div>
    </div>
    <div class="flex-1 min-h-0 px-5">
      <div class="soft-panel overflow-hidden h-full relative" style="--wails-draggable:no-drag">
        <div v-if="editMode" class="page-editor-panel">
          <div class="editor-title">页面编辑</div>
          <label><input v-model="personalization.ShowMetrics" type="checkbox" @change="persistPersonalization" /> 显示统计卡</label>
          <label><input v-model="personalization.ShowFooterLinks" type="checkbox" @change="persistPersonalization" /> 显示底部链接</label>
          <label><input v-model="personalization.ShowClock" type="checkbox" @change="persistPersonalization" /> 显示时间</label>
          <label><input v-model="personalization.ShowMollyCompanion" type="checkbox" @change="persistPersonalization" /> 显示 Molly 挂件</label>
          <label><input v-model="personalization.ShowMollyBackground" type="checkbox" @change="persistPersonalization" /> 显示背景角色</label>
          <label><input v-model="personalization.VisualFeedback" type="checkbox" @change="persistPersonalization" /> 启用视觉反馈</label>
          <div class="editor-actions">
            <NButton size="tiny" secondary @click="addCustomWidget('note')">添加便签</NButton>
            <NButton size="tiny" secondary @click="addCustomWidget('clock')">添加时钟</NButton>
            <NButton size="tiny" secondary @click="addCustomWidget('action')">添加功能按钮</NButton>
          </div>
        </div>
        <button
            v-for="widget in visibleWidgets"
            :key="widget.id"
            class="page-widget"
            :class="{ 'page-widget-editing': editMode }"
            :style="{ left: `${widget.x}px`, top: `${widget.y}px` }"
            @pointerdown="startWidgetDrag($event, widget)"
            @click.stop="runWidget(widget)"
        >
          <span class="page-widget-title">{{ widget.title }}</span>
          <span class="page-widget-content">{{ widget.type === 'clock' ? currentTime : widget.content }}</span>
          <span
              v-if="editMode"
              class="page-widget-remove"
              @pointerdown.stop.prevent
              @click.stop.prevent="removeWidget(widget.id)"
          >
            ×
          </span>
        </button>
        <div v-if="filteredData.length === 0" class="absolute inset-0 flex items-center justify-center pointer-events-none z-10">
          <div class="empty-hint">
            <div class="empty-mark" aria-hidden="true">20</div>
            <div class="mt-3 text-base font-semibold text-[var(--md-text)]">{{ t('index.total_resources', {count: 0}) }}</div>
            <div class="mt-1 text-sm text-[var(--md-text-muted)]">打开抓取后，资源会在这里排队出现。</div>
          </div>
        </div>
        <NDataTable
            :columns="columns"
            :data="filteredData"
            :bordered="false"
            :max-height="tableHeight"
            :row-key="rowKey"
            :virtual-scroll="true"
            :header-height="48"
            :height-for-row="()=> 48"
            :checked-row-keys="checkedRowKeysValue"
            @update:checked-row-keys="handleCheck"
            @update:filters="updateFilters"
        />
      </div>
    </div>
    <div v-if="personalization.ShowFooterLinks" class="mx-5 my-3 soft-panel flex flex-wrap justify-center items-center gap-1 text-[var(--md-primary)] text-sm px-3 py-2" id="bottom">
      <span class="cursor-pointer px-3 py-1.5 rounded-full hover:bg-[var(--md-surface-container)] transition-colors" @click="BrowserOpenURL(certUrl)">{{ t('footer.cert_download') }}</span>
      <span class="cursor-pointer px-3 py-1.5 rounded-full hover:bg-[var(--md-surface-container)] transition-colors" @click="BrowserOpenURL('https://www.apache.org/licenses/LICENSE-2.0')">Apache License 2.0</span>
    </div>
    <Preview v-model:showModal="showPreviewRow" :previewRow="previewRow"/>
    <ShowLoading :loadingText="loadingText" :isLoading="loading"/>
    <ImportJson v-model:showModal="showImport" @submit="handleImport"/>
    <Password v-model:showModal="showPassword" @submit="handlePassword"/>
    <NModal
        v-model:show="showGuide"
        preset="card"
        class="w-[760px] max-w-[92vw]"
        title="欢迎使用 Res Downloader Modified"
        style="--wails-draggable:no-drag"
    >
      <div class="guide-panel">
        <div class="intro-panel">
          <div class="intro-kicker">资源嗅探与下载工具</div>
          <div class="intro-title">它可以帮你发现网页、视频号、小程序或其他应用播放时产生的图片、音频、视频、m3u8 和直播流资源。</div>
          <div class="intro-desc">
            软件通过本地代理分析网络请求，把可下载资源整理成列表，方便预览、复制链接、批量导出或下载。首次使用可能需要安装证书并允许设置系统代理，这是 HTTPS 资源识别和抓取所必需的。
          </div>
        </div>
        <div class="guide-steps">
          <div
              v-for="(step, index) in guideSteps"
              :key="step.title"
              class="guide-step"
          >
            <span class="guide-step-index">{{ index + 1 }}</span>
            <div>
              <div class="guide-step-title">{{ step.title }}</div>
              <div class="guide-step-desc">{{ step.desc }}</div>
            </div>
          </div>
        </div>
        <div class="env-panel">
          <div class="env-head">
            <div>
              <div class="env-title">环境依赖检测</div>
              <div class="env-subtitle">检查代理、证书、保存目录和本地服务状态。</div>
            </div>
            <NButton size="small" secondary type="primary" :loading="envChecking" @click="runEnvCheck">
              重新检测
            </NButton>
          </div>
          <div class="env-list">
            <div
                v-for="item in envChecks"
                :key="item.key"
                class="env-item"
                :class="`env-item-${item.status}`"
            >
              <span class="env-status"></span>
              <div>
                <div class="env-label">{{ item.label }}</div>
                <div class="env-detail">{{ item.detail }}</div>
              </div>
            </div>
          </div>
        </div>
        <div class="guide-actions">
          <NButton secondary @click="BrowserOpenURL(certUrl)">下载证书</NButton>
          <NButton secondary type="primary" @click="open">启动代理</NButton>
          <NButton secondary @click="showGuide = false">我知道了</NButton>
        </div>
      </div>
    </NModal>
  </div>
</template>

<script lang="ts" setup>
import {NButton, NIcon, NImage, NInput, NTooltip, NPopover, NGradientText, NModal} from "naive-ui"
import {computed, h, onMounted, onUnmounted, ref, watch} from "vue"
import type {appType} from "@/types/app"
import type {DataTableRowKey, ImageRenderToolbarProps, DataTableFilterState, DataTableBaseColumn} from "naive-ui"
import Preview from "@/components/Preview.vue"
import ShowLoading from "@/components/ShowLoading.vue"
// @ts-ignore
import {getDecryptionArray} from '@/assets/js/decrypt.js'
import {useIndexStore} from "@/stores"
import appApi from "@/api/app"
import Action from "@/components/Action.vue"
import ActionDesc from "@/components/ActionDesc.vue"
import ImportJson from "@/components/ImportJson.vue"
import {useEventStore} from "@/stores/event"
import {BrowserOpenURL, ClipboardSetText} from "../../wailsjs/runtime"
import Password from "@/components/Password.vue"
import ShowOrEdit from "@/components/ShowOrEdit.vue"
import {useI18n} from 'vue-i18n'
import {
  DownloadOutline,
  ArrowRedoCircleOutline,
  ServerOutline,
  SearchOutline,
  Apps,
  TrashOutline,
  CloseOutline,
  MoonOutline,
  SunnyOutline
} from "@vicons/ionicons5"
import {useDialog} from 'naive-ui'
import * as bind from "../../wailsjs/go/core/Bind"
import {Quit} from "../../wailsjs/runtime"
import {DialogOptions} from "naive-ui/es/dialog/src/DialogProvider"
import {formatSize} from "@/func"

const {t} = useI18n()
const eventStore = useEventStore()
const dialog = useDialog()
const isProxy = computed(() => {
  return store.isProxy
})
const certUrl = computed(() => {
  return store.baseUrl + "/api/cert"
})
const data = ref<any[]>([])
const filterClassify = ref<string[]>([])
const filteredData = computed(() => {
  let result = data.value

  if (filterClassify.value.length > 0) {
    result = result.filter(item => filterClassify.value.includes(item.Classify))
  }

  if (descriptionSearchValue.value) {
    result = result.filter(item => item.Description?.toLowerCase().includes(descriptionSearchValue.value.toLowerCase()))
  }

  if (urlSearchValue.value) {
    result = result.filter(item => item.Url?.toLowerCase().includes(urlSearchValue.value.toLowerCase()))
  }

  return result
})

const runningCount = computed(() => data.value.filter(item => item.Status === "running").length)
const doneCount = computed(() => data.value.filter(item => item.Status === "done").length)

const store = useIndexStore()
const personalization = computed(() => store.globalConfig.Personalization)
const isDarkTheme = computed(() => store.globalConfig.ThemeMode === "dark")
const toggleThemeMode = () => {
  store.setConfig({ ThemeMode: isDarkTheme.value ? "light" : "dark" })
}
const editMode = ref(false)
const now = ref(new Date())
let timeTimer: number | undefined

const currentTime = computed(() =>
  now.value.toLocaleTimeString("zh-CN", {hour: "2-digit", minute: "2-digit", second: "2-digit", hour12: false})
)
const currentDate = computed(() =>
  now.value.toLocaleDateString("zh-CN", {month: "2-digit", day: "2-digit", weekday: "short"})
)
const visibleWidgets = computed(() =>
  (personalization.value.CustomWidgets ?? []).filter(item => item.visible || editMode.value)
)

const persistPersonalization = () => {
  store.setConfig({Personalization: {...personalization.value}})
}

const toggleEditMode = () => {
  editMode.value = !editMode.value
  if (!editMode.value) {
    persistPersonalization()
  }
}

const addCustomWidget = (type: appType.CustomWidget["type"]) => {
  const action = type === "action" ? "toggleProxy" : "none"
  const widget: appType.CustomWidget = {
    id: `widget_${Date.now()}`,
    type,
    title: type === "clock" ? "桌面时钟" : type === "action" ? "快捷按钮" : "灵感便签",
    content: type === "action" ? "启动/关闭代理" : type === "clock" ? "当前时间" : "写点自己的提示",
    action,
    x: 24 + (personalization.value.CustomWidgets?.length ?? 0) * 18,
    y: 24 + (personalization.value.CustomWidgets?.length ?? 0) * 18,
    visible: true,
  }
  personalization.value.CustomWidgets = [...(personalization.value.CustomWidgets ?? []), widget]
  persistPersonalization()
}

const removeWidget = (id: string) => {
  personalization.value.CustomWidgets = (personalization.value.CustomWidgets ?? []).filter(item => item.id !== id)
  persistPersonalization()
}

const runWidget = (widget: appType.CustomWidget) => {
  if (editMode.value) return
  if (widget.action === "toggleProxy") {
    isProxy.value ? close() : open()
  } else if (widget.action === "clearList") {
    clear()
  } else if (widget.action === "toggleThemeMode") {
    toggleThemeMode()
  }
}

const startWidgetDrag = (event: PointerEvent, widget: appType.CustomWidget) => {
  if (!editMode.value) return
  if ((event.target as HTMLElement).closest(".page-widget-remove")) return
  const startX = event.clientX
  const startY = event.clientY
  const originalX = widget.x
  const originalY = widget.y
  const target = event.currentTarget as HTMLElement
  target.setPointerCapture(event.pointerId)

  const move = (moveEvent: PointerEvent) => {
    widget.x = Math.max(8, originalX + moveEvent.clientX - startX)
    widget.y = Math.max(8, originalY + moveEvent.clientY - startY)
  }
  const up = () => {
    target.removeEventListener("pointermove", move)
    target.removeEventListener("pointerup", up)
    persistPersonalization()
  }

  target.addEventListener("pointermove", move)
  target.addEventListener("pointerup", up)
}
const tableHeight = ref(800)
const resourcesType = ref<string[]>(["all"])

const classifyAlias: { [key: string]: any } = {
  image: computed(() => t("index.image")),
  audio: computed(() => t("index.audio")),
  video: computed(() => t("index.video")),
  m3u8: computed(() => t("index.m3u8")),
  live: computed(() => t("index.live")),
  xls: computed(() => t("index.xls")),
  doc: computed(() => t("index.doc")),
  pdf: computed(() => t("index.pdf")),
  stream: computed(() => t("index.stream")),
  font: computed(() => t("index.font"))
}

const dwStatus = computed<any>(() => {
  return {
    ready: t("index.ready"),
    pending: t("index.pending"),
    running: t("index.running"),
    error: t("index.error"),
    done: t("index.done"),
    handle: t("index.handle")
  }
})

const maxConcurrentDownloads = computed(() => {
  return store.globalConfig.DownNumber
})

const classify = ref([
  {
    value: "all",
    label: computed(() => t("index.all")),
  },
])

const descriptionSearchValue = ref("")
const urlSearchValue = ref("")
const rememberChoice = ref(false)
const rememberChoiceTmp = ref(false)

const columns = ref<any[]>([
  {
    type: "selection",
  },
  {
    title: () => {
      if (checkedRowKeysValue.value.length > 0) {
        return h(NGradientText, {type: "success"}, t("index.choice") + `(${checkedRowKeysValue.value.length})`)
      }
      return h('div', {class: 'flex items-center'}, [
        t('index.domain'),
        h(NPopover, {
          style: "--wails-draggable:no-drag",
          trigger: 'click',
          placement: 'bottom',
          showArrow: true,
        }, {
          trigger: () => h(NIcon, {
            size: "18",
            class: `ml-1 cursor-pointer ${urlSearchValue.value ? "text-green-600": "text-gray-500"}`,
            onClick: (e: MouseEvent) => e.stopPropagation()
          }, h(SearchOutline)),
          default: () => h('div', {class: 'p-2 w-64'}, [
            h(NInput, {
              value: urlSearchValue.value,
              'onUpdate:value': (val: string) => urlSearchValue.value = val,
              placeholder: t('index.search_description'),
              clearable: true
            }, {
              prefix: () => h(NIcon, {component: SearchOutline})
            })
          ])
        })
      ])
    },
    key: "Domain",
    width: 90,
    render: (row: appType.MediaInfo) => {
      return h(NTooltip, {
        trigger: 'hover',
        placement: 'top'
      }, {
        trigger: () => h('span', {
          class: 'cursor-default'
        }, row.Domain),
        default: () => row.Url
      })
    }
  },
  {
    title: computed(() => t("index.type")),
    key: "Classify",
    width: 80,
    filterOptions: computed(() => Array.from(classify.value).slice(1)),
    filterMultiple: true,
    filter: (value: string, row: appType.MediaInfo) => {
      return !!~row.Classify.indexOf(String(value))
    },
    render: (row: appType.MediaInfo) => {
      const item = classify.value.find(item => item.value === row.Classify)
      return item ? item.label : row.Classify
    }
  },
  {
    title: computed(() => t("index.preview")),
    key: "Url",
    width: 80,
    render: (row: appType.MediaInfo) => {
      if (row.Classify === "image") {
        return h("div", {
          style: "width: 100%;max-height:80px;overflow:hidden;"
        }, h(NImage, {
          objectFit: "contain",
          lazy: true,
          "render-toolbar": renderToolbar,
          src: row.Url
        }))
      }
      return [
        h(
            NButton,
            {
              strong: true,
              tertiary: true,
              type: "info",
              size: "small",
              style: {
                margin: "2px"
              },
              onClick: () => {
                if (row.Classify === "audio" || row.Classify === "video" || row.Classify === "m3u8" || row.Classify === "live") {
                  previewRow.value = row
                  showPreviewRow.value = true
                }
              }
            },
            {
              default: () => {
                if (row.Classify === "audio" || row.Classify === "video" || row.Classify === "m3u8" || row.Classify === "live") {
                  return t("index.preview")
                }
                return t("index.preview_tip")
              }
            }
        ),
      ]
    }
  },
  {
    title: computed(() => t("index.status")),
    key: "Status",
    width: 80,
    render: (row: appType.MediaInfo, index: number) => {
      let status = "info"
      if (row.Status === "done" || row.Status === "running") {
        status = "success"
      } else if (row.Status === "pending") {
        status = "warning"
      }

      return h(
          NButton,
          {
            tertiary: true,
            type: status as any,
            size: "small",
            style: {
              margin: "2px"
            },
            onClick: () => {
              if (row.SavePath && row.Status === "done") {
                appApi.openFolder({filePath: row.SavePath})
              } else if (row.Status === "ready") {
                download(row, index)
              }
            }
          },
          {
            default: () => {
              return row.Status === "running" ? row.SavePath : dwStatus.value[row.Status as keyof typeof dwStatus]
            }
          }
      )
    }
  },
  {
    title: () => h('div', {class: 'flex items-center'}, [
      t('index.description'),
      h(NPopover, {
        style: "--wails-draggable:no-drag",
        trigger: 'click',
        placement: 'bottom',
        showArrow: true,
      }, {
        trigger: () => h(NIcon, {
          size: "18",
          class: `ml-1 cursor-pointer ${descriptionSearchValue.value ? "text-green-600": "text-gray-500"}`,
          onClick: (e: MouseEvent) => e.stopPropagation()
        }, h(SearchOutline)),
        default: () => h('div', {class: 'p-2 w-64'}, [
          h(NInput, {
            value: descriptionSearchValue.value,
            'onUpdate:value': (val: string) => descriptionSearchValue.value = val,
            placeholder: t('index.search_description'),
            clearable: true
          }, {
            prefix: () => h(NIcon, {component: SearchOutline})
          })
        ])
      })
    ]),
    key: "Description",
    width: 150,
    render: (row: appType.MediaInfo, index: number) => {
      return h(ShowOrEdit, {
        value: row.Description,
        onUpdateValue(v: string) {
          data.value[index].Description = v
          cacheData()
        }
      })
    }
  },
  {
    title: computed(() => t("index.resource_size")),
    key: "Size",
    width: 120,
    sorter: (row1: appType.MediaInfo, row2: appType.MediaInfo) => row1.Size - row2.Size,
    render(row: appType.MediaInfo, index: number) {
      return formatSize(row.Size)
    }
  },
  {
    title: computed(() => t("index.save_path")),
    key: "SavePath",
    render(row: appType.MediaInfo, index: number) {
      return h("a",
          {
            href: "javascript:;",
            class: "ellipsis-2",
            style: {
              color: "#5a95d0"
            },
            onClick: () => {
              if (row.SavePath && row.Status === "done") {
                appApi.openFolder({filePath: row.SavePath})
              }
            }
          },
          row.Status === "running" ? "" : row.SavePath
      )
    }
  },
  {
    key: "actions",
    width: 130,
    render(row: appType.MediaInfo, index: number) {
      return h(Action, {key: index, row: row, index: index, onAction: dataAction})
    },
    title() {
      return h(ActionDesc)
    }
  }
])

const checkedRowKeysValue = ref<DataTableRowKey[]>([])
const showPreviewRow = ref(false)
const previewRow = ref<appType.MediaInfo>()
const loading = ref(false)
const loadingText = ref("")
const showImport = ref(false)
const showPassword = ref(false)
const showGuide = ref(false)
const envChecking = ref(false)
const envChecks = ref<Array<{ key: string; label: string; status: "ok" | "warn"; detail: string }>>([])
const downloadQueue = ref<appType.MediaInfo[]>([])
let activeDownloads = 0
let isOpenProxy = false
let isInstall = false

const guideSteps = [
  {
    title: "确认用途",
    desc: "本工具用于学习研究和下载自己有权访问的资源，请勿用于商业或侵权用途。",
  },
  {
    title: "完成初始化",
    desc: "首次使用请允许证书安装和网络访问；如系统要求权限，请按提示输入密码或以管理员身份运行。",
  },
  {
    title: "设置保存目录",
    desc: "进入系统设置选择资源保存目录，避免下载时因为路径为空而失败。",
  },
  {
    title: "启动代理",
    desc: "回到首页点击“启动代理”，系统代理开启后才能捕获外部页面里的资源请求。",
  },
  {
    title: "抓取资源",
    desc: "在微信视频号、小程序、网页或其他 App 中播放目标内容，资源会自动进入首页列表。",
  },
]

const runEnvCheck = async () => {
  envChecking.value = true
  try {
    const res = await appApi.envCheck()
    if (res.code === 1) {
      envChecks.value = res.data.checks ?? []
    } else {
      window?.$message?.error(res.message)
    }
  } finally {
    envChecking.value = false
  }
}

const openGuide = () => {
  showGuide.value = true
  runEnvCheck()
}

onMounted(() => {
  timeTimer = window.setInterval(() => {
    now.value = new Date()
  }, 1000)
  if (!localStorage.getItem("new-user-guide-seen")) {
    window.setTimeout(() => {
      showGuide.value = true
      runEnvCheck()
      localStorage.setItem("new-user-guide-seen", "1")
    }, 700)
  }
  try {
    window.addEventListener("resize", () => {
      resetTableHeight()
    })
    loading.value = true
    handleInstall().then((is: boolean) => {
      isInstall = true
      loading.value = false
    })

    checkLoading()
    watch(showPassword, () => {
      if (!showPassword.value) {
        checkLoading()
      }
    })
  } catch (e) {
    window.$message?.error(JSON.stringify(e), {duration: 5000})
  }

  buildClassify()

  const temp = localStorage.getItem("resources-type")
  if (temp) {
    resourcesType.value = JSON.parse(temp).res
  } else {
    appApi.setType(resourcesType.value)
  }

  const cache = localStorage.getItem("resources-data")
  if (cache) {
    data.value = JSON.parse(cache)
  }

  const choiceCache = localStorage.getItem("remember-clear-choice")
  if (choiceCache === "1") {
    rememberChoice.value = true
  }

  watch(rememberChoice, (n, o) => {
    if (rememberChoice.value) {
      localStorage.setItem("remember-clear-choice", "1")
    } else {
      localStorage.removeItem("remember-clear-choice")
    }
  })

  resetTableHeight()

  eventStore.addHandle({
    type: "newResources",
    event: (res: appType.MediaInfo) => {
      if (store.globalConfig.InsertTail) {
        data.value.push(res)
      } else {
        data.value.unshift(res)
      }
      cacheData()
    }
  })

  eventStore.addHandle({
    type: "downloadProgress",
    event: (res: { Id: string, SavePath: string, Status: string, Message: string, Size?: number }) => {
      switch (res.Status) {
        case "running":
          updateItem(res.Id, item => {
            item.SavePath = res.Message
            item.Status = 'running'
          })
          break
        case "done":
          updateItem(res.Id, item => {
            item.SavePath = res.SavePath
            item.Status = 'done'
            if (res.Size !== undefined) {
              item.Size = res.Size
            }
          })
          if (activeDownloads > 0) {
            activeDownloads--
          }
          cacheData()
          checkQueue()
          break
        case "error":
          updateItem(res.Id, item => {
            item.SavePath = res.Message
            item.Status = 'error'
            if (res.Size !== undefined) {
              item.Size = res.Size
            }
          })
          if (activeDownloads > 0) {
            activeDownloads--
          }
          cacheData()
          checkQueue()
          break
      }
    }
  })
})

onUnmounted(() => {
  if (timeTimer) {
    window.clearInterval(timeTimer)
  }
})

watch(() => {
  return store.globalConfig.MimeMap
}, () => {
  buildClassify()
})

watch(resourcesType, (n, o) => {
  localStorage.setItem("resources-type", JSON.stringify({res: resourcesType.value}))
  appApi.setType(resourcesType.value)
})

const updateItem = (id: string, updater: (item: any) => void) => {
  const item = data.value.find(i => i.Id === id)
  if (item) updater(item)
}

function cacheData() {
  localStorage.setItem("resources-data", JSON.stringify(data.value))
}

const resetTableHeight = () => {
  try {
    const headerHeight = document.getElementById("header")?.offsetHeight || 0
    const bottomHeight = document.getElementById("bottom")?.offsetHeight || 0
    // @ts-ignore
    const theadHeight = document.getElementsByClassName("n-data-table-thead")[0]?.offsetHeight || 0
    const height = document.documentElement.clientHeight || window.innerHeight
    tableHeight.value = height - headerHeight - bottomHeight - theadHeight - 48
  } catch (e) {
    console.log(e)
  }
}

const buildClassify = () => {
  const mimeMap = store.globalConfig.MimeMap ?? {}
  const seen = new Set()
  classify.value = [
    {value: "all", label: computed(() => t("index.all"))},
    ...Object.values(mimeMap)
        .filter(({Type}) => {
          if (seen.has(Type)) return false
          seen.add(Type)
          return true
        })
        .map(({Type}) => ({
          value: Type,
          label: classifyAlias[Type] ?? Type,
        })),
  ]
}

const dataAction = (row: appType.MediaInfo, index: number, type: string) => {
  switch (type) {
    case "down":
      download(row, index)
      break
    case "cancel":
      if (row.Status === "pending") {
        const queueIndex = downloadQueue.value.findIndex(item => item.Id === row.Id)
        if (queueIndex !== -1) {
          downloadQueue.value.splice(queueIndex, 1)
        }
        updateItem(row.Id, item => {
          item.Status = 'ready'
          item.SavePath = ''
        })
        cacheData()
      } else if (row.Status === "running") {
        appApi.cancel({id: row.Id}).then((res) => {
          updateItem(row.Id, item => {
            item.Status = 'ready'
            item.SavePath = ''
          })
          if (activeDownloads > 0) {
            activeDownloads--
          }
          cacheData()
          checkQueue()
          if (res.code === 0) {
            window?.$message?.error(res.message)
            return
          }
        })
      }
      break
    case "copy":
      ClipboardSetText(row.Url).then((is: boolean) => {
        if (is) {
          window?.$message?.success(t("common.copy_success"))
        } else {
          window?.$message?.error(t("common.copy_fail"))
        }
      })
      break
    case "json":
      ClipboardSetText(encodeURIComponent(JSON.stringify(row))).then((is: boolean) => {
        if (is) {
          window?.$message?.success(t("common.copy_success"))
        } else {
          window?.$message?.error(t("common.copy_fail"))
        }
      })
      break
    case "open":
      BrowserOpenURL(row.Url)
      break
    case "decode":
      decodeWxFile(row, index)
      break
    case "delete":
      if (row.Status === "pending" || row.Status === "running") {
        window?.$message?.error(t("index.delete_tip"))
        return
      }
      appApi.delete({sign: [row.UrlSign]}).then(() => {
        data.value.splice(index, 1)
        cacheData()
      })
      break
  }
}

const renderToolbar = ({nodes}: ImageRenderToolbarProps) => {
  return [
    nodes.rotateCounterclockwise,
    nodes.rotateClockwise,
    nodes.resizeToOriginalSize,
    nodes.zoomOut,
    nodes.zoomIn,
    nodes.close
  ]
}

const rowKey = (row: appType.MediaInfo) => {
  return row.Id
}

const handleCheck = (rowKeys: DataTableRowKey[]) => {
  checkedRowKeysValue.value = rowKeys
}

const updateFilters = (filters: DataTableFilterState, initiatorColumn: DataTableBaseColumn) => {
  filterClassify.value = filters.Classify as string[]
}

const batchDown = async () => {
  if (checkedRowKeysValue.value.length <= 0) {
    window?.$message?.error(t("index.use_data"))
    return
  }

  if (!store.globalConfig.SaveDirectory) {
    window?.$message?.error(t("index.save_path_empty"))
    return
  }

  data.value.forEach((item, index) => {
    if (checkedRowKeysValue.value.includes(item.Id) && item.Classify !== 'live' && item.Classify !== 'm3u8') {
      download(item, index)
    }
  })

  checkedRowKeysValue.value = []
}

const batchCancel = async () => {
  if (checkedRowKeysValue.value.length <= 0) {
    window?.$message?.error(t("index.use_data"))
    return
  }
  loading.value = true
  const cancelTasks: Promise<any>[] = []
  data.value.forEach((item, index) => {
    if (!checkedRowKeysValue.value.includes(item.Id)) {
      return
    }

    if (item.Status === "pending") {
      const queueIndex = downloadQueue.value.findIndex(qItem => qItem.Id === item.Id)
      if (queueIndex !== -1) {
        downloadQueue.value.splice(queueIndex, 1)
      }
      item.Status = 'ready'
      item.SavePath = ''
      return
    }

    if (item.Status === "running") {
      if (activeDownloads > 0) {
        activeDownloads--
      }
      cancelTasks.push(appApi.cancel({id: item.Id}).then(() => {
        item.Status = 'ready'
        item.SavePath = ''
        checkQueue()
      }))
    }
  })
  await Promise.allSettled(cancelTasks)
  loading.value = false
  checkedRowKeysValue.value = []
  cacheData()
}

const batchExport = (type?: string) => {
  if (checkedRowKeysValue.value.length <= 0) {
    window?.$message?.error(t("index.use_data"))
    return
  }

  if (!store.globalConfig.SaveDirectory) {
    window?.$message?.error(t("index.save_path_empty"))
    return
  }

  loadingText.value = t("common.loading")
  loading.value = true

  let jsonData = data.value.filter(item => checkedRowKeysValue.value.includes(item.Id))

  if (type === "url") {
    jsonData = jsonData.map(item => item.Url)
  } else {
    jsonData = jsonData.map(item => encodeURIComponent(JSON.stringify(item)))
  }

  appApi.batchExport({content: jsonData.join("\n")}).then((res: appType.Res) => {
    loading.value = false
    if (res.code === 0) {
      window?.$message?.error(res.message)
      return
    }
    window?.$message?.success(t("index.import_success"))
    window?.$message?.info(t("index.save_path") + "：" + res.data?.file_name, {
      duration: 5000
    })
  })
}

const uint8ArrayToBase64 = (bytes: any) => {
  return window.btoa(Array.from(bytes, (byte: any) => String.fromCharCode(byte)).join(''))
}

const download = (row: appType.MediaInfo, index: number) => {
  if (!store.globalConfig.SaveDirectory) {
    window?.$message?.error(t("index.save_path_empty"))
    return
  }

  if (data.value.some(item => item.Id === row.Id && item.Status === "running")) {
    return
  }

  if (downloadQueue.value.some(item => item.Id === row.Id || item.Url === row.Url)) {
    return
  }

  if (activeDownloads >= maxConcurrentDownloads.value) {
    row.Status = "pending"
    downloadQueue.value.push(row)
    window?.$message?.info(t("index.download_queued", {count: downloadQueue.value.length}))
    return
  }

  startDownload(row, index)
}

const startDownload = (row: appType.MediaInfo, index: number) => {
  activeDownloads++

  const decodeStr = row.DecodeKey
      ? uint8ArrayToBase64(getDecryptionArray(row.DecodeKey))
      : ""

  appApi.download({...row, decodeStr}).then((res: appType.Res) => {
    if (res.code === 0) {
      window?.$message?.error(res.message)
    }
  })
}

const checkQueue = () => {
  if (downloadQueue.value.length > 0 && activeDownloads < maxConcurrentDownloads.value) {
    const nextItem = downloadQueue.value.shift()
    if (nextItem) {
      const index = data.value.findIndex(item => item.Id === nextItem.Id)
      if (index !== -1) {
        startDownload(nextItem, index)
      }
    }
  }
}

const open = () => {
  isOpenProxy = true
  store.openProxy().then((res: appType.Res) => {
    if (res.code === 1) {
      return
    }

    if (["darwin", "linux"].includes(store.envInfo.platform)) {
      showPassword.value = true
    } else {
      window.$message?.error(res.message)
    }
  })
}

const close = () => {
  store.unsetProxy()
}

const clear = async () => {
  const newData = [] as any[]
  const signs: string[] = []
  if (checkedRowKeysValue.value.length > 0) {
    data.value.forEach((item, index) => {
      if (checkedRowKeysValue.value.includes(item.Id) && item.Status !== "pending" && item.Status !== "running") {
        signs.push(item.UrlSign)
      } else {
        newData.push(item)
      }
    })
    checkedRowKeysValue.value = []
  } else {
    data.value.forEach((item, index) => {
      if (item.Status === "pending" || item.Status === "running") {
        newData.push(item)
      } else {
        signs.push(item.UrlSign)
      }
    })
  }
  await appApi.delete({sign: signs})
  data.value = newData
  cacheData()
}

const decodeWxFile = (row: appType.MediaInfo, index: number) => {
  if (!row.DecodeKey) {
    window?.$message?.error(t("index.video_decode_no"))
    return
  }
  appApi.openFileDialog().then((res: appType.Res) => {
    if (res.code === 0) {
      window?.$message?.error(res.message)
      return
    }
    if (res.data.file) {
      loadingText.value = t("index.video_decode_loading")
      loading.value = true
      appApi.wxFileDecode({
        ...row,
        filename: res.data.file,
        decodeStr: uint8ArrayToBase64(getDecryptionArray(row.DecodeKey))
      }).then((res: appType.Res) => {
        loading.value = false
        if (res.code === 0) {
          window?.$message?.error(res.message)
          return
        }
        data.value[index].SavePath = res.data.save_path
        data.value[index].Status = "done"
        cacheData()
        window?.$message?.success(t("index.video_decode_success"))
      })
    }
  })
}

const handleImport = (content: string) => {
  if (!content) {
    window?.$message?.error(t("view.import_empty"))
    return
  }
  let newItems = [] as any[]
  content.split("\n").forEach((line, index) => {
    try {
      let res = JSON.parse(decodeURIComponent(line))
      if (res && res?.Id) {
        res.Id = res.Id + Math.floor(Math.random() * 100000)
        res.SavePath = ""
        res.Status = "ready"
        newItems.push(res)
      }
    } catch (e) {
      console.log(e)
    }
  })
  if (newItems.length > 0) {
    data.value = [...newItems, ...data.value]
    cacheData()
  }
  showImport.value = false
}

const handlePassword = async (password: string, isCache: boolean) => {
  const res = await appApi.setSystemPassword({password, isCache})
  if (res.code === 0) {
    window.$message?.error(res.message)
    return
  }

  if (isOpenProxy) {
    showPassword.value = false
    store.openProxy()
    return
  }

  handleInstall().then((is: boolean) => {
    if (is) {
      showPassword.value = false
    }
  })
}

const handleInstall = async () => {
  isOpenProxy = false
  const res = await appApi.install()
  if (res.code === 1) {
    store.globalConfig.AutoProxy && store.openProxy()
    return true
  }

  window.$message?.error(res.message, {duration: 5000})

  if (store.envInfo.platform === "windows" && res.message.includes("Access is denied")) {
    window.$message?.error(t("index.win_install_tip"))
  } else if (["darwin", "linux"].includes(store.envInfo.platform)) {
    showPassword.value = true
  }
  return false
}

const checkLoading = () => {
  setTimeout(() => {
    if (loading.value && !isInstall && !showPassword.value) {
      dialog.warning({
        title: t("index.start_err_tip"),
        content: t("index.start_err_content"),
        positiveText: t("index.start_err_positiveText"),
        negativeText: t("index.start_err_negativeText"),
        draggable: false,
        closeOnEsc: false,
        closable: false,
        maskClosable: false,
        onPositiveClick: () => {
          bind.ResetApp()
        },
        onNegativeClick: () => {
          Quit()
        }
      } as DialogOptions)
    }
  }, 6000)
}
</script>
<style scoped>
.theme-mode-toggle {
  width: 40px;
  height: 40px;
  border: 1px solid color-mix(in srgb, var(--md-primary) 26%, var(--md-outline));
  background:
    radial-gradient(circle at 35% 30%, color-mix(in srgb, var(--md-accent) 22%, transparent) 0 11px, transparent 12px),
    color-mix(in srgb, var(--md-surface) 88%, var(--md-primary) 8%);
  box-shadow: var(--md-shadow-1);
  transition: transform 180ms ease, box-shadow 180ms ease, border-color 180ms ease;
}

.theme-mode-toggle:hover {
  transform: translateY(-1px);
  border-color: color-mix(in srgb, var(--md-primary) 48%, var(--md-outline));
  box-shadow: var(--md-shadow-2);
}

.live-clock {
  min-width: 112px;
  height: 40px;
  padding: 5px 11px;
  border-radius: 999px;
  display: grid;
  align-content: center;
  background:
    radial-gradient(circle at 14% 30%, color-mix(in srgb, var(--molly-hair) 22%, transparent) 0 13px, transparent 14px),
    color-mix(in srgb, var(--md-surface) 86%, var(--md-primary) 8%);
  border: 1px solid color-mix(in srgb, var(--md-primary) 24%, var(--md-outline));
  box-shadow: var(--md-shadow-1);
  color: var(--md-text);
  line-height: 1;
  text-align: right;
}

.live-clock span {
  font-size: 14px;
  font-weight: 850;
}

.live-clock small {
  margin-top: 3px;
  color: var(--md-text-muted);
  font-size: 10px;
}

.page-editor-panel {
  position: absolute;
  right: 14px;
  top: 14px;
  z-index: 40;
  width: 250px;
  padding: 12px;
  border-radius: 10px;
  display: grid;
  gap: 9px;
  background:
    radial-gradient(circle at 92% 12%, color-mix(in srgb, var(--md-accent) 18%, transparent) 0 34px, transparent 36px),
    color-mix(in srgb, var(--md-surface) 92%, var(--molly) 8%);
  border: 1px solid color-mix(in srgb, var(--md-primary) 28%, var(--md-outline));
  box-shadow: var(--md-shadow-2);
  color: var(--md-text-muted);
  font-size: 12px;
}

.editor-title {
  color: var(--md-text);
  font-size: 14px;
  font-weight: 850;
}

.page-editor-panel label {
  display: flex;
  align-items: center;
  gap: 7px;
}

.page-editor-panel input {
  accent-color: var(--md-primary);
}

.editor-actions {
  display: flex;
  flex-wrap: wrap;
  gap: 6px;
  padding-top: 3px;
}

.page-widget {
  position: absolute;
  z-index: 30;
  width: 160px;
  min-height: 74px;
  padding: 11px;
  border-radius: 10px;
  border: 1px solid color-mix(in srgb, var(--md-primary) 28%, var(--md-outline));
  background:
    radial-gradient(circle at 86% 18%, color-mix(in srgb, var(--md-accent) 18%, transparent) 0 24px, transparent 26px),
    linear-gradient(145deg, color-mix(in srgb, var(--md-surface) 88%, var(--molly) 10%), var(--md-surface));
  color: var(--md-text);
  box-shadow: var(--md-shadow-1);
  text-align: left;
  cursor: pointer;
  transition: transform 160ms ease, box-shadow 180ms ease, border-color 160ms ease;
}

.page-widget:hover {
  transform: translateY(-2px);
  border-color: color-mix(in srgb, var(--md-primary) 48%, var(--md-outline));
  box-shadow: var(--md-shadow-2);
}

.page-widget-editing {
  cursor: grab;
  outline: 1px dashed color-mix(in srgb, var(--md-primary) 52%, transparent);
  outline-offset: 3px;
}

.page-widget-editing:active {
  cursor: grabbing;
}

.page-widget-title,
.page-widget-content {
  display: block;
  overflow: hidden;
  text-overflow: ellipsis;
}

.page-widget-title {
  color: var(--md-primary);
  font-size: 11px;
  font-weight: 850;
  white-space: nowrap;
}

.page-widget-content {
  margin-top: 6px;
  color: var(--md-text);
  font-size: 14px;
  line-height: 1.25;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
}

.page-widget-remove {
  position: absolute;
  right: -7px;
  top: -7px;
  width: 20px;
  height: 20px;
  border-radius: 999px;
  display: grid;
  place-items: center;
  color: white;
  background: var(--md-error);
  box-shadow: var(--md-shadow-1);
}

.guide-panel {
  display: grid;
  gap: 16px;
}

.intro-panel {
  padding: 16px;
  border-radius: 10px;
  background:
    radial-gradient(circle at 92% 16%, color-mix(in srgb, var(--md-accent) 20%, transparent) 0 42px, transparent 44px),
    linear-gradient(135deg, color-mix(in srgb, var(--md-surface) 82%, var(--molly) 12%), var(--md-surface-container));
  border: 1px solid color-mix(in srgb, var(--md-primary) 28%, var(--md-outline));
  box-shadow: var(--md-shadow-1);
}

.intro-kicker {
  color: var(--md-primary);
  font-size: 12px;
  font-weight: 850;
}

.intro-title {
  margin-top: 8px;
  color: var(--md-text);
  font-size: 18px;
  font-weight: 900;
  line-height: 1.45;
}

.intro-desc {
  margin-top: 8px;
  color: var(--md-text-muted);
  font-size: 13px;
  line-height: 1.65;
}

.guide-steps {
  display: grid;
  grid-template-columns: repeat(5, minmax(0, 1fr));
  gap: 10px;
}

.guide-step {
  min-height: 138px;
  padding: 12px;
  border-radius: 10px;
  background:
    radial-gradient(circle at 82% 18%, color-mix(in srgb, var(--md-accent) 18%, transparent) 0 24px, transparent 26px),
    color-mix(in srgb, var(--md-surface-container) 76%, var(--molly) 8%);
  border: 1px solid color-mix(in srgb, var(--md-outline) 72%, var(--md-primary) 20%);
  box-shadow: var(--md-shadow-1);
}

.guide-step-index {
  width: 26px;
  height: 26px;
  border-radius: 999px;
  display: grid;
  place-items: center;
  color: var(--md-on-primary);
  background: var(--md-primary);
  font-size: 12px;
  font-weight: 900;
}

.guide-step-title {
  margin-top: 10px;
  color: var(--md-text);
  font-size: 13px;
  font-weight: 850;
}

.guide-step-desc {
  margin-top: 6px;
  color: var(--md-text-muted);
  font-size: 11px;
  line-height: 1.45;
}

.env-panel {
  padding: 14px;
  border-radius: 10px;
  background: var(--md-surface-container);
  border: 1px solid color-mix(in srgb, var(--md-outline) 76%, var(--md-primary) 18%);
}

.env-head {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 12px;
}

.env-title {
  color: var(--md-text);
  font-size: 15px;
  font-weight: 850;
}

.env-subtitle {
  margin-top: 4px;
  color: var(--md-text-muted);
  font-size: 12px;
}

.env-list {
  display: grid;
  grid-template-columns: repeat(2, minmax(0, 1fr));
  gap: 8px;
  margin-top: 12px;
}

.env-item {
  display: grid;
  grid-template-columns: 10px minmax(0, 1fr);
  gap: 9px;
  align-items: start;
  padding: 10px;
  border-radius: 8px;
  background: var(--md-surface);
  border: 1px solid var(--md-outline);
}

.env-status {
  width: 10px;
  height: 10px;
  margin-top: 4px;
  border-radius: 999px;
  background: var(--md-warning);
  box-shadow: 0 0 0 4px color-mix(in srgb, var(--md-warning) 16%, transparent);
}

.env-item-ok .env-status {
  background: var(--md-accent);
  box-shadow: 0 0 0 4px color-mix(in srgb, var(--md-accent) 16%, transparent);
}

.env-label {
  color: var(--md-text);
  font-size: 12px;
  font-weight: 800;
}

.env-detail {
  margin-top: 3px;
  color: var(--md-text-muted);
  font-size: 11px;
  line-height: 1.35;
  word-break: break-all;
}

.guide-actions {
  display: flex;
  flex-wrap: wrap;
  justify-content: flex-end;
  gap: 8px;
}

@media (max-width: 900px) {
  .guide-steps,
  .env-list {
    grid-template-columns: 1fr;
  }
}

.empty-hint {
  width: min(320px, 80%);
  padding: 28px;
  border-radius: 8px;
  text-align: center;
  background:
    radial-gradient(circle at 18% 12%, color-mix(in srgb, var(--molly-hair) 26%, transparent), transparent 34%),
    color-mix(in srgb, var(--md-surface) 90%, transparent);
  border: 1px dashed color-mix(in srgb, var(--molly-mouth) 34%, var(--md-outline));
  box-shadow: var(--md-shadow-1);
}

.empty-mark {
  position: relative;
  width: 82px;
  height: 58px;
  margin: 0 auto;
  border-radius: 8px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: var(--md-primary);
  font-size: 28px;
  font-weight: 900;
  background:
    linear-gradient(135deg, var(--molly-hair) 0 32%, var(--molly) 33% 58%, var(--md-primary) 59% 80%, var(--molly-eye) 81% 100%);
  border: 1px solid color-mix(in srgb, var(--md-outline) 72%, white);
  box-shadow:
    inset 0 0 0 6px color-mix(in srgb, white 35%, transparent),
    var(--md-shadow-1);
  animation: soft-bob 3s ease-in-out infinite;
}

@keyframes soft-bob {
  0%, 100% {
    transform: translateY(0);
  }
  50% {
    transform: translateY(-5px);
  }
}
</style>
