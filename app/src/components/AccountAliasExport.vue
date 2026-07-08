<template>
    <div class="mb-5">
        <h2>Alias Export</h2>
        <p>
            Download a complete list of your aliases in a CSV file. Includes status, recipients, and descriptions.
        </p>
        <button
            @click="exportAliases"
            :disabled="exporting"
            :aria-busy="exporting"
            class="cta mb-4">
            Export Aliases
        </button>
        <p v-if="error" class="error" role="alert">Error: {{ error }}</p>
    </div>
</template>

<script setup lang="ts">
import { aliasApi } from '../api/alias'
import { ApiError } from '../api/api.ts'
import { ref } from 'vue'

const error = ref('')
const exporting = ref(false)

const exportAliases = async () => {
    exporting.value = true
    try {
        const res = await aliasApi.export()
        error.value = ''
        const url = window.URL.createObjectURL(new Blob([res.data]))
        const link = document.createElement('a')
        link.href = url
        link.setAttribute('download', 'aliases.csv')
        document.body.appendChild(link)
        link.click()
        link.remove()
        URL.revokeObjectURL(url)
    } catch (err) {
        if (err instanceof ApiError) {
            error.value = err.message
        }
    } finally {
        exporting.value = false
    }
}
</script>