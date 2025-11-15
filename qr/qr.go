package qr

import (
	"errors"
	"strings"

	"github.com/skip2/go-qrcode"
)

func QrGen(content string) ([]byte, error) {
	if strings.TrimSpace(content) == "" {
		return nil, errors.New("empty text")
	}
	if len(content) > 4000 {
		return nil, errors.New("maximum text lenght is 4000 characters")
	}
	var png []byte
	png, err := qrcode.Encode(content, qrcode.Medium, 256)
	if err != nil {
		return nil, err
	}
	return png, nil
}
