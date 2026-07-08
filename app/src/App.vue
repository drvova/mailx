<template>
  <router-view v-slot="{ Component }">
      <Transition name="route-fade">
          <component :is="Component" />
      </Transition>
  </router-view>
  <ConfirmDialog />
  <ToastHost />
  <RouteAnnouncer />
</template>

<script setup lang="ts">
import ConfirmDialog from './components/ConfirmDialog.vue'
import ToastHost from './components/ToastHost.vue'
import RouteAnnouncer from './components/RouteAnnouncer.vue'

// Evaluate the initial theme
if (localStorage.theme === 'dark' || (!('theme' in localStorage) && window.matchMedia('(prefers-color-scheme: dark)').matches)) {
  document.documentElement.classList.add('dark')
} else {
  document.documentElement.classList.remove('dark')
}
</script>

<style>
.route-fade-enter-active {
  transition: opacity 120ms var(--ease-smooth);
}

.route-fade-enter-from {
  opacity: 0;
}
</style>