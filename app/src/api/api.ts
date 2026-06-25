const BASE = import.meta.env.VITE_API_URL + '/v1'

export class ApiError extends Error {
    status: number
    data: any

    constructor(status: number, data: any) {
        super(data?.error || `HTTP ${status}`)
        this.status = status
        this.data = data
    }
}

async function request(method: string, path: string, body?: any): Promise<{ data: any; status: number }> {
    const opts: RequestInit = {
        method,
        credentials: 'include',
        headers: {} as Record<string, string>,
    }
    if (body !== undefined) {
        ;(opts.headers as Record<string, string>)['Content-Type'] = 'application/json'
        opts.body = JSON.stringify(body)
    }

    const res = await fetch(BASE + path, opts)
    const text = await res.text()
    let data: any
    try { data = JSON.parse(text) } catch { data = text }

    if (!res.ok) {
        if (res.status === 401 && window.location.pathname.startsWith('/account')) {
            localStorage.removeItem('email')
            window.location.href = '/'
        }
        throw new ApiError(res.status, data)
    }
    return { data, status: res.status }
}

function buildPath(path: string, params?: Record<string, any>): string {
    if (!params) return path
    const qs = new URLSearchParams()
    for (const [k, v] of Object.entries(params)) {
        if (v !== undefined && v !== null) qs.set(k, String(v))
    }
    const s = qs.toString()
    return s ? path + '?' + s : path
}

export const api = {
    get:    (path: string, opts?: { params?: any }) => request('GET', buildPath(path, opts?.params)),
    post:   (path: string, body?: any)  => request('POST', path, body),
    put:    (path: string, body?: any)  => request('PUT', path, body),
    delete: (path: string)              => request('DELETE', path),
}
