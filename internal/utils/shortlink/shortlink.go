package shortlink

import (
	"fmt"
)

func Create(url string, shortID string) string {
	return fmt.Sprintf("%s/link/%s", url, shortID)
}
