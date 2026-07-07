<template>
    <div class="card-primary" role="status" aria-label="Loading">
        <div class="table-container">
            <table>
                <tbody>
                    <tr v-for="r in rows" :key="r">
                        <td v-for="c in cols" :key="c">
                            <span class="sk-bar" :style="{ width: barWidth(r, c) }"></span>
                        </td>
                    </tr>
                </tbody>
            </table>
        </div>
    </div>
</template>

<script setup lang="ts">
withDefaults(defineProps<{ rows?: number, cols?: number }>(), {
    rows: 5,
    cols: 4,
})

// Deterministic width variation: organic look, zero re-render churn
const barWidth = (r: number, c: number) => `${45 + ((r * 17 + c * 31) % 40)}%`
</script>
