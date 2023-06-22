package interceptors

import (
	"context"

	"github.com/a13hander/auth-service-api/internal/service/ratelimiter"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RateLimiterInterceptor struct {
	rateLimiter ratelimiter.RateLimiter
}

func NewRateLimiterInterceptor(rateLimiter ratelimiter.RateLimiter) *RateLimiterInterceptor {
	return &RateLimiterInterceptor{rateLimiter: rateLimiter}
}

func (r *RateLimiterInterceptor) Unary(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	if !r.rateLimiter.Allow() {
		return nil, status.Error(codes.ResourceExhausted, "too many requests")
	}

	return handler(ctx, req)
}
