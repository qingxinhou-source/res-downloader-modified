<template>
  <NModal
      :show="showModal"
      :on-update:show="changeShow"
      style="--wails-draggable:no-drag"
      preset="card"
      class="w-[640px]"
      :title="t('footer.title')"
  >
    <div class="footer-panel rounded-[8px] p-5 bg-[var(--md-surface)]">
      <div class="flex items-start gap-4">
        <div class="footer-molly-mark material-surface-strong shrink-0" aria-hidden="true">RD</div>
        <div class="min-w-0">
          <div class="flex items-end gap-3">
            <div class="text-4xl font-medium text-[var(--md-text)]">{{ store.appInfo.AppName }}</div>
            <div class="text-xs text-[var(--md-text-muted)] pb-1">
              Version {{ store.appInfo.Version }}
            </div>
          </div>
          <div class="text-[var(--md-text-muted)] pt-2 leading-6">
            {{ t('footer.description') }}
          </div>
        </div>
      </div>
      <div class="mt-5 rounded-[8px] bg-[var(--md-surface-container)] border border-[var(--md-outline)] p-4">
        <div class="text-lg font-medium text-[var(--md-text)]">版权与致谢</div>
        <div class="mt-2 text-sm text-[var(--md-text-muted)] leading-6">
          本软件为基于 res-downloader 的二次开发版本，遵循 Apache License 2.0 发布。原项目版权归其原作者及贡献者所有，修改版保留 LICENSE 与 NOTICE，并补充主题、交互、个性化和打包流程相关改动。
        </div>
        <div class="mt-3 grid grid-cols-2 gap-2 text-sm text-[var(--md-text-muted)]">
          <div>Author: houqingxin</div>
          <div>Email: houqingxincampus@hotmail.com</div>
          <div>Modified Build: 2026</div>
          <div>License: Apache-2.0</div>
          <div>Origin: res-downloader</div>
          <div>Attribution: LICENSE / NOTICE</div>
        </div>
      </div>
      <div class="mt-4 rounded-[8px] bg-[var(--md-surface-container)] border border-[var(--md-outline)] p-4">
        <div class="text-lg font-medium text-[var(--md-text)]">更新日志</div>
        <div class="mt-3 changelog-list text-sm text-[var(--md-text-muted)]">
          <div><b>2026-05-16</b><span>加入 Molly 盲盒主题、全局明暗模式、自定义主题工坊。</span></div>
          <div><b>2026-05-16</b><span>新增页面编辑模式、可拖拽自定义组件、时间显示和个性化显示开关。</span></div>
          <div><b>2026-05-16</b><span>优化主题角色卡、视觉反馈、Molly 互动挂件和保留历史版本打包流程。</span></div>
        </div>
      </div>
      <div class="flex w-full flex-wrap gap-2 justify-between pt-6 text-sm text-[var(--md-text-muted)]">
        <div>{{ store.appInfo.Copyright }}</div>
        <div class="flex flex-wrap gap-2">
          <button class="material-chip" @click="toWebsite(certUrl)">{{ t('footer.cert_download') }}</button>
          <button class="material-chip" @click="toWebsite('https://www.apache.org/licenses/LICENSE-2.0')">Apache 2.0</button>
          <button class="material-chip" @click="toWebsite('https://github.com/putyy/res-downloader')">Upstream</button>
        </div>
      </div>
    </div>
  </NModal>
</template>

<script lang="ts" setup>
import {useIndexStore} from "@/stores"
import {BrowserOpenURL} from "../../wailsjs/runtime"
import {computed} from "vue"
import {useI18n} from 'vue-i18n'

const {t} = useI18n()
const store = useIndexStore()
const props = defineProps(["showModal"])
const emits = defineEmits(["update:showModal"])
const certUrl = computed(()=>{
  return store.baseUrl + "/api/cert"
})
const changeShow = (value: boolean) => {
  emits('update:showModal', value)
}

const toWebsite = (url: string) => {
  BrowserOpenURL(url)
}
</script>

<style scoped>
.footer-panel {
  background:
    radial-gradient(circle at 10% 10%, color-mix(in srgb, var(--molly-hair) 26%, transparent), transparent 28%),
    var(--md-surface);
}

.footer-molly-mark {
  width: 72px;
  height: 72px;
  border-radius: 8px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: var(--md-primary);
  font-size: 28px;
  font-weight: 900;
  background:
    linear-gradient(135deg, var(--molly-hair) 0 32%, var(--molly) 33% 58%, var(--md-primary) 59% 80%, var(--molly-eye) 81% 100%);
  box-shadow: inset 0 0 0 6px color-mix(in srgb, white 35%, transparent);
}

.changelog-list {
  display: grid;
  gap: 8px;
}

.changelog-list div {
  display: grid;
  grid-template-columns: 92px minmax(0, 1fr);
  gap: 10px;
}

.changelog-list b {
  color: var(--md-primary);
  font-weight: 800;
}
</style>
