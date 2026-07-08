<template>
    <div class="card-container">
        <header class="head">
            <h2>Admin</h2>
        </header>

        <div class="card-primary">
            <ul class="flex gap-4 border-b mb-6" role="tablist">
                <li><button :class="tab === 'stats' ? 'font-bold border-b-2' : ''" @click="tab = 'stats'" role="tab">Stats</button></li>
                <li><button :class="tab === 'users' ? 'font-bold border-b-2' : ''" @click="tab = 'users'" role="tab">Users</button></li>
                <li><button :class="tab === 'plans' ? 'font-bold border-b-2' : ''" @click="tab = 'plans'" role="tab">Plans</button></li>
                <li><button :class="tab === 'logs' ? 'font-bold border-b-2' : ''" @click="tab = 'logs'" role="tab">System Logs</button></li>
            </ul>

            <div v-if="tab === 'stats'" role="tabpanel">
                <div v-if="stats" class="grid grid-cols-2 md:grid-cols-3 gap-4">
                    <div class="card-secondary text-center">
                        <p class="text-3xl font-bold">{{ stats.total_users }}</p>
                        <p class="text-sm text-gray-500">Total Users</p>
                    </div>
                    <div class="card-secondary text-center">
                        <p class="text-3xl font-bold">{{ stats.active_users }}</p>
                        <p class="text-sm text-gray-500">Active Users</p>
                    </div>
                    <div class="card-secondary text-center">
                        <p class="text-3xl font-bold">{{ stats.total_aliases }}</p>
                        <p class="text-sm text-gray-500">Aliases</p>
                    </div>
                    <div class="card-secondary text-center">
                        <p class="text-3xl font-bold">{{ stats.total_domains }}</p>
                        <p class="text-sm text-gray-500">Domains</p>
                    </div>
                    <div class="card-secondary text-center">
                        <p class="text-3xl font-bold">{{ stats.total_logs }}</p>
                        <p class="text-sm text-gray-500">Log Entries</p>
                    </div>
                    <div class="card-secondary text-center">
                        <p class="text-3xl font-bold">{{ stats.active_plans }}</p>
                        <p class="text-sm text-gray-500">Active Plans</p>
                    </div>
                </div>
                <SkeletonRows v-else :rows="3" />
            </div>

            <div v-if="tab === 'users'" role="tabpanel">
                <div v-if="users.length" class="overflow-x-auto">
                    <table class="table">
                        <thead>
                            <tr>
                                <th>Email</th>
                                <th>Active</th>
                                <th>Admin</th>
                                <th>Joined</th>
                                <th></th>
                            </tr>
                        </thead>
                        <tbody>
                            <tr v-for="u in users" :key="u.id">
                                <td>{{ u.email }}</td>
                                <td>
                                    <span :class="u.is_active ? 'badge badge-success' : 'badge badge-error'">
                                        {{ u.is_active ? 'Active' : 'Suspended' }}
                                    </span>
                                </td>
                                <td>
                                    <span :class="u.is_admin ? 'badge badge-success' : 'badge'">{{ u.is_admin ? 'Admin' : 'User' }}</span>
                                </td>
                                <td>{{ formatDate(u.created_at) }}</td>
                                <td>
                                    <div class="flex gap-1">
                                        <button class="cta cta-tertiary text-sm" @click="toggleUser(u)">
                                            {{ u.is_active ? 'Suspend' : 'Activate' }}
                                        </button>
                                        <button class="cta cta-tertiary text-sm" @click="toggleAdmin(u)">
                                            {{ u.is_admin ? 'Revoke Admin' : 'Make Admin' }}
                                        </button>
                                        <button class="cta cta-tertiary text-sm text-red-500" @click="deleteUser(u)" :disabled="deleting === u.id" :aria-busy="deleting === u.id">
                                            Delete
                                        </button>
                                    </div>
                                </td>
                            </tr>
                        </tbody>
                    </table>
                </div>
                <SkeletonRows v-else :rows="5" />
            </div>

            <div v-if="tab === 'plans'" role="tabpanel">
                <div class="flex justify-between items-center mb-4">
                    <h3>Plans</h3>
                    <button class="cta" @click="showForm = !showForm" v-if="!showForm">Create Plan</button>
                </div>

                <div v-if="showForm" class="card-secondary mb-6">
                    <h4>{{ editing ? 'Edit Plan' : 'New Plan' }}</h4>
                    <div class="grid grid-cols-2 gap-4 mt-4">
                        <div><label for="plan-name">Name</label><input id="plan-name" v-model="form.name" placeholder="pro" /></div>
                        <div><label for="plan-display">Display Name</label><input id="plan-display" v-model="form.display_name" placeholder="Pro Plan" /></div>
                        <div><label for="plan-price">Price (cents)</label><input id="plan-price" type="number" v-model.number="form.price_cents" placeholder="500" /></div>
                        <div><label for="plan-currency">Currency</label><select id="plan-currency" v-model="form.currency"><option value="usd">USD</option><option value="eur">EUR</option><option value="gbp">GBP</option></select></div>
                        <div><label for="plan-interval">Interval</label><select id="plan-interval" v-model="form.interval"><option value="monthly">Monthly</option><option value="yearly">Yearly</option><option value="one_time">One Time</option></select></div>
                        <div><label for="plan-sort">Sort Order</label><input id="plan-sort" type="number" v-model.number="form.sort_order" placeholder="0" /></div>
                        <div><label for="plan-recipients">Max Recipients</label><input id="plan-recipients" type="number" v-model.number="form.max_recipients" /></div>
                        <div><label for="plan-credentials">Max Passkeys</label><input id="plan-credentials" type="number" v-model.number="form.max_credentials" /></div>
                        <div><label for="plan-aliases">Max Daily Aliases</label><input id="plan-aliases" type="number" v-model.number="form.max_daily_aliases" /></div>
                        <div><label for="plan-sendreply">Max Daily Send/Reply</label><input id="plan-sendreply" type="number" v-model.number="form.max_daily_send_reply" /></div>
                        <div><label for="plan-sessions">Max Sessions</label><input id="plan-sessions" type="number" v-model.number="form.max_sessions" /></div>
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
                <div v-else-if="!loadingPlans" class="card-empty"><h3>No plans yet</h3><p>Create your first subscription plan.</p></div>
            </div>

            <div v-if="tab === 'logs'" role="tabpanel">
                <div v-if="logs.length" class="overflow-x-auto">
                    <table class="table">
                        <thead><tr><th>Type</th><th>From</th><th>Destination</th><th>Status</th><th>Date</th></tr></thead>
                        <tbody>
                            <tr v-for="log in logs" :key="log.id">
                                <td><span class="badge">{{ log.log_type }}</span></td>
                                <td>{{ log.from }}</td>
                                <td>{{ log.destination }}</td>
                                <td>{{ log.status }}</td>
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
import { ref, onMounted } from 'vue'
import { planApi, adminApi, type Plan, type AdminUser, type AdminLog, type SystemStats } from '../api/plan'
import SkeletonRows from '../components/SkeletonRows.vue'

const tab = ref('stats')
const stats = ref<SystemStats | null>(null)
const users = ref<AdminUser[]>([])
const plans = ref<Plan[]>([])
const logs = ref<AdminLog[]>([])
const loadingPlans = ref(true)
const deleting = ref('')
const saving = ref(false)
const showForm = ref(false)
const editing = ref(false)
const formError = ref('')

const emptyForm = (): Partial<Plan> => ({ name: '', display_name: '', price_cents: 0, currency: 'usd', interval: 'monthly', max_recipients: 5, max_credentials: 10, max_daily_aliases: 50, max_daily_send_reply: 50, max_sessions: 5, sort_order: 0 })
const form = ref<Partial<Plan>>(emptyForm())

const formatDate = (d: string) => new Date(d).toLocaleDateString()

const fetchStats = async () => { try { stats.value = await adminApi.stats() } catch { /* */ } }
const fetchUsers = async () => { try { users.value = await adminApi.users() } catch { /* */ } }
const fetchPlans = async () => { loadingPlans.value = true; try { plans.value = await planApi.listAll() } catch { /* */ } finally { loadingPlans.value = false } }
const fetchLogs = async () => { try { logs.value = await adminApi.logs() } catch { /* */ } }

const toggleUser = async (u: AdminUser) => {
    try { await adminApi.updateUser({ id: u.id, is_active: !u.is_active }); u.is_active = !u.is_active } catch { /* */ }
}
const toggleAdmin = async (u: AdminUser) => {
    try { await adminApi.updateUser({ id: u.id, is_admin: !u.is_admin }); u.is_admin = !u.is_admin } catch { /* */ }
}
const deleteUser = async (u: AdminUser) => {
    if (!confirm(`Delete user ${u.email}? This removes all their data.`)) return
    deleting.value = u.id
    try { await adminApi.deleteUser(u.id); users.value = users.value.filter(x => x.id !== u.id) } catch { /* */ } finally { deleting.value = '' }
}

const savePlan = async () => {
    saving.value = true; formError.value = ''
    try {
        if (editing.value && form.value.id) { await planApi.update(form.value.id, form.value) } else { await planApi.create(form.value) }
        resetForm(); await fetchPlans()
    } catch (err: any) { formError.value = err?.message || 'Failed' } finally { saving.value = false }
}
const editPlan = (plan: Plan) => { editing.value = true; showForm.value = true; form.value = { ...plan } }
const deletePlan = async (plan: Plan) => {
    if (!confirm(`Deactivate plan "${plan.display_name}"?`)) return
    deleting.value = plan.id
    try { await planApi.delete(plan.id); await fetchPlans() } catch { /* */ } finally { deleting.value = '' }
}
const resetForm = () => { showForm.value = false; editing.value = false; form.value = emptyForm(); formError.value = '' }

onMounted(() => { fetchStats(); fetchUsers(); fetchPlans(); fetchLogs() })
</script>
