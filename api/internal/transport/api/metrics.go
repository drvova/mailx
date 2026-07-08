package api

import (
	"strconv"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
)

// Metrics holds anonymous endpoint hit counters — no PII, no cookies,
// no client-side JS. Counters are keyed by "METHOD /route-pattern" and
// bucketed by day (UTC). A sliding window of 90 days is retained.
type Metrics struct {
	mu      sync.RWMutex
	buckets map[string]int64 // key: "2006-01-02|GET|/v1/aliases"
}

func NewMetrics() *Metrics {
	return &Metrics{buckets: make(map[string]int64)}
}

// Record increments the counter for the given route+method on today's bucket.
func (m *Metrics) Record(route, method string) {
	key := time.Now().UTC().Format("2006-01-02") + "|" + method + "|" + route
	m.mu.Lock()
	m.buckets[key]++
	m.mu.Unlock()
}

// Snapshot returns a copy of all counters, optionally filtered to the
// last N days. The map key is "date|method|route".
func (m *Metrics) Snapshot(days int) map[string]int64 {
	cutoff := time.Now().UTC().AddDate(0, 0, -days)
	m.mu.RLock()
	defer m.mu.RUnlock()
	out := make(map[string]int64, len(m.buckets))
	for k, v := range m.buckets {
		dateStr := k[:10]
		t, err := time.Parse("2006-01-02", dateStr)
		if err != nil || t.Before(cutoff) {
			continue
		}
		out[k] = v
	}
	return out
}

// Prune removes buckets older than 90 days. Called periodically.
func (m *Metrics) Prune() {
	cutoff := time.Now().UTC().AddDate(0, 0, -90)
	m.mu.Lock()
	for k := range m.buckets {
		dateStr := k[:10]
		t, err := time.Parse("2006-01-02", dateStr)
		if err != nil || t.Before(cutoff) {
			delete(m.buckets, k)
		}
	}
	m.mu.Unlock()
}

// MetricsMiddleware records a hit for each request using the matched
// route pattern (not the raw URL, so /alias/123 and /alias/456 both
// count under /alias/:id — no path-parameter PII leaks into metrics).
func (m *Metrics) Middleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Next()
		route := c.Route().Path
		if route == "" {
			route = "unknown"
		}
		m.Record(route, c.Method())
		return nil
	}
}

// HandleMetrics returns the aggregated counters as JSON. Protected by
// the same basic-auth guard as /docs.
func (m *Metrics) HandleMetrics(c *fiber.Ctx) error {
	days := 30
	if d := c.Query("days"); d != "" {
		if n, err := strconv.Atoi(d); err == nil && n > 0 && n <= 90 {
			days = n
		}
	}
	return c.JSON(m.Snapshot(days))
}
