<template>
    <div>
        <div v-bind:id="'modal-verify-domain' + domain.id" class="hs-overlay hidden">
            <div>
                <div>
                    <header>
                        <button @click="close" class="close">
                            <i class="icon arrow-left-line icon-primary"></i>
                        </button>
                        <h4>VERIFY DNS RECORDS · Step 2 of 2</h4>
                    </header>
                    <article>
                        <div>
                            <div class="mb-5">
                                <p>
                                    Set the following DNS records for your domain. It may take some time for the DNS changes to propagate.
                                </p>
                            </div>
                            <div class="mb-5">
                                <table class="sm desktop">
                                    <thead>
                                        <tr>
                                            <th>Type</th>
                                            <th>Host</th>
                                            <th>Value</th>
                                        </tr>
                                    </thead>
                                    <tbody>
                                        <template v-for="(mx_host, index) in config.mx_hosts" :key="mx_host">
                                            <tr>
                                                <td>MX {{ 10 * (index + 1) }}</td>
                                                <td>
                                                    <div class="hs-tooltip inline-block">
                                                        <div class="hs-tooltip-toggle">
                                                            <button class="plain truncate max-w-[320px] text-[13px] p-0" @click="copy('@').catch(() => toast('Failed to copy', 'error'))">
                                                                @
                                                            </button>
                                                            <span class="hs-tooltip-content hs-tooltip-shown:opacity-100 hs-tooltip-shown:visible" role="tooltip">
                                                                {{ copied === '@' ? 'Copied' : 'Click to copy' }}
                                                            </span>
                                                        </div>
                                                    </div>
                                                </td>
                                                <td>
                                                    <div class="hs-tooltip inline-block">
                                                        <div class="hs-tooltip-toggle">
                                                            <button class="plain truncate max-w-[320px] text-[13px] p-0" @click="copy(mx_host + '.').catch(() => toast('Failed to copy', 'error'))">
                                                                {{ mx_host }}.
                                                            </button>
                                                            <span class="hs-tooltip-content hs-tooltip-shown:opacity-100 hs-tooltip-shown:visible" role="tooltip">
                                                                {{ copied === mx_host + '.' ? 'Copied' : 'Click to copy' }}
                                                            </span>
                                                        </div>
                                                    </div>
                                                </td>
                                            </tr>
                                        </template>
                                        <tr>
                                            <td>TXT</td>
                                            <td>
                                                <div class="hs-tooltip inline-block">
                                                    <div class="hs-tooltip-toggle">
                                                        <button class="plain truncate max-w-[320px] text-[13px] p-0" @click="copy('@').catch(() => toast('Failed to copy', 'error'))">
                                                            @
                                                        </button>
                                                        <span class="hs-tooltip-content hs-tooltip-shown:opacity-100 hs-tooltip-shown:visible" role="tooltip">
                                                            {{ copied === '@' ? 'Copied' : 'Click to copy' }}
                                                        </span>
                                                    </div>
                                                </div>
                                            </td>
                                            <td>
                                                <div class="hs-tooltip inline-block">
                                                        <div class="hs-tooltip-toggle">
                                                        <button class="plain truncate max-w-[320px] text-[13px] p-0" @click="copy('v=spf1 include:spf.' + config.domain + ' -all').catch(() => toast('Failed to copy', 'error'))">
                                                            v=spf1 include:spf.{{ config.domain }} -all
                                                        </button>
                                                        <span class="hs-tooltip-content hs-tooltip-shown:opacity-100 hs-tooltip-shown:visible" role="tooltip">
                                                            {{ copied === 'v=spf1 include:spf.' + config.domain + ' -all' ? 'Copied' : 'Click to copy' }}
                                                        </span>
                                                    </div>
                                                </div>
                                            </td>
                                        </tr>
                                        <template v-for="selector in config.dkim_selectors" :key="selector">
                                            <tr>
                                                <td>CNAME</td>
                                                <td>
                                                    <div class="hs-tooltip inline-block">
                                                        <div class="hs-tooltip-toggle">
                                                            <button class="plain truncate max-w-[320px] text-[13px] p-0" @click="copy(selector + '._domainkey').catch(() => toast('Failed to copy', 'error'))">
                                                                {{ selector }}._domainkey
                                                            </button>
                                                            <span class="hs-tooltip-content hs-tooltip-shown:opacity-100 hs-tooltip-shown:visible" role="tooltip">
                                                                {{ copied === selector + '._domainkey' ? 'Copied' : 'Click to copy' }}
                                                            </span>
                                                        </div>
                                                    </div>
                                                </td>
                                                <td>
                                                    <div class="hs-tooltip inline-block">
                                                        <div class="hs-tooltip-toggle">
                                                            <button class="plain truncate max-w-[320px] text-[13px] p-0" @click="copy(selector + '._domainkey.' + config.domain + '.').catch(() => toast('Failed to copy', 'error'))">
                                                                {{ selector }}._domainkey.{{ config.domain }}.
                                                            </button>
                                                            <span class="hs-tooltip-content hs-tooltip-shown:opacity-100 hs-tooltip-shown:visible" role="tooltip">
                                                                {{ copied === selector + '._domainkey.' + config.domain + '.' ? 'Copied' : 'Click to copy' }}
                                                            </span>
                                                        </div>
                                                    </div>
                                                </td>
                                            </tr>
                                        </template>
                                        <tr>
                                            <td>TXT</td>
                                            <td>
                                                <div class="hs-tooltip inline-block">
                                                    <div class="hs-tooltip-toggle">
                                                        <button class="plain truncate max-w-[320px] text-[13px] p-0" @click="copy('_dmarc').catch(() => toast('Failed to copy', 'error'))">
                                                            _dmarc
                                                        </button>
                                                        <span class="hs-tooltip-content hs-tooltip-shown:opacity-100 hs-tooltip-shown:visible" role="tooltip">
                                                            {{ copied === '_dmarc' ? 'Copied' : 'Click to copy' }}
                                                        </span>
                                                    </div>
                                                </div>
                                            </td>
                                            <td>
                                                <div class="hs-tooltip inline-block">
                                                    <div class="hs-tooltip-toggle">
                                                        <button class="plain truncate max-w-[320px] text-[13px] p-0" @click="copy('v=DMARC1; p=quarantine; adkim=s').catch(() => toast('Failed to copy', 'error'))">
                                                            v=DMARC1; p=quarantine; adkim=s
                                                        </button>
                                                        <span class="hs-tooltip-content hs-tooltip-shown:opacity-100 hs-tooltip-shown:visible" role="tooltip">
                                                            {{ copied === 'v=DMARC1; p=quarantine; adkim=s' ? 'Copied' : 'Click to copy' }}
                                                        </span>
                                                    </div>
                                                </div>
                                            </td>
                                        </tr>
                                    </tbody>
                                </table>
                                <div class="tablet">
                                    <p class="font-secondary text-sm leading-[2rem] text-black dark:text-white">
                                        <template v-for="(mx_host, index) in config.mx_hosts" :key="mx_host">
                                            MX {{ 10 * (index + 1) }} {{ mx_host }}.<br>
                                        </template>
                                        TXT @ v=spf1 include:spf.{{ config.domain }}. -all <br>
                                        <template v-for="selector in config.dkim_selectors" :key="selector">
                                            CNAME {{ selector }}._domainkey {{ selector }}._domainkey.{{ config.domain }}. <br>
                                        </template>
                                        TXT _dmarc v=DMARC1; p=quarantine; adkim=s
                                    </p>
                                </div>
                            </div>
                        </div>
                    </article>
                    <footer>
                        <nav>
                            <button @click.stop="verifyDomain" :disabled="saving" :aria-busy="saving" class="cta">
                                Verify DNS Records
                            </button>
                            <button @click="close" class="cancel">
                                Cancel
                            </button>
                        </nav>
                        <p v-if="error" class="error px-5" role="alert">Error: {{ error }}</p>
                    </footer>
                </div>
            </div>
        </div>
    </div> 
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { ApiError } from '../api/api.ts'
import overlay from '@preline/overlay'
import { domainApi } from '../api/domain.ts'
import events from '../events.ts'
import tooltip from '@preline/tooltip'
import { useClipboard } from '../composables/useClipboard.ts'
import { toast } from '../composables/useToast.ts'

const props = defineProps(['domain'])
const domain = ref(props.domain)
const error = ref('')
const saving = ref(false)
const { copied, copy } = useClipboard(2000)

const config = ref({
    verify: '',
    domain: '',
    dkim_selectors: [] as string[],
    mx_hosts: [] as string[],
})

const getConfig = async () => {
    try {
        const res = await domainApi.getConfig()
        config.value = res.data
        setTimeout(() => {
            tooltip.autoInit()
        }, 0)
    } catch (err) {
        if (err instanceof ApiError) {
            error.value = err.data?.error || err.message || err.message
        }
    }
}

const verifyDomain = async () => {
    saving.value = true
    try {
        await domainApi.verifyDns(domain.value.id)
        error.value = ''
        close()
    } catch (err) {
        if (err instanceof ApiError) {
            error.value = err.data?.error || err.message || err.message
        }
    } finally {
        saving.value = false
    }
}

const close = () => {
    error.value = ''
    const modal = document.querySelector('#modal-verify-domain' + domain.value.id) as any
    overlay.close(modal)
    events.emit('domain.reload', {})
}

const addEvents = () => {
    const modal = overlay.getInstance('#modal-verify-domain' + domain.value.id as any, true) as any
    modal.element.on('close', () => {
        close()
    })
    modal.element.on('open', () => {
        tooltip.autoInit()
    })
}

onMounted(() => {
    overlay.autoInit()
    getConfig()
    addEvents()
})
</script>