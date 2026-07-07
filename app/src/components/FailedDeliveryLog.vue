<template>
    <div>
        <div v-bind:id="'modal-delivery-log' + log.id" class="hs-overlay hidden">
            <div>
                <div>
                    <header>
                        <button @click="close" class="close">
                            <i class="icon arrow-left-line icon-primary"></i>
                        </button>
                        <h4>FAILED DELIVERY LOG</h4>
                    </header>
                    <article>
                        <label for="log_text">
                            Full log:
                        </label>
                        <textarea
                            v-model="log_text"
                            id="log_text"
                            disabled
                        >
                        </textarea>
                    </article>
                    <footer>
                        <nav>
                            <button @click="close" class="cancel">
                                Close
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
import { logApi } from '../api/log.ts'

const props = defineProps(['log'])
const log = ref(props.log)
const log_text = ref('')
const error = ref('')

const getLog = async () => {
    try {
        const res = await logApi.getFile(log.value.id)
        log_text.value = res.data
        error.value = ''
    } catch (err) {
        if (err instanceof ApiError) {
            error.value = err.data?.error || err.message || err.message
        }
    }
}

const close = () => {
    error.value = ''
    const modal = document.querySelector('#modal-delivery-log' + log.value.id) as any
    overlay.close(modal)
}

const addEvents = () => {
    const modal = overlay.getInstance('#modal-delivery-log' + log.value.id as any, true) as any
    modal.element.on('close', () => {
        close()
    })
    modal.element.on('open', () => {
        getLog()
    })
}

onMounted(() => {
    overlay.autoInit()
    addEvents()
})
</script>