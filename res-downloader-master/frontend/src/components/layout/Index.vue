<template>
<div class="w-full h-full p-3 overflow-hidden">
  <div id="content" class="app-shell h-full w-full flex overflow-hidden">
    <Sider />
    <NLayoutContent class="workspace-page flex-1 min-w-0">
      <RouterView v-slot="{ Component, route }">
        <keep-alive v-if="route.meta.keepAlive">
          <component :is="Component" />
        </keep-alive>
        <component v-else :is="Component" />
      </RouterView>
    </NLayoutContent>
  </div>
</div>
</template>
<script lang="ts" setup>
import Sider from "./Sider.vue"
</script>
