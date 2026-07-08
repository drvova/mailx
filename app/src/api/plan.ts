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
    // Alias moderation
    aliases: async (search?: string) => (await api.get('/admin/aliases', { params: search ? { search } : undefined })).data as { aliases: AdminAlias[]; total: number },
    deleteAlias: async (id: string) => api.delete(`/admin/alias/${id}`),
    toggleAlias: async (id: string, enabled: boolean) => api.put(`/admin/alias/${id}/toggle`, { enabled }),
    // Domain moderation
    domains: async () => (await api.get('/admin/domains')).data as AdminDomain[],
    deleteDomain: async (id: string) => api.delete(`/admin/domain/${id}`),
    toggleDomain: async (id: string, enabled: boolean) => api.put(`/admin/domain/${id}/toggle`, { enabled }),
    // Recipient moderation
    recipients: async (search?: string) => (await api.get('/admin/recipients', { params: search ? { search } : undefined })).data as { recipients: AdminRecipient[]; total: number },
    deleteRecipient: async (id: string) => api.delete(`/admin/recipient/${id}`),
    // Log filtering
    logsFiltered: async (type?: string) => (await api.get('/admin/logs/filter', { params: type ? { type } : undefined })).data as { logs: AdminLog[]; total: number },
    // User search + detail
    searchUsers: async (search: string) => (await api.get('/admin/users/search', { params: { search } })).data as { users: AdminUser[]; total: number },
    userDetail: async (id: string) => (await api.get(`/admin/user/${id}/detail`)).data as { user: AdminUser; subscription: any; aliases: AdminAlias[]; recipients: AdminRecipient[]; domains: AdminDomain[] },
    // Access key moderation
    accessKeys: async () => (await api.get('/admin/accesskeys')).data as { keys: AdminAccessKey[]; total: number },
    deleteAccessKey: async (id: string) => api.delete(`/admin/accesskey/${id}`),
    // Session moderation
    sessions: async () => (await api.get('/admin/sessions')).data as { sessions: AdminSession[]; total: number },
    deleteSession: async (id: string) => api.delete(`/admin/session/${id}`),
    forceLogout: async (userId: string) => api.delete(`/admin/user/${userId}/sessions`),
    // Credential moderation
    credentials: async () => (await api.get('/admin/credentials')).data as { credentials: AdminCredential[]; total: number },
    deleteCredential: async (id: string) => api.delete(`/admin/credential/${id}`),
    // Subscription override
    updateSubscription: async (data: { user_id: string; tier?: string; is_active: boolean; active_until?: string }) => api.put('/admin/subscription', data),
    // Bulk operations
    bulkUpdateUsers: async (userIds: string[], isActive: boolean) => api.post('/admin/users/bulk', { user_ids: userIds, is_active: isActive }),
}

export interface AdminAlias {
    id: string
    name: string
    user_id: string
    enabled: boolean
    description: string
    catch_all: boolean
    created_at: string
}

export interface AdminDomain {
    id: string
    name: string
    user_id: string
    enabled: boolean
    description: string
    owner_verified_at: string | null
    mx_verified_at: string | null
    created_at: string
}

export interface AdminRecipient {
    id: string
    email: string
    user_id: string
    is_active: boolean
    pgp_enabled: boolean
    created_at: string
}

export interface AdminAccessKey {
    id: string
    user_id: string
    name: string
    expires_at: string | null
    created_at: string
}

export interface AdminSession {
    id: string
    token: string
    expires_at: string
    created_at: string
}

export interface AdminCredential {
    id: string
    user_id: string
    created_at: string
}
