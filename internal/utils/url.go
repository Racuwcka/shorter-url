package shortid

import (
	"errors"
	"fmt"
	"strings"
)

func GetShortId(link string) (string, error) {
	index := strings.Index(link, "/link/")
	if index == -1 {
		return "", errors.New("url is not correct")
	}

	shortId := link[index+6:]
	if shortId == "" {
		return "", errors.New("link is not correct")
	}

	return shortId, nil
}

func CreateShortLink(url string, shortID string) string {
	return fmt.Sprintf("%s/link/%s", url, shortID)
}
