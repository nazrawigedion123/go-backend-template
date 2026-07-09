package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type clientRate struct {
	tokens int
	last   time.Time
}

type rateLimiter struct {
	mu       sync.Mutex
	clients  map[string]*clientRate
	rate     int
	interval time.Duration
	burst    int
}

func newRateLimiter(rate int, interval time.Duration, burst int) *rateLimiter {
	rl := &rateLimiter{
		clients:  make(map[string]*clientRate),
		rate:     rate,
		interval: interval,
		burst:    burst,
	}
	go rl.cleanup()
	return rl
}

func (rl *rateLimiter) cleanup() {
	for {
		time.Sleep(10 * time.Minute)
		rl.mu.Lock()
		for ip, cr := range rl.clients {
			if time.Since(cr.last) > rl.interval*2 {
				delete(rl.clients, ip)
			}
		}
		rl.mu.Unlock()
	}
}

func (rl *rateLimiter) allow(ip string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	cr, exists := rl.clients[ip]
	now := time.Now()

	if !exists {
		rl.clients[ip] = &clientRate{tokens: rl.burst - 1, last: now}
		return true
	}

	elapsed := now.Sub(cr.last)
	cr.last = now
	cr.tokens += int(elapsed / rl.interval)
	if cr.tokens > rl.burst {
		cr.tokens = rl.burst
	}

	if cr.tokens <= 0 {
		return false
	}
	cr.tokens--
	return true
}

func RateLimiterMiddleware() gin.HandlerFunc {
	rl := newRateLimiter(5, time.Second, 5)
	return func(c *gin.Context) {
		if !rl.allow(c.ClientIP()) {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "too many requests"})
			c.Abort()
			return
		}
		c.Next()
	}
}
