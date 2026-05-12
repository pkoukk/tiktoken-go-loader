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

var benchFiles = []string{
	"cl100k_base.tiktoken",
	"o200k_base.tiktoken",
	"p50k_base.tiktoken",
	"r50k_base.tiktoken",
}

func BenchmarkLoadTiktokenBpe(b *testing.B) {
	loader := NewOfflineLoader()
	for _, name := range benchFiles {
		b.Run(name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				m, err := loader.LoadTiktokenBpe(name)
				if err != nil {
					b.Fatal(err)
				}
				if len(m) == 0 {
					b.Fatal("empty")
				}
			}
		})
	}
}
