<template>
    <button type="button" :aria-label="ariaLabel" @click.stop="copy(text)">
        <slot :copied="isCopied">
            <span aria-hidden="true">{{ isCopied ? 'copied' : 'copy' }}</span>
        </slot>
        <span class="sr-only" role="status">{{ isCopied ? 'Copied to clipboard' : '' }}</span>
    </button>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { useClipboard } from '../composables/useClipboard'

const props = defineProps<{
    text: string
    label?: string
}>()

const { copied, copy } = useClipboard()
const isCopied = computed(() => copied.value === props.text)
const ariaLabel = computed(() => props.label ?? 'Copy ' + props.text + ' to clipboard')
</script>
