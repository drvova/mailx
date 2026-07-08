import { api } from './api'

export interface Plan {
    id: string
    name: string
    display_name: string
    price_cents: number
    currency: string
    interval: string
    max_recipients: number
    max_credentials: number
    max_daily_aliases: number
    max_daily_send_reply: number
    max_sessions: number
    is_active: boolean
    sort_order: number
}

export const planApi = {
    list: async () => (await api.get('/plans')).data as Plan[],
    listAll: async () => (await api.get('/admin/plans')).data as Plan[],
    create: async (data: Partial<Plan>) => (await api.post('/admin/plan', data)).data as Plan,
    update: async (id: string, data: Partial<Plan>) => (await api.put(`/admin/plan/${id}`, data)).data as Plan,
    delete: async (id: string) => api.delete(`/admin/plan/${id}`),
}

export const billingApi = {
    checkout: async (planId: string) => (await api.post('/billing/checkout', { plan_id: planId })).data as { url: string },
}
