package service

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"image"
	_ "image/jpeg"
	"image/png"
	"strings"

	"github.com/aws/aws-sdk-go/service/rekognition"
	"github.com/aws/aws-sdk-go/service/textract"
	"github.com/texttheater/golang-levenshtein/levenshtein"

	"github.com/Barboxxa/DojoLebre/internal/domain"
)

type Upload interface {
	GetSign(ctx context.Context, payload domain.SignRequest) (string, error)
}

type uploadService struct {
	rekoClient     *rekognition.Rekognition
	textractClient *textract.Textract
}

func NewUploadService(
	rekoClient *rekognition.Rekognition,
	textractClient *textract.Textract,
) Upload {
	return &uploadService{rekoClient, textractClient}
}

type returnImage struct {
	Height float64
	Left   float64
	Top    float64
	Width  float64
}

func (u *uploadService) GetSign(ctx context.Context, payload domain.SignRequest) (string, error) {
	decodedImage, _ := base64.StdEncoding.DecodeString(payload.Image)

	// TODO: usar detect labels: é um documento?
	detectLabelsInput := &rekognition.DetectLabelsInput{
		Image: &rekognition.Image{
			Bytes: decodedImage,
		},
	}

	detectLabels, err := u.rekoClient.DetectLabels(detectLabelsInput)

	if err != nil {
		return "", err
	}

	hasDocument := false
	for _, label := range detectLabels.Labels {
		if label.Name != nil && *label.Name == "Document" {
			hasDocument = true
			break
		}
	}

	if !hasDocument {
		return "", fmt.Errorf("this image is not a document")
	}

	// TODO: Usar textract, para identificar assinatura;

	// Criar uma solicitação para detectar texto em uma imagem em base64
	textractInput := &textract.DetectDocumentTextInput{
		Document: &textract.Document{
			Bytes: decodedImage,
		},
	}

	// Chamar a operação DetectDocumentText
	result, err := u.textractClient.DetectDocumentText(textractInput)

	if err != nil {
		// fmt.Println("Erro ao detectar texto:", err)
		return "", err
	}

	textRg := "ASSINATURA DO TITULAR"
	textCNH := "ASSINATURA DO PORTADOR"
	positionArray := 0
	for _, item := range result.Blocks {
		if item.Text != nil && (*item.BlockType == "LINE" || *item.BlockType == "WORD" || *item.BlockType == "PAGE") {
			fmt.Println(*item.Text)
			distanceRG := levenshtein.DistanceForStrings([]rune(textRg), []rune(*item.Text), levenshtein.DefaultOptions)
			distanceCNH := levenshtein.DistanceForStrings([]rune(textCNH), []rune(*item.Text), levenshtein.DefaultOptions)
			if distanceRG <= 5 || distanceCNH <= 5 {
				fmt.Println(*item.Geometry.BoundingBox.Height)
				fmt.Println(*item.Geometry.BoundingBox.Width)
				fmt.Println(*item.Geometry.BoundingBox.Left)
				fmt.Println(*item.Geometry.BoundingBox.Top)
				break
			}
		}
		positionArray++
	}

	obj := result.Blocks[positionArray-1]
	Coordinates := returnImage{
		Height: *obj.Geometry.BoundingBox.Height,
		Width:  *obj.Geometry.BoundingBox.Width,
		Left:   *obj.Geometry.BoundingBox.Left,
		Top:    *obj.Geometry.BoundingBox.Top,
	}

	fmt.Println(Coordinates.Height)
	fmt.Println(Coordinates.Width)
	fmt.Println(Coordinates.Left)
	fmt.Println(Coordinates.Top)

	// TODO: recortar assinatura;
	// Criar um leitor de imagem a partir dos dados decodificados
	// imageReader := strings.NewReader("data:image/png;base64," + payload.Image)

	// imageReader := bytes.NewReader([]byte("data:image/png;base64," + payload.Image))

	imageReader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(payload.Image))

	// Decodificar a imagem (por exemplo, no formato PNG)
	img, _, err := image.Decode(imageReader)
	if err != nil {
		// fmt.Println("Erro ao decodificar a imagem:", err)
		return "", err
	}

	width := img.Bounds().Dx()
	height := img.Bounds().Dy()

	xInitial := float64(width) * Coordinates.Left
	yInitial := float64(height) * Coordinates.Top
	xFinal := xInitial + (float64(width) * Coordinates.Width)
	yFinal := yInitial + (float64(height) * Coordinates.Height)

	img, err = cropImage(img, image.Rect(
		int(xInitial),
		int(yInitial),
		int(xFinal),
		int(yFinal),
	))

	if err != nil {
		return "", err
	}

	buffer := new(bytes.Buffer)

	if err := png.Encode(buffer, img); err != nil {
		return "", err
	}

	base64String := base64.StdEncoding.EncodeToString(buffer.Bytes())

	// TODO: remover fundo;

	// TODO: retornar assinatura em base 64;
	return base64String, nil
}

// cropImage takes an image and crops it to the specified rectangle.
func cropImage(img image.Image, crop image.Rectangle) (image.Image, error) {
	type subImager interface {
		SubImage(r image.Rectangle) image.Image
	}

	// img is an Image interface. This checks if the underlying value has a
	// method called SubImage. If it does, then we can use SubImage to crop the
	// image.
	simg, ok := img.(subImager)
	if !ok {
		return nil, fmt.Errorf("image does not support cropping")
	}

	return simg.SubImage(crop), nil
}
