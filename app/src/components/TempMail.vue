<template>
    <div class="hyper p-6 min-h-[70vh]">
        <header class="flex flex-wrap items-center justify-between gap-4 mb-6">
            <div>
                <h2 class="text-xl font-semibold m-0">Temp Mail</h2>
                <p class="hy-muted text-sm m-0 mt-1">Disposable inboxes. Mail is stored here, never forwarded, and self-destructs.</p>
            </div>
            <div class="flex items-center gap-3">
                <select v-model="newDomain" class="hy-select" aria-label="Domain">
                    <option v-for="d in domains" :key="d" :value="d">@{{ d }}</option>
                </select>
                <select v-model.number="newTTL" class="hy-select" aria-label="Lifetime">
                    <option :value="1">1 hour</option>
                    <option :value="24">24 hours</option>
                    <option :value="168">7 days</option>
                    <option :value="720">30 days</option>
                </select>
                <button class="hy-btn hy-btn--primary" :disabled="creating" @click="createInbox">
                    {{ creating ? 'Generating…' : 'Generate temp email' }}
                </button>
            </div>
        </header>

        <p v-if="error" class="hy-error mb-4" role="alert">{{ error }}</p>

        <div v-if="loaded && !inboxes.length" class="hy-card p-10 text-center">
            <p class="hy-muted m-0">No temp inboxes yet. Generate one and use it anywhere you don't trust with your real address.</p>
        </div>

        <div v-else class="grid gap-4 lg:grid-cols-[280px_320px_1fr] md:grid-cols-[280px_1fr]">
            <!-- Inbox list -->
            <div class="hy-card p-2 hy-scroll max-h-[70vh]">
                <p class="hy-label px-3 pt-2 pb-1 m-0">Inboxes</p>
                <div v-for="inbox in inboxes" :key="inbox.id"
                    class="hy-row" :class="{ 'hy-row--active': inbox.id === selectedInbox?.id }">
                    <button type="button" class="hy-row__select" @click="selectInbox(inbox)">
                        <span class="hy-mono text-xs block truncate">{{ inbox.name }}</span>
                        <span class="hy-dim text-xs block mt-0.5">{{ expiresIn(inbox.expires_at) }}</span>
                    </button>
                    <CopyButton class="hy-badge" :text="inbox.name" :label="'Copy ' + inbox.name + ' to clipboard'" />
                </div>
            </div>

            <!-- Message list -->
            <div class="hy-card p-2 hy-scroll max-h-[70vh]">
                <p class="hy-label px-3 pt-2 pb-1 m-0">Messages</p>
                <p v-if="selectedInbox && messagesLoaded && !messages.length" class="hy-dim text-sm px-3 py-4 m-0">
                    Empty. Mail sent to <span class="hy-mono hy-accent-text">{{ selectedInbox.name }}</span> lands here.
                </p>
                <p v-else-if="!selectedInbox" class="hy-dim text-sm px-3 py-4 m-0">Select an inbox.</p>
                <button v-for="msg in messages" :key="msg.id"
                    class="hy-row" :class="{ 'hy-row--active': msg.id === openMessage?.id }"
                    @click="openMsg(msg.id)">
                    <span v-if="!msg.read" class="hy-unread" aria-label="Unread"></span>
                    <span class="flex-1 min-w-0">
                        <span class="block text-sm truncate" :class="{ 'font-semibold': !msg.read }">{{ msg.subject || '(no subject)' }}</span>
                        <span class="hy-dim text-xs block truncate mt-0.5">{{ msg.from_name || msg.from }} · {{ shortDate(msg.created_at) }}</span>
                    </span>
                </button>
            </div>

            <!-- Reader -->
            <div class="hy-card p-4 hy-scroll max-h-[70vh]">
                <p v-if="!openMessage" class="hy-dim text-sm m-0">Select a message.</p>
                <template v-else>
                    <div class="flex items-start justify-between gap-4 mb-3">
                        <div class="min-w-0">
                            <h3 class="text-base font-semibold m-0 break-words">{{ openMessage.subject || '(no subject)' }}</h3>
                            <p class="hy-muted text-xs m-0 mt-1">
                                <span class="hy-mono">{{ openMessage.from_name ? openMessage.from_name + ' <' + openMessage.from + '>' : openMessage.from }}</span>
                                · {{ fullDate(openMessage.created_at) }}
                            </p>
                        </div>
                        <button class="hy-btn hy-btn--danger hy-btn--sm" :disabled="deleting" :aria-busy="deleting" @click="deleteMsg(openMessage.id)">Delete</button>
                    </div>
                    <div v-if="openMessage.attachments?.length" class="flex flex-wrap gap-2 mb-3">
                        <span v-for="a in openMessage.attachments" :key="a.name" class="hy-badge">
                            {{ a.name }} ({{ formatSize(a.size) }})
                        </span>
                    </div>
                    <iframe v-if="openMessage.html" class="hy-reader-frame" sandbox=""
                        :srcdoc="frameDoc(openMessage.html)" :style="{ height: '52vh' }" title="Message content"></iframe>
                    <pre v-else class="hy-mono text-sm whitespace-pre-wrap break-words m-0">{{ openMessage.text }}</pre>
                </template>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { onMounted, onUnmounted, ref } from 'vue'
import { ApiError } from '../api/api.ts'
import { aliasApi } from '../api/alias.ts'
import { inboxApi } from '../api/inbox.ts'
import { useClipboard } from '../composables/useClipboard.ts'
import { toast } from '../composables/useToast.ts'
import CopyButton from './CopyButton.vue'

interface InboxAlias {
    id: string
    name: string
    type: number
    expires_at?: string
}

interface InboxMessage {
    id: number
    created_at: string
    from: string
    from_name: string
    subject: string
    read: boolean
    size: number
}

interface RenderedMessage {
    id: number
    created_at: string
    from: string
    from_name: string
    subject: string
    html: string
    text: string
    attachments: { name: string; size: number }[]
}

const domains = import.meta.env.VITE_DOMAINS.split(',')

const inboxes = ref<InboxAlias[]>([])
const messages = ref<InboxMessage[]>([])
const selectedInbox = ref<InboxAlias | null>(null)
const openMessage = ref<RenderedMessage | null>(null)
const newDomain = ref(domains[0])
const newTTL = ref(24)
const creating = ref(false)
const deleting = ref(false)
const loaded = ref(false)
const messagesLoaded = ref(false)
const error = ref('')

const { copy } = useClipboard()

const handle = (err: unknown) => {
    error.value = err instanceof ApiError ? (err.data?.error || err.message) : String(err)
}

const getInboxes = async () => {
    try {
        const response = await aliasApi.getList({ limit: 0, page: 0 })
        inboxes.value = (response.data.aliases || []).filter((a: InboxAlias) => a.type === 1)
        loaded.value = true
        error.value = ''
        if (!selectedInbox.value && inboxes.value.length) {
            selectInbox(inboxes.value[0])
        }
    } catch (err) {
        handle(err)
    }
}

const createInbox = async () => {
    creating.value = true
    try {
        const response = await aliasApi.create({
            type: 'inbox',
            ttl_hours: newTTL.value,
            domain: newDomain.value,
            enabled: true,
            format: 'random',
        })
        error.value = ''
        await getInboxes()
        const created = inboxes.value.find((a) => a.id === response.data.alias?.id)
        if (created) selectInbox(created)
        await copy(response.data.alias?.name || '')
        toast('Temp inbox ready — address copied')
    } catch (err) {
        handle(err)
    } finally {
        creating.value = false
    }
}

const selectInbox = async (inbox: InboxAlias) => {
    selectedInbox.value = inbox
    openMessage.value = null
    messagesLoaded.value = false
    try {
        const response = await inboxApi.getMessages(inbox.id)
        messages.value = response.data || []
        messagesLoaded.value = true
        error.value = ''
    } catch (err) {
        handle(err)
    }
}

const openMsg = async (id: number) => {
    try {
        const response = await inboxApi.getMessage(id)
        openMessage.value = response.data
        const row = messages.value.find((m) => m.id === id)
        if (row) row.read = true
        error.value = ''
    } catch (err) {
        handle(err)
    }
}

const deleteMsg = async (id: number) => {
    deleting.value = true
    try {
        await inboxApi.deleteMessage(id)
        messages.value = messages.value.filter((m) => m.id !== id)
        openMessage.value = null
        error.value = ''
        toast('Message deleted')
    } catch (err) {
        handle(err)
    } finally {
        deleting.value = false
    }
}

const expiresIn = (expiresAt?: string) => {
    if (!expiresAt) return 'no expiry'
    const ms = new Date(expiresAt).getTime() - Date.now()
    if (ms <= 0) return 'expired'
    const hours = Math.floor(ms / 3_600_000)
    if (hours >= 48) return 'expires in ' + Math.floor(hours / 24) + 'd'
    if (hours >= 1) return 'expires in ' + hours + 'h'
    return 'expires in ' + Math.max(1, Math.floor(ms / 60_000)) + 'm'
}

const shortDate = (iso: string) => new Date(iso).toLocaleString(undefined, { month: 'short', day: 'numeric', hour: '2-digit', minute: '2-digit' })
const fullDate = (iso: string) => new Date(iso).toLocaleString()

const formatSize = (bytes: number) => {
    if (bytes < 1024) return bytes + ' B'
    if (bytes < 1_048_576) return (bytes / 1024).toFixed(1) + ' KB'
    return (bytes / 1_048_576).toFixed(1) + ' MB'
}

// Sanitized server-side; sandbox="" blocks scripts as a second layer.
const frameDoc = (html: string) =>
    '<!doctype html><html><head><meta charset="utf-8"><style>body{font-family:system-ui,sans-serif;font-size:14px;line-height:1.6;color:oklch(0.218 0 0);margin:12px;word-break:break-word}a{color:oklch(0.556 0.141 44.92)}</style></head><body>' + html + '</body></html>'

// Auto-refresh: temp inboxes exist to wait for incoming mail (Doherty threshold)
let pollTimer: number | undefined

const refreshMessages = async () => {
    if (document.hidden || !selectedInbox.value) return
    try {
        const response = await inboxApi.getMessages(selectedInbox.value.id)
        messages.value = response.data || []
        messagesLoaded.value = true
    } catch (err) {
        // Poll errors stay off the banner; manual actions surface errors
        console.warn('inbox poll failed:', err)
    }
}

onMounted(() => {
    getInboxes()
    pollTimer = window.setInterval(refreshMessages, 15000)
})

onUnmounted(() => window.clearInterval(pollTimer))
</script>
