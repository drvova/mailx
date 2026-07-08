<template>
    <div class="mb-5">
        <h2>Account</h2>
        <p v-if="sub.id" class="text-sm">
            <span v-if="isActive()" class="badge success">Active</span>
            <span v-if="!isActive()" class="badge">Inactive</span>
        </p>
        <div class="mb-3">
            <h4>Account email:</h4>
            <p class="mb-3">
                {{ email }}
            </p>
        </div>
        <div v-if="isActive()" class="mb-3">
            <h4>Subscription active until:</h4>
            <p class="mb-3">
                {{ activeUntilDate() }}
            </p>
        </div>
        <div v-if="isLimited()" class="card-tertiary">
            <footer>
                <div>
                    <i class="icon info icon-primary"></i>
                </div>
                <div>
                    <h4>Limited Access Mode</h4>
                    <p>
                        Existing aliases forward normally. New aliases are disabled.
                    </p>
                </div>
            </footer>
        </div>
        <div v-if="isOutage()" class="card-tertiary">
            <footer>
                <div>
                    <i class="icon info icon-primary"></i>
                </div>
                <div>
                    <h4>Out of sync</h4>
                    <p>
                        Your last account status update was {{ updatedAtDate() }}.
                    </p>
                </div>
            </footer>
        </div>
        <p v-if="error" class="error" role="alert">Error: {{ error }}</p>
        <p v-if="success" class="success">{{ success }}</p>
    </div>
</template>

<script setup lang="ts">
import { onMounted, onUnmounted, ref } from 'vue'
import { formatDistanceToNow } from 'date-fns'
import { ApiError } from '../api/api.ts'
import { subscriptionApi } from '../api/subscription.ts'
import events from '../events.ts'

const sub = ref({
    id: '',
    updated_at: '',
    active_until: '',
    status: '',
    outage: false,
    type: '',
})
const error = ref('')
const success = ref('')
const email = ref(localStorage.getItem('email'))

const getSubscription = async () => {
    try {
        const res = await subscriptionApi.get()
        sub.value = res.data
    } catch (err) {
        if (err instanceof ApiError) {
            error.value = err.data?.error || err.message || err.message
        }
    }
}

const isActive = () => {
    return sub.value.status === 'active' || sub.value.status === 'grace_period'
}

const isLimited = () => {
    return sub.value.status === 'limited_access'
}

const activeUntilDate = () => {
    return formatDistanceToNow(new Date(sub.value.active_until)) + ' left'
}

const updatedAtDate = () => {
    return new Date(sub.value.updated_at).toLocaleString()
}

const onUpdateEmail = (event: any) => {
    email.value = event.email
}

const isOutage = () => {
    return sub.value.outage
}

onMounted(() => {
    getSubscription()
    events.on('user.update', onUpdateEmail)
})

onUnmounted(() => {
    events.off('user.update', onUpdateEmail)
})
</script>
