<template>
    <div class="page center">
        <div></div>
        <form class="card-tertiary center" @submit.prevent="">
            <article>
                <div>
                    <div v-if="passkeySupported" id="tabs-with-underline-1" role="tabpanel" aria-labelledby="tabs-with-underline-item-1">
                        <h1 class="flex justify-center text-accent mb-8">
                            <span class="logo"></span>
                        </h1>
                        <h4 class="text-center mb-8">Sign up with Passkey</h4>
                        <div v-if="!apiSuccess">
                            <div class="mb-7">
                                <input
                                    v-model="emailAuthn"
                                    v-bind:class="{ 'error': emailAuthnError }"
                                    placeholder="Email Address"
                                    aria-label="Email address"
                                    id="email_authn"
                                    type="email"
                                    autocomplete="email"
                                    class="email"
                                    @keypress.enter.prevent
                                >
                                <p v-if="emailAuthnError" class="error" role="alert">Required</p>
                            </div>
                            <div class="flex items-center w-full">
                                <button @click="registerWithPasskey" :disabled="isLoading" :aria-busy="isLoading" class="cta full">
                                    Sign Up with Passkey
                                </button>
                            </div>
                            <p v-if="apiError" class="error mt-6" role="alert">Error: {{ apiError }}</p>
                        </div>
                    </div>
                    <div
                        id="tabs-with-underline-2"
                        v-bind:class="{ 'hidden': passkeySupported }"
                        role="tabpanel"
                        aria-labelledby="tabs-with-underline-item-2">
                        <h1 class="flex justify-center text-accent mb-8">
                            <span class="logo"></span>
                        </h1>
                        <h4 class="text-center mb-8">Sign up with email and password</h4>
                        <div v-if="!apiSuccess">
                            <div class="mb-7">
                                <input
                                    v-model="email"
                                    v-bind:class="{ 'error': emailError }"
                                    placeholder="Email Address"
                                    aria-label="Email address"
                                    id="email"
                                    type="email"
                                    autocomplete="email"
                                    class="email"
                                    @blur="validateEmail"
                                    @keypress.enter.prevent
                                >
                                <p v-if="emailError" class="error" role="alert">Required</p>
                            </div>
                            <div class="mb-7">
                                <PasswordInput
                                    v-model="password"
                                    v-bind:class="{ 'error': passwordError }"
                                    placeholder="Password"
                                    aria-label="Password"
                                    id="password"
                                    autocomplete="new-password"
                                    class="password"
                                    @blur="validatePassword"
                                    @keypress.enter.prevent
                                />
                                <p v-if="passwordError" class="error" role="alert">{{ passwordError }}</p>
                            </div>
                            <p class="text-sm mb-5">Must be 12+ characters and contain uppercase, lowercase, number, and special character (e.g. -_+=~!@#$%^&*(),;.?":{}|<>)</p>
                            <div class="flex items-center w-full">
                                <button @click="register" :disabled="isLoading" :aria-busy="isLoading" class="cta full">
                                    Sign Up
                                </button>
                            </div>
                            <p v-if="apiError" class="error mt-5" role="alert">Error: {{ apiError }}</p>
                        </div>
                    </div>
                </div>
                <nav v-if="passkeySupported" aria-label="Tabs" role="tablist" aria-orientation="horizontal" class="tabs-router">
                    <button
                        class="active"
                        id="tabs-with-underline-item-1" aria-selected="true" data-hs-tab="#tabs-with-underline-1"
                        aria-controls="tabs-with-underline-1" role="tab">
                        Use Passkey instead
                    </button>
                    <button
                        id="tabs-with-underline-item-2" aria-selected="false" data-hs-tab="#tabs-with-underline-2"
                        aria-controls="tabs-with-underline-2" role="tab">
                        Use Password instead
                    </button>
                </nav>
                <div v-if="apiSuccess">
                    <p class="success mb-6">{{ apiSuccess }}</p>
                    <router-link to="/login" tag="button" class="cta full">
                        Proceed to Log In
                    </router-link>
                </div>
            </article>
        </form>
        <Footer />
    </div>
</template>

<script setup lang="ts">
import { ref, onMounted, onUpdated } from 'vue'
import { ApiError } from '../api/api.ts'
import { userApi } from '../api/user.ts'
import { startRegistration, browserSupportsWebAuthn } from '@simplewebauthn/browser'
import tabs from '@preline/tabs'
import Footer from './Footer.vue'
import PasswordInput from './PasswordInput.vue'

const email = ref('')
const emailAuthn = ref('')
const password = ref('')
const emailError = ref(false)
const emailAuthnError = ref(false)
const passwordError = ref('')
const apiSuccess = ref('')
const apiError = ref('')
const isLoading = ref(false)
const passkeySupported = ref(false)

const validateEmail = () => {
    emailError.value = !email.value
    return !emailError.value
}

const validateEmailAuthn = () => {
    emailAuthnError.value = !emailAuthn.value
    return !emailAuthnError.value
}

const validatePassword = () => {
    passwordError.value = ''

    if (!password.value) {
        passwordError.value = 'Required'
        return !passwordError.value
    }

    if (password.value.length < 12) {
        passwordError.value = 'Password must be 12+ characters'
        return !passwordError.value
    }

    if (!/^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[-_+=~!@#$%^&*(),;.?":{}|<>])/.test(password.value)) {
        passwordError.value = 'Password must contain uppercase, lowercase, number, and special character'
        return !passwordError.value
    }

    return !passwordError.value
}

const validate = () => {
    const validEmail = validateEmail()
    const validPass = validatePassword()
    return validEmail && validPass
}

const register = async () => {
    if (!validate()) return

    isLoading.value = true
    const data = {
        email: email.value,
        password: password.value,
    }

    try {
        await userApi.register(data)
        apiError.value = ''
        window.location.href = '/signup-complete'
    } catch (err) {
        apiSuccess.value = ''
        if (err instanceof ApiError) {
            apiError.value = err.data?.error || err.message || err.message

            if (err.status === 429) {
                apiError.value = 'Too many requests, please try again later.'
            }
        }
    } finally {
        isLoading.value = false
    }
}

const registerWithPasskey = async () => {
    if (!validateEmailAuthn()) return

    isLoading.value = true

    const data = {
        email: emailAuthn.value,
    }

    try {
        let res = await userApi.registerBegin(data)
        const creds = await startRegistration({ optionsJSON: res.data['publicKey'] })
        res = await userApi.registerFinish(creds)
        apiError.value = ''
        localStorage.setItem('email', data.email)
        window.location.href = '/account'
    } catch (err) {
        if (err instanceof ApiError) {
            apiError.value = err.data?.error || err.message || err.message

            if (err.status === 429) {
                apiError.value = 'Too many requests, please try again later.'
            }
        }
    } finally {
        isLoading.value = false
    }
}

const isLoggedIn = (): boolean => {
    const email = localStorage.getItem('email')
    return email !== null && email.trim() !== ''
}

onMounted(() => {
    if (isLoggedIn()) {
        window.location.href = '/account'
    }
    
    passkeySupported.value = browserSupportsWebAuthn()
    tabs.autoInit()
})

onUpdated(() => {
    tabs.autoInit()
})
</script>
