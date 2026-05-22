<template>
  <div class="flex items-center">
    <span>
      {{ t('index.operation') }}
    </span>
    <NPopover trigger="hover">
      <template #trigger>
        <NIcon size="18" class="ml-1 text-gray-500">
          <HelpCircleOutline/>
        </NIcon>
      </template>
      <div class="flex flex-col gap-1">
        <div class="flex items-center justify-start p-1.5">
          <n-icon size="28"
                  class="material-icon-button text-[var(--md-secondary)] bg-[var(--md-secondary)]/10 p-1.5 cursor-pointer hover:bg-[var(--md-secondary)]/20">
            <DownloadOutline/>
          </n-icon>
          <span class="ml-1">{{ t("index.direct_download") }}</span>
        </div>

        <div class="flex items-center justify-start p-1.5">
          <n-icon size="28"
                  class="material-icon-button text-[var(--md-error)] bg-[var(--md-error)]/10 p-1.5 cursor-pointer hover:bg-[var(--md-error)]/20">
            <CloseOutline/>
          </n-icon>
          <span class="ml-1">{{ t("index.cancel_down") }}</span>
        </div>

        <div class="flex items-center justify-start p-1.5">
          <n-icon
              size="28"
              class="material-icon-button text-[var(--md-primary)] bg-[var(--md-primary)]/10 p-1.5 cursor-pointer hover:bg-[var(--md-primary)]/20"
          >
            <LinkOutline/>
          </n-icon>
          <span class="ml-1">{{ t("index.copy_link") }}</span>
        </div>

        <div class="flex items-center justify-start p-1.5">
          <n-icon
              size="28"
              class="material-icon-button text-[var(--md-primary)] bg-[var(--md-primary)]/10 p-1.5 cursor-pointer hover:bg-[var(--md-primary)]/20"
          >
            <GlobeOutline/>
          </n-icon>
          <span class="ml-1">{{ t("index.open_link") }}</span>
        </div>

        <div class="flex items-center justify-start p-1.5">
          <n-icon
              size="28"
              class="material-icon-button text-[var(--md-warning)] bg-[var(--md-warning)]/10 p-1.5 cursor-pointer hover:bg-[var(--md-warning)]/20"
          >
            <LockOpenSharp/>
          </n-icon>
          <span class="ml-1">{{ t("index.video_decode") }}</span>
        </div>

        <div class="flex items-center justify-start p-1.5">
          <n-icon
              size="28"
              class="material-icon-button text-[var(--md-primary)] bg-[var(--md-primary)]/10 p-1.5 cursor-pointer hover:bg-[var(--md-primary)]/20"
          >
            <CopyOutline/>
          </n-icon>
          <span class="ml-1">{{ t("index.copy_data") }}</span>
        </div>

        <div class="flex items-center justify-start p-1.5">
          <n-icon
              size="28"
              class="material-icon-button text-[var(--md-error)] bg-[var(--md-error)]/10 p-1.5 cursor-pointer hover:bg-[var(--md-error)]/20"
          >
            <TrashOutline/>
          </n-icon>
          <span class="ml-1">{{ t("index.delete_row") }}</span>
        </div>

        <div class="flex items-center justify-start p-1.5">
          <n-icon
              size="28"
              class="material-icon-button text-[var(--md-primary)] bg-[var(--md-primary)]/10 p-2 cursor-pointer hover:bg-[var(--md-primary)]/20"
          >
            <GridSharp/>
          </n-icon>
          <span class="ml-1">{{ t("index.more_operation") }}</span>
        </div>
      </div>
    </NPopover>
  </div>
</template>
<script setup lang="ts">
import {useI18n} from "vue-i18n"
import {
  CopyOutline,
  DownloadOutline,
  GlobeOutline,
  HelpCircleOutline,
  LinkOutline,
  LockOpenSharp,
  GridSharp,
  CloseOutline,
  TrashOutline
} from "@vicons/ionicons5"

const {t} = useI18n()
</script>
