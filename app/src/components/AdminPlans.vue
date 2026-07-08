<template>
    <div class="card-container">
        <header class="head">
            <h2>Admin</h2>
        </header>

        <div class="card-primary">
            <ul class="flex gap-4 border-b mb-6 overflow-x-auto" role="tablist">
                <li><button :class="tab === 'stats' ? 'font-bold border-b-2' : ''" @click="tab = 'stats'" role="tab">Stats</button></li>
                <li><button :class="tab === 'users' ? 'font-bold border-b-2' : ''" @click="tab = 'users'" role="tab">Users</button></li>
                <li><button :class="tab === 'aliases' ? 'font-bold border-b-2' : ''" @click="tab = 'aliases'" role="tab">Aliases</button></li>
                <li><button :class="tab === 'domains' ? 'font-bold border-b-2' : ''" @click="tab = 'domains'" role="tab">Domains</button></li>
                <li><button :class="tab === 'recipients' ? 'font-bold border-b-2' : ''" @click="tab = 'recipients'" role="tab">Recipients</button></li>
                <li><button :class="tab === 'plans' ? 'font-bold border-b-2' : ''" @click="tab = 'plans'" role="tab">Plans</button></li>
                <li><button :class="tab === 'keys' ? 'font-bold border-b-2' : ''" @click="tab = 'keys'" role="tab">API Keys</button></li>
                <li><button :class="tab === 'sessions' ? 'font-bold border-b-2' : ''" @click="tab = 'sessions'" role="tab">Sessions</button></li>
                <li><button :class="tab === 'passkeys' ? 'font-bold border-b-2' : ''" @click="tab = 'passkeys'" role="tab">Passkeys</button></li>
                <li><button :class="tab === 'inbox' ? 'font-bold border-b-2' : ''" @click="tab = 'inbox'" role="tab">Inbox</button></li>
                <li><button :class="tab === 'messages' ? 'font-bold border-b-2' : ''" @click="tab = 'messages'" role="tab">Messages</button></li>
                <li><button :class="tab === 'subs' ? 'font-bold border-b-2' : ''" @click="tab = 'subs'" role="tab">Subscriptions</button></li>
                <li><button :class="tab === 'system' ? 'font-bold border-b-2' : ''" @click="tab = 'system'" role="tab">System</button></li>
                <li><button :class="tab === 'logs' ? 'font-bold border-b-2' : ''" @click="tab = 'logs'" role="tab">Logs</button></li>
                <li><button :class="tab === 'audit' ? 'font-bold border-b-2' : ''" @click="tab = 'audit'" role="tab">Audit</button></li>
            </ul>

            <!-- STATS -->
            <div v-if="tab === 'stats'" role="tabpanel">
                <div v-if="stats" class="grid grid-cols-2 md:grid-cols-4 gap-4">
                    <div class="card-secondary text-center"><p class="text-3xl font-bold">{{ stats.total_users }}</p><p class="text-sm text-gray-500">Total Users</p></div>
                    <div class="card-secondary text-center"><p class="text-3xl font-bold">{{ stats.active_users }}</p><p class="text-sm text-gray-500">Active Users</p></div>
                    <div class="card-secondary text-center"><p class="text-3xl font-bold text-red-500">{{ stats.suspended_users }}</p><p class="text-sm text-gray-500">Suspended</p></div>
                    <div class="card-secondary text-center"><p class="text-3xl font-bold text-blue-500">{{ stats.admin_users }}</p><p class="text-sm text-gray-500">Admins</p></div>
                    <div class="card-secondary text-center"><p class="text-3xl font-bold">{{ stats.total_aliases }}</p><p class="text-sm text-gray-500">Aliases</p></div>
                    <div class="card-secondary text-center"><p class="text-3xl font-bold">{{ stats.total_domains }}</p><p class="text-sm text-gray-500">Domains</p></div>
                    <div class="card-secondary text-center"><p class="text-3xl font-bold">{{ stats.total_recipients }}</p><p class="text-sm text-gray-500">Recipients</p></div>
                    <div class="card-secondary text-center"><p class="text-3xl font-bold">{{ stats.total_inbox_messages }}</p><p class="text-sm text-gray-500">Inbox Msgs</p></div>
                    <div class="card-secondary text-center"><p class="text-3xl font-bold">{{ stats.total_subscriptions }}</p><p class="text-sm text-gray-500">Subscriptions</p></div>
                    <div class="card-secondary text-center"><p class="text-3xl font-bold text-green-500">{{ stats.active_subscriptions }}</p><p class="text-sm text-gray-500">Active Subs</p></div>
                    <div class="card-secondary text-center"><p class="text-3xl font-bold">{{ stats.total_logs }}</p><p class="text-sm text-gray-500">Log Entries</p></div>
                    <div class="card-secondary text-center"><p class="text-3xl font-bold">{{ stats.active_plans }}</p><p class="text-sm text-gray-500">Active Plans</p></div>
                    <div class="col-span-2 md:col-span-4 flex gap-2 mt-2">
                        <button class="cta cta-tertiary text-sm" @click="exportUsers">Export Users CSV</button>
                        <button class="cta cta-tertiary text-sm" @click="exportAliases">Export Aliases CSV</button>
                        <button class="cta cta-tertiary text-sm" @click="exportRecipients">Export Recipients CSV</button>
                        <button class="cta cta-tertiary text-sm" @click="exportSubscriptions">Export Subscriptions CSV</button>
                        <button class="cta cta-tertiary text-sm" @click="exportDomains">Export Domains CSV</button>
                        <button class="cta cta-tertiary text-sm" @click="exportLogs">Export Logs CSV</button>
                        <button class="cta cta-tertiary text-sm" @click="exportInbox">Export Inbox CSV</button>
                        <button class="cta cta-tertiary text-sm" @click="exportMessages">Export Messages CSV</button>
                        <button class="cta cta-tertiary text-sm" @click="exportUsersEnriched">Export Users + Subs CSV</button>
                    </div>
                    <div v-if="subStats" class="grid grid-cols-3 gap-4 mb-4">
                        <div class="card-secondary text-center"><p class="text-2xl font-bold text-green-500">{{ subStats.active }}</p><p class="text-sm text-gray-500">Active Subs</p></div>
                        <div class="card-secondary text-center"><p class="text-2xl font-bold text-yellow-500">{{ subStats.grace_period }}</p><p class="text-sm text-gray-500">Grace Period</p></div>
                        <div class="card-secondary text-center"><p class="text-2xl font-bold text-red-500">{{ subStats.expired }}</p><p class="text-sm text-gray-500">Expired</p></div>
                    </div>
                </div>
                <SkeletonRows v-else :rows="3" />
            </div>

            <!-- USERS -->
            <div v-if="tab === 'users'" role="tabpanel">
                <div class="flex gap-2 mb-4">
                    <input v-model="userSearch" placeholder="Search by email..." @input="searchUsersDeb" class="flex-1" />
                    <button class="cta cta-tertiary text-sm" @click="createUser">Create User</button>
                    <button class="cta cta-tertiary text-sm" @click="bulkSuspend">Bulk Suspend</button>
                    <button class="cta cta-tertiary text-sm" @click="bulkActivate">Bulk Activate</button>
                    <button class="cta cta-tertiary text-sm text-red-500" @click="bulkDeleteUsers">Bulk Delete</button>
                </div>
                <div v-if="users.length" class="overflow-x-auto">
                    <table class="table">
                        <thead><tr><th><input type="checkbox" @click="toggleAllUsers($event)" /></th><th>Email</th><th>Active</th><th>Admin</th><th>Joined</th><th></th></tr></thead>
                        <tbody>
                            <tr v-for="u in users" :key="u.id">
                                <td><input type="checkbox" v-model="(u as any)._selected" /></td>
                                <td><button class="text-blue-500 hover:underline" @click="viewUser(u.id)">{{ u.email }}</button></td>
                                <td><span :class="u.is_active ? 'badge badge-success' : 'badge badge-error'">{{ u.is_active ? 'Active' : 'Suspended' }}</span></td>
                                <td><span :class="u.is_admin ? 'badge badge-success' : 'badge'">{{ u.is_admin ? 'Admin' : 'User' }}</span></td>
                                <td>{{ formatDate(u.created_at) }}</td>
                                <td><div class="flex gap-1">
                                    <button class="cta cta-tertiary text-sm" @click="toggleUser(u)">{{ u.is_active ? 'Suspend' : 'Activate' }}</button>
                                    <button class="cta cta-tertiary text-sm" @click="toggleAdmin(u)">{{ u.is_admin ? 'Revoke' : 'Make Admin' }}</button>
                                    <button class="cta cta-tertiary text-sm text-red-500" @click="deleteUser(u)" :disabled="deleting === u.id" :aria-busy="deleting === u.id">Delete</button>
                                </div></td>
                            </tr>
                        </tbody>
                    </table>
                    <div v-if="usersTotal > 50" class="flex items-center justify-between mt-4">
                        <button class="cta cta-tertiary text-sm" :disabled="usersOffset === 0" @click="prevUsers">Prev</button>
                        <span class="text-sm text-gray-500">{{ usersOffset + 1 }}-{{ Math.min(usersOffset + 50, usersTotal) }} of {{ usersTotal }}</span>
                        <button class="cta cta-tertiary text-sm" :disabled="usersOffset + 50 >= usersTotal" @click="nextUsers">Next</button>
                    </div>
                </div>
                <SkeletonRows v-else :rows="5" />

                <!-- User detail modal -->
                <div v-if="userDetail" class="fixed inset-0 bg-black/50 flex items-center justify-center z-50 p-4" @click.self="userDetail = null">
                    <div class="card-primary max-w-2xl w-full max-h-[80vh] overflow-y-auto">
                        <div class="flex justify-between items-center mb-4">
                            <h3>{{ userDetail.user.email }}</h3>
                            <button class="cta cta-tertiary" @click="userDetail = null">Close</button>
                        </div>
                        <div class="grid grid-cols-2 gap-4 mb-4">
                            <div><p class="text-sm text-gray-500">Status</p><span :class="userDetail.user.is_active ? 'badge badge-success' : 'badge badge-error'">{{ userDetail.user.is_active ? 'Active' : 'Suspended' }}</span></div>
                            <div><p class="text-sm text-gray-500">Role</p><span :class="userDetail.user.is_admin ? 'badge badge-success' : 'badge'">{{ userDetail.user.is_admin ? 'Admin' : 'User' }}</span></div>
                            <div><p class="text-sm text-gray-500">Subscription Tier</p><p>{{ userDetail.subscription.tier || 'self-hosted' }}</p></div>
                            <div><p class="text-sm text-gray-500">Active Until</p><p>{{ userDetail.subscription.active_until ? formatDate(userDetail.subscription.active_until) : 'N/A' }}</p></div>
                            <div><p class="text-sm text-gray-500">2FA</p><span :class="userDetail.user.totp_enabled ? 'badge badge-success' : 'badge'">{{ userDetail.user.totp_enabled ? 'Enabled' : 'Disabled' }}</span></div>
                            <div><p class="text-sm text-gray-500">Created</p><p>{{ formatDate(userDetail.user.created_at) }}</p></div>
                            <div><p class="text-sm text-gray-500">Notes</p><p class="text-xs">{{ userDetail.user.notes || 'none' }}</p><button class="cta cta-tertiary text-sm mt-1" @click="editUserNotes(userDetail.user)">Edit Notes</button></div>
                        </div>
                        <div v-if="userStats" class="grid grid-cols-5 gap-2 mb-4 text-center">
                            <div class="card-secondary"><p class="text-lg font-bold">{{ userStats.forwards }}</p><p class="text-xs text-gray-500">Forwards</p></div>
                            <div class="card-secondary"><p class="text-lg font-bold">{{ userStats.blocks }}</p><p class="text-xs text-gray-500">Blocks</p></div>
                            <div class="card-secondary"><p class="text-lg font-bold">{{ userStats.replies }}</p><p class="text-xs text-gray-500">Replies</p></div>
                            <div class="card-secondary"><p class="text-lg font-bold">{{ userStats.sends }}</p><p class="text-xs text-gray-500">Sends</p></div>
                            <div class="card-secondary"><p class="text-lg font-bold">{{ userStats.aliases }}</p><p class="text-xs text-gray-500">Aliases</p></div>
                        </div>
                        <div class="mb-4">
                            <label class="text-sm text-gray-500">Assign Plan</label>
                            <select v-model="selectedPlan" @change="assignPlan(userDetail.user.id)" class="max-w-xs">
                                <option value="">-- Select Plan --</option>
                                <option v-for="p in plans" :key="p.id" :value="p.id">{{ p.display_name }} ({{ p.price_cents === 0 ? 'Free' : p.price_cents / 100 }})</option>
                            </select>
                        </div>
                        <div class="mb-4 flex gap-2">
                            <button class="cta cta-tertiary text-sm" @click="forceLogout(userDetail.user.id)">Force Logout</button>
                            <button class="cta cta-tertiary text-sm" @click="overrideSub(userDetail.user.id)">Override Subscription</button>
                            <button class="cta cta-tertiary text-sm" @click="disableTotp(userDetail.user.id)">Disable 2FA</button>
                            <button class="cta cta-tertiary text-sm" @click="resetPassword(userDetail.user.id)">Reset Password</button>
                            <button class="cta cta-tertiary text-sm text-red-500" @click="purgeInbox(userDetail.user.id)">Purge Inbox</button>
                            <button class="cta cta-tertiary text-sm" @click="impersonate(userDetail.user.id)">Login As</button>
                            <button class="cta cta-tertiary text-sm" @click="changeEmail(userDetail.user.id)">Change Email</button>
                            <button class="cta cta-tertiary text-sm" @click="createRecipientForUser(userDetail.user.id)">Add Recipient</button>
                            <button class="cta cta-tertiary text-sm" @click="createDomainForUser(userDetail.user.id)">Add Domain</button>
                            <button class="cta cta-tertiary text-sm" @click="createAliasForUser(userDetail.user.id)">Add Alias</button>
                            <button class="cta cta-tertiary text-sm" @click="createKeyForUser(userDetail.user.id)">Add API Key</button>
                        </div>
                        <h4 class="mb-2">Aliases ({{ userDetail.aliases.length }})</h4>
                        <div v-if="userDetail.aliases.length" class="overflow-x-auto mb-4">
                            <table class="table"><thead><tr><th>Name</th><th>Enabled</th></tr></thead><tbody>
                                <tr v-for="a in userDetail.aliases" :key="a.id"><td>{{ a.name }}</td><td><span :class="a.enabled ? 'badge badge-success' : 'badge'">{{ a.enabled ? 'Yes' : 'No' }}</span></td></tr>
                            </tbody></table>
                        </div>
                        <h4 class="mb-2">Recipients ({{ userDetail.recipients.length }})</h4>
                        <div v-if="userDetail.recipients.length" class="overflow-x-auto mb-4">
                            <table class="table"><thead><tr><th>Email</th><th>Active</th><th>PGP</th></tr></thead><tbody>
                                <tr v-for="r in userDetail.recipients" :key="r.id"><td>{{ r.email }}</td><td><span :class="r.is_active ? 'badge badge-success' : 'badge'">{{ r.is_active ? 'Yes' : 'No' }}</span></td><td>{{ r.pgp_enabled ? 'Yes' : 'No' }}</td></tr>
                            </tbody></table>
                        </div>
                        <h4 class="mb-2">Domains ({{ userDetail.domains.length }})</h4>
                        <div v-if="userDetail.domains.length" class="overflow-x-auto mb-4">
                            <table class="table"><thead><tr><th>Name</th><th>Enabled</th><th>Verified</th></tr></thead><tbody>
                                <tr v-for="d in userDetail.domains" :key="d.id"><td>{{ d.name }}</td><td><span :class="d.enabled ? 'badge badge-success' : 'badge'">{{ d.enabled ? 'Yes' : 'No' }}</span></td><td>{{ d.mx_verified_at ? 'Yes' : 'No' }}</td></tr>
                            </tbody></table>
                        </div>
                        <h4 class="mb-2">User Settings</h4>
                        <div v-if="userSettings" class="grid grid-cols-2 gap-2 text-sm">
                            <div><span class="text-gray-500">Domain:</span> {{ userSettings.domain || 'default' }}</div>
                            <div><span class="text-gray-500">Recipient:</span> {{ userSettings.recipient || 'default' }}</div>
                            <div><span class="text-gray-500">From Name:</span> {{ userSettings.from_name || 'none' }}</div>
                            <div><span class="text-gray-500">Alias Format:</span> {{ userSettings.alias_format || 'random' }}</div>
                            <div><span class="text-gray-500">Log Issues:</span> {{ userSettings.log_issues ? 'Yes' : 'No' }}</div>
                            <div><span class="text-gray-500">Remove Header:</span> {{ userSettings.remove_header ? 'Yes' : 'No' }}</div>
                        </div>
                        <div v-else class="text-sm text-gray-500 mb-4">Loading settings...</div>
                    </div>
                </div>
            </div>

            <!-- ALIASES -->
            <div v-if="tab === 'aliases'" role="tabpanel">
                <div class="flex gap-2 mb-4">
                    <input v-model="aliasSearch" placeholder="Search aliases..." @input="searchAliasesDeb" class="flex-1" />
                    <button class="cta cta-tertiary text-sm text-red-500" @click="bulkDeleteAliases">Bulk Delete Selected</button>
                </div>
                <div v-if="aliases.length" class="overflow-x-auto">
                    <table class="table">
                        <thead><tr><th><input type="checkbox" @click="toggleAllAliases($event)" /></th><th>Alias</th><th>Enabled</th><th>Catch-All</th><th>Created</th><th></th></tr></thead>
                        <tbody>
                            <tr v-for="a in aliases" :key="a.id">
                                <td><input type="checkbox" v-model="(a as any)._selected" /></td>
                                <td>{{ a.name }}</td>
                                <td><span :class="a.enabled ? 'badge badge-success' : 'badge'">{{ a.enabled ? 'Yes' : 'No' }}</span></td>
                                <td>{{ a.catch_all ? 'Yes' : 'No' }}</td>
                                <td>{{ formatDate(a.created_at) }}</td>
                                <td><div class="flex gap-1">
                                    <button class="cta cta-tertiary text-sm" @click="toggleAlias(a)">{{ a.enabled ? 'Disable' : 'Enable' }}</button>
                                    <button class="cta cta-tertiary text-sm" @click="editAlias(a)">Edit</button>
                                    <button class="cta cta-tertiary text-sm" @click="transferAlias(a)">Transfer</button>
                                    <button class="cta cta-tertiary text-sm" @click="setAliasExpiry(a)">Expiry</button>
                                    <button class="cta cta-tertiary text-sm text-red-500" @click="deleteAlias(a)">Delete</button>
                                </div></td>
                            </tr>
                        </tbody>
                    </table>
                    <div v-if="aliasesTotal > 50" class="flex items-center justify-between mt-4">
                        <button class="cta cta-tertiary text-sm" :disabled="aliasesOffset === 0" @click="aliasesOffset = Math.max(0, aliasesOffset - 50); fetchAliases()">Prev</button>
                        <span class="text-sm text-gray-500">{{ aliasesOffset + 1 }}-{{ Math.min(aliasesOffset + 50, aliasesTotal) }} of {{ aliasesTotal }}</span>
                        <button class="cta cta-tertiary text-sm" :disabled="aliasesOffset + 50 >= aliasesTotal" @click="aliasesOffset += 50; fetchAliases()">Next</button>
                    </div>
                </div>
                <SkeletonRows v-else :rows="5" />
            </div>

            <!-- DOMAINS -->
            <div v-if="tab === 'domains'" role="tabpanel">
                <div class="flex gap-2 mb-4">
                    <input v-model="domainSearch" placeholder="Search domains..." @input="searchDomainsDeb" class="flex-1" />
                    <button class="cta cta-tertiary text-sm text-red-500" @click="bulkDeleteDomains">Bulk Delete Selected</button>
                </div>
                <div v-if="domains.length" class="overflow-x-auto">
                    <table class="table">
                        <thead><tr><th><input type="checkbox" @click="toggleAllDomains($event)" /></th><th>Domain</th><th>Enabled</th><th>Verified</th><th>Created</th><th></th></tr></thead>
                        <tbody>
                            <tr v-for="d in domains" :key="d.id">
                                <td><input type="checkbox" v-model="(d as any)._selected" /></td>
                                <td>{{ d.name }}</td>
                                <td><span :class="d.enabled ? 'badge badge-success' : 'badge'">{{ d.enabled ? 'Yes' : 'No' }}</span></td>
                                <td><span :class="d.mx_verified_at ? 'badge badge-success' : 'badge'">{{ d.mx_verified_at ? 'Yes' : 'No' }}</span></td>
                                <td>{{ formatDate(d.created_at) }}</td>
                                <td><div class="flex gap-1">
                                    <button class="cta cta-tertiary text-sm" @click="toggleDomain(d)">{{ d.enabled ? 'Disable' : 'Enable' }}</button>
                                    <button class="cta cta-tertiary text-sm" @click="verifyDomain(d)">{{ d.mx_verified_at ? 'Unverify' : 'Verify' }}</button>
                                    <button class="cta cta-tertiary text-sm" @click="editDomain(d)">Edit</button>
                                    <button class="cta cta-tertiary text-sm" @click="viewDomainDNS(d)">DNS</button>
                                    <button class="cta cta-tertiary text-sm" @click="transferDomain(d)">Transfer</button>
                                    <button class="cta cta-tertiary text-sm text-red-500" @click="deleteDomain(d)">Delete</button>
                                </div></td>
                            </tr>
                        </tbody>
                    </table>
                </div>
                <SkeletonRows v-else :rows="5" />
            </div>

            <!-- RECIPIENTS -->
            <div v-if="tab === 'recipients'" role="tabpanel">
                <div class="flex gap-2 mb-4">
                    <input v-model="recipientSearch" placeholder="Search by email..." @input="searchRecipientsDeb" class="flex-1" />
                    <button class="cta cta-tertiary text-sm text-red-500" @click="bulkDeleteRecipients">Bulk Delete Selected</button>
                </div>
                <div v-if="recipients.length" class="overflow-x-auto">
                    <table class="table">
                        <thead><tr><th><input type="checkbox" @click="toggleAllRecipients($event)" /></th><th>Email</th><th>Active</th><th>PGP</th><th>Created</th><th></th></tr></thead>
                        <tbody>
                            <tr v-for="r in recipients" :key="r.id">
                                <td><input type="checkbox" v-model="(r as any)._selected" /></td>
                                <td>{{ r.email }}</td>
                                <td><span :class="r.is_active ? 'badge badge-success' : 'badge'">{{ r.is_active ? 'Yes' : 'No' }}</span></td>
                                <td>{{ r.pgp_enabled ? 'Yes' : 'No' }}</td>
                                <td>{{ formatDate(r.created_at) }}</td>
                                <td><div class="flex gap-1">
                                    <button class="cta cta-tertiary text-sm" @click="toggleRecipient(r)">{{ r.is_active ? 'Suspend' : 'Activate' }}</button>
                                    <button v-if="r.pgp_enabled" class="cta cta-tertiary text-sm" @click="toggleRecipientPGP(r)">Disable PGP</button>
                                    <button v-if="r.pgp_enabled" class="cta cta-tertiary text-sm text-red-500" @click="removePGPKey(r)">Remove PGP Key</button>
                                    <button class="cta cta-tertiary text-sm" @click="editRecipient(r)">Edit</button>
                                    <button class="cta cta-tertiary text-sm" @click="uploadPGP(r)">Upload PGP</button>
                                    <button class="cta cta-tertiary text-sm text-red-500" @click="deleteRecipient(r)">Delete</button>
                                </div></td>
                            </tr>
                        </tbody>
                    </table>
                    <div v-if="recipientsTotal > 50" class="flex items-center justify-between mt-4">
                        <button class="cta cta-tertiary text-sm" :disabled="recipientsOffset === 0" @click="recipientsOffset = Math.max(0, recipientsOffset - 50); fetchRecipients()">Prev</button>
                        <span class="text-sm text-gray-500">{{ recipientsOffset + 1 }}-{{ Math.min(recipientsOffset + 50, recipientsTotal) }} of {{ recipientsTotal }}</span>
                        <button class="cta cta-tertiary text-sm" :disabled="recipientsOffset + 50 >= recipientsTotal" @click="recipientsOffset += 50; fetchRecipients()">Next</button>
                    </div>
                </div>
                <SkeletonRows v-else :rows="5" />
            </div>

            <!-- PLANS -->
            <div v-if="tab === 'plans'" role="tabpanel">
                <div class="flex justify-between items-center mb-4">
                    <h3>Plans</h3>
                    <button class="cta" @click="showForm = !showForm" v-if="!showForm">Create Plan</button>
                </div>
                <div v-if="showForm" class="card-secondary mb-6">
                    <h4>{{ editing ? 'Edit Plan' : 'New Plan' }}</h4>
                    <div class="grid grid-cols-2 gap-4 mt-4">
                        <div><label>Name</label><input v-model="form.name" placeholder="pro" /></div>
                        <div><label>Display Name</label><input v-model="form.display_name" placeholder="Pro Plan" /></div>
                        <div><label>Price (cents)</label><input type="number" v-model.number="form.price_cents" placeholder="500" /></div>
                        <div><label>Currency</label><select v-model="form.currency"><option value="usd">USD</option><option value="eur">EUR</option><option value="gbp">GBP</option></select></div>
                        <div><label>Interval</label><select v-model="form.interval"><option value="monthly">Monthly</option><option value="yearly">Yearly</option><option value="one_time">One Time</option></select></div>
                        <div><label>Sort Order</label><input type="number" v-model.number="form.sort_order" placeholder="0" /></div>
                        <div><label>Max Recipients</label><input type="number" v-model.number="form.max_recipients" /></div>
                        <div><label>Max Passkeys</label><input type="number" v-model.number="form.max_credentials" /></div>
                        <div><label>Max Daily Aliases</label><input type="number" v-model.number="form.max_daily_aliases" /></div>
                        <div><label>Max Daily Send/Reply</label><input type="number" v-model.number="form.max_daily_send_reply" /></div>
                        <div><label>Max Sessions</label><input type="number" v-model.number="form.max_sessions" /></div>
                    </div>
                    <div class="flex gap-2 mt-4">
                        <button class="cta" @click="savePlan" :disabled="saving" :aria-busy="saving">{{ editing ? 'Update' : 'Create' }}</button>
                        <button class="cta cta-tertiary" @click="resetForm">Cancel</button>
                    </div>
                    <p v-if="formError" role="alert" class="text-red-500 mt-2">{{ formError }}</p>
                </div>
                <table v-if="plans.length" class="table">
                    <thead><tr><th>Name</th><th>Price</th><th>Interval</th><th>Recipients</th><th>Aliases/day</th><th>Active</th><th></th></tr></thead>
                    <tbody>
                        <tr v-for="plan in plans" :key="plan.id">
                            <td>{{ plan.display_name }} <span class="text-xs text-gray-500">({{ plan.name }})</span></td>
                            <td>{{ plan.price_cents === 0 ? 'Free' : (plan.price_cents / 100) + ' ' + plan.currency }}</td>
                            <td>{{ plan.interval }}</td>
                            <td>{{ plan.max_recipients }}</td>
                            <td>{{ plan.max_daily_aliases }}</td>
                            <td><span :class="plan.is_active ? 'badge badge-success' : 'badge badge-error'">{{ plan.is_active ? 'Active' : 'Inactive' }}</span></td>
                            <td><div class="flex gap-1"><button class="cta cta-tertiary text-sm" @click="editPlan(plan)">Edit</button><button class="cta cta-tertiary text-sm" @click="deletePlan(plan)" :disabled="deleting === plan.id" :aria-busy="deleting === plan.id">Delete</button></div></td>
                        </tr>
                    </tbody>
                </table>
                <div v-else-if="!loadingPlans" class="card-empty"><h3>No plans yet</h3></div>
            </div>

            <!-- API KEYS -->
            <div v-if="tab === 'keys'" role="tabpanel">
                <div class="flex gap-2 mb-4">
                    <button class="cta cta-tertiary text-sm text-red-500" @click="bulkDeleteKeys">Bulk Revoke Selected</button>
                </div>
                <div v-if="accessKeys.length" class="overflow-x-auto">
                    <table class="table">
                        <thead><tr><th><input type="checkbox" @click="toggleAllKeys($event)" /></th><th>Name</th><th>User ID</th><th>Expires</th><th>Created</th><th></th></tr></thead>
                        <tbody>
                            <tr v-for="k in accessKeys" :key="k.id">
                                <td><input type="checkbox" v-model="(k as any)._selected" /></td>
                                <td>{{ k.name }}</td>
                                <td class="text-xs text-gray-500">{{ k.user_id.slice(0,8) }}...</td>
                                <td>{{ k.expires_at ? formatDate(k.expires_at) : 'Never' }}</td>
                                <td>{{ formatDate(k.created_at) }}</td>
                                <td><div class="flex gap-1">
                                    <button class="cta cta-tertiary text-sm" @click="setKeyExpiry(k)">Expiry</button>
                                    <button class="cta cta-tertiary text-sm text-red-500" @click="deleteAccessKey(k)">Revoke</button>
                                </div></td>
                            </tr>
                        </tbody>
                    </table>
                    <div v-if="keysTotal > 50" class="flex items-center justify-between mt-4">
                        <button class="cta cta-tertiary text-sm" :disabled="keysOffset === 0" @click="keysOffset = Math.max(0, keysOffset - 50); fetchAccessKeys()">Prev</button>
                        <span class="text-sm text-gray-500">{{ keysOffset + 1 }}-{{ Math.min(keysOffset + 50, keysTotal) }} of {{ keysTotal }}</span>
                        <button class="cta cta-tertiary text-sm" :disabled="keysOffset + 50 >= keysTotal" @click="keysOffset += 50; fetchAccessKeys()">Next</button>
                    </div>
                </div>
                <SkeletonRows v-else :rows="5" />
            </div>

            <!-- SESSIONS -->
            <div v-if="tab === 'sessions'" role="tabpanel">
                <div v-if="sessions.length" class="overflow-x-auto">
                    <table class="table">
                        <thead><tr><th>Token</th><th>Expires</th><th>Created</th><th></th></tr></thead>
                        <tbody>
                            <tr v-for="s in sessions" :key="s.id">
                                <td class="text-xs text-gray-500 font-mono">{{ s.token.slice(0,16) }}...</td>
                                <td>{{ formatDate(s.expires_at) }}</td>
                                <td>{{ formatDate(s.created_at) }}</td>
                                <td><div class="flex gap-1">
                                    <button class="cta cta-tertiary text-sm" @click="viewSessionData(s)">Data</button>
                                    <button class="cta cta-tertiary text-sm text-red-500" @click="deleteSession(s)">Terminate</button>
                                </div></td>
                            </tr>
                        </tbody>
                    </table>
                    <div v-if="sessionsTotal > 50" class="flex items-center justify-between mt-4">
                        <button class="cta cta-tertiary text-sm" :disabled="sessionsOffset === 0" @click="sessionsOffset = Math.max(0, sessionsOffset - 50); fetchSessions()">Prev</button>
                        <span class="text-sm text-gray-500">{{ sessionsOffset + 1 }}-{{ Math.min(sessionsOffset + 50, sessionsTotal) }} of {{ sessionsTotal }}</span>
                        <button class="cta cta-tertiary text-sm" :disabled="sessionsOffset + 50 >= sessionsTotal" @click="sessionsOffset += 50; fetchSessions()">Next</button>
                    </div>
                </div>
                <SkeletonRows v-else :rows="5" />
            </div>

            <!-- PASSKEYS -->
            <div v-if="tab === 'passkeys'" role="tabpanel">
                <div class="flex gap-2 mb-4">
                    <button class="cta cta-tertiary text-sm text-red-500" @click="bulkDeleteCredentials">Bulk Remove Selected</button>
                </div>
                <div v-if="credentials.length" class="overflow-x-auto">
                    <table class="table">
                        <thead><tr><th><input type="checkbox" @click="toggleAllCredentials($event)" /></th><th>User ID</th><th>Created</th><th></th></tr></thead>
                        <tbody>
                            <tr v-for="cr in credentials" :key="cr.id">
                                <td><input type="checkbox" v-model="(cr as any)._selected" /></td>
                                <td class="text-xs text-gray-500">{{ cr.user_id.slice(0,8) }}...</td>
                                <td>{{ formatDate(cr.created_at) }}</td>
                                <td><button class="cta cta-tertiary text-sm text-red-500" @click="deleteCredential(cr)">Remove</button></td>
                            </tr>
                        </tbody>
                    </table>
                    <div v-if="passkeysTotal > 50" class="flex items-center justify-between mt-4">
                        <button class="cta cta-tertiary text-sm" :disabled="passkeysOffset === 0" @click="passkeysOffset = Math.max(0, passkeysOffset - 50); fetchCredentials()">Prev</button>
                        <span class="text-sm text-gray-500">{{ passkeysOffset + 1 }}-{{ Math.min(passkeysOffset + 50, passkeysTotal) }} of {{ passkeysTotal }}</span>
                        <button class="cta cta-tertiary text-sm" :disabled="passkeysOffset + 50 >= passkeysTotal" @click="passkeysOffset += 50; fetchCredentials()">Next</button>
                    </div>
                </div>
                <SkeletonRows v-else :rows="5" />
            </div>

            <!-- INBOX -->
            <div v-if="tab === 'inbox'" role="tabpanel">
                <div class="flex gap-2 mb-4">
                    <input v-model="inboxSearch" placeholder="Search by from or subject..." @input="searchInboxDeb" class="flex-1" />
                    <button class="cta cta-tertiary text-sm text-red-500" @click="bulkDeleteInbox">Bulk Delete Selected</button>
                    <button class="cta cta-tertiary text-sm text-red-500" @click="purgeAllInbox">Purge All</button>
                </div>
                <div v-if="inboxMessages.length" class="overflow-x-auto">
                    <table class="table">
                        <thead><tr><th><input type="checkbox" @click="toggleAllInbox($event)" /></th><th>From</th><th>Subject</th><th>Alias</th><th>Size</th><th>Date</th><th></th></tr></thead>
                        <tbody>
                            <tr v-for="m in inboxMessages" :key="m.id">
                                <td><input type="checkbox" v-model="(m as any)._selected" /></td>
                                <td>{{ m.from_name || m.from }}</td>
                                <td class="max-w-xs truncate">{{ m.subject }}</td>
                                <td class="text-xs text-gray-500">{{ m.alias_id.slice(0,8) }}...</td>
                                <td>{{ Math.round(m.size / 1024) }}KB</td>
                                <td>{{ formatDate(m.created_at) }}</td>
                                <td><div class="flex gap-1">
                                    <button class="cta cta-tertiary text-sm" @click="viewInboxRaw(m)">View</button>
                                    <button class="cta cta-tertiary text-sm" @click="markInboxRead(m)">{{ m.read ? 'Unread' : 'Read' }}</button>
                                    <button class="cta cta-tertiary text-sm text-red-500" @click="deleteInboxMsg(m)">Delete</button>
                                </div></td>
                            </tr>
                        </tbody>
                    </table>
                    <div v-if="inboxTotal > 50" class="flex items-center justify-between mt-4">
                        <button class="cta cta-tertiary text-sm" :disabled="inboxOffset === 0" @click="inboxOffset = Math.max(0, inboxOffset - 50); fetchInboxMessages()">Prev</button>
                        <span class="text-sm text-gray-500">{{ inboxOffset + 1 }}-{{ Math.min(inboxOffset + 50, inboxTotal) }} of {{ inboxTotal }}</span>
                        <button class="cta cta-tertiary text-sm" :disabled="inboxOffset + 50 >= inboxTotal" @click="inboxOffset += 50; fetchInboxMessages()">Next</button>
                    </div>
                </div>
                <!-- Raw message modal -->
                <div v-if="viewingRaw !== null" class="fixed inset-0 bg-black/50 flex items-center justify-center z-50 p-4" @click.self="viewingRaw = null">
                    <div class="card-container max-w-3xl w-full max-h-[80vh] overflow-y-auto">
                        <div class="flex items-center justify-between mb-4">
                            <h3>Raw Message</h3>
                            <button class="cta cta-tertiary" @click="viewingRaw = null">Close</button>
                        </div>
                        <pre class="text-xs whitespace-pre-wrap break-all">{{ viewingRaw }}</pre>
                    </div>
                </div>
                <SkeletonRows v-else :rows="5" />
            </div>

            <!-- MESSAGES -->
            <div v-if="tab === 'messages'" role="tabpanel">
                <div class="flex gap-2 mb-4">
                    <input v-model="msgSearch" placeholder="Search by user/alias ID..." @input="searchMessagesDeb" class="flex-1" />
                    <select v-model="msgTypeFilter" @change="messagesOffset = 0; fetchMessages()" class="max-w-xs">
                        <option value="">All Types</option>
                        <option value="0">Forward</option>
                        <option value="1">Block</option>
                        <option value="2">Reply</option>
                        <option value="3">Send</option>
                        <option value="4">Bounce</option>
                        <option value="5">Inbox</option>
                    </select>
                    <button class="cta cta-tertiary text-sm text-red-500" @click="bulkDeleteMessages">Bulk Delete Selected</button>
                </div>
                <div v-if="messages.length" class="overflow-x-auto">
                    <table class="table">
                        <thead><tr><th><input type="checkbox" @click="toggleAllMessages($event)" /></th><th>Type</th><th>User ID</th><th>Alias ID</th><th>Date</th></tr></thead>
                        <tbody>
                            <tr v-for="m in messages" :key="m.id">
                                <td><input type="checkbox" v-model="(m as any)._selected" /></td>
                                <td><span class="badge">{{ ['Forward','Block','Reply','Send','Bounce','Inbox'][m.type] }}</span></td>
                                <td class="text-xs text-gray-500">{{ m.user_id?.slice(0,8) || '-' }}...</td>
                                <td class="text-xs text-gray-500">{{ m.alias_id?.slice(0,8) || '-' }}...</td>
                                <td>{{ formatDate(m.created_at) }}</td>
                            </tr>
                        </tbody>
                    </table>
                    <div v-if="messagesTotal > 50" class="flex items-center justify-between mt-4">
                        <button class="cta cta-tertiary text-sm" :disabled="messagesOffset === 0" @click="messagesOffset = Math.max(0, messagesOffset - 50); fetchMessages()">Prev</button>
                        <span class="text-sm text-gray-500">{{ messagesOffset + 1 }}-{{ Math.min(messagesOffset + 50, messagesTotal) }} of {{ messagesTotal }}</span>
                        <button class="cta cta-tertiary text-sm" :disabled="messagesOffset + 50 >= messagesTotal" @click="messagesOffset += 50; fetchMessages()">Next</button>
                    </div>
                </div>
                <SkeletonRows v-else :rows="5" />
            </div>

            <!-- SUBSCRIPTIONS -->
            <div v-if="tab === 'subs'" role="tabpanel">
                <div class="flex gap-2 mb-4">
                    <select v-model="subFilter" @change="fetchSubscriptions" class="max-w-xs">
                        <option value="">All Tiers</option>
                        <option value="self-hosted">Self-hosted</option>
                        <option value="self">Self</option>
                        <option value="pro">Pro</option>
                        <option value="free">Free</option>
                    </select>
                    <button class="cta cta-tertiary text-sm" @click="bulkExtendSubs">Bulk Extend</button>
                </div>
                <div v-if="subscriptions.length" class="overflow-x-auto">
                    <table class="table">
                        <thead><tr><th><input type="checkbox" @click="toggleAllSubs($event)" /></th><th>User ID</th><th>Tier</th><th>Type</th><th>Active</th><th>Active Until</th><th>Created</th><th></th></tr></thead>
                        <tbody>
                            <tr v-for="s in subscriptions" :key="s.id">
                                <td><input type="checkbox" v-model="(s as any)._selected" /></td>
                                <td class="text-xs text-gray-500">{{ s.user_id.slice(0,8) }}...</td>
                                <td><span class="badge">{{ s.tier || 'none' }}</span></td>
                                <td>{{ s.type }}</td>
                                <td><span :class="s.is_active ? 'badge badge-success' : 'badge badge-error'">{{ s.is_active ? 'Yes' : 'No' }}</span></td>
                                <td>{{ s.active_until ? formatDate(s.active_until) : 'N/A' }}</td>
                                <td>{{ formatDate(s.created_at) }}</td>
                                <td><div class="flex gap-1">
                                    <button class="cta cta-tertiary text-sm" @click="extendSub(s)">Extend</button>
                                    <button class="cta cta-tertiary text-sm text-red-500" @click="deleteSub(s)">Delete</button>
                                </div></td>
                            </tr>
                        </tbody>
                    </table>
                    <div v-if="subsTotal > 50" class="flex items-center justify-between mt-4">
                        <button class="cta cta-tertiary text-sm" :disabled="subsOffset === 0" @click="subsOffset = Math.max(0, subsOffset - 50); fetchSubscriptions()">Prev</button>
                        <span class="text-sm text-gray-500">{{ subsOffset + 1 }}-{{ Math.min(subsOffset + 50, subsTotal) }} of {{ subsTotal }}</span>
                        <button class="cta cta-tertiary text-sm" :disabled="subsOffset + 50 >= subsTotal" @click="subsOffset += 50; fetchSubscriptions()">Next</button>
                    </div>
                </div>
                <SkeletonRows v-else :rows="5" />
            </div>

            <!-- SYSTEM -->
            <div v-if="tab === 'system'" role="tabpanel">
                <div v-if="tableSizes" class="grid grid-cols-2 md:grid-cols-4 gap-4 mb-6">
                    <div v-for="(count, table) in tableSizes" :key="table" class="card-secondary text-center">
                        <p class="text-2xl font-bold">{{ count }}</p>
                        <p class="text-sm text-gray-500">{{ table }}</p>
                    </div>
                </div>
                <div class="card-secondary">
                    <h3 class="mb-3 font-bold">Recent Signups (7 days)</h3>
                    <div v-if="recentSignups.length" class="overflow-x-auto">
                        <table class="table">
                            <thead><tr><th>Email</th><th>Active</th><th>Admin</th><th>Joined</th></tr></thead>
                            <tbody>
                                <tr v-for="u in recentSignups" :key="u.id">
                                    <td>{{ u.email }}</td>
                                    <td><span :class="u.is_active ? 'badge badge-success' : 'badge badge-error'">{{ u.is_active ? 'Active' : 'Suspended' }}</span></td>
                                    <td>{{ u.is_admin ? 'Admin' : 'User' }}</td>
                                    <td>{{ formatDate(u.created_at) }}</td>
                                </tr>
                            </tbody>
                        </table>
                    </div>
                    <p v-else class="text-gray-500">No recent signups.</p>
                </div>
                <div v-if="configInfo" class="card-secondary mt-4">
                    <h3 class="mb-3 font-bold">System Configuration</h3>
                    <div class="grid grid-cols-2 gap-2 text-sm">
                        <div><span class="text-gray-500">FQDN:</span> {{ configInfo.fqdn || 'not set' }}</div>
                        <div><span class="text-gray-500">Port:</span> {{ configInfo.port }}</div>
                        <div><span class="text-gray-500">Domains:</span> {{ configInfo.domains }}</div>
                        <div><span class="text-gray-500">Token Expiration:</span> {{ configInfo.token_expiration }}</div>
                        <div><span class="text-gray-500">Admin Emails:</span> {{ (configInfo.admin_emails || []).join(', ') || 'none' }}</div>
                        <div><span class="text-gray-500">CORS Origin:</span> {{ configInfo.api_allow_origin }}</div>
                        <div><span class="text-gray-500">Preauth URL:</span> {{ configInfo.preauth_url_set ? 'configured' : 'not set' }}</div>
                        <div><span class="text-gray-500">Preauth PSK:</span> {{ configInfo.preauth_psk_set ? 'configured' : 'not set' }}</div>
                        <div><span class="text-gray-500">Signup Webhook:</span> {{ configInfo.signup_webhook_set ? 'configured' : 'not set' }}</div>
                        <div><span class="text-gray-500">SMTP:</span> {{ configInfo.smtp_configured ? 'configured' : 'not set' }}</div>
                        <div><span class="text-gray-500">Oxapay:</span> {{ configInfo.oxapay_configured ? 'configured' : 'not set' }}</div>
                    </div>
                </div>
            </div>

            <!-- LOGS -->
            <div v-if="tab === 'logs'" role="tabpanel">
                <div class="flex gap-2 mb-4">
                    <select v-model="logFilter" @change="logsOffset = 0; fetchLogsFiltered()" class="max-w-xs">
                        <option value="">All Types</option>
                        <option value="bounce">Bounce</option>
                        <option value="disabled_alias">Disabled Alias</option>
                        <option value="disabled_domain">Disabled Domain</option>
                        <option value="unauthorised_send">Unauthorised Send</option>
                        <option value="inactive_subscription">Inactive Subscription</option>
                    </select>
                    <input v-model="logSearch" placeholder="Search from/destination/message..." @input="searchLogsDeb" class="flex-1" />
                    <input v-model="logFrom" type="date" @change="logsOffset = 0; fetchLogsFiltered()" class="max-w-[150px]" title="From date" />
                    <input v-model="logTo" type="date" @change="logsOffset = 0; fetchLogsFiltered()" class="max-w-[150px]" title="To date" />
                    <button class="cta cta-tertiary text-sm text-red-500" @click="purgeLogs">Purge Old Logs</button>
                </div>
                <div v-if="logs.length" class="overflow-x-auto">
                    <table class="table">
                        <thead><tr><th>Type</th><th>From</th><th>Destination</th><th>Status</th><th>Message</th><th>Date</th></tr></thead>
                        <tbody>
                            <tr v-for="log in logs" :key="log.id">
                                <td><span class="badge">{{ log.log_type }}</span></td>
                                <td>{{ log.from }}</td>
                                <td>{{ log.destination }}</td>
                                <td>{{ log.status }}</td>
                                <td class="max-w-xs truncate">{{ log.message }}</td>
                                <td>{{ formatDate(log.created_at) }}</td>
                                <td><div class="flex gap-1">
                                    <button class="cta cta-tertiary text-sm" @click="viewLog(log)">View</button>
                                    <button class="cta cta-tertiary text-sm text-red-500" @click="deleteLog(log)">Delete</button>
                                </div></td>
                            </tr>
                        </tbody>
                    </table>
                    <div v-if="logsTotal > 100" class="flex items-center justify-between mt-4">
                        <button class="cta cta-tertiary text-sm" :disabled="logsOffset === 0" @click="logsOffset = Math.max(0, logsOffset - 100); (logSearch ? searchLogsDeb() : fetchLogsFiltered())">Prev</button>
                        <span class="text-sm text-gray-500">{{ logsOffset + 1 }}-{{ Math.min(logsOffset + 100, logsTotal) }} of {{ logsTotal }}</span>
                        <button class="cta cta-tertiary text-sm" :disabled="logsOffset + 100 >= logsTotal" @click="logsOffset += 100; (logSearch ? searchLogsDeb() : fetchLogsFiltered())">Next</button>
                    </div>
                </div>
                <SkeletonRows v-else :rows="5" />
                <!-- Log detail modal -->
                <div v-if="logDetail" class="fixed inset-0 bg-black/50 flex items-center justify-center z-50 p-4" @click.self="logDetail = null">
                    <div class="card-container max-w-2xl w-full">
                        <div class="flex items-center justify-between mb-4">
                            <h3>Log Detail</h3>
                            <button class="cta cta-tertiary" @click="logDetail = null">Close</button>
                        </div>
                        <div class="grid grid-cols-1 gap-2 text-sm">
                            <div><span class="text-gray-500">ID:</span> {{ logDetail.id }}</div>
                            <div><span class="text-gray-500">Type:</span> <span class="badge">{{ logDetail.log_type }}</span></div>
                            <div><span class="text-gray-500">From:</span> {{ logDetail.from }}</div>
                            <div><span class="text-gray-500">To:</span> {{ logDetail.destination }}</div>
                            <div><span class="text-gray-500">Status:</span> {{ logDetail.status }}</div>
                            <div><span class="text-gray-500">Remote MTA:</span> {{ logDetail.remote_mta || 'N/A' }}</div>
                            <div><span class="text-gray-500">Date:</span> {{ formatDate(logDetail.created_at) }}</div>
                            <div><span class="text-gray-500">Attempted:</span> {{ formatDate(logDetail.attempted_at) }}</div>
                            <div class="mt-2"><span class="text-gray-500">Message:</span><pre class="mt-1 text-xs bg-gray-50 p-2 rounded whitespace-pre-wrap">{{ logDetail.message }}</pre></div>
                        </div>
                    </div>
                </div>
            </div>

            <!-- AUDIT -->
            <div v-if="tab === 'audit'" role="tabpanel">
                <div v-if="auditEntries.length" class="overflow-x-auto">
                    <table class="table">
                        <thead><tr><th>Admin</th><th>Action</th><th>Target</th><th>Details</th><th>Date</th></tr></thead>
                        <tbody>
                            <tr v-for="a in auditEntries" :key="a.id">
                                <td>{{ a.admin_email }}</td>
                                <td><span class="badge">{{ a.action }}</span></td>
                                <td class="max-w-xs truncate">{{ a.target }}</td>
                                <td class="max-w-xs truncate">{{ a.details }}</td>
                                <td>{{ formatDate(a.created_at) }}</td>
                            </tr>
                        </tbody>
                    </table>
                    <div v-if="auditTotal > 50" class="flex items-center justify-between mt-4">
                        <button class="cta cta-tertiary text-sm" :disabled="auditOffset === 0" @click="auditOffset = Math.max(0, auditOffset - 50); fetchAuditLog()">Prev</button>
                        <span class="text-sm text-gray-500">{{ auditOffset + 1 }}-{{ Math.min(auditOffset + 50, auditTotal) }} of {{ auditTotal }}</span>
                        <button class="cta cta-tertiary text-sm" :disabled="auditOffset + 50 >= auditTotal" @click="auditOffset += 50; fetchAuditLog()">Next</button>
                    </div>
                </div>
                <SkeletonRows v-else :rows="5" />
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref, onMounted, watch } from 'vue'
import { planApi, adminApi, type Plan, type AdminUser, type AdminLog, type AdminAlias, type AdminDomain, type AdminRecipient, type AdminAccessKey, type AdminSession, type AdminCredential, type AdminInboxMessage, type AdminSubscription, type AdminAudit, type SystemStats } from '../api/plan'
import SkeletonRows from '../components/SkeletonRows.vue'

const tab = ref('stats')
const stats = ref<SystemStats | null>(null)
const users = ref<AdminUser[]>([])
const aliases = ref<AdminAlias[]>([])
const domains = ref<AdminDomain[]>([])
const recipients = ref<AdminRecipient[]>([])
const plans = ref<Plan[]>([])
const logs = ref<AdminLog[]>([])
const accessKeys = ref<AdminAccessKey[]>([])
const sessions = ref<AdminSession[]>([])
const credentials = ref<AdminCredential[]>([])
const inboxMessages = ref<AdminInboxMessage[]>([])
const subscriptions = ref<AdminSubscription[]>([])
const subFilter = ref('')
const tableSizes = ref<Record<string, number> | null>(null)
const recentSignups = ref<AdminUser[]>([])
const configInfo = ref<Record<string, any> | null>(null)
const loadingPlans = ref(true)
const deleting = ref('')
const saving = ref(false)
const showForm = ref(false)
const editing = ref(false)
const formError = ref('')
const userSearch = ref('')
const usersTotal = ref(0)
const usersOffset = ref(0)
const aliasSearch = ref('')
const recipientSearch = ref('')
const logFilter = ref('')
const userDetail = ref<any>(null)
const selectedPlan = ref('')

let searchTimer: ReturnType<typeof setTimeout>

const emptyForm = (): Partial<Plan> => ({ name: '', display_name: '', price_cents: 0, currency: 'usd', interval: 'monthly', max_recipients: 5, max_credentials: 10, max_daily_aliases: 50, max_daily_send_reply: 50, max_sessions: 5, sort_order: 0 })
const form = ref<Partial<Plan>>(emptyForm())
const formatDate = (d: string) => new Date(d).toLocaleDateString()

const fetchStats = async () => { try { stats.value = await adminApi.stats() } catch { /* */ } }
const fetchUsers = async () => { try { const r = await adminApi.usersPaginated(50, usersOffset.value, userSearch.value || undefined); users.value = r.users; usersTotal.value = r.total } catch { /* */ } }
const nextUsers = () => { usersOffset.value += 50; fetchUsers() }
const prevUsers = () => { usersOffset.value = Math.max(0, usersOffset.value - 50); fetchUsers() }
const fetchAliases = async () => { try { const r = await adminApi.aliases(aliasSearch.value || undefined, aliasesOffset.value); aliases.value = r.aliases; aliasesTotal.value = r.total } catch { /* */ } }
const fetchDomains = async () => { try { domains.value = await adminApi.domains() } catch { /* */ } }
const fetchRecipients = async () => { try { const r = await adminApi.recipients(recipientSearch.value || undefined, recipientsOffset.value); recipients.value = r.recipients; recipientsTotal.value = r.total } catch { /* */ } }
const fetchPlans = async () => { loadingPlans.value = true; try { plans.value = await planApi.listAll() } catch { /* */ } finally { loadingPlans.value = false } }
const fetchLogs = async () => { try { logs.value = await adminApi.logs() } catch { /* */ } }
const fetchLogsFiltered = async () => { try { if (logFrom.value || logTo.value) { const r = await adminApi.logsDateRange(logFrom.value, logTo.value, logFilter.value || undefined, logsOffset.value); logs.value = r.logs; logsTotal.value = r.total } else { const r = await adminApi.logsFiltered(logFilter.value || undefined, logsOffset.value); logs.value = r.logs; logsTotal.value = r.total } } catch { /* */ } }
const fetchAccessKeys = async () => { try { const r = await adminApi.accessKeys(keysOffset.value); accessKeys.value = r.keys; keysTotal.value = r.total } catch { /* */ } }
const fetchSessions = async () => { try { const r = await adminApi.sessions(sessionsOffset.value); sessions.value = r.sessions; sessionsTotal.value = r.total } catch { /* */ } }
const fetchCredentials = async () => { try { const r = await adminApi.credentials(passkeysOffset.value); credentials.value = r.credentials; passkeysTotal.value = r.total } catch { /* */ } }
const fetchSubscriptions = async () => { try { const r = await adminApi.subscriptions(subFilter.value || undefined, subsOffset.value); subscriptions.value = r.subscriptions; subsTotal.value = r.total } catch { /* */ } }
const fetchMessages = async () => { try { const r = await adminApi.messages(msgTypeFilter.value || undefined, messagesOffset.value); messages.value = r.messages; messagesTotal.value = r.total } catch { /* */ } }
const fetchInboxMessages = async () => { try { const r = await adminApi.inboxMessages(inboxOffset.value); inboxMessages.value = r.messages; inboxTotal.value = r.total } catch { /* */ } }
const fetchTableSizes = async () => { try { tableSizes.value = await adminApi.tableSizes() } catch { /* */ } }
const fetchRecentSignups = async () => { try { const r = await adminApi.recentSignups(7); recentSignups.value = r.users } catch { /* */ } }
const fetchConfig = async () => { try { configInfo.value = await adminApi.getConfig() } catch { /* */ } }
const searchMessagesDeb = () => { clearTimeout(msgSearchTimer); msgSearchTimer = setTimeout(async () => { try { const r = await adminApi.searchMessages(msgSearch.value || undefined, msgTypeFilter.value || undefined); messages.value = r.messages } catch { /* */ } }, 300) }
const toggleRecipientPGP = async (r: AdminRecipient) => { if (!confirm('Disable PGP for this recipient?')) return; try { await adminApi.toggleRecipientPGP(r.id, false); r.pgp_enabled = false } catch { /* */ } }
const removePGPKey = async (r: AdminRecipient) => { if (!confirm('Permanently remove PGP key?')) return; try { await adminApi.removeRecipientPGPKey(r.id); r.pgp_enabled = false } catch { /* */ } }
const editAlias = async (a: AdminAlias) => {
    const desc = prompt('Description:', a.description)
    if (desc === null) return
    const rcps = prompt('Recipients (comma-separated):', a.recipients)
    if (rcps === null) return
    const fn = prompt('From name:', a.from_name)
    if (fn === null) return
    try { await adminApi.updateAlias(a.id, { description: desc, recipients: rcps, from_name: fn }); a.description = desc; a.recipients = rcps; a.from_name = fn } catch { /* */ } }
const editDomain = async (d: AdminDomain) => {
    const desc = prompt('Description:', d.description)
    if (desc === null) return
    const rcp = prompt('Recipient:', d.recipient || '')
    if (rcp === null) return
    const fn = prompt('From name:', d.from_name || '')
    if (fn === null) return
    try { await adminApi.updateDomain(d.id, { description: desc, recipient: rcp, from_name: fn }); d.description = desc; d.recipient = rcp; d.from_name = fn } catch { /* */ } }
const markInboxRead = async (m: AdminInboxMessage) => { try { await adminApi.markInboxRead(m.id, !m.read); m.read = !m.read } catch { /* */ } }
const createRecipientForUser = async (userId: string) => {
    const email = prompt('Enter recipient email:')
    if (!email) return
    try { await adminApi.createRecipient(userId, email); alert('Recipient created'); viewUser(userId) } catch { /* */ } }
const createDomainForUser = async (userId: string) => {
    const name = prompt('Enter domain name:')
    if (!name) return
    try { await adminApi.createDomain(userId, name); alert('Domain created'); viewUser(userId) } catch { /* */ } }
const createAliasForUser = async (userId: string) => {
    const name = prompt('Enter alias name:')
    if (!name) return
    try { await adminApi.createAlias(userId, name, true); alert('Alias created'); viewUser(userId) } catch { /* */ } }
const editRecipient = async (r: AdminRecipient) => {
    const email = prompt('Enter new email:', r.email)
    if (!email || email === r.email) return
    try { await adminApi.updateRecipient(r.id, email); r.email = email } catch { /* */ } }
const fetchAuditLog = async () => { try { const r = await adminApi.auditLog(auditOffset.value); auditEntries.value = r.entries; auditTotal.value = r.total } catch { /* */ } }
const viewSessionData = async (s: AdminSession) => { try { const d = await adminApi.sessionData(s.id); alert(JSON.stringify(d, null, 2)) } catch { alert('Unable to load session data') } }
const viewLog = (log: AdminLog) => { logDetail.value = log }
const deleteLog = async (log: AdminLog) => { if (!confirm('Delete this log entry?')) return; try { await adminApi.deleteLog(log.id); logs.value = logs.value.filter(x => x.id !== log.id) } catch { /* */ } }
const setKeyExpiry = async (k: AdminAccessKey) => {
    const date = prompt('Expiry date (YYYY-MM-DD), empty to clear:', k.expires_at ? k.expires_at.slice(0,10) : '')
    if (date === null) return
    if (date && !/^\d{4}-\d{2}-\d{2}$/.test(date)) { alert('Use YYYY-MM-DD format'); return }
    try { await adminApi.setAccessKeyExpiry(k.id, date); k.expires_at = date || null; alert(date ? `Expires ${date}` : 'Expiry cleared') } catch { /* */ } }
const toggleAllInbox = (e: Event) => { const checked = (e.target as HTMLInputElement).checked; inboxMessages.value.forEach(m => { (m as any)._selected = checked }) }
const toggleAllKeys = (e: Event) => { const checked = (e.target as HTMLInputElement).checked; accessKeys.value.forEach(k => { (k as any)._selected = checked }) }
const toggleAllCredentials = (e: Event) => { const checked = (e.target as HTMLInputElement).checked; credentials.value.forEach(c => { (c as any)._selected = checked }) }
const toggleAllSubs = (e: Event) => { const checked = (e.target as HTMLInputElement).checked; subscriptions.value.forEach(s => { (s as any)._selected = checked }) }
const toggleAllMessages = (e: Event) => { const checked = (e.target as HTMLInputElement).checked; messages.value.forEach(m => { (m as any)._selected = checked }) }
const bulkDeleteMessages = async () => {
    const selected = messages.value.filter(m => (m as any)._selected)
    if (!selected.length) { alert('Select messages first'); return }
    if (!confirm(`Delete ${selected.length} messages?`)) return
    try { await adminApi.bulkDeleteMessages(selected.map(m => m.id)); messages.value = messages.value.filter(m => !selected.includes(m)) } catch { /* */ } }
const bulkDeleteKeys = async () => {
    const selected = accessKeys.value.filter(k => (k as any)._selected)
    if (!selected.length) { alert('Select keys first'); return }
    if (!confirm(`Revoke ${selected.length} keys?`)) return
    try { await adminApi.bulkDeleteAccessKeys(selected.map(k => k.id)); accessKeys.value = accessKeys.value.filter(k => !selected.includes(k)) } catch { /* */ } }
const bulkDeleteCredentials = async () => {
    const selected = credentials.value.filter(c => (c as any)._selected)
    if (!selected.length) { alert('Select passkeys first'); return }
    if (!confirm(`Remove ${selected.length} passkeys?`)) return
    try { await adminApi.bulkDeleteCredentials(selected.map(c => c.id)); credentials.value = credentials.value.filter(c => !selected.includes(c)) } catch { /* */ } }
const bulkExtendSubs = async () => {
    const selected = subscriptions.value.filter(s => (s as any)._selected)
    if (!selected.length) { alert('Select subscriptions first'); return }
    const days = parseInt(prompt('Extend by how many days?', '30') || '0')
    if (!days) return
    if (!confirm(`Extend ${selected.length} subscriptions by ${days} days?`)) return
    try { await adminApi.bulkExtendSubscriptions(selected.map(s => s.id), days); subscriptions.value = selected; fetchSubscriptions() } catch { /* */ } }
const exportUsersEnriched = () => { window.open(`${import.meta.env.VITE_API_URL}/v1/admin/export/users-enriched`, '_blank') }
const bulkDeleteInbox = async () => {
    const selected = inboxMessages.value.filter(m => (m as any)._selected)
    if (!selected.length) { alert('Select messages first'); return }
    if (!confirm(`Delete ${selected.length} messages?`)) return
    try { await adminApi.bulkDeleteInbox(selected.map(m => m.id)); inboxMessages.value = inboxMessages.value.filter(m => !selected.includes(m)) } catch { /* */ } }
const extendSub = async (s: AdminSubscription) => {
    const days = parseInt(prompt('Extend by how many days?', '30') || '0')
    if (!days) return
    try { await adminApi.extendSubscription(s.id, days); alert(`Extended by ${days} days`); fetchSubscriptions() } catch { /* */ } }
const transferAlias = async (a: AdminAlias) => {
    const newUserId = prompt('Enter new owner user ID:')
    if (!newUserId) return
    if (!confirm(`Transfer alias ${a.name} to user ${newUserId}?`)) return
    try { await adminApi.transferAlias(a.id, newUserId); a.user_id = newUserId; alert('Alias transferred') } catch { /* */ } }
const transferDomain = async (d: AdminDomain) => {
    const newUserId = prompt('Enter new owner user ID:')
    if (!newUserId) return
    if (!confirm(`Transfer domain ${d.name} to user ${newUserId}?`)) return
    try { await adminApi.transferDomain(d.id, newUserId); d.user_id = newUserId; alert('Domain transferred') } catch { /* */ } }
const createKeyForUser = async (userId: string) => {
    const name = prompt('Enter key name:')
    if (!name) return
    try { const r = await adminApi.createAccessKey(userId, name); prompt('Access key (copy now, shown once):', r.key) } catch { /* */ } }
const purgeLogs = async () => {
    const days = parseInt(prompt('Purge logs older than how many days?', '30') || '0')
    if (!days) return
    if (!confirm(`Delete ALL logs older than ${days} days?`)) return
    try { const r = await adminApi.purgeLogs(days); alert(r.message); fetchLogs() } catch { /* */ } }
const purgeAllInbox = async () => {
    if (!confirm('Delete ALL inbox messages for ALL users? This is irreversible.')) return
    if (!confirm('Are you absolutely sure? Every temp mail message will be permanently deleted.')) return
    try { const r = await adminApi.purgeAllInbox(); alert(r.message); inboxMessages.value = [] } catch { /* */ } }
const createUser = async () => {
    const email = prompt('Enter email address:')
    if (!email || !/^[^@]+@[^@]+\.[^@]+$/.test(email)) { alert('Invalid email'); return }
    const password = prompt('Enter password (min 12 chars):')
    if (!password || password.length < 12) { alert('Password must be 12+ characters'); return }
    try { await adminApi.createUser(email, password); alert('User created'); usersOffset.value = 0; fetchUsers() } catch { /* */ } }
const viewInboxRaw = async (m: AdminInboxMessage) => { try { viewingRaw.value = await adminApi.inboxRaw(m.id) } catch { alert('Unable to load message') } }
const setAliasExpiry = async (a: AdminAlias) => {
    const date = prompt('Expiry date (YYYY-MM-DD), empty to clear:', '')
    if (date === null) return
    if (date && !/^\d{4}-\d{2}-\d{2}$/.test(date)) { alert('Use YYYY-MM-DD format'); return }
    try { await adminApi.setAliasExpiry(a.id, date); alert(date ? `Expires ${date}` : 'Expiry cleared') } catch { /* */ } }
let logSearchTimer: any
const searchLogsDeb = () => { clearTimeout(logSearchTimer); logSearchTimer = setTimeout(async () => { logsOffset.value = 0; try { const r = await adminApi.searchLogs(logSearch.value, logFilter.value || undefined, logsOffset.value); logs.value = r.logs; logsTotal.value = r.total } catch { /* */ } }, 300) }
let domainSearchTimer: any
const searchDomainsDeb = () => { clearTimeout(domainSearchTimer); domainSearchTimer = setTimeout(async () => { try { const r = await adminApi.searchDomains(domainSearch.value); domains.value = r.domains } catch { /* */ } }, 300) }

const searchUsersDeb = () => { clearTimeout(searchTimer); searchTimer = setTimeout(() => { usersOffset.value = 0; fetchUsers() }, 300) }
const searchAliasesDeb = () => { clearTimeout(searchTimer); searchTimer = setTimeout(() => { aliasesOffset.value = 0; fetchAliases() }, 300) }
const searchRecipientsDeb = () => { clearTimeout(searchTimer); searchTimer = setTimeout(() => { recipientsOffset.value = 0; fetchRecipients() }, 300) }

const viewUser = async (id: string) => { try { userDetail.value = await adminApi.userDetail(id); selectedPlan.value = userDetail.value.subscription?.plan_id || ''; userStats.value = await adminApi.userStats(id); userSettings.value = await adminApi.getSettings(id) } catch { /* */ } }
const assignPlan = async (userId: string) => { if (!selectedPlan.value) return; try { await adminApi.assignPlan(userId, selectedPlan.value); viewUser(userId) } catch { /* */ } }

const toggleUser = async (u: AdminUser) => { if (u.email === localStorage.getItem('email')) { alert('Cannot suspend your own account'); return } try { await adminApi.updateUser({ id: u.id, is_active: !u.is_active }); u.is_active = !u.is_active } catch { /* */ } }
const toggleAdmin = async (u: AdminUser) => { if (u.email === localStorage.getItem('email')) { alert('Cannot modify your own admin role'); return } try { await adminApi.updateUser({ id: u.id, is_admin: !u.is_admin }); u.is_admin = !u.is_admin } catch { /* */ } }
const deleteUser = async (u: AdminUser) => { if (u.email === localStorage.getItem('email')) { alert('Cannot delete your own account'); return } if (!confirm(`Delete ${u.email}? Removes all data.`)) return; deleting.value = u.id; try { await adminApi.deleteUser(u.id); users.value = users.value.filter(x => x.id !== u.id) } catch { /* */ } finally { deleting.value = '' } }

const toggleAlias = async (a: AdminAlias) => { try { await adminApi.toggleAlias(a.id, !a.enabled); a.enabled = !a.enabled } catch { /* */ } }
const deleteAlias = async (a: AdminAlias) => { if (!confirm(`Delete alias ${a.name}?`)) return; try { await adminApi.deleteAlias(a.id); aliases.value = aliases.value.filter(x => x.id !== a.id) } catch { /* */ } }

const toggleDomain = async (d: AdminDomain) => { try { await adminApi.toggleDomain(d.id, !d.enabled); d.enabled = !d.enabled } catch { /* */ } }
const deleteDomain = async (d: AdminDomain) => { if (!confirm(`Delete domain ${d.name}?`)) return; try { await adminApi.deleteDomain(d.id); domains.value = domains.value.filter(x => x.id !== d.id) } catch { /* */ } }

const deleteRecipient = async (r: AdminRecipient) => { if (!confirm(`Delete recipient ${r.email}?`)) return; try { await adminApi.deleteRecipient(r.id); recipients.value = recipients.value.filter(x => x.id !== r.id) } catch { /* */ } }

const deleteAccessKey = async (k: AdminAccessKey) => { if (!confirm(`Revoke key "${k.name}"?`)) return; try { await adminApi.deleteAccessKey(k.id); accessKeys.value = accessKeys.value.filter(x => x.id !== k.id) } catch { /* */ } }
const deleteSession = async (s: AdminSession) => { if (!confirm('Terminate this session?')) return; try { await adminApi.deleteSession(s.id); sessions.value = sessions.value.filter(x => x.id !== s.id) } catch { /* */ } }
const deleteCredential = async (cr: AdminCredential) => { if (!confirm('Remove this passkey?')) return; try { await adminApi.deleteCredential(cr.id); credentials.value = credentials.value.filter(x => x.id !== cr.id) } catch { /* */ } }
const forceLogout = async (userId: string) => { if (!confirm('Terminate ALL sessions for this user?')) return; try { await adminApi.forceLogout(userId); alert('All sessions terminated') } catch { /* */ } }
const overrideSub = async (userId: string) => {
    const tier = prompt('Enter tier (e.g. pro, free, self-hosted):')
    if (!tier) return
    const activeUntil = prompt('Enter active_until date (YYYY-MM-DD) or leave empty:') || ''
    try { await adminApi.updateSubscription({ user_id: userId, tier, is_active: true, active_until: activeUntil }); viewUser(userId) } catch { /* */ } }
const bulkSuspend = async () => {
    const selected = users.value.filter(u => (u as any)._selected)
    if (!selected.length) { alert('Select users first'); return }
    if (!confirm(`Suspend ${selected.length} users?`)) return
    try { await adminApi.bulkUpdateUsers(selected.map(u => u.id), false); selected.forEach(u => u.is_active = false) } catch { /* */ } }
const bulkActivate = async () => {
    const selected = users.value.filter(u => (u as any)._selected)
    if (!selected.length) { alert('Select users first'); return }
    if (!confirm(`Activate ${selected.length} users?`)) return
    try { await adminApi.bulkUpdateUsers(selected.map(u => u.id), true); selected.forEach(u => u.is_active = true) } catch { /* */ } }
const toggleAllUsers = (e: Event) => {
    const checked = (e.target as HTMLInputElement).checked
    users.value.forEach(u => { (u as any)._selected = checked })
}
const deleteInboxMsg = async (m: AdminInboxMessage) => { if (!confirm(`Delete message "${m.subject}"?`)) return; try { await adminApi.deleteInboxMessage(m.id); inboxMessages.value = inboxMessages.value.filter(x => x.id !== m.id) } catch { /* */ } }
const disableTotp = async (userId: string) => { if (!confirm('Disable 2FA for this user?')) return; try { await adminApi.disableTotp(userId); alert('2FA disabled') } catch { /* */ } }
const resetPassword = async (userId: string) => {
    const pw = prompt('Enter new password (min 12 chars):')
    if (!pw || pw.length < 12) { alert('Password must be 12+ characters'); return }
    try { await adminApi.resetPassword(userId, pw); alert('Password reset') } catch { /* */ } }
const purgeInbox = async (userId: string) => { if (!confirm('Delete ALL inbox messages for this user?')) return; try { await adminApi.purgeInbox(userId); alert('Inbox purged') } catch { /* */ } }
const exportUsers = () => { window.open(`${import.meta.env.VITE_API_URL}/admin/export/users`, '_blank') }
const exportAliases = () => { window.open(`${import.meta.env.VITE_API_URL}/admin/export/aliases`, '_blank') }
const deleteSub = async (s: AdminSubscription) => { if (!confirm('Delete this subscription?')) return; try { await adminApi.deleteSubscription(s.id); subscriptions.value = subscriptions.value.filter(x => x.id !== s.id) } catch { /* */ } }
const bulkDeleteAliases = async () => { const selected = aliases.value.filter(a => (a as any)._selected); if (!selected.length) { alert('Select aliases first'); return } if (!confirm(`Delete ${selected.length} aliases?`)) return; try { await adminApi.bulkDeleteAliases(selected.map(a => a.id)); aliases.value = aliases.value.filter(a => !selected.includes(a)) } catch { /* */ } }
const bulkDeleteDomains = async () => { const selected = domains.value.filter(d => (d as any)._selected); if (!selected.length) { alert('Select domains first'); return } if (!confirm(`Delete ${selected.length} domains?`)) return; try { await adminApi.bulkDeleteDomains(selected.map(d => d.id)); domains.value = domains.value.filter(d => !selected.includes(d)) } catch { /* */ } }
const bulkDeleteRecipients = async () => { const selected = recipients.value.filter(r => (r as any)._selected); if (!selected.length) { alert('Select recipients first'); return } if (!confirm(`Delete ${selected.length} recipients?`)) return; try { await adminApi.bulkDeleteRecipients(selected.map(r => r.id)); recipients.value = recipients.value.filter(r => !selected.includes(r)) } catch { /* */ } }
const toggleAllAliases = (e: Event) => { const checked = (e.target as HTMLInputElement).checked; aliases.value.forEach(a => { (a as any)._selected = checked }) }
const toggleAllDomains = (e: Event) => { const checked = (e.target as HTMLInputElement).checked; domains.value.forEach(d => { (d as any)._selected = checked }) }
const toggleAllRecipients = (e: Event) => { const checked = (e.target as HTMLInputElement).checked; recipients.value.forEach(r => { (r as any)._selected = checked }) }
const toggleRecipient = async (r: AdminRecipient) => { try { await adminApi.toggleRecipient(r.id, !r.is_active); r.is_active = !r.is_active } catch { /* */ } }
const exportRecipients = () => { window.open(`${import.meta.env.VITE_API_URL}/v1/admin/export/recipients`, '_blank') }
const exportSubscriptions = () => { window.open(`${import.meta.env.VITE_API_URL}/v1/admin/export/subscriptions`, '_blank') }
const exportDomains = () => { window.open(`${import.meta.env.VITE_API_URL}/v1/admin/export/domains`, '_blank') }
const exportLogs = () => { window.open(`${import.meta.env.VITE_API_URL}/v1/admin/export/logs`, '_blank') }
const exportInbox = () => { window.open(`${import.meta.env.VITE_API_URL}/v1/admin/export/inbox`, '_blank') }
const exportMessages = () => { window.open(`${import.meta.env.VITE_API_URL}/v1/admin/export/messages`, '_blank') }
const bulkDeleteUsers = async () => { const selected = users.value.filter(u => (u as any)._selected); if (!selected.length) { alert('Select users first'); return } if (!confirm(`DELETE ${selected.length} users and ALL their data? This is irreversible.`)) return; try { await adminApi.bulkDeleteUsers(selected.map(u => u.id)); users.value = users.value.filter(u => !selected.includes(u)) } catch { /* */ } }
const changeEmail = async (userId: string) => { const email = prompt('Enter new email address:'); if (!email || !/^[^@]+@[^@]+\.[^@]+$/.test(email)) { alert('Invalid email'); return } try { await adminApi.changeEmail(userId, email); alert('Email changed'); viewUser(userId) } catch { /* */ } }
const verifyDomain = async (d: AdminDomain) => { try { await adminApi.verifyDomain(d.id, !d.mx_verified_at); d.mx_verified_at = d.mx_verified_at ? null : new Date().toISOString() as any } catch { /* */ } }
const impersonate = async (userId: string) => { if (!confirm('Login as this user? You will get a 24h session token.')) return; try { const r = await adminApi.impersonate(userId); document.cookie = `authn=${r.token}; path=/; secure; max-age=86400`; window.open('/account/aliases', '_blank') } catch { /* */ } }
const inboxSearch = ref('')
const logSearch = ref('')
const domainSearch = ref('')
const msgTypeFilter = ref('')
const msgSearch = ref('')
const userStats = ref<any>(null)
const userSettings = ref<any>(null)
const aliasesTotal = ref(0)
const aliasesOffset = ref(0)
const recipientsTotal = ref(0)
const recipientsOffset = ref(0)
const inboxTotal = ref(0)
const inboxOffset = ref(0)
const viewingRaw = ref<string | null>(null)
let msgSearchTimer: any
const keysOffset = ref(0)
const sessionsOffset = ref(0)
const passkeysOffset = ref(0)
const subsOffset = ref(0)
const messagesOffset = ref(0)
const logsOffset = ref(0)
const keysTotal = ref(0)
const sessionsTotal = ref(0)
const passkeysTotal = ref(0)
const subsTotal = ref(0)
const messagesTotal = ref(0)
const logsTotal = ref(0)
const logDetail = ref<AdminLog | null>(null)
const messages = ref<any[]>([])
const auditEntries = ref<AdminAudit[]>([])
const auditTotal = ref(0)
const auditOffset = ref(0)
const logFrom = ref('')
const logTo = ref('')
let inboxSearchTimer: any
const searchInboxDeb = () => { clearTimeout(inboxSearchTimer); inboxSearchTimer = setTimeout(async () => { try { const r = await adminApi.searchInbox(inboxSearch.value); inboxMessages.value = r.messages } catch { /* */ } }, 300) }

const savePlan = async () => { saving.value = true; formError.value = ''; try { if (editing.value && form.value.id) { await planApi.update(form.value.id, form.value) } else { await planApi.create(form.value) }; resetForm(); await fetchPlans() } catch (err: any) { formError.value = err?.message || 'Failed' } finally { saving.value = false } }
const editPlan = (plan: Plan) => { editing.value = true; showForm.value = true; form.value = { ...plan } }
const deletePlan = async (plan: Plan) => { if (!confirm(`Deactivate "${plan.display_name}"?`)) return; deleting.value = plan.id; try { await planApi.delete(plan.id); await fetchPlans() } catch { /* */ } finally { deleting.value = '' } }
const resetForm = () => { showForm.value = false; editing.value = false; form.value = emptyForm(); formError.value = '' }

watch(tab, (t) => {
    if (t === 'aliases' && !aliases.value.length) fetchAliases()
    if (t === 'domains' && !domains.value.length) fetchDomains()
    if (t === 'recipients' && !recipients.value.length) fetchRecipients()
    if (t === 'keys' && !accessKeys.value.length) fetchAccessKeys()
    if (t === 'sessions' && !sessions.value.length) fetchSessions()
    if (t === 'passkeys' && !credentials.value.length) fetchCredentials()
    if (t === 'inbox' && !inboxMessages.value.length) fetchInboxMessages()
    if (t === 'messages') fetchMessages()
    if (t === 'subs') fetchSubscriptions()
    if (t === 'system') { fetchTableSizes(); fetchRecentSignups(); fetchConfig() }
    if (t === 'logs') fetchLogsFiltered()
    if (t === 'audit' && !auditEntries.value.length) fetchAuditLog()
})

const uploadPGP = async (r: AdminRecipient) => {
    const key = prompt('Paste PGP public key:')
    if (key === null) return
    try { await adminApi.setRecipientPGP(r.id, key); r.pgp_enabled = true; alert('PGP key set') } catch { /* */ } }
const viewDomainDNS = async (d: AdminDomain) => {
    try { const dns = await adminApi.domainDNS(d.id); alert(`Domain: ${dns.domain}\nMX: ${(dns.mx_hosts||[]).join(', ') || 'N/A'}\nDKIM: ${(dns.dkim_selectors||[]).join(', ') || 'N/A'}`) } catch { alert('Unable to load DNS') } }
const editUserNotes = async (u: AdminUser) => {
    const notes = prompt('Notes:', u.notes || '')
    if (notes === null) return
    try { await adminApi.updateUserNotes(u.id, notes); u.notes = notes } catch { /* */ } }
const subStats = ref<any>(null)
const fetchSubStats = async () => { try { subStats.value = await adminApi.subscriptionStats() } catch { /* */ } }

onMounted(() => { fetchStats(); fetchUsers(); fetchPlans(); fetchLogs(); fetchSubStats() })
</script>
