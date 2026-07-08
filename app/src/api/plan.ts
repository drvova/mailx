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

export interface SystemStats {
    total_users: number
    active_users: number
    total_aliases: number
    total_domains: number
    total_logs: number
    active_plans: number
}

export interface AdminUser {
    id: string
    email: string
    is_active: boolean
    is_admin: boolean
    created_at: string
}

export interface AdminLog {
    id: string
    created_at: string
    attempted_at: string
    log_type: string
    from: string
    destination: string
    message: string
    status: string
    remote_mta: string
}

export const adminApi = {
    users: async () => (await api.get('/admin/users')).data as AdminUser[],
    stats: async () => (await api.get('/admin/stats')).data as SystemStats,
    logs: async () => (await api.get('/admin/logs')).data as AdminLog[],
    updateUser: async (data: { id: string; is_active?: boolean; is_admin?: boolean }) => api.put('/admin/user', data),
    deleteUser: async (id: string) => api.delete(`/admin/user/${id}`),
    assignPlan: async (userId: string, planId: string) => api.post('/admin/user/assign-plan', { user_id: userId, plan_id: planId }),
}
