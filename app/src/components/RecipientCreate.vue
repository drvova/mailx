<template>
    <div>
        <button v-bind:data-hs-overlay="'#modal-create-recipient'" class="cta">
            New Recipient
        </button>
        <div v-bind:id="'modal-create-recipient'" class="hs-overlay hidden">
            <div>
                <div>
                    <header>
                        <button @click="close" class="close">
                            <i class="icon arrow-left-line icon-primary"></i>
                        </button>
                        <h4>NEW RECIPIENT</h4>
                    </header>
                    <article>
                        <div class="mb-5">
                            <p>
                                Add a new email address to receive forwarded emails. New addresses need a one-time verification before use.
                            </p>
                        </div>
                        <div class="mb-5">
                            <label for="recipient_email">
                                Email:
                            </label>
                            <input
                                v-model="recipient.email"
                                v-bind:class="{ 'error': emailError }"
                                id="recipient_email"
                                placeholder="name@example.net"
                                type="text"
                            >
                            <p v-if="emailError" class="error" role="alert">Required</p>
                        </div>
                    </article>
                    <footer>
                        <nav>
                            <button @click="postRecipient" :disabled="saving" :aria-busy="saving" class="cta">
                                Add Recipient
                            </button>
                            <button @click="close" class="cancel">
                                Cancel
                            </button>
                        </nav>
                        <p v-if="error" class="error px-5" role="alert">Error: {{ error }}</p>
                        <p class="px-5">
                            Note: Unverified recipient email addresses are automatically deleted 7 days after creation. You can add up to 10 recipients.
                        </p>
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

const recipient = ref({
    email: '',
})
const error = ref('')
const emailError = ref(false)
const saving = ref(false)

const validateEmail = () => {
    emailError.value = !recipient.value.email
    return !emailError.value
}

const postRecipient = async () => {
    if (!validateEmail()) {
        return
    }

    try {
        saving.value = true
        await recipientApi.create(recipient.value)
        toast('Recipient added')
        error.value = ''
        events.emit('recipient.create', {})
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
    recipient.value = {} as any
    error.value = ''
    emailError.value = false
    document.removeEventListener('keydown', handleKeydown)
    const modal = document.querySelector('#modal-create-recipient') as any
    overlay.close(modal)
}

const addEvents = () => {
    const modal = overlay.getInstance('#modal-create-recipient' as any, true) as any
    modal.element.on('close', () => {
        close()
    })
    modal.element.on('open', () => {
        document.addEventListener('keydown', handleKeydown)
        focusFirstInput()
    })
}

const focusFirstInput = () => {
    const input = document.getElementById('recipient_email')
    input?.focus()
}

const handleKeydown = (event: KeyboardEvent) => {
    if (event.key === 'Enter') {
        event.preventDefault()
        postRecipient()
    }
}

onMounted(() => {
    overlay.autoInit()
    addEvents()
})
</script>