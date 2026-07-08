<template>
    <Teleport to="body">
        <div class="fixed bottom-4 right-4 z-50 flex flex-col gap-2 items-end">
            <TransitionGroup name="toast-fade">
                <div v-for="t in toasts" :key="t.id" :role="t.kind === 'error' ? 'alert' : 'status'"
                    class="toast-panel text-sm px-4 py-3" :class="{ error: t.kind === 'error' }">
                    {{ t.message }}
                </div>
            </TransitionGroup>
        </div>
    </Teleport>
</template>

<script setup lang="ts">
import { toasts } from '../composables/useToast.ts'
</script>

<style scoped>
.toast-panel {
    background: var(--shell-bg);
    border: 1px solid var(--sk-border);
    border-radius: var(--radius);
    box-shadow: var(--raised-shadow);
    font-family: var(--font-mono);
}

.toast-panel.error {
    border-color: var(--danger);
    color: var(--danger);
}

.toast-fade-enter-active {
    transition: opacity 120ms var(--ease-smooth);
}

.toast-fade-enter-from {
    opacity: 0;
}
</style>
