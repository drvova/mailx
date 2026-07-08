<template>
    <div class="card-container">
        <header class="head">
            <h2>Account Settings</h2>
        </header>
        <div class="card-primary">
            <nav aria-label="Account sections" role="tablist" aria-orientation="horizontal" class="tabs-router mb-2"
                @keydown.left.prevent="focusTab(-1)" @keydown.right.prevent="focusTab(1)"
                @keydown.home.prevent="focusTab(0, true)" @keydown.end.prevent="focusTab(0, false, true)">
                <button id="account-tab-1" aria-selected="true" data-hs-tab="#account-panel-1"
                    aria-controls="account-panel-1" role="tab" tabindex="0">
                    Profile
                </button>
                <button id="account-tab-2" aria-selected="false" data-hs-tab="#account-panel-2"
                    aria-controls="account-panel-2" role="tab" tabindex="-1">
                    Security
                </button>
                <button id="account-tab-3" aria-selected="false" data-hs-tab="#account-panel-3"
                    aria-controls="account-panel-3" role="tab" tabindex="-1">
                    Encryption
                </button>
                <button id="account-tab-4" aria-selected="false" data-hs-tab="#account-panel-4"
                    aria-controls="account-panel-4" role="tab" tabindex="-1">
                    Data
                </button>
                <button id="account-tab-5" aria-selected="false" data-hs-tab="#account-panel-5"
                    aria-controls="account-panel-5" role="tab" tabindex="-1">
                    Danger
                </button>
            </nav>

            <div id="account-panel-1" role="tabpanel" aria-labelledby="account-tab-1">
                <AccountVerify />
                <hr>
                <AccountSubscription />
                <hr>
                <AccountChangeEmail />
            </div>
            <div id="account-panel-2" class="hidden" role="tabpanel" aria-labelledby="account-tab-2">
                <AccountChangePassword />
                <hr>
                <AccountTotp />
                <hr>
                <AccountPasskeys />
            </div>
            <div id="account-panel-3" class="hidden" role="tabpanel" aria-labelledby="account-tab-3">
                <AccountAccessKeys />
            </div>
            <div id="account-panel-4" class="hidden" role="tabpanel" aria-labelledby="account-tab-4">
                <AccountAliasExport />
            </div>
            <div id="account-panel-5" class="hidden" role="tabpanel" aria-labelledby="account-tab-5">
                <AccountDelete />
            </div>

            <hr>
            <h2>Support</h2>
            <p class="mt-0">
                Support:
                <a href="mailto:support@freethemail.net">Email</a> /
                <a href="/faq">FAQ</a>
            </p>
        </div>
    </div>
</template>

<script setup lang="ts">
import { onMounted, onUpdated, ref } from 'vue'
import tabs from '@preline/tabs'
import AccountVerify from './AccountVerify.vue'
import AccountSubscription from './AccountSubscription.vue'
import AccountChangeEmail from './AccountChangeEmail.vue'
import AccountChangePassword from './AccountChangePassword.vue'
import AccountTotp from './AccountTotp.vue'
import AccountPasskeys from './AccountPasskeys.vue'
import AccountAccessKeys from './AccountAccessKeys.vue'
import AccountAliasExport from './AccountAliasExport.vue'
import AccountDelete from './AccountDelete.vue'

const activeTab = ref(0)

// ARIA tablist keyboard navigation (WCAG 2.1.1)
const focusTab = (dir: number, home = false, end = false) => {
    const tabsEls = [...document.querySelectorAll('[role="tab"]')] as HTMLButtonElement[]
    const current = tabsEls.findIndex(t => t.id === `account-tab-${activeTab.value + 1}`)
    let next: number
    if (home) next = 0
    else if (end) next = tabsEls.length - 1
    else next = (current + dir + tabsEls.length) % tabsEls.length
    tabsEls[next]?.focus()
}

onMounted(() => {
    tabs.autoInit()
    // Track which tab is active for tabindex management
    document.querySelectorAll('[role="tab"]').forEach((tab, i) => {
        tab.addEventListener('click', () => {
            activeTab.value = i
            updateTabindex()
        })
    })
})

const updateTabindex = () => {
    document.querySelectorAll('[role="tab"]').forEach((tab, i) => {
        (tab as HTMLButtonElement).tabIndex = i === activeTab.value ? 0 : -1
    })
}

onUpdated(() => {
    tabs.autoInit()
    updateTabindex()
})
</script>
