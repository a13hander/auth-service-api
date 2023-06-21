package usecase

import (
	"context"
	"time"

	"github.com/a13hander/auth-service-api/internal/domain/errs"
	"github.com/a13hander/auth-service-api/internal/domain/util"
)

type refreshTokenGenerator struct {
	repo       UserRepo
	secretKey  []byte
	expiration time.Duration
}

func NewRefreshTokenGenerator(repo UserRepo, secretKey []byte, expiration time.Duration) *refreshTokenGenerator {
	return &refreshTokenGenerator{repo: repo, secretKey: secretKey, expiration: expiration}
}

func (g *refreshTokenGenerator) Generate(ctx context.Context, username string, password string) (string, error) {
	user, err := g.repo.Get(ctx, username)
	if err != nil {
		return "", errs.NotFoundErr
	}

	if !util.VerifyPassword(user.Password, password) {
		return "", errs.NotFoundErr
	}

	return util.GenerateToken(user, g.secretKey, g.expiration)
}
