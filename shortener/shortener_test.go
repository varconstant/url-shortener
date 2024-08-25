package shortener

import (
	"testing"
)

type MockStorageService struct {
}

func (s *MockStorageService) Connect() error {
	return nil
}

func (s *MockStorageService) Save(k string, val any) error {
	return nil
}
func (s *MockStorageService) Fetch(k string) (any, error) {
	return "", nil
}

func TestUrlShortener_Shorten(t *testing.T) {
	s := NewUrlShortener()
	encoded, err := s.Shorten(ShortenRequest{"https://www.google.com", "user1"})
	if err != nil {
		t.Error("Error while shortening url", err)
	}
	if encoded == "" {
		t.Error("encoding error: encoded url blank")
	}
	if len(encoded) != 8 {
		t.Error("encoding error: encoded short code length not equal to 8")
	}
	println(encoded)
}
