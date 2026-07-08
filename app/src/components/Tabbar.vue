<template>
    <header class="bg-secondary fixed bottom-0 left-0 right-0 z-10">
        <nav class="mobile" role="navigation" aria-label="Primary">
            <div class="flex flex-row items-center">
                <router-link v-bind:class="{ 'active': route == '/account' && !route.startsWith('/account/') }" :aria-current="route == '/account' && !route.startsWith('/account/') ? 'page' : undefined" to="/account" title="Aliases" aria-label="Aliases">
                    <i class="icon at icon-primary"></i>
                </router-link>
                <router-link v-bind:class="{ 'active': route == '/account/inbox' }" :aria-current="route == '/account/inbox' ? 'page' : undefined" to="/account/inbox" title="Temp Mail" aria-label="Temp Mail">
                    <i class="icon inbox icon-primary"></i>
                </router-link>
                <router-link v-bind:class="{ 'active': route == '/account/recipients' }" :aria-current="route == '/account/recipients' ? 'page' : undefined" to="/account/recipients" title="Recipients" aria-label="Recipients">
                    <i class="icon user icon-primary"></i>
                </router-link>
                <router-link v-bind:class="{ 'active': route == '/account/stats' }" :aria-current="route == '/account/stats' ? 'page' : undefined" to="/account/stats" title="Stats" aria-label="Stats">
                    <i class="icon chart icon-primary"></i>
                </router-link>
                <button
                    type="button"
                    class="tabbar-more"
                    :class="{ 'active': isSecondaryActive }"
                    :aria-current="isSecondaryActive ? 'page' : undefined"
                    :aria-expanded="sheetOpen"
                    aria-haspopup="dialog"
                    aria-label="More navigation"
                    data-hs-overlay="#tabbar-more-sheet"
                    @click="onMoreClick">
                    <i class="icon more icon-primary"></i>
                </button>
            </div>
        </nav>

        <div id="tabbar-more-sheet" class="hs-overlay hidden" role="dialog" aria-modal="true" aria-label="More navigation">
            <div>
                <div>
                    <header>
                        <h4>More</h4>
                        <button type="button" class="close" aria-label="Close" @click="closeSheet">
                            <i class="icon close icon-primary"></i>
                        </button>
                    </header>
                    <article>
                        <nav class="tabbar-sheet-nav" role="navigation" aria-label="Secondary">
                            <router-link v-bind:class="{ 'active': route == '/account/wildcard' }" :aria-current="route == '/account/wildcard' ? 'page' : undefined" to="/account/wildcard" @click="closeSheet">
                                <i class="icon scan icon-primary"></i>
                                Wildcard
                            </router-link>
                            <router-link v-bind:class="{ 'active': route == '/account/diagnostics' }" :aria-current="route == '/account/diagnostics' ? 'page' : undefined" to="/account/diagnostics" @click="closeSheet">
                                <i class="icon alert icon-primary"></i>
                                Diagnostics
                            </router-link>
                            <router-link v-bind:class="{ 'active': route == '/account/settings' }" :aria-current="route == '/account/settings' ? 'page' : undefined" to="/account/settings" @click="closeSheet">
                                <i class="icon settings icon-primary"></i>
                                Settings
                            </router-link>
                            <router-link v-bind:class="{ 'active': route == '/account/profile' }" :aria-current="route == '/account/profile' ? 'page' : undefined" to="/account/profile" @click="closeSheet">
                                <i class="icon user icon-primary"></i>
                                Account
                            </router-link>
                            <router-link v-if="isAdmin" v-bind:class="{ 'active': route == '/account/admin' }" :aria-current="route == '/account/admin' ? 'page' : undefined" to="/account/admin" @click="closeSheet">
                                <i class="icon key icon-primary"></i>
                                Admin
                            </router-link>
                        </nav>
                    </article>
                </div>
            </div>
        </div>
    </header>
</template>

<script setup lang="ts">
import { computed, nextTick, onMounted, ref, watch } from 'vue'
import { useRoute } from 'vue-router'
import overlay from '@preline/overlay'
import { userApi } from '../api/user.ts'

const route = ref('/')
const currentRoute = useRoute()
const sheetOpen = ref(false)
const isAdmin = ref(false)
let previousFocus: HTMLElement | null = null

const secondaryPaths = ['/account/wildcard', '/account/diagnostics', '/account/settings', '/account/profile', '/account/admin']
const isSecondaryActive = computed(() => secondaryPaths.includes(route.value))

watch(currentRoute, (newRoute) => {
    route.value = newRoute.path
}, { immediate: true })

onMounted(async () => {
    try {
        const res = await userApi.get()
        isAdmin.value = res.data?.is_admin ?? false
    } catch { /* not logged in */ }
})

const closeSheet = () => {
    const sheet = document.querySelector('#tabbar-more-sheet') as HTMLElement | null
    if (sheet) overlay.close(sheet)
    sheetOpen.value = false
    previousFocus?.focus()
    previousFocus = null
}

const onMoreClick = async () => {
    previousFocus = document.activeElement as HTMLElement | null
    sheetOpen.value = true
    await nextTick()
    const firstLink = document.querySelector('#tabbar-more-sheet .tabbar-sheet-nav a') as HTMLElement | null
    firstLink?.focus()
}
</script>
