import { api } from './api'

export const inboxApi = {
    getMessages: (aliasId: string) => api.get('/alias/' + aliasId + '/inbox'),
    getMessage: (id: number) => api.get('/inbox/message/' + id),
    deleteMessage: (id: number) => api.delete('/inbox/message/' + id),
}
