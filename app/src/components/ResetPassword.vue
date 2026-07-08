<template>
    <div class="page center">
        <div></div>
        <form class="card-tertiary center" @submit.prevent="resetPassword">
            <article>
                <h1 class="flex justify-center text-accent mb-8">
                    <span class="logo"></span>
                </h1>
                <h4 class="text-center mb-8">Set new password</h4>
                <div v-if="!apiSuccess">
                    <div class="mb-5">
                        <PasswordInput
                            v-model="password"
                            v-bind:class="{ 'error': passwordError }"
                            id="password-new"
                            autocomplete="new-password"
                            placeholder="New Password"
                            aria-label="New password"
                            class="password"
                        />
                    </div>
                    <div class="mb-3">
                        <PasswordInput
                            v-model="passwordConfirm"
                            v-bind:class="{ 'error': passwordError }"
                            id="password-new-confirm"
                            autocomplete="new-password"
                            placeholder="Confirm Password"
                            aria-label="Confirm password"
                            class="password"
                        />
                    </div>
                    <p class="text-sm mb-5">
                        Must be 12+ characters and contain uppercase, lowercase, number, and special character (e.g. -_+=~!@#$%^&*(),;.?":{}|<>)
                    </p>
                    <div class="flex items-center justify-between">
                        <button :disabled="isLoading" :aria-busy="isLoading" class="cta full">
                            Update password
                        </button>
                    </div>
                    <p v-if="passwordError" class="error mt-5" role="alert">Error: {{ passwordError }}</p>
                    <p v-if="apiError" class="error mt-5" role="alert">Error: {{ apiError }}</p>
                </div>
                <div v-if="apiSuccess">
                    <p class="success mb-5">{{ apiSuccess }}</p>
                </div>
                <nav role="tablist" class="tabs-router">
                    <router-link to="/login" custom v-slot="{ navigate }">
                        <button @click="navigate">Back to Log In</button>
                    </router-link>
                </nav>
            </article>
        </form>
        <Footer />
    </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { ApiError } from '../api/api.ts'
import { useRoute } from 'vue-router'
import { userApi } from '../api/user.ts'
import Footer from './Footer.vue'
import PasswordInput from './PasswordInput.vue'

const password = ref('')
const passwordConfirm = ref('')
const passwordError = ref('')
const apiError = ref('')
const apiSuccess = ref('')
const isLoading = ref(false)
const otp = ref('')

const validatePassword = () => {
    apiSuccess.value = ''
    passwordError.value = ''

    if (!password.value || !passwordConfirm.value) {
        passwordError.value = 'Please fill required fields'
    }

    if (password.value !== passwordConfirm.value) {
        passwordError.value = 'Passwords do not match'
    }

    return !passwordError.value
}

const resetPassword = async () => {
    if (!validatePassword()) return

    isLoading.value = true // Start loading

    const req = {
        password: password.value,
        otp: otp.value
    }

    try {
        const res = await userApi.resetPassword(req)
        apiSuccess.value = res.data.message
        apiError.value = ''
    } catch (err) {
        apiSuccess.value = ''
        if (err instanceof ApiError) {
            apiError.value = err.data?.error || err.message || err.message

            if (err.status === 429) {
                apiError.value = 'Too many requests, please try again later.'
            }
        }
    } finally {
        isLoading.value = false // End loading
    }
}

onMounted(() => {
    const route = useRoute()
    otp.value = route.params.token as string
})
</script>