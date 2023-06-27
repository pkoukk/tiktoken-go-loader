package tiktoken_loader

import (
	"testing"
)

func TestOffline(t *testing.T) {
	loader := NewOfflineLoader()
	bpeRanks, err := loader.LoadTiktokenBpe("https://openaipublic.blob.core.windows.net/encodings/cl100k_base.tiktoken")
	if err != nil {
		t.Error(err)
	}
	if len(bpeRanks) == 0 {
		t.Error("bpeRanks is empty")
	}
}
