import { reactive } from 'vue'

interface ConfirmOptions {
    title?: string
    confirmLabel?: string
}

interface ConfirmState {
    open: boolean
    title: string
    message: string
    confirmLabel: string
    resolve: ((value: boolean) => void) | null
}

// Single dialog instance state - one confirmation can be open at a time,
// which mirrors how native confirm() behaved.
export const confirmState = reactive<ConfirmState>({
    open: false,
    title: 'Are you sure?',
    message: '',
    confirmLabel: 'Delete',
    resolve: null,
})

/**
 * Promise-based replacement for native confirm(). Resolves true when the
 * destructive action is confirmed, false on cancel/escape/backdrop.
 */
export function appConfirm(message: string, opts: ConfirmOptions = {}): Promise<boolean> {
    // Settle any dangling request before opening a new one
    confirmState.resolve?.(false)
    confirmState.title = opts.title ?? 'Are you sure?'
    confirmState.message = message
    confirmState.confirmLabel = opts.confirmLabel ?? 'Delete'
    confirmState.open = true
    return new Promise(resolve => { confirmState.resolve = resolve })
}

export function settleConfirm(value: boolean) {
    confirmState.open = false
    confirmState.resolve?.(value)
    confirmState.resolve = null
}
