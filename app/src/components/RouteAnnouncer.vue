<template>
    <span class="sr-only" aria-live="polite">{{ message }}</span>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'

const message = ref('')

useRouter().afterEach((to, _from, failure) => {
    if (!failure) message.value = String(to.name ?? '')
})
</script>
