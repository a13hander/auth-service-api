package usecase

import "context"

type CheckEndpoint interface {
	Check(ctx context.Context, url string) bool
}

type checkEndpoint struct {
}

func NewCheckEndpoint() *checkEndpoint {
	return &checkEndpoint{}
}

func (c checkEndpoint) Check(ctx context.Context, url string) bool {
	return true
}
