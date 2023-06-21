package usecase

import (
	"context"
	"time"

	"github.com/a13hander/auth-service-api/internal/domain/errs"
	"github.com/a13hander/auth-service-api/internal/domain/util"
)

type accessTokenGenerator struct {
	repo             UserRepo
	refreshSecretKey []byte
	accessSecretKey  []byte
	expiration       time.Duration
}

func NewAccessTokenGenerator(repo UserRepo, refreshSecretKey []byte, accessSecretKey []byte, expiration time.Duration) *accessTokenGenerator {
	return &accessTokenGenerator{repo: repo, refreshSecretKey: refreshSecretKey, accessSecretKey: accessSecretKey, expiration: expiration}
}

func (g *accessTokenGenerator) Generate(ctx context.Context, refreshToken string) (string, error) {
	claims, err := util.VerifyToken(refreshToken, g.refreshSecretKey)
	if err != nil {
		return "", err
	}

	user, err := g.repo.Get(ctx, claims.Username)
	if err != nil {
		return "", err
	}

	if user == nil {
		return "", errs.NotFoundErr
	}

	return util.GenerateToken(user, g.accessSecretKey, g.expiration)
}
