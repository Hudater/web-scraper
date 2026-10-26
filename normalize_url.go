package main

import (
	"errors"
	"net/url"
)

func normalizeURL(ogURL string) (string, error) {
	normalizedURL, err := url.Parse(ogURL)
	if err != nil {
		return "", errors.New("could not parse URL")
	}
	return normalizedURL.Host+normalizedURL.Path, nil
}
