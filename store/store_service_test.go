package store

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewStoreService(t *testing.T) {
	assert.True(t, NewRedis(context.Background()) != nil)
}

func TestStorageService_Connect(t *testing.T) {
	storeService := NewRedis(context.Background())
	err := storeService.Connect()
	if err != nil {
		t.Error("Error while connecting to database store", err)
	}
}
