<template>
    <div class="card-container">
        <header class="head">
            <h2>Account Settings</h2>
        </header>
        <div class="card-primary">
            <nav aria-label="Account sections" role="tablist" aria-orientation="horizontal" class="tabs-router mb-2">
                <button id="account-tab-1" aria-selected="true" data-hs-tab="#account-panel-1"
                    aria-controls="account-panel-1" role="tab">
                    Profile
                </button>
                <button id="account-tab-2" aria-selected="false" data-hs-tab="#account-panel-2"
                    aria-controls="account-panel-2" role="tab">
                    Security
                </button>
                <button id="account-tab-3" aria-selected="false" data-hs-tab="#account-panel-3"
                    aria-controls="account-panel-3" role="tab">
                    Encryption
                </button>
                <button id="account-tab-4" aria-selected="false" data-hs-tab="#account-panel-4"
                    aria-controls="account-panel-4" role="tab">
                    Data
                </button>
                <button id="account-tab-5" aria-selected="false" data-hs-tab="#account-panel-5"
                    aria-controls="account-panel-5" role="tab">
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
import { onMounted, onUpdated } from 'vue'
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

onMounted(() => tabs.autoInit())
onUpdated(() => tabs.autoInit())
</script>
