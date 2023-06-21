package usecase

import (
	"context"

	"github.com/a13hander/auth-service-api/internal/domain/model"
	"github.com/a13hander/auth-service-api/internal/domain/util"
)

type AccessRepo interface {
	Get(ctx context.Context, endpoint string) (*model.Access, error)
}

type CheckEndpoint interface {
	Check(ctx context.Context, accessToken string, url string) bool
}

type checkEndpoint struct {
	accessSecretKey []byte
	accessRepo      AccessRepo
}

func NewCheckEndpoint(accessSecretKey []byte, repo AccessRepo) *checkEndpoint {
	return &checkEndpoint{
		accessSecretKey: accessSecretKey,
		accessRepo:      repo,
	}
}

func (c *checkEndpoint) Check(ctx context.Context, accessToken string, url string) bool {
	token, err := util.VerifyToken(accessToken, c.accessSecretKey)
	if err != nil {
		return false
	}

	access, err := c.accessRepo.Get(ctx, url)
	if err != nil {
		return false
	}

	return token.Role == access.Role
}
