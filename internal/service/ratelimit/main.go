package ratelimit

import (
	"context"
	"time"
)

var _ RateLimiter = (*tokenBucketLimiter)(nil)

type RateLimiter interface {
	Allow() bool
}

type tokenBucketLimiter struct {
	tokenBucketCh chan struct{}
}

func NewLimiter(ctx context.Context, limit int, period time.Duration) RateLimiter {
	limiter := &tokenBucketLimiter{
		tokenBucketCh: make(chan struct{}, limit),
	}

	for i := 0; i < limit; i++ {
		limiter.tokenBucketCh <- struct{}{}
	}

	replenishmentInterval := period.Nanoseconds() / int64(limit)
	go limiter.startPeriodicReplenishment(ctx, time.Duration(replenishmentInterval))

	return limiter
}

func (l *tokenBucketLimiter) startPeriodicReplenishment(ctx context.Context, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			l.tokenBucketCh <- struct{}{}
		}
	}
}

func (l *tokenBucketLimiter) Allow() bool {
	select {
	case <-l.tokenBucketCh:
		return true
	default:
		return false
	}
}
