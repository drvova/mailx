<template>
    <div class="relative">
        <input
            v-bind="$attrs"
            :type="visible ? 'text' : 'password'"
            :value="modelValue"
            class="pr-11"
            @input="$emit('update:modelValue', ($event.target as HTMLInputElement).value)"
        >
        <button
            type="button"
            class="absolute top-0 right-0 bottom-0 px-3 flex items-center text-tertiary"
            aria-label="Show password"
            :aria-pressed="visible"
            @click="visible = !visible">
            <svg v-if="visible" class="shrink-0 size-4" xmlns="http://www.w3.org/2000/svg" width="24" height="24"
                viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"
                stroke-linejoin="round" aria-hidden="true">
                <path d="M10.733 5.076a10.744 10.744 0 0 1 11.205 6.575 1 1 0 0 1 0 .696 10.747 10.747 0 0 1-1.444 2.49"></path>
                <path d="M14.084 14.158a3 3 0 0 1-4.242-4.242"></path>
                <path d="M17.479 17.499a10.75 10.75 0 0 1-15.417-5.151 1 1 0 0 1 0-.696 10.75 10.75 0 0 1 4.446-5.143"></path>
                <path d="m2 2 20 20"></path>
            </svg>
            <svg v-else class="shrink-0 size-4" xmlns="http://www.w3.org/2000/svg" width="24" height="24"
                viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"
                stroke-linejoin="round" aria-hidden="true">
                <path d="M2.062 12.348a1 1 0 0 1 0-.696 10.75 10.75 0 0 1 19.876 0 1 1 0 0 1 0 .696 10.75 10.75 0 0 1-19.876 0"></path>
                <circle cx="12" cy="12" r="3"></circle>
            </svg>
        </button>
    </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

defineOptions({ inheritAttrs: false })
defineProps<{ modelValue: string }>()
defineEmits(['update:modelValue'])

const visible = ref(false)
</script>
