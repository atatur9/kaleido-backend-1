package admanager

import (
	"aurora-backend/lib/functions/lib/ad"
	"context"
	"testing"
)

func TestNewProvider(t *testing.T) {
	t.Run("connect succeed", func(t *testing.T) {
		_, err := NewProvider()
		if err != nil {
			t.Error(err)
		}
	})
}

func TestDisplayByIndex(t *testing.T) {
	t.Run("display enabled", func(t *testing.T) {
		p, _ := NewProvider()
		meta, err := p.DisplayByMetadata(context.Background(), ad.GetInput{
			Account:  "0xb5bE22F33D8f0b1Cc131674C562069D1B5912147",
			Metadata: "",
		})
		if err != nil {
			t.Error(err)
		}
		if meta == "" {
			t.Error("metadata is empty string")
		}
	})
}
