package usecase

import (
	"context"
	"strings"

	"google.golang.org/grpc/metadata"

	"github.com/a13hander/auth-service-api/internal/domain/model"
	"github.com/a13hander/auth-service-api/internal/domain/util"
)

const authPrefix = "Bearer "

type AccessRepo interface {
	Get(ctx context.Context, endpoint string) (*model.Access, error)
}

type CheckEndpoint interface {
	Check(ctx context.Context, url string) bool
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

func (c *checkEndpoint) Check(ctx context.Context, url string) bool {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return false
	}

	authHeader, ok := md["authorization"]
	if !ok || len(authHeader) == 0 {
		return false
	}

	if !strings.HasPrefix(authHeader[0], authPrefix) {
		return false
	}

	accessToken := strings.TrimPrefix(authHeader[0], authPrefix)
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
