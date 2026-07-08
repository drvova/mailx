<template>
    <tr class="desktop-lg">
        <td>
            <div class="flex items-center hs-tooltip">
                <input
                    @change="updateAlias"
                    v-bind:checked="alias.enabled && !isDomainUnverified"
                    v-bind:disabled="!alias.recipients.length || isDomainUnverified"
                    type="checkbox"
                >
                <span v-if="isDomainUnverified" class="hs-tooltip-content hs-tooltip-shown:opacity-100 hs-tooltip-shown:visible" role="tooltip">
                    Domain not verified or disabled. Address is not forwarding mail.
                </span>
                <span v-else-if="!alias.recipients.length" class="hs-tooltip-content hs-tooltip-shown:opacity-100 hs-tooltip-shown:visible" role="tooltip">
                    Disabled
                </span>
            </div>
        </td>
        <td class="whitespace-normal">
            <div class="block break-all hs-tooltip">
                <p class="hs-tooltip-toggle m-0">{{ truncatedDescription }}</p>
                <span v-if="alias.description && alias.description.length > 45" class="hs-tooltip-content hs-tooltip-shown:opacity-100 hs-tooltip-shown:visible" role="tooltip">{{ alias.description }}</span>
            </div>
        </td>
        <td>
            <div class="hs-tooltip inline-block">
                <p class="hs-tooltip-toggle m-0">
                    <button class="plain text-wrap text-start text-sm p-0" @click="copyAlias(alias.name)">
                        {{ alias.name }}
                    </button>
                    <span class="hs-tooltip-content hs-tooltip-shown:opacity-100 hs-tooltip-shown:visible" role="tooltip">
                        {{ copied === alias.name ? 'Copied' : 'Click to copy' }}: {{ alias.name }}
                    </span>
                </p>
                <span class="sr-only" role="status">{{ copied === alias.name ? 'Copied to clipboard' : '' }}</span>
            </div>
        </td>
        <td>
            <div class="flex items-center gap-3 mb-1">
                <p class="flex items-center gap-1 hs-tooltip">
                    {{ alias.stats.forwards }}
                    <i class="icon forward text-xs icon-tertiary"></i>
                    <span class="hs-tooltip-content hs-tooltip-shown:opacity-100 hs-tooltip-shown:visible" role="tooltip">{{ alias.stats.forwards }} Forwards</span>
                </p>
                <p class="flex items-center gap-1 hs-tooltip">
                    {{ alias.stats.blocks }}
                    <i class="icon block text-xs icon-tertiary"></i>
                    <span class="hs-tooltip-content hs-tooltip-shown:opacity-100 hs-tooltip-shown:visible" role="tooltip">{{ alias.stats.blocks }} Blocks</span>
                </p>
            </div>
            <div class="flex items-center gap-3 mt-1">
                <p class="flex items-center gap-1 hs-tooltip">
                    {{ alias.stats.replies }}
                    <i class="icon reply text-xs icon-tertiary"></i>
                    <span class="hs-tooltip-content hs-tooltip-shown:opacity-100 hs-tooltip-shown:visible" role="tooltip">{{ alias.stats.replies }} Replies</span>
                </p>
                <p class="flex items-center gap-1 hs-tooltip">
                    {{ alias.stats.sends }}
                    <i class="icon send text-xs icon-tertiary"></i>
                    <span class="hs-tooltip-content hs-tooltip-shown:opacity-100 hs-tooltip-shown:visible" role="tooltip">{{ alias.stats.sends }} Sends</span>
                </p>
            </div>
        </td>
        <td>
            <div class="mt-1 flex items-center gap-2">
                <p :title="new Date(alias.created_at).toLocaleString()">{{ formatDistanceToNow(new Date(alias.created_at)) }} ago</p>
            </div>
        </td>
        <td>
            <div class="hs-dropdown [--offset:0]">
                <button v-bind:id="'hs-dropdown-alias-edit-' + alias.id" aria-label="Alias actions">
                    <i class="icon icon-secondary more text-lg"></i>
                </button>
                <div
                    class="hs-dropdown-menu hs-dropdown-open:opacity-100 hidden"
                    v-bind:aria-labelledby="'hs-dropdown-alias-edit-' + alias.id"
                >
                    <button
                        v-bind:disabled="!alias.recipients.length"
                        v-bind:data-hs-overlay="'#modal-send-alias' + alias.id"
                        v-bind:class="{ 'hide': alias.catch_all }"
                        >
                        <i class="icon icon-primary send text-xs"></i>
                        Send
                    </button>
                    <button v-bind:data-hs-overlay="'#modal-alias-edit' + alias.id" :aria-label="'Edit ' + alias.name">
                        <i class="icon icon-primary edit text-xs"></i>
                        Edit
                    </button>
                    <button @click.stop="deleteAlias" class="delete" :aria-label="'Delete ' + alias.name">
                        <i class="icon icon-error trash text-xs"></i>
                        Delete
                    </button>
                </div>
            </div>
        </td>
    </tr>
    <tr class="tablet-lg">
        <td>
            <div class="flex gap-2 justify-between">
                <div class="text-start">
                    <div>
                        <p class="mb-3" :title="new Date(alias.created_at).toLocaleString()">{{ formatDistanceToNow(new Date(alias.created_at)) }} ago</p>
                    </div>
                    <div>
                        <div class="hs-tooltip inline-block mb-5 break-all">
                            <p class="hs-tooltip-toggle mb-0">
                                <button class="plain truncate text-sm p-0 text-wrap text-start" @click="copyAlias(alias.name)">
                                    <span v-if="alias.description" class="block break-words">{{ truncatedDescription }}</span>
                                    <span class="block text-sm break-all">{{ alias.name }}</span>
                                </button>
                                <span class="hs-tooltip-content hs-tooltip-shown:opacity-100 hs-tooltip-shown:visible" role="tooltip">
                                    {{ copied === alias.name ? 'Copied' : 'Click to copy' }}: {{ alias.name }}
                                </span>
                            </p>
                        </div>
                    </div>
                    <div class="flex items-center hs-tooltip">
                        <input
                            @change="updateAlias"
                            v-bind:checked="alias.enabled && !isDomainUnverified"
                            v-bind:disabled="!alias.recipients.length || isDomainUnverified"
                            type="checkbox"
                        >
                        <span v-if="isDomainUnverified" class="hs-tooltip-content hs-tooltip-shown:opacity-100 hs-tooltip-shown:visible" role="tooltip">
                            Domain not verified. Address is not forwarding mail.
                        </span>
                        <span v-else-if="!alias.recipients.length" class="hs-tooltip-content hs-tooltip-shown:opacity-100 hs-tooltip-shown:visible" role="tooltip">
                            Disabled
                        </span>
                    </div>
                </div>
                <div>
                    <div class="hs-dropdown [--offset:0] mb-3">
                        <button class="py-0" v-bind:id="'hs-dropdown-alias-edit-' + alias.id" aria-label="Alias actions">
                            <i class="icon icon-secondary more text-lg"></i>
                        </button>
                        <div
                            class="hs-dropdown-menu hs-dropdown-open:opacity-100 hidden"
                            v-bind:aria-labelledby="'hs-dropdown-alias-edit-' + alias.id"
                        >
                            <button
                                v-bind:disabled="!alias.recipients.length"
                                v-bind:data-hs-overlay="'#modal-send-alias' + alias.id"
                                v-bind:class="{ 'hide': alias.catch_all }"
                                >
                                <i class="icon icon-primary send text-xs"></i>
                                Send
                            </button>
                            <button v-bind:data-hs-overlay="'#modal-alias-edit' + alias.id" :aria-label="'Edit ' + alias.name">
                                <i class="icon icon-primary edit text-xs"></i>
                                Edit
                            </button>
                            <button @click.stop="deleteAlias" class="delete" :aria-label="'Delete ' + alias.name">
                                <i class="icon icon-error trash text-xs"></i>
                                Delete
                            </button>
                        </div>
                    </div>
                    <div>
                        <div class="flex items-center gap-3 mb-1">
                            <p class="flex items-center gap-1 hs-tooltip mb-2">
                                {{ alias.stats.forwards }}
                                <i class="icon forward text-xs icon-tertiary"></i>
                                <span class="hs-tooltip-content hs-tooltip-shown:opacity-100 hs-tooltip-shown:visible" role="tooltip">{{ alias.stats.forwards }} Forwards</span>
                            </p>
                            <p class="flex items-center gap-1 hs-tooltip mb-2">
                                {{ alias.stats.blocks }}
                                <i class="icon block text-xs icon-tertiary"></i>
                                <span class="hs-tooltip-content hs-tooltip-shown:opacity-100 hs-tooltip-shown:visible" role="tooltip">{{ alias.stats.blocks }} Blocks</span>
                            </p>
                        </div>
                        <div class="flex items-center gap-3 mt-1">
                            <p class="flex items-center gap-1 hs-tooltip mb-2">
                                {{ alias.stats.replies }}
                                <i class="icon reply text-xs icon-tertiary"></i>
                                <span class="hs-tooltip-content hs-tooltip-shown:opacity-100 hs-tooltip-shown:visible" role="tooltip">{{ alias.stats.replies }} Replies</span>
                            </p>
                            <p class="flex items-center gap-1 hs-tooltip mb-2">
                                {{ alias.stats.sends }}
                                <i class="icon send text-xs icon-tertiary"></i>
                                <span class="hs-tooltip-content hs-tooltip-shown:opacity-100 hs-tooltip-shown:visible" role="tooltip">{{ alias.stats.sends }} Sends</span>
                            </p>
                        </div>
                    </div>
                </div>
            </div>
            <hr>
        </td>
    </tr>
    <AliasSend :alias="alias" />
    <AliasEdit :alias="alias" :recipients="recipients" :key="rowKey" />
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import tooltip from '@preline/tooltip'
import AliasEdit from './AliasEdit.vue'
import AliasSend from './AliasSend.vue'
import { aliasApi } from '../api/alias.ts'
import events from '../events.ts'
import { formatDistanceToNow } from 'date-fns'
import dropdown from '@preline/dropdown'
import { appConfirm } from '../composables/useConfirm.ts'
import { toast } from '../composables/useToast.ts'
import { useClipboard } from '../composables/useClipboard.ts'

const props = defineProps(['alias', 'recipients', 'catchAll'])
const alias = ref(props.alias)
const recipients = ref(props.recipients)
const isDomainUnverified = computed(() => alias.value.is_custom_domain === true && (alias.value.is_domain_verified === false || alias.value.is_domain_enabled === false))
const truncatedDescription = computed(() => {
    const desc = alias.value.description
    if (!desc) return ''
    return desc.length > 45 ? desc.slice(0, 45) + '...' : desc
})
const rowKey = ref(0)

const updateAlias = async () => {
    alias.value.enabled = !alias.value.enabled
    try {
        await aliasApi.update(alias.value.id, alias.value)
        renderRow()
        toast(alias.value.enabled ? 'Alias enabled' : 'Alias disabled')
    } catch {
        alias.value.enabled = !alias.value.enabled // revert the optimistic toggle
        toast('Failed to update alias', 'error')
    }
}

const deleteAlias = async () => {
    const errMessage = props.catchAll ? 'WARNING: You will not be able to create the same catch-all alias in the next 90 days. Are you sure you want to delete alias? ' : 'Are you sure you want to delete alias?'
    if (!(await appConfirm(errMessage, { confirmLabel: 'Delete alias' }))) return

    events.emit('alias.delete', { id: alias.value.id, catchAll: props.catchAll })
}

const { copied, copy } = useClipboard(2000)

const copyAlias = (alias: string) => {
    copy(alias).catch(() => toast('Failed to copy', 'error'))
}

const renderRow = () => {
    tooltip.autoInit()
}

onMounted(() => {
    tooltip.autoInit()
    dropdown.autoInit()
})
</script>
