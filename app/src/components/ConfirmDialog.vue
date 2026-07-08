<template>
    <Teleport to="body">
        <div v-if="confirmState.open" class="confirm-backdrop" @click.self="settleConfirm(false)">
            <div ref="panel" class="confirm-panel" role="alertdialog" aria-modal="true"
                aria-labelledby="confirm-title" aria-describedby="confirm-message" @keydown="onKeydown">
                <h4 id="confirm-title" class="m-0 mb-2">{{ confirmState.title }}</h4>
                <p id="confirm-message" class="m-0 mb-6 text-secondary text-sm">{{ confirmState.message }}</p>
                <div class="flex justify-end gap-3">
                    <button ref="cancelBtn" type="button" class="cta cancel" @click="settleConfirm(false)">
                        Cancel
                    </button>
                    <button ref="confirmBtn" type="button" class="cta delete" @click="settleConfirm(true)">
                        {{ confirmState.confirmLabel }}
                    </button>
                </div>
            </div>
        </div>
    </Teleport>
</template>

<script setup lang="ts">
import { nextTick, ref, watch } from 'vue'
import { confirmState, settleConfirm } from '../composables/useConfirm'

const panel = ref<HTMLElement>()
const cancelBtn = ref<HTMLButtonElement>()
const confirmBtn = ref<HTMLButtonElement>()

let previousFocus: HTMLElement | null = null

watch(() => confirmState.open, async (open) => {
    if (open) {
        previousFocus = document.activeElement as HTMLElement | null
        await nextTick()
        // Safe default: focus lands on Cancel, not the destructive action
        cancelBtn.value?.focus()
    } else {
        previousFocus?.focus()
        previousFocus = null
    }
})

const onKeydown = (e: KeyboardEvent) => {
    if (e.key === 'Escape') {
        e.stopPropagation()
        settleConfirm(false)
        return
    }
    // Two focusable controls: cycle Tab between them
    if (e.key === 'Tab') {
        e.preventDefault()
        const next = document.activeElement === cancelBtn.value ? confirmBtn.value : cancelBtn.value
        next?.focus()
    }
}
</script>

<style scoped>
.confirm-backdrop {
    position: fixed;
    inset: 0;
    z-index: 90;
    display: flex;
    align-items: center;
    justify-content: center;
    padding: 20px;
    background: oklch(0 0 0 / 0.6);
}

.confirm-panel {
    width: 100%;
    max-width: 26rem;
    padding: 24px;
    border: 1px solid var(--sk-border);
    border-radius: var(--radius);
    background: oklch(1 0 0);
    box-shadow: var(--raised-shadow);
}

html.dark .confirm-panel {
    background: oklch(0.213 0 0);
}
</style>
