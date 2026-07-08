<template>
    <div class="card-container">
        <header class="head">
            <h2>Admin - Plans</h2>
        </header>

        <div class="card-primary">
            <div class="flex justify-between items-center mb-4">
                <h3>Subscription Plans</h3>
                <button class="cta" @click="showForm = !showForm" v-if="!showForm">Create Plan</button>
            </div>

            <div v-if="showForm" class="card-secondary mb-6">
                <h4>{{ editing ? 'Edit Plan' : 'New Plan' }}</h4>
                <div class="grid grid-cols-2 gap-4 mt-4">
                    <div>
                        <label for="plan-name">Name (unique key)</label>
                        <input id="plan-name" v-model="form.name" placeholder="pro" />
                    </div>
                    <div>
                        <label for="plan-display">Display Name</label>
                        <input id="plan-display" v-model="form.display_name" placeholder="Pro Plan" />
                    </div>
                    <div>
                        <label for="plan-price">Price (cents)</label>
                        <input id="plan-price" type="number" v-model.number="form.price_cents" placeholder="500" />
                    </div>
                    <div>
                        <label for="plan-currency">Currency</label>
                        <select id="plan-currency" v-model="form.currency">
                            <option value="usd">USD</option>
                            <option value="eur">EUR</option>
                            <option value="gbp">GBP</option>
                        </select>
                    </div>
                    <div>
                        <label for="plan-interval">Interval</label>
                        <select id="plan-interval" v-model="form.interval">
                            <option value="monthly">Monthly</option>
                            <option value="yearly">Yearly</option>
                            <option value="one_time">One Time</option>
                        </select>
                    </div>
                    <div>
                        <label for="plan-sort">Sort Order</label>
                        <input id="plan-sort" type="number" v-model.number="form.sort_order" placeholder="0" />
                    </div>
                    <div>
                        <label for="plan-recipients">Max Recipients</label>
                        <input id="plan-recipients" type="number" v-model.number="form.max_recipients" placeholder="25" />
                    </div>
                    <div>
                        <label for="plan-credentials">Max Passkeys</label>
                        <input id="plan-credentials" type="number" v-model.number="form.max_credentials" placeholder="10" />
                    </div>
                    <div>
                        <label for="plan-aliases">Max Daily Aliases</label>
                        <input id="plan-aliases" type="number" v-model.number="form.max_daily_aliases" placeholder="100" />
                    </div>
                    <div>
                        <label for="plan-sendreply">Max Daily Send/Reply</label>
                        <input id="plan-sendreply" type="number" v-model.number="form.max_daily_send_reply" placeholder="100" />
                    </div>
                    <div>
                        <label for="plan-sessions">Max Sessions</label>
                        <input id="plan-sessions" type="number" v-model.number="form.max_sessions" placeholder="10" />
                    </div>
                </div>
                <div class="flex gap-2 mt-4">
                    <button class="cta" @click="savePlan" :disabled="saving" :aria-busy="saving">{{ editing ? 'Update' : 'Create' }}</button>
                    <button class="cta cta-tertiary" @click="resetForm">Cancel</button>
                </div>
                <p v-if="formError" role="alert" class="text-red-500 mt-2">{{ formError }}</p>
            </div>

            <table v-if="plans.length" class="table">
                <thead>
                    <tr>
                        <th>Name</th>
                        <th>Price</th>
                        <th>Interval</th>
                        <th>Recipients</th>
                        <th>Aliases/day</th>
                        <th>Active</th>
                        <th></th>
                    </tr>
                </thead>
                <tbody>
                    <tr v-for="plan in plans" :key="plan.id">
                        <td>{{ plan.display_name }} <span class="text-xs text-gray-500">({{ plan.name }})</span></td>
                        <td>{{ plan.price_cents === 0 ? 'Free' : (plan.price_cents / 100) + ' ' + plan.currency }}</td>
                        <td>{{ plan.interval }}</td>
                        <td>{{ plan.max_recipients }}</td>
                        <td>{{ plan.max_daily_aliases }}</td>
                        <td>
                            <span :class="plan.is_active ? 'badge badge-success' : 'badge badge-error'">
                                {{ plan.is_active ? 'Active' : 'Inactive' }}
                            </span>
                        </td>
                        <td>
                            <div class="flex gap-1">
                                <button class="cta cta-tertiary text-sm" @click="editPlan(plan)">Edit</button>
                                <button class="cta cta-tertiary text-sm" @click="deletePlan(plan)" :disabled="deleting === plan.id" :aria-busy="deleting === plan.id">Delete</button>
                            </div>
                        </td>
                    </tr>
                </tbody>
            </table>

            <div v-else-if="!loading" class="card-empty">
                <span class="icon text-4xl icon-settings"></span>
                <h3>No plans yet</h3>
                <p>Create your first subscription plan to get started.</p>
            </div>

            <SkeletonRows v-else :rows="3" />
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { planApi, type Plan } from '../api/plan'
import SkeletonRows from '../components/SkeletonRows.vue'

const plans = ref<Plan[]>([])
const loading = ref(true)
const showForm = ref(false)
const editing = ref(false)
const saving = ref(false)
const deleting = ref('')
const formError = ref('')

const emptyForm = (): Partial<Plan> => ({
    name: '', display_name: '', price_cents: 0, currency: 'usd',
    interval: 'monthly', max_recipients: 5, max_credentials: 10,
    max_daily_aliases: 50, max_daily_send_reply: 50, max_sessions: 5,
    sort_order: 0,
})
const form = ref<Partial<Plan>>(emptyForm())

const fetchPlans = async () => {
    loading.value = true
    try {
        plans.value = await planApi.listAll()
    } catch { plans.value = [] } finally { loading.value = false }
}

const savePlan = async () => {
    saving.value = true
    formError.value = ''
    try {
        if (editing.value && form.value.id) {
            await planApi.update(form.value.id, form.value)
        } else {
            await planApi.create(form.value)
        }
        resetForm()
        await fetchPlans()
    } catch (err: any) {
        formError.value = err?.message || 'Failed to save plan'
    } finally { saving.value = false }
}

const editPlan = (plan: Plan) => {
    editing.value = true
    showForm.value = true
    form.value = { ...plan }
}

const deletePlan = async (plan: Plan) => {
    if (!confirm(`Deactivate plan "${plan.display_name}"?`)) return
    deleting.value = plan.id
    try { await planApi.delete(plan.id); await fetchPlans() }
    catch { /* toast */ } finally { deleting.value = '' }
}

const resetForm = () => {
    showForm.value = false
    editing.value = false
    form.value = emptyForm()
    formError.value = ''
}

onMounted(fetchPlans)
</script>
