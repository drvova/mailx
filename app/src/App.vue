<template>
  <router-view />
  <ConfirmDialog />
</template>

<script setup lang="ts">
import ConfirmDialog from './components/ConfirmDialog.vue'

// Evaluate the initial theme
if (localStorage.theme === 'dark' || (!('theme' in localStorage) && window.matchMedia('(prefers-color-scheme: dark)').matches)) {
  document.documentElement.classList.add('dark')
} else {
  document.documentElement.classList.remove('dark')
}
</script>