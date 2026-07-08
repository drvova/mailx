<template>
    <div class="card-container">
        <header class="head">
            <h2>Plans</h2>
        </header>

        <div v-if="loading" class="grid grid-cols-1 md:grid-cols-3 gap-6 mt-6">
            <SkeletonRows v-for="i in 3" :key="i" :rows="4" />
        </div>

        <div v-else-if="plans.length" class="grid grid-cols-1 md:grid-cols-3 gap-6 mt-6">
            <div v-for="plan in plans" :key="plan.id" class="card-primary text-center">
                <h3>{{ plan.display_name }}</h3>
                <p class="text-3xl font-bold mt-4">
                    {{ plan.price_cents === 0 ? 'Free' : (plan.price_cents / 100) + ' ' + plan.currency }}
                    <span v-if="plan.interval !== 'one_time'" class="text-sm font-normal text-gray-500">/{{ plan.interval }}</span>
                </p>
                <ul class="text-left mt-6 space-y-2">
                    <li>{{ plan.max_recipients }} recipients</li>
                    <li>{{ plan.max_daily_aliases }} aliases/day</li>
                    <li>{{ plan.max_daily_send_reply }} send/reply/day</li>
                    <li>{{ plan.max_credentials }} passkeys</li>
                    <li>{{ plan.max_sessions }} sessions</li>
                </ul>
                <button
                    class="cta w-full mt-6"
                    @click="choose(plan)"
                    :disabled="choosing === plan.id"
                    :aria-busy="choosing === plan.id">
                    {{ plan.price_cents === 0 ? 'Sign Up Free' : 'Choose ' + plan.display_name }}
                </button>
            </div>
        </div>

        <div v-else class="card-empty">
            <h3>No plans available</h3>
            <p>Check back later for subscription options.</p>
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { planApi, billingApi, type Plan } from '../api/plan'
import SkeletonRows from '../components/SkeletonRows.vue'

const plans = ref<Plan[]>([])
const loading = ref(true)
const choosing = ref('')

const fetchPlans = async () => {
    loading.value = true
    try { plans.value = await planApi.list() }
    catch { plans.value = [] } finally { loading.value = false }
}

const choose = async (plan: Plan) => {
    if (plan.price_cents === 0) {
        window.location.href = '/signup'
        return
    }
    choosing.value = plan.id
    try {
        const { url } = await billingApi.checkout(plan.id)
        window.location.href = url
    } catch { choosing.value = '' }
}

onMounted(fetchPlans)
</script>
