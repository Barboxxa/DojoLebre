package service

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/rekognition"

	"github.com/Barboxxa/DojoLebre/internal/domain"
)

type Upload interface {
	GetSign(ctx context.Context, payload domain.SignRequest) (string, error)
}

type uploadService struct {
	rekoClient *rekognition.Client
}

func NewUploadService(rekoClient *rekognition.Client) Upload {
	return &uploadService{rekoClient}
}

func (p *uploadService) GetSign(ctx context.Context, payload domain.SignRequest) (string, error) {

	// TODO: usar detect labels: é um documento?
	input := &rekognition.DetectLabelsInput{
		// Image: &rekognition.{
		// 		Bytes: decodedImage,
		// },
	}

	_, err := p.rekoClient.DetectLabels(ctx, input)

	if err != nil {
		return "", err
	}

	// TODO: Usar text track, para identificar assinatura;

	// TODO: recortar assinatura e remover fundo;

	// TODO: retornar assinatura em base 64;

	return "", nil
}
