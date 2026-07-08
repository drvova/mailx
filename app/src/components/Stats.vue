<template>
    <div class="card-container">
        <header class="head">
            <h2>Stats</h2>
        </header>
        <SkeletonRows v-if="!loaded && !error" :rows="4" :cols="4" />
        <div v-if="loaded || error" class="card-primary pt-7">
            <h3>Messages in last 7 days</h3>
            <div id="chart" role="img" :aria-label="chartLabel" class="mb-5"></div>
            <h3>Messages in last 90 days</h3>
            <div class="grid grid-cols-2 md:grid-cols-4 gap-4 text-center mb-8">
                <div class="p-4 border-r border-secondary">
                    <h1 class="mb-2">{{ stats.forwards }}</h1>
                    <p class="m-0">Forwards</p>
                </div>
                <div class="p-4 border-r border-transparent md:border-secondary">
                    <h1 class="mb-2">{{ stats.blocks }}</h1>
                    <p class="m-0">Blocks</p>
                </div>
                <div class="p-4 border-r border-secondary">
                    <h1 class="mb-2">{{ stats.replies }}</h1>
                    <p class="m-0">Replies</p>
                </div>
                <div class="p-4 border-r border-transparent">
                    <h1 class="mb-2">{{ stats.sends }}</h1>
                    <p class="m-0">Sends</p>
                </div>
            </div>
            <p v-if="error" class="error" role="alert">Error: {{ error }}</p>
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted, onUnmounted } from 'vue'
import { ApiError } from '../api/api.ts'
import { userApi } from '../api/user.ts'
import ApexCharts from 'apexcharts'
import events from '../events.ts'
import SkeletonRows from './SkeletonRows.vue'

interface Message {
    created_at: string
    type: number
}

interface CountData {
    name: string
    data: number[]
}

const stats = ref({
    forwards: 0,
    blocks: 0,
    replies: 0,
    sends: 0,
    aliases: 0,
    messages: [],
})
const error = ref('')
const loaded = ref(false)

const getStats = async () => {
    try {
        const response = await userApi.stats()
        stats.value = response.data
        if (!stats.value.messages) {
            stats.value.messages = []
        }
        error.value = ''
        initChart()
    } catch (err) {
        if (err instanceof ApiError) {
            error.value = err.data?.error || err.message || err.message
        }
    } finally {
        loaded.value = true
    }
}

const initChart = () => {
    const options = {
        chart: {
            type: 'bar',
            height: 350,
            toolbar: {
                show: false,
            },
            background: 'transparent',
        },
        dataLabels: {
            enabled: false,
        },
        series: getLast7DaysCounts(stats.value.messages),
        xaxis: {
            categories: getLast7Days(),
        },
        yaxis: {
            forceNiceScale: true,
        },
        theme: {
            mode: getTheme()
        },
    }

    chart = new ApexCharts(document.querySelector('#chart'), options)
    chart.render()
}

const chartLabel = computed(() => {
    const total = stats.value.forwards + stats.value.blocks + stats.value.replies + stats.value.sends
    return `Bar chart: messages in last 7 days. ${total} total: ${stats.value.forwards} forwards, ${stats.value.blocks} blocks, ${stats.value.replies} replies, ${stats.value.sends} sends.`
})

let chart: ApexCharts | undefined

const onThemeChange = ({ mode }: { mode: 'light' | 'dark' }) => {
    chart?.updateOptions({ theme: { mode } })
}

function getTheme(): string {
    if (localStorage.theme === 'dark' || (!('theme' in localStorage) && window.matchMedia('(prefers-color-scheme: dark)').matches)) {
        return 'dark'
    } else {
        return 'light'
    }
}

function getLast7Days(): string[] {
    const result: string[] = []

    for (let i = 6; i >= 0; i--) {
        const today = new Date()
        today.setDate(today.getDate() - i)
        result.push((today).toLocaleDateString('en-US', { weekday: 'short' }))
    }

    return result
}

function getLast7DaysCounts(messages: Message[]): CountData[] {
    const typeNames = ['Forwards', 'Blocks', 'Replies', 'Sends']
    const last7Days = getLast7Days()

    const days: { [key: string]: number[] } = {
        Forwards: Array(7).fill(0),
        Blocks: Array(7).fill(0),
        Replies: Array(7).fill(0),
        Sends: Array(7).fill(0),
    }

    const now = new Date()
    const nowDay = (now).toLocaleDateString('en-US', { weekday: 'short' })
    const sevenDaysAgo = new Date(now)
    sevenDaysAgo.setDate(sevenDaysAgo.getDate() - 7)

    messages.forEach((msg) => {
        const messageDate = new Date(msg.created_at)
        const messageDay = (messageDate).toLocaleDateString('en-US', { weekday: 'short' })

        if (nowDay === messageDay && now.getDate() !== messageDate.getDate()) return

        if (messageDate >= sevenDaysAgo) {
            const typeIndex = msg.type
            const dayIndex = last7Days.indexOf(messageDay)

            if (typeIndex >= 0 && typeIndex < 4 && dayIndex >= 0 && dayIndex < 7) {
                days[typeNames[typeIndex]][dayIndex]++
            }
        }
    })

    const result: CountData[] = []
    for (let i = 0; i < 4; i++) {
        result.push({
            name: typeNames[i],
            data: days[typeNames[i]],
        })
    }

    return result
}

onMounted(() => {
    getStats()
    events.on('theme.change', onThemeChange)
})

onUnmounted(() => {
    events.off('theme.change', onThemeChange)
    chart?.destroy()
})
</script>