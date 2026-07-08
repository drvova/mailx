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
    suspended_users: number
    admin_users: number
    total_aliases: number
    total_domains: number
    total_recipients: number
    total_logs: number
    total_inbox_messages: number
    total_subscriptions: number
    active_subscriptions: number
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
    aliases: async (search?: string, offset = 0) => (await api.get('/admin/aliases', { params: { offset, ...(search ? { search } : {}) } })).data as { aliases: AdminAlias[]; total: number },
    deleteAlias: async (id: string) => api.delete(`/admin/alias/${id}`),
    toggleAlias: async (id: string, enabled: boolean) => api.put(`/admin/alias/${id}/toggle`, { enabled }),
    // Domain moderation
    domains: async () => (await api.get('/admin/domains')).data as AdminDomain[],
    deleteDomain: async (id: string) => api.delete(`/admin/domain/${id}`),
    toggleDomain: async (id: string, enabled: boolean) => api.put(`/admin/domain/${id}/toggle`, { enabled }),
    // Recipient moderation
    recipients: async (search?: string, offset = 0) => (await api.get('/admin/recipients', { params: { offset, ...(search ? { search } : {}) } })).data as { recipients: AdminRecipient[]; total: number },
    deleteRecipient: async (id: string) => api.delete(`/admin/recipient/${id}`),
    // Log filtering
    logsFiltered: async (type?: string, offset = 0) => (await api.get('/admin/logs/filter', { params: { offset, ...(type ? { type } : {}) } })).data as { logs: AdminLog[]; total: number },
    // User search + detail
    searchUsers: async (search: string) => (await api.get('/admin/users/search', { params: { search } })).data as { users: AdminUser[]; total: number },
    userDetail: async (id: string) => (await api.get(`/admin/user/${id}/detail`)).data as { user: AdminUser; subscription: any; aliases: AdminAlias[]; recipients: AdminRecipient[]; domains: AdminDomain[] },
    // Access key moderation
    accessKeys: async (offset = 0) => (await api.get('/admin/accesskeys', { params: { offset } })).data as { keys: AdminAccessKey[]; total: number },
    deleteAccessKey: async (id: string) => api.delete(`/admin/accesskey/${id}`),
    // Session moderation
    sessions: async (offset = 0) => (await api.get('/admin/sessions', { params: { offset } })).data as { sessions: AdminSession[]; total: number },
    deleteSession: async (id: string) => api.delete(`/admin/session/${id}`),
    forceLogout: async (userId: string) => api.delete(`/admin/user/${userId}/sessions`),
    // Credential moderation
    credentials: async (offset = 0) => (await api.get('/admin/credentials', { params: { offset } })).data as { credentials: AdminCredential[]; total: number },
    deleteCredential: async (id: string) => api.delete(`/admin/credential/${id}`),
    // Subscription override
    updateSubscription: async (data: { user_id: string; tier?: string; is_active: boolean; active_until?: string }) => api.put('/admin/subscription', data),
    // Bulk operations
    bulkUpdateUsers: async (userIds: string[], isActive: boolean) => api.post('/admin/users/bulk', { user_ids: userIds, is_active: isActive }),
    // Inbox moderation
    inboxMessages: async (offset = 0) => (await api.get('/admin/inbox', { params: { offset } })).data as { messages: AdminInboxMessage[]; total: number },
    deleteInboxMessage: async (id: number) => api.delete(`/admin/inbox/message/${id}`),
    purgeInbox: async (userId: string) => api.delete(`/admin/inbox/purge/${userId}`),
    // TOTP and password management
    disableTotp: async (userId: string) => api.delete(`/admin/user/${userId}/totp`),
    resetPassword: async (userId: string, password: string) => api.post('/admin/user/reset-password', { user_id: userId, password }),
    // Settings override
    getSettings: async (userId: string) => (await api.get(`/admin/user/${userId}/settings`)).data as AdminSettings,
    updateSettings: async (data: { user_id: string; domain?: string; recipient?: string; from_name?: string; alias_format?: string; log_issues?: boolean; remove_header?: boolean }) => api.put('/admin/user/settings', data),
    // CSV export
    exportUsers: () => `${import.meta.env.VITE_API_URL}/v1/admin/export/users`,
    exportAliases: () => `${import.meta.env.VITE_API_URL}/v1/admin/export/aliases`,
    // Subscription management
    subscriptions: async (tier?: string, offset = 0) => (await api.get('/admin/subscriptions', { params: { offset, ...(tier ? { tier } : {}) } })).data as { subscriptions: AdminSubscription[]; total: number },
    deleteSubscription: async (id: string) => api.delete(`/admin/subscription/${id}`),
    // Bulk delete
    bulkDeleteAliases: async (ids: string[]) => api.post('/admin/aliases/bulk-delete', { ids }),
    bulkDeleteDomains: async (ids: string[]) => api.post('/admin/domains/bulk-delete', { ids }),
    bulkDeleteRecipients: async (ids: string[]) => api.post('/admin/recipients/bulk-delete', { ids }),
    // System health
    tableSizes: async () => (await api.get('/admin/system/tables')).data as Record<string, number>,
    recentSignups: async (days?: number) => (await api.get('/admin/system/recent-signups', { params: days ? { days } : undefined })).data as { users: AdminUser[]; count: number },
    // Domain verification
    verifyDomain: async (id: string, verified: boolean) => api.put(`/admin/domain/${id}/verify`, { verified }),
    // Impersonation
    impersonate: async (userId: string) => (await api.post(`/admin/user/${userId}/impersonate`)).data as { token: string; message: string },
    // Search keys/sessions/inbox
    searchAccessKeys: async (userId?: string) => (await api.get('/admin/accesskeys/search', { params: userId ? { user_id: userId } : undefined })).data as { keys: AdminAccessKey[]; total: number },
    searchSessions: async (userId?: string) => (await api.get('/admin/sessions/search', { params: userId ? { user_id: userId } : undefined })).data as { sessions: AdminSession[]; total: number },
    searchInbox: async (search?: string) => (await api.get('/admin/inbox/search', { params: search ? { search } : undefined })).data as { messages: AdminInboxMessage[]; total: number },
    // Messages
    messages: async (type?: string, offset = 0) => (await api.get('/admin/messages', { params: { offset, ...(type ? { type } : {}) } })).data as { messages: any[]; total: number },
    // User stats
    userStats: async (id: string) => (await api.get(`/admin/user/${id}/stats`)).data as { forwards: number; blocks: number; replies: number; sends: number; aliases: number },
    // Log search (text + type)
    searchLogs: async (search: string, type?: string, offset = 0) => (await api.get('/admin/logs/search', { params: { offset, search, ...(type ? { type } : {}) } })).data as { logs: AdminLog[]; total: number },
    logsDateRange: async (from: string, to: string, type?: string, offset = 0) => (await api.get('/admin/logs/date-range', { params: { offset, from, to, ...(type ? { type } : {}) } })).data as { logs: AdminLog[]; total: number },
    // Recipient toggle
    toggleRecipient: async (id: string, isActive: boolean) => api.put(`/admin/recipient/${id}/toggle`, { is_active: isActive }),
    // Domain search
    searchDomains: async (search: string) => (await api.get('/admin/domains/search', { params: { search } })).data as { domains: AdminDomain[]; total: number },
    // CSV export recipients and subscriptions
    exportRecipients: () => `${import.meta.env.VITE_API_URL}/v1/admin/export/recipients`,
    exportSubscriptions: () => `${import.meta.env.VITE_API_URL}/v1/admin/export/subscriptions`,
    // Email change
    changeEmail: async (userId: string, newEmail: string) => api.put('/admin/user/email', { user_id: userId, new_email: newEmail }),
    // Bulk delete users
    bulkDeleteUsers: async (userIds: string[]) => api.post('/admin/users/bulk-delete', { user_ids: userIds }),
    // CSV export domains and logs
    exportDomains: () => `${import.meta.env.VITE_API_URL}/v1/admin/export/domains`,
    exportLogs: () => `${import.meta.env.VITE_API_URL}/v1/admin/export/logs`,
    // Config viewer
    getConfig: async () => (await api.get('/admin/system/config')).data as Record<string, any>,
    // Message search
    searchMessages: async (search?: string, type?: string) => (await api.get('/admin/messages/search', { params: { ...(search ? { search } : {}), ...(type ? { type } : {}) } })).data as { messages: any[]; total: number },
    // Recipient PGP management
    toggleRecipientPGP: async (id: string, pgpEnabled: boolean) => api.put(`/admin/recipient/${id}/pgp`, { pgp_enabled: pgpEnabled }),
    removeRecipientPGPKey: async (id: string) => api.delete(`/admin/recipient/${id}/pgp`),
    // Alias edit
    updateAlias: async (id: string, data: { description?: string; recipients?: string; from_name?: string }) => api.put(`/admin/alias/${id}`, data),
    // Domain edit
    updateDomain: async (id: string, data: { description?: string; recipient?: string; from_name?: string }) => api.put(`/admin/domain/${id}`, data),
    // Inbox mark as read
    markInboxRead: async (id: number, isRead: boolean) => api.put(`/admin/inbox/message/${id}/read`, { is_read: isRead }),
    // Paginated users
    usersPaginated: async (limit: number, offset: number, search?: string) => (await api.get('/admin/users/paginated', { params: { limit, offset, ...(search ? { search } : {}) } })).data as { users: AdminUser[]; total: number; limit: number; offset: number },
    // Create resources for users
    createRecipient: async (userId: string, email: string) => api.post('/admin/recipient', { user_id: userId, email }),
    createDomain: async (userId: string, name: string) => api.post('/admin/domain', { user_id: userId, name }),
    // CSV exports for inbox and messages
    exportInbox: async () => (await api.get('/admin/export/inbox')).data as string,
    exportMessages: async () => (await api.get('/admin/export/messages')).data as string,
    // Enriched export
    exportUsersEnriched: async () => (await api.get('/admin/export/users-enriched')).data as string,
    // Bulk revoke keys
    bulkDeleteAccessKeys: async (ids: string[]) => api.post('/admin/accesskeys/bulk-revoke', { ids }),
    // Bulk remove credentials
    bulkDeleteCredentials: async (ids: string[]) => api.post('/admin/credentials/bulk-remove', { ids }),
    // Bulk extend subscriptions
    bulkExtendSubscriptions: async (ids: string[], days: number) => (await api.post('/admin/subscriptions/bulk-extend', { ids, days })).data as { message: string },
    // Alias creation
    createAlias: async (userId: string, name: string, enabled: boolean) => api.post('/admin/alias', { user_id: userId, name, enabled }),
    // Recipient edit
    updateRecipient: async (id: string, email: string) => api.put(`/admin/recipient/${id}`, { email }),
    // Log deletion
    deleteLog: async (id: string) => api.delete(`/admin/log/${id}`),
    // Inbox bulk delete
    bulkDeleteInbox: async (ids: number[]) => api.post('/admin/inbox/bulk-delete', { ids }),
    // Subscription extend
    extendSubscription: async (id: string, days: number) => api.post(`/admin/subscription/${id}/extend`, { days }),
    // Access key creation
    createAccessKey: async (userId: string, name: string) => (await api.post('/admin/accesskey', { user_id: userId, name })).data as { key: string; message: string },
    // Ownership transfer
    transferAlias: async (id: string, newUserId: string) => api.post(`/admin/alias/${id}/transfer`, { new_user_id: newUserId }),
    transferDomain: async (id: string, newUserId: string) => api.post(`/admin/domain/${id}/transfer`, { new_user_id: newUserId }),
    // Purge operations
    purgeLogs: async (days: number, logType?: string) => (await api.post('/admin/logs/purge', { days, ...(logType ? { log_type: logType } : {}) })).data as { message: string },
    purgeAllInbox: async () => (await api.delete('/admin/inbox/purge-all')).data as { message: string },
    // User creation
    createUser: async (email: string, password: string) => (await api.post('/admin/user/create', { email, password })).data as { message: string; user: AdminUser },
    // Inbox raw content
    inboxRaw: async (id: number) => (await api.get(`/admin/inbox/message/${id}/raw`)).data as string,
    // Alias expiry
    setAliasExpiry: async (id: string, expiresAt: string) => api.put(`/admin/alias/${id}/expiry`, { expires_at: expiresAt }),
    setAccessKeyExpiry: async (id: string, expiresAt: string) => api.put(`/admin/accesskey/${id}/expiry`, { expires_at: expiresAt }),
    // Audit log
    auditLog: async (offset = 0) => (await api.get('/admin/audit', { params: { offset } })).data as { entries: AdminAudit[]; total: number },
    // Session data
    sessionData: async (id: string) => (await api.get(`/admin/session/${id}/data`)).data as any,
}

export interface AdminAlias {
    id: string
    name: string
    user_id: string
    enabled: boolean
    description: string
    recipients: string
    from_name: string
    catch_all: boolean
    created_at: string
}

export interface AdminDomain {
    id: string
    name: string
    user_id: string
    enabled: boolean
    description: string
    recipient: string
    from_name: string
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

export interface AdminInboxMessage {
    id: number
    user_id: string
    alias_id: string
    from: string
    from_name: string
    subject: string
    read: boolean
    size: number
    created_at: string
}

export interface AdminSettings {
    id: string
    user_id: string
    domain: string
    recipient: string
    from_name: string
    alias_format: string
    log_issues: boolean
    remove_header: boolean
}

export interface AdminSubscription {
    id: string
    user_id: string
    type: string
    is_active: boolean
    tier: string
    active_until: string
    plan_id: string | null
    created_at: string
}

export interface AdminAudit {
    id: number
    admin_email: string
    action: string
    target: string
    details: string
    created_at: string
}
