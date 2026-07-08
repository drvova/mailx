<template>
    <div>
        <div v-bind:id="'modal-add-key-recipient' + recipient.id" class="hs-overlay hidden">
            <div>
                <div>
                    <header>
                        <button @click="close" class="close">
                            <i class="icon arrow-left-line icon-primary"></i>
                        </button>
                        <h4>ADD PGP PUBLIC KEY</h4>
                    </header>
                    <article>
                        <label for="recipient_pgp">
                            Enter your public PGP key:
                        </label>
                        <textarea
                            v-model="pgp_key"
                            v-bind:class="{ 'error': pgpError }"
                            id="recipient_pgp"
                            placeholder="Starts with '-----BEGIN PGP PUBLIC KEY BLOCK-----'"
                        >
                        </textarea>
                        <p v-if="pgpError" class="error" role="alert">Required</p>
                    </article>
                    <footer>
                        <nav>
                            <button @click="addKey" :disabled="saving" :aria-busy="saving" class="cta">
                                Add PGP Public Key
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
import { recipientApi } from '../api/recipient.ts'
import events from '../events.ts'
import { toast } from '../composables/useToast.ts'

const props = defineProps(['recipient'])
const recipient = ref(props.recipient)
const pgp_key = ref('')
const error = ref('')
const pgpError = ref(false)
const saving = ref(false)

const validatePgp = () => {
    pgpError.value = !pgp_key.value
    return !pgpError.value
}

const addKey = async () => {
    if (!validatePgp()) {
        return
    }

    const payload = {
        id: recipient.value.id,
        pgp_enabled: true,
        pgp_key: pgp_key.value.trim()
    }

    saving.value = true
    try {
        await recipientApi.update(payload)
        error.value = ''
        toast('PGP key added')
        events.emit('recipient.update', {})
        close()
    } catch (err) {
        if (err instanceof ApiError) {
            error.value = err.data?.error || err.message || err.message

            if (err.status === 429) {
                error.value = 'Too many requests, please try again later.'
            }
        }
    } finally {
        saving.value = false
    }
}

const close = () => {
    error.value = ''
    pgpError.value = false
    pgp_key.value = ''
    const modal = document.querySelector('#modal-add-key-recipient' + recipient.value.id) as any
    overlay.close(modal)
}

const addEvents = () => {
    const modal = overlay.getInstance('#modal-add-key-recipient' + recipient.value.id as any, true) as any
    modal.element.on('close', () => {
        close()
    })
}

onMounted(() => {
    overlay.autoInit()
    addEvents()
})
</script>