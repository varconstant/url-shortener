package shortener

import (
	"errors"
	"fmt"
	"math/big"
	"urlShortner/util"
)

type ShortenRequest struct {
	Url    string
	UserId string
}

type Shortener interface {
	Shorten(req ShortenRequest) (string, error)
}

type UrlShortener struct {
}

func NewUrlShortener() *UrlShortener {
	return &UrlShortener{}
}

func (shortener *UrlShortener) Shorten(req ShortenRequest) (string, error) {
	if req.Url == "" || req.UserId == "" {
		return "", errors.New("url is empty")
	}
	bytes := util.Sha256Of(req.Url + req.UserId)
	generatedNumber := new(big.Int).SetBytes(bytes).Uint64()
	encoded, err := util.Base58Encoded([]byte(fmt.Sprintf("%d", generatedNumber)))
	if err != nil {
		return "", err
	}
	return encoded[:8], nil
}
