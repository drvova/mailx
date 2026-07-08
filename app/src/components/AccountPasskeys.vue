<template>
    <div class="mb-5">
        <h2>Passkeys</h2>
        <div v-if="passkeySupported">
            <p>
                Add or remove Passkeys associated with your account.<br>
            </p>
            <div class="flex justify-start items-center gap-x-3 mb-3">
                <button @click="addPasskey" :disabled="adding" :aria-busy="adding" class="cta">
                    New Passkey
                </button>
            </div>
            <p v-if="error" class="error mt-6 mb-4" role="alert">Error: {{ error }}</p>
        </div>
        <div v-if="!passkeySupported">
            <p>
                Your browser/device does not support adding Passkeys.<br>
            </p>
        </div>
        <div v-if="list.length" class="table-container">
            <table>
                <thead class="desktop">
                    <tr>
                        <th>Created</th>
                        <th>ID</th>
                        <th>Actions</th>
                    </tr>
                </thead>
                <tbody>
                    <tr v-for="cred in list" :key="`desktop-${cred.id}`" class="desktop">
                        <td>
                            {{ formatDistanceToNow(new Date(cred.created_at)) }} ago
                        </td>
                        <td>
                            {{ cred.id }}
                        </td>
                        <td>
                            <button @click.stop="deleteCred(cred.id)" class="delete w-full flex items-center gap-x-2 py-2 place-content-end">
                                <i class="icon icon-error trash text-xs"></i>
                                Delete
                            </button>
                        </td>
                    </tr>
                    <tr v-for="cred in list" :key="`tablet-${cred.id}`" class="tablet">
                        <hr>
                        <div class="flex gap-2 justify-between">
                            <div class="text-start">
                                <p class="mb-4 text-sm">{{ formatDistanceToNow(new Date(cred.created_at)) }} ago</p>
                                <div>
                                    <p class="mb-1 text-sm">ID:</p>
                                    {{ cred.id }}
                                </div>
                            </div>
                            <div class="text-end">
                                    <button @click.stop="deleteCred(cred.id)" class="delete w-full flex items-center gap-x-2 py-2 place-content-end">
                                        <i class="icon icon-error trash text-xs"></i>
                                    </button>
                            </div>
                        </div>
                    </tr>
                </tbody>
            </table>
        </div>
    </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { formatDistanceToNow } from 'date-fns'
import { ApiError } from '../api/api.ts'
import { userApi } from '../api/user.ts'
import { startRegistration, browserSupportsWebAuthn } from '@simplewebauthn/browser'
import { appConfirm } from '../composables/useConfirm.ts'
import { toast } from '../composables/useToast.ts'

const credential = {
    id: '',
    created_at: '',
}

const list = ref([] as typeof credential[])
const error = ref('')
const passkeySupported = ref(false)
const adding = ref(false)

const getList = async () => {
    try {
        const res = await userApi.getCredentials()
        list.value = res.data
        error.value = ''
    } catch (err) {
        if (err instanceof ApiError) {
            error.value = err.message
        }
    }
}

const deleteCred = async (id: string) => {
    if (!(await appConfirm('You will no longer be able to log in with this passkey.', { title: 'Delete this passkey?', confirmLabel: 'Delete passkey' }))) return

    try {
        await userApi.deleteCredential(id)
        list.value = list.value.filter((cred: any) => cred.id !== id)
        error.value = ''
        toast('Passkey deleted')
    } catch (err) {
        if (err instanceof ApiError) {
            error.value = err.message
        }
    }
}

const addPasskey = async () => {
    try {
        adding.value = true
        const res = await userApi.registerAdd()
        startAddPasskey(res)
    } catch (err) {
        if (err instanceof ApiError) {
            error.value = err.data?.error || err.message || err.message

            if (err.status === 429) {
                error.value = 'Too many requests, please try again later.'
            }
        }
    } finally {
        adding.value = false
    }
}

const startAddPasskey = async (res: any) => {
    try {
        const creds = await startRegistration({ optionsJSON: res.data['publicKey'] })
        res = await userApi.registerAddFinish(creds)
        error.value = ''
        getList()
    } catch (err: any) {
        if (err instanceof ApiError) {
            error.value = err.data?.error || err.message || err.message

            if (err.status === 429) {
                error.value = 'Too many requests, please try again later.'
            }
        } else {
            error.value = 'The operation was aborted or failed.'
        }
    }
}

onMounted(() => {
    getList()
    passkeySupported.value = browserSupportsWebAuthn()
})
</script>