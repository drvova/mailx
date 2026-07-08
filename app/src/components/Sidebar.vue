<template>
    <header class="bg-secondary flex flex-col justify-between h-full">
        <nav>
            <router-link to="/account" class="p-0">
                <h1 class="pl-6 pr-5 m-0 text-accent head flex items-center justify-between">
                    <span class="logo"></span>
                </h1>
            </router-link>
            <div class="flex flex-col items-center">
                <router-link v-bind:class="{ 'active': route == '/account' && !route.startsWith('/account/') }" :aria-current="route == '/account' && !route.startsWith('/account/') ? 'page' : undefined" to="/account">
                    <i class="icon at icon-primary"></i>
                    Aliases
                </router-link>
                <router-link v-bind:class="{ 'active': route == '/account/wildcard' }" :aria-current="route == '/account/wildcard' ? 'page' : undefined" to="/account/wildcard">
                    <i class="icon scan icon-primary"></i>
                    Wildcard
                </router-link>
                <router-link v-bind:class="{ 'active': route == '/account/inbox' }" :aria-current="route == '/account/inbox' ? 'page' : undefined" to="/account/inbox">
                    <i class="icon inbox icon-primary"></i>
                    Temp Mail
                </router-link>
                <router-link v-bind:class="{ 'active': route == '/account/recipients' }" :aria-current="route == '/account/recipients' ? 'page' : undefined" to="/account/recipients">
                    <i class="icon inbox icon-primary"></i>
                    Recipients
                </router-link>
                <!-- <router-link v-bind:class="{ 'active': route == '/account/domains' }" to="/account/domains">
                    <i class="icon global icon-primary"></i>
                    Domains
                </router-link> -->
                <router-link v-bind:class="{ 'active': route == '/account/stats' }" :aria-current="route == '/account/stats' ? 'page' : undefined" to="/account/stats">
                    <i class="icon chart icon-primary"></i>
                    Stats
                </router-link>
                <router-link v-bind:class="{ 'active': route == '/account/diagnostics' }" :aria-current="route == '/account/diagnostics' ? 'page' : undefined" to="/account/diagnostics">
                    <i class="icon alert icon-primary"></i>
                    Diagnostics
                </router-link>
                <router-link v-bind:class="{ 'active': route == '/account/settings' }" :aria-current="route == '/account/settings' ? 'page' : undefined" to="/account/settings">
                    <i class="icon settings icon-primary"></i>
                    Settings
                </router-link>
                <router-link v-bind:class="{ 'active': route == '/account/profile' }" :aria-current="route == '/account/profile' ? 'page' : undefined" to="/account/profile">
                    <i class="icon user icon-primary"></i>
                    Account
                </router-link>
            </div>
        </nav>
        <div>
            <nav>
                <div class="flex items-center py-5 pb-3 pr-5">
                    <a @click.stop="logout">
                        <i class="icon logout icon-primary"></i>
                        Log out
                    </a>
                    <ThemeSwitch />
                </div>
            </nav>
            <p class="px-5 mt-0 pl-6 text-sm">
                Support:
                <a href="mailto:support@freethemail.net">Email</a> /
                <a href="/faq">FAQ</a>
            </p>
        </div>
    </header>
</template>

<script setup lang="ts">
import { ref, onMounted, onUnmounted, watch } from 'vue'
import { useRoute } from 'vue-router'
import { userApi } from '../api/user.ts'
import events from '../events.ts'
import ThemeSwitch from './ThemeSwitch.vue'
import { appConfirm } from '../composables/useConfirm.ts'

const route = ref('/')
const currentRoute = useRoute()
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

watch(currentRoute, (newRoute) => {
    route.value = newRoute.path
}, { immediate: true })
</script>