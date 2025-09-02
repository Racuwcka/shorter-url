package shortid

import (
	"errors"
	"strings"
)

func Get(link string) (string, error) {
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
