<template>
    <header class="flex items-center justify-between py-5">
        <router-link to="/account" class="p-0">
            <h1 class="px-5 md:px-8 m-0 text-accent head flex items-center justify-between">
                <span class="logo"></span>
            </h1>
        </router-link>
        <nav>
            <a @click.stop="logout" class="p-0 px-5 md:px-8">
                <i class="icon logout icon-primary"></i>
                Log out
            </a>
        </nav>
    </header>
</template>

<script setup lang="ts">
import { ref, onMounted, onUnmounted } from 'vue'
import { userApi } from '../api/user.ts'
import events from '../events.ts'
import { appConfirm } from '../composables/useConfirm.ts'

const email = ref(localStorage.getItem('email'))

const logout = async () => {
    if (!await appConfirm('End your session?')) return
    try {
        await userApi.logout()
        userApi.clearSession()
    } catch { }
}

const onUpdateEmail = (event: any) => {
    email.value = event.email
}

onMounted(() => {
    events.on('user.update', onUpdateEmail)
})

onUnmounted(() => {
    events.off('user.update', onUpdateEmail)
})
</script>