<script setup lang="ts">
import { computed } from 'vue'
import { userApi } from '../api/user.ts'
import { appConfirm } from '../composables/useConfirm.ts'

const signupUrl = import.meta.env.VITE_PRICING_URL

const isLoggedIn = computed(() => {
    const email = localStorage.getItem('email')
    return email !== null && email.trim() !== ''
})

const logout = async () => {
    if (!(await appConfirm('End your current session?', { title: 'Log out', confirmLabel: 'Log out' }))) return
    try {
        await userApi.logout()
        userApi.clearSession()
    } catch { }
}

const heroFeatures = [
    'Forward emails through isolated aliases with FreeTheMail.',
    'Send and reply without exposing your primary email address.',
    'Generate self-destructing temp inboxes for one-off signups.',
    'PGP encryption, custom domains, wildcards, and catch-all.',
]
</script>

<template>
    <div class="landing-page">
        <header class="landing-header">
            <router-link to="/" class="landing-logo-link">
                <span class="landing-logo"></span>
            </router-link>
            <div class="landing-actions">
                <template v-if="isLoggedIn">
                    <router-link to="/account" class="landing-btn landing-btn-signup">
                        [DASHBOARD]
                    </router-link>
                    <a @click="logout" class="landing-btn landing-btn-login" style="cursor:pointer">
                        [LOG OUT]
                    </a>
                </template>
                <template v-else>
                    <a :href="signupUrl" class="landing-btn landing-btn-signup" target="_blank">
                        [SIGN UP]
                    </a>
                    <router-link to="/login" class="landing-btn landing-btn-login">
                        [LOG IN]
                    </router-link>
                </template>
            </div>
        </header>

        <section class="hero-section">
            <div class="hero-content">
                <div class="hero-text">
                    <h1 class="hero-title">
                        Resist email<br>surveillance.
                    </h1>
                    <div class="hero-features">
                        <div v-for="text in heroFeatures" :key="text" class="feature-item">
                            <span class="feature-bullet">&gt;</span>
                            <p class="feature-text">{{ text }}</p>
                        </div>
                    </div>
                </div>
                <div class="hero-cta">
                    <a :href="signupUrl" class="cta-btn cta-primary" target="_blank">
                        ./GET_ACCESS
                    </a>
                    <a href="https://github.com/freethemail/freethemail" target="_blank" class="cta-btn cta-secondary">
                        ./VIEW_SOURCE
                    </a>
                </div>
                <div class="hero-background light-only" aria-hidden="true"></div>
                <div class="hero-background dark-only" aria-hidden="true"></div>
            </div>
            <div class="hero-screenshot">
                <div class="screenshot-container">
                    <div class="screenshot-border">
                        <img
                            src="../assets/landing-screenshot-light2.png"
                            srcset="../assets/landing-screenshot-light2.png 1x, ../assets/landing-screenshot-light2@2x.png 2x"
                            alt="FreeTheMail Application Screenshot"
                            class="screenshot-image light-only"
                        />
                        <img
                            src="../assets/landing-screenshot-dark2.png"
                            srcset="../assets/landing-screenshot-dark2.png 1x, ../assets/landing-screenshot-dark2@2x.png 2x"
                            alt="FreeTheMail Application Screenshot"
                            class="screenshot-image dark-only"
                        />
                    </div>
                </div>
            </div>
        </section>

        <section class="landing-footer-strip-section">
            <div class="landing-footer-strip">
                <span class="landing-footer-strip-text">2026 FreeTheMail</span>
                <span class="landing-footer-strip-divider" aria-hidden="true"></span>
                <a href="https://github.com/freethemail/freethemail" target="_blank" class="landing-footer-strip-text landing-footer-strip-link">
                    ./SOURCE
                </a>
                <a href="https://github.com/freethemail/freethemail#security-audit" target="_blank" class="landing-footer-strip-text landing-footer-strip-link">
                    ./AUDIT
                </a>
                <a href="/privacy" class="landing-footer-strip-text landing-footer-strip-link">
                    ./PRIVACY
                </a>
                <a href="/tos" class="landing-footer-strip-text landing-footer-strip-link">
                    ./TERMS
                </a>
                <a href="/faq" class="landing-footer-strip-text landing-footer-strip-link">
                    ./FAQ
                </a>
                <a href="https://github.com/freethemail/freethemail/blob/main/LICENSE.md" target="_blank" class="landing-footer-strip-text landing-footer-strip-link">
                    ./LICENSE
                </a>
            </div>
        </section>
    </div>
</template>

<style scoped>
.landing-page {
    @apply w-full;
    @apply bg-white dark:bg-[oklch(0.145_0_0)];
    @apply min-h-screen;
    @apply flex flex-col;

    * {
        font-family: 'Roboto Mono', monospace;
    }
}

.landing-header {
    @apply flex items-center justify-between;
    @apply border-b border-solid;
    @apply border-light-gray-6 dark:border-dark-neutral-3;
    @apply py-3;
    @apply px-4 md:px-8 lg:px-[110px];
    @apply w-full;
}

.landing-logo-link {
    @apply p-0 flex items-center;
}

.landing-logo {
    @apply block bg-no-repeat bg-contain bg-center;
    @apply bg-[url("../assets/freethemail-light.png")] dark:bg-[url("../assets/freethemail-dark.png")];
    @apply w-[98px] h-5 md:w-[118px] md:h-[25px];
}

.landing-actions {
    @apply flex gap-2 items-center;
}

.landing-btn {
    @apply flex items-center justify-center;
    @apply font-medium text-sm leading-none;
    @apply py-[11px] whitespace-nowrap;
    @apply transition-colors;
}

.landing-btn-signup {
    @apply px-2;
    @apply text-[oklch(0.659_0.192_40.23)] dark:text-[oklch(0.659_0.192_40.23)];
}

.landing-btn-signup:hover {
    @apply opacity-80;
}

.landing-btn-login {
    @apply px-3;
    @apply bg-[oklch(0.659_0.192_40.23)] dark:bg-[oklch(0.659_0.192_40.23)];
    @apply text-[oklch(0.998_0.002_325.59)] dark:text-[oklch(0.145_0_0)];
}

.landing-btn-login:hover {
    @apply opacity-90;
}

/* Hero */
.hero-section {
    @apply w-full flex flex-col flex-1;
    @apply relative overflow-hidden;
    @apply pb-8;
}

.hero-background {
    @apply absolute pointer-events-none;
    @apply hidden md:block;
    top: -30px;
    right: 0;
    bottom: -30px;
    aspect-ratio: 1 / 1;
    background-size: contain;
    background-repeat: no-repeat;
    background-position: center right;
}

.hero-background.light-only {
    background-image: url('../assets/hero-bg-light-tablet.svg');
}

.hero-background.dark-only {
    background-image: url('../assets/hero-bg-dark-tablet.svg');
}

@media (min-width: 1024px) {
    .hero-background.light-only {
        background-image: url('../assets/hero-bg-light-desktop.svg');
    }

    .hero-background.dark-only {
        background-image: url('../assets/hero-bg-dark-desktop.svg');
    }
}

.hero-content {
    @apply flex flex-col;
    @apply py-8;
    @apply max-w-[1060px] mx-auto;
    @apply px-4 md:px-8;
    @apply relative;
    @apply w-full;
}

.hero-text {
    @apply flex flex-col gap-5 md:gap-6 z-10;
    @apply relative;
}

.hero-title {
    @apply text-[oklch(0.226_0_0)] dark:text-[oklch(0.998_0.002_325.59)];
    @apply text-[42px] leading-[46px] md:text-[52px] md:leading-[54px] lg:text-[64px] lg:leading-[66px];
    @apply m-0;
    font-weight: 600;
    letter-spacing: -0.03em;
}

.landing-page .hero-title,
.landing-page .hero-title * {
    font-family: var(--font-serif);
}

.hero-features {
    @apply flex flex-col gap-3;
}

.feature-item {
    @apply flex gap-2 items-center;
    @apply relative z-10;
}

.feature-bullet {
    @apply text-[oklch(0.659_0.192_40.23)] dark:text-[oklch(0.692_0.189_44.533)];
    @apply text-base;
    @apply shrink-0;
    @apply flex items-center justify-center;
    @apply w-6 h-6;
}

.feature-text {
    @apply text-black dark:text-[oklch(0.998_0.002_325.59)];
    @apply text-base leading-[21px] md:leading-4;
    @apply m-0;
    @apply flex-1;
}

.hero-cta {
    @apply flex gap-2 items-center mt-8 z-10;
    @apply flex-col md:flex-row w-full md:w-auto;
    @apply relative;
}

.cta-btn {
    @apply flex items-center justify-center;
    @apply px-4 py-[13px];
    @apply font-semibold text-sm leading-4;
    @apply whitespace-nowrap;
    @apply w-full md:w-auto;
    border-radius: var(--radius);
    transition: box-shadow 200ms var(--ease-smooth), background 200ms var(--ease-smooth),
        border-color 200ms var(--ease-smooth), color 200ms var(--ease-smooth),
        transform 100ms var(--ease-smooth);
}

.cta-btn:active {
    transform: scale(var(--press-scale));
}

.cta-primary {
    @apply text-[oklch(0.998_0.002_325.59)] dark:text-[oklch(0.998_0.002_325.59)];
    border: 1px solid var(--accent);
    background: var(--accent-grad);
    box-shadow: var(--raised-shadow);
    text-shadow: var(--accent-text-glow);
}

.cta-primary:hover {
    background: var(--accent-grad-hover);
    box-shadow: var(--raised-shadow), var(--accent-glow);
}

.cta-secondary {
    @apply px-[15px] py-[12px];
    @apply text-[oklch(0.659_0.192_40.23)] dark:text-[oklch(0.692_0.189_44.533)];
    border: 1px solid var(--sk-border);
    background: transparent;
}

.cta-secondary:hover {
    background: var(--inset-bg);
    box-shadow: var(--inset-shadow);
}

/* Screenshot */
.hero-screenshot {
    @apply relative z-10;
    @apply max-w-[1060px] mx-auto;
    @apply px-4 md:px-8;
    @apply w-full;
}

.screenshot-container {
    @apply w-full;
    @apply bg-[oklch(0.967_0.002_247.839)] dark:bg-[oklch(0.185_0.003_17.454)];
    @apply p-2;
}

.screenshot-border {
    @apply w-full;
    @apply bg-white dark:bg-[oklch(0.213_0_0)];
    @apply border border-solid;
    @apply border-[oklch(0.902_0.009_258.337)] dark:border-[oklch(0.269_0_0)];
    @apply p-1;
    @apply overflow-hidden;
}

.screenshot-image {
    @apply w-full;
    @apply block;
    @apply h-auto;
}

/* Theme visibility */
img.light-only {
    @apply dark:hidden !important;
}

img.dark-only {
    @apply hidden dark:block !important;
}

.hero-background.light-only {
    @apply dark:hidden md:block;
}

.hero-background.dark-only {
    @apply hidden dark:md:block;
}

/* Footer strip */
.landing-footer-strip-section {
    @apply w-full;
    @apply bg-[oklch(1_0_0)] dark:bg-[oklch(0.121_0.004_245.473)];
}

.landing-footer-strip {
    @apply w-full;
    @apply flex flex-wrap items-center justify-center;
    @apply gap-x-3 gap-y-2;
    @apply py-4;
    @apply px-4 md:px-8 lg:px-0;
}

.landing-footer-strip-text {
    @apply text-sm leading-[14px];
    @apply text-[oklch(0.45_0.13_40.121)];
    @apply whitespace-nowrap;
}

.landing-footer-strip-link {
    @apply no-underline;
}

.landing-footer-strip-link:hover {
    @apply opacity-80;
}

.landing-footer-strip-divider {
    @apply hidden md:block;
    @apply w-0 h-2;
    @apply border-l border-solid;
    @apply border-[oklch(0.806_0.022_252.529)] dark:border-[oklch(0.353_0.035_248.846)];
}
</style>
