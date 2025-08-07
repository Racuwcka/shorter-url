package shortener

import (
	"crypto/sha256"
	"fmt"
)

type Hash struct{}

func (h *Hash) Generate(link string) string {
	sum := sha256.Sum256([]byte(link))
	return fmt.Sprintf("%x", sum[:6])
}
