import { reactive } from 'vue'

interface Toast {
    id: number
    message: string
}

let nextId = 0

// Module-level singleton list, mirrors the useConfirm pattern.
export const toasts = reactive<Toast[]>([])

export function toast(message: string) {
    const id = nextId++
    toasts.push({ id, message })
    setTimeout(() => {
        const index = toasts.findIndex(t => t.id === id)
        if (index !== -1) toasts.splice(index, 1)
    }, 4000)
}
