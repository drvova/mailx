<script setup lang="ts">
import { computed } from 'vue'
import { userApi } from '../api/user.ts'

const signupUrl = import.meta.env.VITE_PRICING_URL

const isLoggedIn = computed(() => {
    const email = localStorage.getItem('email')
    return email !== null && email.trim() !== ''
})

const logout = async () => {
    if (!confirm('Do you want to proceed with the logout?')) return

    try {
        await userApi.logout()
        userApi.clearSession()
    } catch { }
}

// --- Hero -------------------------------------------------------------------
const heroFeatures = [
    'Forward emails through isolated aliases with FreeTheMail.',
    'Send and reply without exposing your primary email address.',
    'Supports PGP encryption, domain choice, multiple',
    'recipients, and catch-all aliases.',
]

// --- Constraints ------------------------------------------------------------
const constraints = [
    "FreeTheMail is not an email provider, it can't replace your primary email",
    'Messages are visible to FreeTheMail servers during relay (use PGP)',
    'No IMAP/POP3 access. Email forwarding only.',
    'Not designed for protection against targeted surveillance',
]

</script>

<template>
    <div class="landing-page">
        <!-- Section 1: Header -->
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

        <!-- Section 2: Hero Content -->
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
                    <a href="https://github.com/ivpn/mailx" target="_blank" class="cta-btn cta-secondary">
                        ./VIEW_SOURCE
                    </a>
                </div>
                <!-- Background patterns - hidden on mobile, shown on tablet/desktop -->
                <div class="hero-background light-only" aria-hidden="true"></div>
                <div class="hero-background dark-only" aria-hidden="true"></div>
            </div>
            <div class="hero-screenshot">
                <div class="screenshot-container">
                    <div class="screenshot-border">
                        <!-- Light mode screenshot -->
                        <img 
                            src="../assets/landing-screenshot-light2.png" 
                            srcset="../assets/landing-screenshot-light2.png 1x, ../assets/landing-screenshot-light2@2x.png 2x"
                            alt="FreeTheMail Application Screenshot"
                            class="screenshot-image light-only"
                        />
                        <!-- Dark mode screenshot -->
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

        <!-- Section 3: How It Works -->
        <section class="how-it-works-section">
            <div class="section-container">
                <!-- Title and Subtitle -->
                <div class="section-header">
                    <h2 class="how-it-works-title">Why aliases?</h2>
                    <p class="how-it-works-subtitle">FreeTheMail generates unique aliases for each service and sender.</p>
                </div>

                <!-- Feature Cards -->
                <div class="feature-cards">
                    <div class="feature-card">
                        <div class="feature-card-header">
                            <img src="../assets/icons/feature-icon-lock-light.svg" class="feature-icon light-only" />
                            <img src="../assets/icons/feature-icon-lock-dark.svg" class="feature-icon dark-only" />
                            <h3 class="feature-card-title">Block spam and phishing</h3>
                        </div>
                        <p class="feature-card-text">Disable or delete any alias instantly. Even if one gets compromised, your real email remains unexposed.</p>
                    </div>

                    <div class="feature-card">
                        <div class="feature-card-header">
                            <img src="../assets/icons/feature-icon-data-leak-light.svg" class="feature-icon light-only" />
                            <img src="../assets/icons/feature-icon-data-leak-dark.svg" class="feature-icon dark-only" />
                            <h3 class="feature-card-title">Detect data leaks</h3>
                        </div>
                        <p class="feature-card-text">One alias per service means you'll know exactly who shared your email without your consent.</p>
                    </div>

                    <div class="feature-card">
                        <div class="feature-card-header">
                            <img src="../assets/icons/feature-icon-users-light.svg" class="feature-icon light-only" />
                            <img src="../assets/icons/feature-icon-users-dark.svg" class="feature-icon dark-only" />
                            <h3 class="feature-card-title">Prevent identity mapping</h3>
                        </div>
                        <p class="feature-card-text">Each service gets a unique alias. Advertisers and data brokers can't connect your accounts across platforms.</p>
                    </div>
                </div>
            </div>
        </section>

        <!-- Section 5: Verifiable Privacy -->
        <section class="verifiable-privacy-section">
            <div class="section-container">
                <div class="section-header">
                    <div class="section-command">/etc/freethemail/trust.cfg</div>
                    <h2 class="section-title">
                        Verifiable privacy
                    </h2>
                </div>

                <div class="verifiable-privacy-grid">
                    <div class="verifiable-privacy-card trust-card-accountable">
                        <div class="verifiable-privacy-card-content">
                            <div class="verifiable-privacy-card-title-row">
                                <img src="../assets/icons/trust-accountable-light.svg" class="verifiable-privacy-icon light-only" />
                                <img src="../assets/icons/trust-accountable-dark.svg" class="verifiable-privacy-icon dark-only" />
                                <h3 class="verifiable-privacy-card-title">ACCOUNTABLE_OPERATORS</h3>
                            </div>
                            <p class="verifiable-privacy-card-text">Built by the public team behind IVPN, with a 15-year history in operating privacy services.</p>
                        </div>
                        <a href="https://www.ivpn.net/en/team/" target="_blank" class="verifiable-privacy-card-link">./MEET_TEAM</a>
                    </div>

                    <div class="verifiable-privacy-card trust-card-open-source">
                        <div class="verifiable-privacy-card-content">
                            <div class="verifiable-privacy-card-title-row">
                                <img src="../assets/icons/trust-open-source-light.svg" class="verifiable-privacy-icon light-only" />
                                <img src="../assets/icons/trust-open-source-dark.svg" class="verifiable-privacy-icon dark-only" />
                                <h3 class="verifiable-privacy-card-title">OPEN_SOURCE</h3>
                            </div>
                            <p class="verifiable-privacy-card-text">The entire FreeTheMail project is open-source. Our implementation is public and available for review.</p>
                        </div>
                        <a href="https://github.com/ivpn/mailx" target="_blank" class="verifiable-privacy-card-link">./VIEW_SOURCE</a>
                    </div>

                    <div class="verifiable-privacy-card trust-card-security">
                        <div class="verifiable-privacy-card-content">
                            <div class="verifiable-privacy-card-title-row">
                                <img src="../assets/icons/trust-audit-light.svg" class="verifiable-privacy-icon light-only" />
                                <img src="../assets/icons/trust-audit-dark.svg" class="verifiable-privacy-icon dark-only" />
                                <h3 class="verifiable-privacy-card-title">SECURITY_AUDIT</h3>
                            </div>
                            <p class="verifiable-privacy-card-text">FreeTheMail has undergone a third-party security audit to validate our claims and architecture.</p>
                        </div>
                        <a href="https://www.ivpn.net/resources/IVP-07-report.pdf" class="verifiable-privacy-card-link">./READ_AUDIT</a>
                    </div>

                    <div class="verifiable-privacy-card trust-card-no-tracking">
                        <div class="verifiable-privacy-card-content">
                            <div class="verifiable-privacy-card-title-row">
                                <img src="../assets/icons/trust-no-tracking-light.svg" class="verifiable-privacy-icon light-only" />
                                <img src="../assets/icons/trust-no-tracking-dark.svg" class="verifiable-privacy-icon dark-only" />
                                <h3 class="verifiable-privacy-card-title">NO_TRACKING</h3>
                            </div>
                            <p class="verifiable-privacy-card-text">Your IP address is never logged. Forwarded emails are automatically deleted after delivery.</p>
                        </div>
                        <router-link to="/privacy" class="verifiable-privacy-card-link">./PRIVACY_POLICY</router-link>
                    </div>
                </div>
            </div>
        </section>

        <!-- Section 7: Constraints -->
        <section class="constraints-section">
            <div class="section-container">
                <div class="section-header">
                    <div class="constraints-command">/etc/freethemail/limitations.txt</div>
                    <h2 class="constraints-title">Constraints</h2>
                </div>

                <div class="constraints-list">
                    <div v-for="text in constraints" :key="text" class="constraints-item">
                        <span class="constraints-bullet">&gt;</span>
                        <p class="constraints-text">{{ text }}</p>
                    </div>
                </div>
            </div>
        </section>

        <!-- Section 8: Get Access -->
        <section class="get-access-section">
            <div class="section-container">
                <div class="section-header">
                    <div class="section-command">freethemail auth --init</div>
                    <h2 class="section-title">Get access</h2>
                </div>

                <p class="get-access-text">Included with IVPN subscriptions. One account for aliases, wildcards, temp inboxes and forwarding — no separate billing.</p>

                <div class="get-access-actions">
                    <a :href="signupUrl" class="cta-btn cta-primary" target="_blank">./SIGNUP</a>
                    <a href="https://github.com/ivpn/mailx" target="_blank" class="cta-btn cta-secondary">./VIEW_SOURCE</a>
                </div>
            </div>
        </section>

        <!-- Section 9: Footer Strip -->
        <section class="landing-footer-strip-section">
            <div class="landing-footer-strip">
                <span class="landing-footer-strip-text">2026 FreeTheMail</span>
                <span class="landing-footer-strip-divider" aria-hidden="true"></span>
                <a
                    href="/privacy"
                    class="landing-footer-strip-text landing-footer-strip-link"
                >
                    ./PRIVACY
                </a>
                <a
                    href="/tos"
                    class="landing-footer-strip-text landing-footer-strip-link"
                >
                    ./TERMS
                </a>
                <a
                    href="/faq"
                    class="landing-footer-strip-text landing-footer-strip-link"
                >
                    ./FAQ
                </a>
                <a
                    href="https://github.com/ivpn/mailx/blob/main/LICENSE.md"
                    target="_blank"
                    class="landing-footer-strip-text landing-footer-strip-link"
                >
                    ./LICENSE
                </a>
            </div>
        </section>
        
    </div>
</template>

<style scoped>
/* =============================================================================
   Shared section layout helpers
   Used across: How It Works, Verifiable Privacy, Constraints, Get Access.
   ========================================================================== */

.section-container {
    @apply w-full flex flex-col gap-8;
    @apply max-w-[1060px] mx-auto;
    @apply px-4 md:px-8;
}

.section-header {
    @apply flex flex-col gap-4;
}

/* Blue command / path badge */
.section-command {
    @apply text-xs leading-3 px-3 py-3 w-fit;
    @apply bg-[rgba(238,93,31,0.12)] dark:bg-[#12161b];
    @apply text-[#ee5d1f] dark:text-[#f76c1d];
}

/* Blue section title (32 px mobile -> 36 px desktop) */
.section-title {
    @apply m-0 uppercase font-bold;
    @apply text-[#ee5d1f] dark:text-[#f76c1d];
    @apply text-[32px] leading-[32px] md:text-[36px] md:leading-[40px];
}

.landing-page {
    @apply w-full;
    @apply bg-white dark:bg-[#0a0a0a];
    @apply min-h-screen;

    * {
        font-family: 'Roboto Mono', monospace;
    }
}

.landing-header {
    @apply flex items-center justify-between;
    @apply border-b border-solid;
    @apply border-light-gray-6 dark:border-dark-neutral-3;
    @apply py-3;
    /* Responsive padding: 110px desktop, 32px tablet, 16px mobile */
    @apply px-4 md:px-8 lg:px-[110px];
    @apply w-full;
}

.landing-logo-link {
    @apply p-0 flex items-center;
}

.landing-logo {
    @apply block bg-no-repeat bg-contain bg-center;
    @apply bg-[url("../assets/freethemail-light.png")] dark:bg-[url("../assets/freethemail-dark.png")];
    /* Desktop/Tablet: 118px x 25px, Mobile: 98px x 20px */
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
    @apply text-[#ee5d1f] dark:text-[#ee5d1f];
}

.landing-btn-signup:hover {
    @apply opacity-80;
}

.landing-btn-login {
    @apply px-3;
    @apply bg-[#ee5d1f] dark:bg-[#ee5d1f];
    @apply text-[#fffeff] dark:text-[#0a0a0a];
}

.landing-btn-login:hover {
    @apply opacity-90;
}

/* Hero Section */
.hero-section {
    @apply w-full flex flex-col;
    @apply relative overflow-hidden;
}

/* Background patterns - hidden on mobile, shown on tablet/desktop */
.hero-background {
    @apply absolute pointer-events-none;
    @apply hidden md:block;
    /* Tether to top, right, bottom - resize with container while maintaining 1:1 ratio */
    top: -30px;
    right: 0;
    bottom: -30px;
    /* Use aspect ratio to maintain 1:1 while filling height */
    aspect-ratio: 1 / 1;
    background-size: contain;
    background-repeat: no-repeat;
    background-position: center right;
}

/* Light mode backgrounds */
.hero-background.light-only {
    background-image: url('../assets/hero-bg-light-tablet.svg');
}

@media (min-width: 1024px) {
    .hero-background.light-only {
        background-image: url('../assets/hero-bg-light-desktop.svg');
    }
}

/* Dark mode backgrounds */
.hero-background.dark-only {
    background-image: url('../assets/hero-bg-dark-tablet.svg');
}

@media (min-width: 1024px) {
    .hero-background.dark-only {
        background-image: url('../assets/hero-bg-dark-desktop.svg');
    }
}

.hero-content {
    @apply flex flex-col;
    @apply py-8 md:py-8 lg:py-8;
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
    @apply text-[#1c1c1c] dark:text-[#fffeff];
    /* Responsive font sizes: 60px desktop, 52px tablet, 42px mobile */
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
    @apply text-[#ee5d1f] dark:text-[#f76c1d];
    @apply text-base;
    @apply shrink-0;
    @apply flex items-center justify-center;
    @apply w-6 h-6;
}

.feature-text {
    @apply text-black dark:text-[#fffeff];
    @apply text-base leading-[21px] md:leading-4;
    @apply m-0;
    @apply flex-1;
}

.hero-cta {
    @apply flex gap-2 items-center mt-8 z-10;
    /* Full width buttons on mobile, auto width on desktop */
    @apply flex-col md:flex-row w-full md:w-auto;
    @apply relative;
}

.cta-btn {
    @apply flex items-center justify-center;
    @apply px-4 py-[13px];
    @apply font-semibold text-sm leading-4;
    @apply whitespace-nowrap;
    /* Full width on mobile */
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
    @apply text-[#fffeff] dark:text-[#fffeff];
    border: 1px solid var(--accent);
    background: var(--accent-grad);
    box-shadow: var(--raised-shadow);
    text-shadow: var(--accent-text-glow);
}

.cta-primary:hover {
    background: var(--accent-grad-hover);
    box-shadow: var(--raised-shadow), var(--accent-glow);
}

.cta-primary:hover {
    @apply opacity-90;
}

.cta-secondary {
    @apply px-[15px] py-[12px];
    @apply text-[#ee5d1f] dark:text-[#f76c1d];
    border: 1px solid var(--sk-border);
    background: transparent;
}

.cta-secondary:hover {
    background: var(--inset-bg);
    box-shadow: var(--inset-shadow);
}

.cta-secondary:hover {
    @apply opacity-80;
}

/* Screenshot Section */
.hero-screenshot {
    @apply relative z-10;
    @apply max-w-[1060px] mx-auto;
    @apply px-4 md:px-8;
    @apply w-full;
}

.screenshot-container {
    @apply w-full;
    @apply bg-[#f3f4f5] dark:bg-[#141212];
    @apply p-2;
}

.screenshot-border {
    @apply w-full;
    @apply bg-white dark:bg-[#191919];
    @apply border border-solid;
    @apply border-[#dbdfe5] dark:border-[#262626];
    @apply p-1;
    @apply overflow-hidden;
}

.screenshot-image {
    @apply w-full;
    @apply block;
    @apply h-auto;
}

/* Show/hide images based on theme */
img.light-only {
    @apply dark:hidden !important;
}

img.dark-only {
    @apply hidden dark:block !important;
}

/* Override theme visibility for backgrounds on mobile - force hide */
.hero-background.light-only {
    @apply dark:hidden md:block;
}

.hero-background.dark-only {
    @apply hidden dark:md:block;
}

/* Section 3: How It Works */
.how-it-works-section {
    @apply w-full;
    @apply bg-white dark:bg-[#0a0a0a];
    @apply py-[72px];
}

/* .section-container / .section-header cover this section */

.how-it-works-title {
    @apply font-bold uppercase;
    @apply text-[#ee5d1f] dark:text-[#f76c1d];
    @apply text-[36px] leading-[40px];
    @apply m-0;
}

.landing-page .how-it-works-subtitle {
    @apply text-light-neutral-11 dark:text-dark-neutral-11;
    @apply text-base m-0 max-w-[587px];
    font-family: var(--font-sans);
    line-height: 1.65;
}

/* Flow Diagrams */


.instruction-item {
    @apply flex gap-2 items-start;
}

.instruction-number {
    @apply text-[#ee5d1f] dark:text-[#f76c1d];
    @apply text-sm leading-4;
    @apply w-5 h-5;
    @apply flex items-center justify-center;
    @apply shrink-0;
}

.instruction-text {
    @apply text-black dark:text-[#fffeff];
    @apply text-sm leading-5;
    @apply flex-1;
}

/* Feature Cards */
.feature-cards {
    @apply border border-solid;
    @apply border-[#dbdfe5] dark:border-[#282727];
    @apply flex flex-col lg:flex-row;
}

.feature-card {
    @apply flex-1;
    @apply p-6;
    @apply flex flex-col gap-4;
    @apply border-[#dbdfe5] dark:border-[#282727];
    @apply border-b lg:border-b-0 lg:border-r;
    @apply last:border-b-0 lg:last:border-r-0;
}

.feature-card-header {
    @apply flex gap-3 items-center;
}

.feature-icon {
    @apply w-6 h-6;
    @apply shrink-0;
}

.feature-card-title {
    @apply text-[#ee5d1f] dark:text-[#f76c1d];
    @apply text-base leading-4;
    @apply font-medium uppercase;
    @apply m-0;
}

.landing-page .feature-card-text {
    @apply text-light-neutral-11 dark:text-dark-neutral-11;
    @apply text-sm m-0;
    font-family: var(--font-sans);
    line-height: 1.65;
}

/* Section 4: Feature Set */

/* .section-container / .section-header / .section-command / .section-title cover this section */


/* Section 5: Verifiable Privacy */
.verifiable-privacy-section {
    @apply w-full;
    @apply bg-white dark:bg-[#050607];
    @apply py-6 md:py-[52px] lg:py-[72px];
}

/* .section-container / .section-header / .section-command / .section-title cover this section */

.verifiable-privacy-break {
    @apply md:hidden;
}

.verifiable-privacy-grid {
    @apply w-full;
    @apply flex flex-col md:flex-row md:flex-wrap;
    @apply pb-px md:pb-0 md:pr-px;
}

.verifiable-privacy-card {
    @apply w-full md:w-1/2;
    @apply border border-solid;
    @apply border-[#dbdfe5] dark:border-[#282727];
    @apply -mb-px md:mb-0 md:-mr-px;
    @apply px-4 py-5 md:p-6;
    @apply flex flex-col gap-4;
    @apply md:mb-[-1px] lg:mb-[-1px];
}

.verifiable-privacy-card-content {
    @apply flex flex-col gap-3;
}

.verifiable-privacy-card-title-row {
    @apply flex items-center gap-3;
}

.verifiable-privacy-icon {
    @apply w-6 h-6;
    @apply shrink-0;
}

.verifiable-privacy-card-title {
    @apply m-0;
    @apply text-base leading-4 font-medium uppercase;
    @apply text-[#ee5d1f] dark:text-[#f76c1d];
}

.landing-page .verifiable-privacy-card-text {
    @apply text-sm m-0;
    @apply text-light-neutral-11 dark:text-dark-neutral-11;
    font-family: var(--font-sans);
    line-height: 1.65;
}

.verifiable-privacy-card-link {
    @apply m-0;
    @apply inline-block;
    @apply pt-2;
    @apply text-sm leading-4 font-medium;
    @apply text-[#ee5d1f] dark:text-[#ee5d1f];
    @apply no-underline;
}

.verifiable-privacy-card-link:hover {
    @apply opacity-80;
}

.trust-card-accountable {
    @apply order-1;
}

.trust-card-open-source {
    @apply order-2 md:order-3 lg:order-2;
}

.trust-card-security {
    @apply order-3 md:order-2 lg:order-3;
}

.trust-card-no-tracking {
    @apply order-4;
}


/* Section 7: Constraints */
.constraints-section {
    @apply w-full;
    @apply bg-[#fffeff] dark:bg-[#0a0a0a];
    @apply py-6 md:py-[52px] lg:py-[72px];
}

/* .section-container / .section-header cover this section */

.constraints-command {
    @apply bg-[rgba(195,31,31,0.05)] dark:bg-[rgba(214,64,64,0.05)];
    @apply lg:bg-[rgba(195,31,31,0.1)] lg:dark:bg-[rgba(214,64,64,0.05)];
    @apply text-[#c31f1f] dark:text-[#d64040];
    @apply text-xs leading-3;
    @apply px-3 py-3;
    @apply w-fit;
}

.constraints-title {
    @apply m-0;
    @apply uppercase font-bold;
    @apply text-[#c31f1f] dark:text-[#d64040];
    @apply text-[32px] leading-[32px] md:text-[36px] md:leading-[40px];
}

.constraints-list {
    @apply flex flex-col gap-3;
}

.constraints-item {
    @apply flex items-center gap-2;
}

.constraints-bullet {
    @apply text-[#c31f1f] dark:text-[#d64040];
    @apply text-base;
    @apply w-6 h-6;
    @apply shrink-0;
    @apply flex items-center justify-center;
}

.constraints-text {
    @apply m-0;
    @apply flex-1;
    @apply text-[#8d1719] dark:text-[#eb7676];
    @apply text-sm leading-5 md:text-base md:leading-4;
}

/* Section 8: Get Access */
.get-access-section {
    @apply w-full;
    @apply bg-[#f3f4f5] dark:bg-[rgba(255,255,255,0.02)];
    @apply py-6 md:py-[52px] lg:py-[72px];
    @apply md:border md:border-solid;
    @apply md:border-[#dbdfe5] md:dark:border-[#282727];
}

/* .section-container / .section-header / .section-command / .section-title cover this section */

/* Section 9: Footer Strip */
.landing-footer-strip-section {
    @apply w-full;
    @apply bg-[#ffffff] dark:bg-[#050607];
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
    @apply text-[#8e3510];
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
    @apply border-[#b6c1ce] dark:border-[#2d3d4d];
}
.landing-page .section-title,
.landing-page .how-it-works-title,
.landing-page .constraints-title {
    @apply text-[#1c1c1c] dark:text-[#fffeff];
    font-family: var(--font-serif);
    font-weight: 600;
    letter-spacing: -0.02em;
    text-transform: none;
}
.landing-page .get-access-text {
    @apply text-light-neutral-11 dark:text-dark-neutral-11;
    @apply text-base m-0 max-w-[52ch];
    font-family: var(--font-sans);
    line-height: 1.65;
}

.get-access-actions {
    @apply flex flex-col md:flex-row gap-3;
}
</style>
