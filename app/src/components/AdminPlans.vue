<template>
    <div class="card-container">
        <header class="head">
            <h2>Admin</h2>
        </header>

        <div class="card-primary">
            <ul class="flex gap-4 border-b mb-6 overflow-x-auto" role="tablist">
                <li><button :class="tab === 'stats' ? 'font-bold border-b-2' : ''" @click="tab = 'stats'" role="tab">Stats</button></li>
                <li><button :class="tab === 'users' ? 'font-bold border-b-2' : ''" @click="tab = 'users'" role="tab">Users</button></li>
                <li><button :class="tab === 'aliases' ? 'font-bold border-b-2' : ''" @click="tab = 'aliases'" role="tab">Aliases</button></li>
                <li><button :class="tab === 'domains' ? 'font-bold border-b-2' : ''" @click="tab = 'domains'" role="tab">Domains</button></li>
                <li><button :class="tab === 'recipients' ? 'font-bold border-b-2' : ''" @click="tab = 'recipients'" role="tab">Recipients</button></li>
                <li><button :class="tab === 'plans' ? 'font-bold border-b-2' : ''" @click="tab = 'plans'" role="tab">Plans</button></li>
                <li><button :class="tab === 'keys' ? 'font-bold border-b-2' : ''" @click="tab = 'keys'" role="tab">API Keys</button></li>
                <li><button :class="tab === 'sessions' ? 'font-bold border-b-2' : ''" @click="tab = 'sessions'" role="tab">Sessions</button></li>
                <li><button :class="tab === 'passkeys' ? 'font-bold border-b-2' : ''" @click="tab = 'passkeys'" role="tab">Passkeys</button></li>
                <li><button :class="tab === 'inbox' ? 'font-bold border-b-2' : ''" @click="tab = 'inbox'" role="tab">Inbox</button></li>
                <li><button :class="tab === 'logs' ? 'font-bold border-b-2' : ''" @click="tab = 'logs'" role="tab">Logs</button></li>
            </ul>

            <!-- STATS -->
            <div v-if="tab === 'stats'" role="tabpanel">
                <div v-if="stats" class="grid grid-cols-2 md:grid-cols-4 gap-4">
                    <div class="card-secondary text-center"><p class="text-3xl font-bold">{{ stats.total_users }}</p><p class="text-sm text-gray-500">Total Users</p></div>
                    <div class="card-secondary text-center"><p class="text-3xl font-bold">{{ stats.active_users }}</p><p class="text-sm text-gray-500">Active Users</p></div>
                    <div class="card-secondary text-center"><p class="text-3xl font-bold text-red-500">{{ stats.suspended_users }}</p><p class="text-sm text-gray-500">Suspended</p></div>
                    <div class="card-secondary text-center"><p class="text-3xl font-bold text-blue-500">{{ stats.admin_users }}</p><p class="text-sm text-gray-500">Admins</p></div>
                    <div class="card-secondary text-center"><p class="text-3xl font-bold">{{ stats.total_aliases }}</p><p class="text-sm text-gray-500">Aliases</p></div>
                    <div class="card-secondary text-center"><p class="text-3xl font-bold">{{ stats.total_domains }}</p><p class="text-sm text-gray-500">Domains</p></div>
                    <div class="card-secondary text-center"><p class="text-3xl font-bold">{{ stats.total_recipients }}</p><p class="text-sm text-gray-500">Recipients</p></div>
                    <div class="card-secondary text-center"><p class="text-3xl font-bold">{{ stats.total_inbox_messages }}</p><p class="text-sm text-gray-500">Inbox Msgs</p></div>
                    <div class="card-secondary text-center"><p class="text-3xl font-bold">{{ stats.total_subscriptions }}</p><p class="text-sm text-gray-500">Subscriptions</p></div>
                    <div class="card-secondary text-center"><p class="text-3xl font-bold text-green-500">{{ stats.active_subscriptions }}</p><p class="text-sm text-gray-500">Active Subs</p></div>
                    <div class="card-secondary text-center"><p class="text-3xl font-bold">{{ stats.total_logs }}</p><p class="text-sm text-gray-500">Log Entries</p></div>
                    <div class="card-secondary text-center"><p class="text-3xl font-bold">{{ stats.active_plans }}</p><p class="text-sm text-gray-500">Active Plans</p></div>
                    <div class="col-span-2 md:col-span-4 flex gap-2 mt-2">
                        <button class="cta cta-tertiary text-sm" @click="exportUsers">Export Users CSV</button>
                        <button class="cta cta-tertiary text-sm" @click="exportAliases">Export Aliases CSV</button>
                    </div>
                </div>
                <SkeletonRows v-else :rows="3" />
            </div>

            <!-- USERS -->
            <div v-if="tab === 'users'" role="tabpanel">
                <div class="flex gap-2 mb-4">
                    <input v-model="userSearch" placeholder="Search by email..." @input="searchUsersDeb" class="flex-1" />
                    <button class="cta cta-tertiary text-sm" @click="bulkSuspend">Bulk Suspend Selected</button>
                </div>
                <div v-if="users.length" class="overflow-x-auto">
                    <table class="table">
                        <thead><tr><th><input type="checkbox" @click="toggleAllUsers($event)" /></th><th>Email</th><th>Active</th><th>Admin</th><th>Joined</th><th></th></tr></thead>
                        <tbody>
                            <tr v-for="u in users" :key="u.id">
                                <td><input type="checkbox" v-model="(u as any)._selected" /></td>
                                <td><button class="text-blue-500 hover:underline" @click="viewUser(u.id)">{{ u.email }}</button></td>
                                <td><span :class="u.is_active ? 'badge badge-success' : 'badge badge-error'">{{ u.is_active ? 'Active' : 'Suspended' }}</span></td>
                                <td><span :class="u.is_admin ? 'badge badge-success' : 'badge'">{{ u.is_admin ? 'Admin' : 'User' }}</span></td>
                                <td>{{ formatDate(u.created_at) }}</td>
                                <td><div class="flex gap-1">
                                    <button class="cta cta-tertiary text-sm" @click="toggleUser(u)">{{ u.is_active ? 'Suspend' : 'Activate' }}</button>
                                    <button class="cta cta-tertiary text-sm" @click="toggleAdmin(u)">{{ u.is_admin ? 'Revoke' : 'Make Admin' }}</button>
                                    <button class="cta cta-tertiary text-sm text-red-500" @click="deleteUser(u)" :disabled="deleting === u.id" :aria-busy="deleting === u.id">Delete</button>
                                </div></td>
                            </tr>
                        </tbody>
                    </table>
                </div>
                <SkeletonRows v-else :rows="5" />

                <!-- User detail modal -->
                <div v-if="userDetail" class="fixed inset-0 bg-black/50 flex items-center justify-center z-50 p-4" @click.self="userDetail = null">
                    <div class="card-primary max-w-2xl w-full max-h-[80vh] overflow-y-auto">
                        <div class="flex justify-between items-center mb-4">
                            <h3>{{ userDetail.user.email }}</h3>
                            <button class="cta cta-tertiary" @click="userDetail = null">Close</button>
                        </div>
                        <div class="grid grid-cols-2 gap-4 mb-4">
                            <div><p class="text-sm text-gray-500">Status</p><span :class="userDetail.user.is_active ? 'badge badge-success' : 'badge badge-error'">{{ userDetail.user.is_active ? 'Active' : 'Suspended' }}</span></div>
                            <div><p class="text-sm text-gray-500">Role</p><span :class="userDetail.user.is_admin ? 'badge badge-success' : 'badge'">{{ userDetail.user.is_admin ? 'Admin' : 'User' }}</span></div>
                            <div><p class="text-sm text-gray-500">Subscription Tier</p><p>{{ userDetail.subscription.tier || 'self-hosted' }}</p></div>
                            <div><p class="text-sm text-gray-500">Active Until</p><p>{{ userDetail.subscription.active_until ? formatDate(userDetail.subscription.active_until) : 'N/A' }}</p></div>
                        </div>
                        <div class="mb-4">
                            <label class="text-sm text-gray-500">Assign Plan</label>
                            <select v-model="selectedPlan" @change="assignPlan(userDetail.user.id)" class="max-w-xs">
                                <option value="">-- Select Plan --</option>
                                <option v-for="p in plans" :key="p.id" :value="p.id">{{ p.display_name }} ({{ p.price_cents === 0 ? 'Free' : p.price_cents / 100 }})</option>
                            </select>
                        </div>
                        <div class="mb-4 flex gap-2">
                            <button class="cta cta-tertiary text-sm" @click="forceLogout(userDetail.user.id)">Force Logout</button>
                            <button class="cta cta-tertiary text-sm" @click="overrideSub(userDetail.user.id)">Override Subscription</button>
                            <button class="cta cta-tertiary text-sm" @click="disableTotp(userDetail.user.id)">Disable 2FA</button>
                            <button class="cta cta-tertiary text-sm" @click="resetPassword(userDetail.user.id)">Reset Password</button>
                            <button class="cta cta-tertiary text-sm text-red-500" @click="purgeInbox(userDetail.user.id)">Purge Inbox</button>
                        </div>
                        <h4 class="mb-2">Aliases ({{ userDetail.aliases.length }})</h4>
                        <div v-if="userDetail.aliases.length" class="overflow-x-auto mb-4">
                            <table class="table"><thead><tr><th>Name</th><th>Enabled</th></tr></thead><tbody>
                                <tr v-for="a in userDetail.aliases" :key="a.id"><td>{{ a.name }}</td><td><span :class="a.enabled ? 'badge badge-success' : 'badge'">{{ a.enabled ? 'Yes' : 'No' }}</span></td></tr>
                            </tbody></table>
                        </div>
                        <h4 class="mb-2">Recipients ({{ userDetail.recipients.length }})</h4>
                        <div v-if="userDetail.recipients.length" class="overflow-x-auto mb-4">
                            <table class="table"><thead><tr><th>Email</th><th>Active</th><th>PGP</th></tr></thead><tbody>
                                <tr v-for="r in userDetail.recipients" :key="r.id"><td>{{ r.email }}</td><td><span :class="r.is_active ? 'badge badge-success' : 'badge'">{{ r.is_active ? 'Yes' : 'No' }}</span></td><td>{{ r.pgp_enabled ? 'Yes' : 'No' }}</td></tr>
                            </tbody></table>
                        </div>
                        <h4 class="mb-2">Domains ({{ userDetail.domains.length }})</h4>
                        <div v-if="userDetail.domains.length" class="overflow-x-auto">
                            <table class="table"><thead><tr><th>Name</th><th>Enabled</th><th>Verified</th></tr></thead><tbody>
                                <tr v-for="d in userDetail.domains" :key="d.id"><td>{{ d.name }}</td><td><span :class="d.enabled ? 'badge badge-success' : 'badge'">{{ d.enabled ? 'Yes' : 'No' }}</span></td><td>{{ d.mx_verified_at ? 'Yes' : 'No' }}</td></tr>
                            </tbody></table>
                        </div>
                    </div>
                </div>
            </div>

            <!-- ALIASES -->
            <div v-if="tab === 'aliases'" role="tabpanel">
                <div class="flex gap-2 mb-4">
                    <input v-model="aliasSearch" placeholder="Search aliases..." @input="searchAliasesDeb" class="flex-1" />
                </div>
                <div v-if="aliases.length" class="overflow-x-auto">
                    <table class="table">
                        <thead><tr><th>Alias</th><th>Enabled</th><th>Catch-All</th><th>Created</th><th></th></tr></thead>
                        <tbody>
                            <tr v-for="a in aliases" :key="a.id">
                                <td>{{ a.name }}</td>
                                <td><span :class="a.enabled ? 'badge badge-success' : 'badge'">{{ a.enabled ? 'Yes' : 'No' }}</span></td>
                                <td>{{ a.catch_all ? 'Yes' : 'No' }}</td>
                                <td>{{ formatDate(a.created_at) }}</td>
                                <td><div class="flex gap-1">
                                    <button class="cta cta-tertiary text-sm" @click="toggleAlias(a)">{{ a.enabled ? 'Disable' : 'Enable' }}</button>
                                    <button class="cta cta-tertiary text-sm text-red-500" @click="deleteAlias(a)">Delete</button>
                                </div></td>
                            </tr>
                        </tbody>
                    </table>
                </div>
                <SkeletonRows v-else :rows="5" />
            </div>

            <!-- DOMAINS -->
            <div v-if="tab === 'domains'" role="tabpanel">
                <div v-if="domains.length" class="overflow-x-auto">
                    <table class="table">
                        <thead><tr><th>Domain</th><th>Enabled</th><th>Verified</th><th>Created</th><th></th></tr></thead>
                        <tbody>
                            <tr v-for="d in domains" :key="d.id">
                                <td>{{ d.name }}</td>
                                <td><span :class="d.enabled ? 'badge badge-success' : 'badge'">{{ d.enabled ? 'Yes' : 'No' }}</span></td>
                                <td><span :class="d.mx_verified_at ? 'badge badge-success' : 'badge'">{{ d.mx_verified_at ? 'Yes' : 'No' }}</span></td>
                                <td>{{ formatDate(d.created_at) }}</td>
                                <td><div class="flex gap-1">
                                    <button class="cta cta-tertiary text-sm" @click="toggleDomain(d)">{{ d.enabled ? 'Disable' : 'Enable' }}</button>
                                    <button class="cta cta-tertiary text-sm text-red-500" @click="deleteDomain(d)">Delete</button>
                                </div></td>
                            </tr>
                        </tbody>
                    </table>
                </div>
                <SkeletonRows v-else :rows="5" />
            </div>

            <!-- RECIPIENTS -->
            <div v-if="tab === 'recipients'" role="tabpanel">
                <div class="flex gap-2 mb-4">
                    <input v-model="recipientSearch" placeholder="Search by email..." @input="searchRecipientsDeb" class="flex-1" />
                </div>
                <div v-if="recipients.length" class="overflow-x-auto">
                    <table class="table">
                        <thead><tr><th>Email</th><th>Active</th><th>PGP</th><th>Created</th><th></th></tr></thead>
                        <tbody>
                            <tr v-for="r in recipients" :key="r.id">
                                <td>{{ r.email }}</td>
                                <td><span :class="r.is_active ? 'badge badge-success' : 'badge'">{{ r.is_active ? 'Yes' : 'No' }}</span></td>
                                <td>{{ r.pgp_enabled ? 'Yes' : 'No' }}</td>
                                <td>{{ formatDate(r.created_at) }}</td>
                                <td><button class="cta cta-tertiary text-sm text-red-500" @click="deleteRecipient(r)">Delete</button></td>
                            </tr>
                        </tbody>
                    </table>
                </div>
                <SkeletonRows v-else :rows="5" />
            </div>

            <!-- PLANS -->
            <div v-if="tab === 'plans'" role="tabpanel">
                <div class="flex justify-between items-center mb-4">
                    <h3>Plans</h3>
                    <button class="cta" @click="showForm = !showForm" v-if="!showForm">Create Plan</button>
                </div>
                <div v-if="showForm" class="card-secondary mb-6">
                    <h4>{{ editing ? 'Edit Plan' : 'New Plan' }}</h4>
                    <div class="grid grid-cols-2 gap-4 mt-4">
                        <div><label>Name</label><input v-model="form.name" placeholder="pro" /></div>
                        <div><label>Display Name</label><input v-model="form.display_name" placeholder="Pro Plan" /></div>
                        <div><label>Price (cents)</label><input type="number" v-model.number="form.price_cents" placeholder="500" /></div>
                        <div><label>Currency</label><select v-model="form.currency"><option value="usd">USD</option><option value="eur">EUR</option><option value="gbp">GBP</option></select></div>
                        <div><label>Interval</label><select v-model="form.interval"><option value="monthly">Monthly</option><option value="yearly">Yearly</option><option value="one_time">One Time</option></select></div>
                        <div><label>Sort Order</label><input type="number" v-model.number="form.sort_order" placeholder="0" /></div>
                        <div><label>Max Recipients</label><input type="number" v-model.number="form.max_recipients" /></div>
                        <div><label>Max Passkeys</label><input type="number" v-model.number="form.max_credentials" /></div>
                        <div><label>Max Daily Aliases</label><input type="number" v-model.number="form.max_daily_aliases" /></div>
                        <div><label>Max Daily Send/Reply</label><input type="number" v-model.number="form.max_daily_send_reply" /></div>
                        <div><label>Max Sessions</label><input type="number" v-model.number="form.max_sessions" /></div>
                    </div>
                    <div class="flex gap-2 mt-4">
                        <button class="cta" @click="savePlan" :disabled="saving" :aria-busy="saving">{{ editing ? 'Update' : 'Create' }}</button>
                        <button class="cta cta-tertiary" @click="resetForm">Cancel</button>
                    </div>
                    <p v-if="formError" role="alert" class="text-red-500 mt-2">{{ formError }}</p>
                </div>
                <table v-if="plans.length" class="table">
                    <thead><tr><th>Name</th><th>Price</th><th>Interval</th><th>Recipients</th><th>Aliases/day</th><th>Active</th><th></th></tr></thead>
                    <tbody>
                        <tr v-for="plan in plans" :key="plan.id">
                            <td>{{ plan.display_name }} <span class="text-xs text-gray-500">({{ plan.name }})</span></td>
                            <td>{{ plan.price_cents === 0 ? 'Free' : (plan.price_cents / 100) + ' ' + plan.currency }}</td>
                            <td>{{ plan.interval }}</td>
                            <td>{{ plan.max_recipients }}</td>
                            <td>{{ plan.max_daily_aliases }}</td>
                            <td><span :class="plan.is_active ? 'badge badge-success' : 'badge badge-error'">{{ plan.is_active ? 'Active' : 'Inactive' }}</span></td>
                            <td><div class="flex gap-1"><button class="cta cta-tertiary text-sm" @click="editPlan(plan)">Edit</button><button class="cta cta-tertiary text-sm" @click="deletePlan(plan)" :disabled="deleting === plan.id" :aria-busy="deleting === plan.id">Delete</button></div></td>
                        </tr>
                    </tbody>
                </table>
                <div v-else-if="!loadingPlans" class="card-empty"><h3>No plans yet</h3></div>
            </div>

            <!-- API KEYS -->
            <div v-if="tab === 'keys'" role="tabpanel">
                <div v-if="accessKeys.length" class="overflow-x-auto">
                    <table class="table">
                        <thead><tr><th>Name</th><th>User ID</th><th>Expires</th><th>Created</th><th></th></tr></thead>
                        <tbody>
                            <tr v-for="k in accessKeys" :key="k.id">
                                <td>{{ k.name }}</td>
                                <td class="text-xs text-gray-500">{{ k.user_id.slice(0,8) }}...</td>
                                <td>{{ k.expires_at ? formatDate(k.expires_at) : 'Never' }}</td>
                                <td>{{ formatDate(k.created_at) }}</td>
                                <td><button class="cta cta-tertiary text-sm text-red-500" @click="deleteAccessKey(k)">Revoke</button></td>
                            </tr>
                        </tbody>
                    </table>
                </div>
                <SkeletonRows v-else :rows="5" />
            </div>

            <!-- SESSIONS -->
            <div v-if="tab === 'sessions'" role="tabpanel">
                <div v-if="sessions.length" class="overflow-x-auto">
                    <table class="table">
                        <thead><tr><th>Token</th><th>Expires</th><th>Created</th><th></th></tr></thead>
                        <tbody>
                            <tr v-for="s in sessions" :key="s.id">
                                <td class="text-xs text-gray-500 font-mono">{{ s.token.slice(0,16) }}...</td>
                                <td>{{ formatDate(s.expires_at) }}</td>
                                <td>{{ formatDate(s.created_at) }}</td>
                                <td><button class="cta cta-tertiary text-sm text-red-500" @click="deleteSession(s)">Terminate</button></td>
                            </tr>
                        </tbody>
                    </table>
                </div>
                <SkeletonRows v-else :rows="5" />
            </div>

            <!-- PASSKEYS -->
            <div v-if="tab === 'passkeys'" role="tabpanel">
                <div v-if="credentials.length" class="overflow-x-auto">
                    <table class="table">
                        <thead><tr><th>User ID</th><th>Created</th><th></th></tr></thead>
                        <tbody>
                            <tr v-for="cr in credentials" :key="cr.id">
                                <td class="text-xs text-gray-500">{{ cr.user_id.slice(0,8) }}...</td>
                                <td>{{ formatDate(cr.created_at) }}</td>
                                <td><button class="cta cta-tertiary text-sm text-red-500" @click="deleteCredential(cr)">Remove</button></td>
                            </tr>
                        </tbody>
                    </table>
                </div>
                <SkeletonRows v-else :rows="5" />
            </div>

            <!-- INBOX -->
            <div v-if="tab === 'inbox'" role="tabpanel">
                <div v-if="inboxMessages.length" class="overflow-x-auto">
                    <table class="table">
                        <thead><tr><th>From</th><th>Subject</th><th>Alias</th><th>Size</th><th>Date</th><th></th></tr></thead>
                        <tbody>
                            <tr v-for="m in inboxMessages" :key="m.id">
                                <td>{{ m.from_name || m.from }}</td>
                                <td class="max-w-xs truncate">{{ m.subject }}</td>
                                <td class="text-xs text-gray-500">{{ m.alias_id.slice(0,8) }}...</td>
                                <td>{{ Math.round(m.size / 1024) }}KB</td>
                                <td>{{ formatDate(m.created_at) }}</td>
                                <td><button class="cta cta-tertiary text-sm text-red-500" @click="deleteInboxMsg(m)">Delete</button></td>
                            </tr>
                        </tbody>
                    </table>
                </div>
                <SkeletonRows v-else :rows="5" />
            </div>

            <!-- LOGS -->
            <div v-if="tab === 'logs'" role="tabpanel">
                <div class="flex gap-2 mb-4">
                    <select v-model="logFilter" @change="fetchLogsFiltered" class="max-w-xs">
                        <option value="">All Types</option>
                        <option value="bounce">Bounce</option>
                        <option value="disabled_alias">Disabled Alias</option>
                        <option value="disabled_domain">Disabled Domain</option>
                        <option value="unauthorised_send">Unauthorised Send</option>
                        <option value="inactive_subscription">Inactive Subscription</option>
                    </select>
                </div>
                <div v-if="logs.length" class="overflow-x-auto">
                    <table class="table">
                        <thead><tr><th>Type</th><th>From</th><th>Destination</th><th>Status</th><th>Message</th><th>Date</th></tr></thead>
                        <tbody>
                            <tr v-for="log in logs" :key="log.id">
                                <td><span class="badge">{{ log.log_type }}</span></td>
                                <td>{{ log.from }}</td>
                                <td>{{ log.destination }}</td>
                                <td>{{ log.status }}</td>
                                <td class="max-w-xs truncate">{{ log.message }}</td>
                                <td>{{ formatDate(log.created_at) }}</td>
                            </tr>
                        </tbody>
                    </table>
                </div>
                <SkeletonRows v-else :rows="5" />
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref, onMounted, watch } from 'vue'
import { planApi, adminApi, type Plan, type AdminUser, type AdminLog, type AdminAlias, type AdminDomain, type AdminRecipient, type AdminAccessKey, type AdminSession, type AdminCredential, type AdminInboxMessage, type SystemStats } from '../api/plan'
import SkeletonRows from '../components/SkeletonRows.vue'

const tab = ref('stats')
const stats = ref<SystemStats | null>(null)
const users = ref<AdminUser[]>([])
const aliases = ref<AdminAlias[]>([])
const domains = ref<AdminDomain[]>([])
const recipients = ref<AdminRecipient[]>([])
const plans = ref<Plan[]>([])
const logs = ref<AdminLog[]>([])
const accessKeys = ref<AdminAccessKey[]>([])
const sessions = ref<AdminSession[]>([])
const credentials = ref<AdminCredential[]>([])
const inboxMessages = ref<AdminInboxMessage[]>([])
const loadingPlans = ref(true)
const deleting = ref('')
const saving = ref(false)
const showForm = ref(false)
const editing = ref(false)
const formError = ref('')
const userSearch = ref('')
const aliasSearch = ref('')
const recipientSearch = ref('')
const logFilter = ref('')
const userDetail = ref<any>(null)
const selectedPlan = ref('')

let searchTimer: ReturnType<typeof setTimeout>

const emptyForm = (): Partial<Plan> => ({ name: '', display_name: '', price_cents: 0, currency: 'usd', interval: 'monthly', max_recipients: 5, max_credentials: 10, max_daily_aliases: 50, max_daily_send_reply: 50, max_sessions: 5, sort_order: 0 })
const form = ref<Partial<Plan>>(emptyForm())
const formatDate = (d: string) => new Date(d).toLocaleDateString()

const fetchStats = async () => { try { stats.value = await adminApi.stats() } catch { /* */ } }
const fetchUsers = async () => { try { users.value = await adminApi.users() } catch { /* */ } }
const fetchAliases = async () => { try { const r = await adminApi.aliases(aliasSearch.value || undefined); aliases.value = r.aliases } catch { /* */ } }
const fetchDomains = async () => { try { domains.value = await adminApi.domains() } catch { /* */ } }
const fetchRecipients = async () => { try { const r = await adminApi.recipients(recipientSearch.value || undefined); recipients.value = r.recipients } catch { /* */ } }
const fetchPlans = async () => { loadingPlans.value = true; try { plans.value = await planApi.listAll() } catch { /* */ } finally { loadingPlans.value = false } }
const fetchLogs = async () => { try { logs.value = await adminApi.logs() } catch { /* */ } }
const fetchLogsFiltered = async () => { try { const r = await adminApi.logsFiltered(logFilter.value || undefined); logs.value = r.logs } catch { /* */ } }
const fetchAccessKeys = async () => { try { const r = await adminApi.accessKeys(); accessKeys.value = r.keys } catch { /* */ } }
const fetchSessions = async () => { try { const r = await adminApi.sessions(); sessions.value = r.sessions } catch { /* */ } }
const fetchCredentials = async () => { try { const r = await adminApi.credentials(); credentials.value = r.credentials } catch { /* */ } }
const fetchInboxMessages = async () => { try { const r = await adminApi.inboxMessages(); inboxMessages.value = r.messages } catch { /* */ } }

const searchUsersDeb = () => { clearTimeout(searchTimer); searchTimer = setTimeout(async () => { if (!userSearch.value) { fetchUsers(); return }; try { const r = await adminApi.searchUsers(userSearch.value); users.value = r.users } catch { /* */ } }, 300) }
const searchAliasesDeb = () => { clearTimeout(searchTimer); searchTimer = setTimeout(fetchAliases, 300) }
const searchRecipientsDeb = () => { clearTimeout(searchTimer); searchTimer = setTimeout(fetchRecipients, 300) }

const viewUser = async (id: string) => { try { userDetail.value = await adminApi.userDetail(id); selectedPlan.value = userDetail.value.subscription?.plan_id || '' } catch { /* */ } }
const assignPlan = async (userId: string) => { if (!selectedPlan.value) return; try { await adminApi.assignPlan(userId, selectedPlan.value); viewUser(userId) } catch { /* */ } }

const toggleUser = async (u: AdminUser) => { try { await adminApi.updateUser({ id: u.id, is_active: !u.is_active }); u.is_active = !u.is_active } catch { /* */ } }
const toggleAdmin = async (u: AdminUser) => { try { await adminApi.updateUser({ id: u.id, is_admin: !u.is_admin }); u.is_admin = !u.is_admin } catch { /* */ } }
const deleteUser = async (u: AdminUser) => { if (!confirm(`Delete ${u.email}? Removes all data.`)) return; deleting.value = u.id; try { await adminApi.deleteUser(u.id); users.value = users.value.filter(x => x.id !== u.id) } catch { /* */ } finally { deleting.value = '' } }

const toggleAlias = async (a: AdminAlias) => { try { await adminApi.toggleAlias(a.id, !a.enabled); a.enabled = !a.enabled } catch { /* */ } }
const deleteAlias = async (a: AdminAlias) => { if (!confirm(`Delete alias ${a.name}?`)) return; try { await adminApi.deleteAlias(a.id); aliases.value = aliases.value.filter(x => x.id !== a.id) } catch { /* */ } }

const toggleDomain = async (d: AdminDomain) => { try { await adminApi.toggleDomain(d.id, !d.enabled); d.enabled = !d.enabled } catch { /* */ } }
const deleteDomain = async (d: AdminDomain) => { if (!confirm(`Delete domain ${d.name}?`)) return; try { await adminApi.deleteDomain(d.id); domains.value = domains.value.filter(x => x.id !== d.id) } catch { /* */ } }

const deleteRecipient = async (r: AdminRecipient) => { if (!confirm(`Delete recipient ${r.email}?`)) return; try { await adminApi.deleteRecipient(r.id); recipients.value = recipients.value.filter(x => x.id !== r.id) } catch { /* */ } }

const deleteAccessKey = async (k: AdminAccessKey) => { if (!confirm(`Revoke key "${k.name}"?`)) return; try { await adminApi.deleteAccessKey(k.id); accessKeys.value = accessKeys.value.filter(x => x.id !== k.id) } catch { /* */ } }
const deleteSession = async (s: AdminSession) => { if (!confirm('Terminate this session?')) return; try { await adminApi.deleteSession(s.id); sessions.value = sessions.value.filter(x => x.id !== s.id) } catch { /* */ } }
const deleteCredential = async (cr: AdminCredential) => { if (!confirm('Remove this passkey?')) return; try { await adminApi.deleteCredential(cr.id); credentials.value = credentials.value.filter(x => x.id !== cr.id) } catch { /* */ } }
const forceLogout = async (userId: string) => { if (!confirm('Terminate ALL sessions for this user?')) return; try { await adminApi.forceLogout(userId); alert('All sessions terminated') } catch { /* */ } }
const overrideSub = async (userId: string) => {
    const tier = prompt('Enter tier (e.g. pro, free, self-hosted):')
    if (!tier) return
    const activeUntil = prompt('Enter active_until date (YYYY-MM-DD) or leave empty:') || ''
    try { await adminApi.updateSubscription({ user_id: userId, tier, is_active: true, active_until: activeUntil }); viewUser(userId) } catch { /* */ } }
const bulkSuspend = async () => {
    const selected = users.value.filter(u => (u as any)._selected)
    if (!selected.length) { alert('Select users first'); return }
    if (!confirm(`Suspend ${selected.length} users?`)) return
    try { await adminApi.bulkUpdateUsers(selected.map(u => u.id), false); selected.forEach(u => u.is_active = false) } catch { /* */ } }
const toggleAllUsers = (e: Event) => {
    const checked = (e.target as HTMLInputElement).checked
    users.value.forEach(u => { (u as any)._selected = checked })
}
const deleteInboxMsg = async (m: AdminInboxMessage) => { if (!confirm(`Delete message "${m.subject}"?`)) return; try { await adminApi.deleteInboxMessage(m.id); inboxMessages.value = inboxMessages.value.filter(x => x.id !== m.id) } catch { /* */ } }
const disableTotp = async (userId: string) => { if (!confirm('Disable 2FA for this user?')) return; try { await adminApi.disableTotp(userId); alert('2FA disabled') } catch { /* */ } }
const resetPassword = async (userId: string) => {
    const pw = prompt('Enter new password (min 12 chars):')
    if (!pw || pw.length < 12) { alert('Password must be 12+ characters'); return }
    try { await adminApi.resetPassword(userId, pw); alert('Password reset') } catch { /* */ } }
const purgeInbox = async (userId: string) => { if (!confirm('Delete ALL inbox messages for this user?')) return; try { await adminApi.purgeInbox(userId); alert('Inbox purged') } catch { /* */ } }
const exportUsers = () => { window.open(`${import.meta.env.VITE_API_URL}/admin/export/users`, '_blank') }
const exportAliases = () => { window.open(`${import.meta.env.VITE_API_URL}/admin/export/aliases`, '_blank') }

const savePlan = async () => { saving.value = true; formError.value = ''; try { if (editing.value && form.value.id) { await planApi.update(form.value.id, form.value) } else { await planApi.create(form.value) }; resetForm(); await fetchPlans() } catch (err: any) { formError.value = err?.message || 'Failed' } finally { saving.value = false } }
const editPlan = (plan: Plan) => { editing.value = true; showForm.value = true; form.value = { ...plan } }
const deletePlan = async (plan: Plan) => { if (!confirm(`Deactivate "${plan.display_name}"?`)) return; deleting.value = plan.id; try { await planApi.delete(plan.id); await fetchPlans() } catch { /* */ } finally { deleting.value = '' } }
const resetForm = () => { showForm.value = false; editing.value = false; form.value = emptyForm(); formError.value = '' }

watch(tab, (t) => {
    if (t === 'aliases' && !aliases.value.length) fetchAliases()
    if (t === 'domains' && !domains.value.length) fetchDomains()
    if (t === 'recipients' && !recipients.value.length) fetchRecipients()
    if (t === 'keys' && !accessKeys.value.length) fetchAccessKeys()
    if (t === 'sessions' && !sessions.value.length) fetchSessions()
    if (t === 'passkeys' && !credentials.value.length) fetchCredentials()
    if (t === 'inbox' && !inboxMessages.value.length) fetchInboxMessages()
})

onMounted(() => { fetchStats(); fetchUsers(); fetchPlans(); fetchLogs() })
</script>
