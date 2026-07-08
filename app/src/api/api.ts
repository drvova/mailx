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

const REQUEST_TIMEOUT_MS = 15000

// Human messages for failures the server cannot explain itself.
// Server-provided error messages always take precedence.
function friendlyFailure(status: number, data: any): string | undefined {
    if (data && typeof data === 'object' && data.error) return undefined
    if (status === 429) return 'Too many requests. Wait a moment and try again.'
    if (status >= 500) return 'Something went wrong on our end. Try again in a moment.'
    return undefined
}

async function request(method: string, path: string, body?: any): Promise<{ data: any; status: number }> {
    const opts: RequestInit = {
        method,
        credentials: 'include',
        headers: {} as Record<string, string>,
        signal: AbortSignal.timeout(REQUEST_TIMEOUT_MS),
    }
    if (body !== undefined) {
        ;(opts.headers as Record<string, string>)['Content-Type'] = 'application/json'
        opts.body = JSON.stringify(body)
    }

    let res: Response
    try {
        res = await fetch(BASE + path, opts)
    } catch (err) {
        const timedOut = err instanceof DOMException && err.name === 'TimeoutError'
        throw new ApiError(0, {
            error: timedOut
                ? 'The server is taking too long to respond. Try again.'
                : "Can't reach the server. Check your connection and try again.",
        })
    }

    const text = await res.text()
    let data: any
    try { data = JSON.parse(text) } catch { data = text }

    if (!res.ok) {
        if (res.status === 401 && window.location.pathname.startsWith('/account')) {
            // Session expired mid-task: preserve where the user was and tell
            // them why they landed on the login page.
            localStorage.removeItem('email')
            sessionStorage.setItem('session_expired', '1')
            window.location.href = '/login?redirect=' + encodeURIComponent(window.location.pathname + window.location.search)
        }
        const friendly = friendlyFailure(res.status, data)
        throw new ApiError(res.status, friendly ? { error: friendly } : data)
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
