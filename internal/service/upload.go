package service

import (
	"context"

	"github.com/Barboxxa/DojoLebre/internal/domain"
)

type Upload interface {
	GetSign(ctx context.Context, payload domain.SignRequest) (string, error)
}

type uploadService struct {
}

func NewUploadService() Upload {
	return &uploadService{}
}

func (p *uploadService) GetSign(ctx context.Context, payload domain.SignRequest) (string, error) {

	return "", nil
}
