import { ref } from 'vue'

/**
 * Shared clipboard interaction: copies text and exposes a timed
 * `copied` state for visual feedback. Single source of truth for
 * the copy pattern across the app.
 */
export function useClipboard(resetMs = 1600) {
    const copied = ref('')
    let timer: number | undefined

    const copy = async (text: string) => {
        if (!text) return
        await navigator.clipboard.writeText(text)
        copied.value = text
        window.clearTimeout(timer)
        timer = window.setTimeout(() => { copied.value = '' }, resetMs)
    }

    return { copied, copy }
}
