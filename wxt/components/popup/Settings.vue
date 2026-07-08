<template>
    <div class="p-4 pb-16">
        <h2 class="m-0">Settings</h2>
        <hr class="my-5">
        <p class="text-sm my-4">Show FreeTheMail button on email input fields:</p>
        <div class="flex items-center">
            <input @change="toggleInputButton(($event.target as HTMLInputElement).checked)" v-bind:checked="preferences.input_button" type="checkbox" aria-label="Show FreeTheMail button on email input fields">
        </div>
        <hr class="my-5">
        <p class="text-sm my-4">Add website description when creating alias via input button:</p>
        <div class="flex items-center">
            <input @change="toggleAddDescription(($event.target as HTMLInputElement).checked)" v-bind:checked="preferences.add_description" type="checkbox" aria-label="Add website description when creating alias via input button">
        </div>
        <hr class="my-5">
        <p class="text-sm my-4">Refresh recipients, domains and defaults:</p>
        <button @click="refreshDefaults" class="cta sm">Refresh Defaults</button>
        <hr class="my-5">
        <p class="text-sm my-4">Log out / delete session:</p>
        <button @click="logout" class="cta sm">Log Out</button>
        <p v-if="error" class="error my-4 mt-5">Error: {{ error }}</p>
        <p v-if="success" class="success my-4 mt-5">{{ success }}</p>
    </div>
</template>

<script lang="ts" setup>
import { api } from '@/lib/api'
import { appConfirm } from '@/lib/useConfirm'
import { store } from '@/lib/store'
import { Defaults, Preferences } from '@/lib/types'

const props = defineProps<{
    apiToken: string
    defaults: Defaults
    preferences: Preferences
}>()

const success = ref('')
const error = ref('')

const refreshDefaults = async () => {
    try {
        const res = await api.fetchDefaults(props.apiToken)
        processResponse(res)
        success.value = 'Defaults refreshed successfully'
        error.value = ''
    } catch (err) {
        console.error('Fetch defaults error:', err)
        success.value = ''
        error.value = err instanceof Error ? err.message : 'An unexpected error occurred'
    }
}

const logout = async () => {
    if (!(await appConfirm('You will need to log in again to use the extension.', { title: 'Log out?', confirmLabel: 'Log out' }))) return

    try {
        await api.logout(props.apiToken)
        store.clearAll()
        error.value = ''
        success.value = 'You have been logged out.'
    } catch (err) {
        console.error('Logout error:', err)
        success.value = ''
        error.value = err instanceof Error ? err.message : 'An unexpected error occurred during logout'
    }
}

const processResponse = (res: any) => {
    const defaults = {
        domain: res.domain,
        domains: res.domains,
        recipient: res.recipient,
        recipients: res.recipients,
        alias_format: res.alias_format,
    }
    store.setDefaults(defaults)
}

const toggleInputButton = async (enabled: boolean) => {
    const newPreferences = { ...props.preferences, input_button: enabled }
    store.setPreferences(newPreferences)
}

const toggleAddDescription = async (enabled: boolean) => {
    const newPreferences = { ...props.preferences, add_description: enabled }
    store.setPreferences(newPreferences)
}
</script>