package tiktoken_loader

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"fmt"
	"path"
	"strconv"

	"github.com/pkoukk/tiktoken-go-loader/assets"
)

type OfflineLoader struct{}

func (l *OfflineLoader) LoadTiktokenBpe(tiktokenBpeFile string) (map[string]int, error) {
	name := path.Base(tiktokenBpeFile)
	f, err := assets.Assets.Open(name)
	if err != nil {
		return nil, fmt.Errorf("failed to open file %q: %w", name, err)
	}
	defer f.Close()

	bpeRanks := make(map[string]int)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Bytes()
		if len(line) == 0 {
			continue
		}
		tokenStr, rankStr, ok := bytes.Cut(line, []byte(" "))
		if !ok {
			return nil, fmt.Errorf("invalid line in asset %q: %q", name, line)
		}
		token := make([]byte, base64.StdEncoding.DecodedLen(len(tokenStr)))
		n, err := base64.StdEncoding.Decode(token, tokenStr)
		if err != nil {
			return nil, fmt.Errorf("decoding token from asset: %w", err)
		}
		token = token[:n]
		rank, err := strconv.Atoi(string(rankStr))
		if err != nil {
			return nil, fmt.Errorf("parsing rank from asset: %w", err)
		}
		bpeRanks[string(token)] = rank
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("reading asset %q: %w", name, err)
	}
	return bpeRanks, nil
}

func NewOfflineLoader() *OfflineLoader {
	return &OfflineLoader{}
}
