package parser

import (
	"fmt"
	"net/url"
	"strings"
)

func ReadUrl(pageUrl string, parseURLs []string) (NormalizeUrl []string, err error) {
	for _, url := range parseURLs {
		clearnURl, err := normalizeURL(pageUrl, url)
		if err == nil {
			NormalizeUrl = append(NormalizeUrl, clearnURl)
		}
	}
	return
}

func normalizeURL(baseUrl string, href string) (string, error) {

	base, err := url.Parse(baseUrl)
	if err != nil {
		return "", err
	}
	path, err := url.Parse(href)
	if err != nil {
		return "", err
	}

	if path.Scheme != "" && path.Scheme != "http" && path.Scheme != "https" {
		// fmt.Println("wrong scheme")
		return "", fmt.Errorf("wrong url scheme")
	}

	fullURL := base.ResolveReference(path)
	fullURL.Fragment = ""
	host := strings.ToLower(fullURL.Hostname())

	port := fullURL.Port()

	if port != "" ||
		(fullURL.Scheme != "http" && port != "80") ||
		(fullURL.Scheme != "https" && port != "443") {
		fullURL.Host = host

	} else {
		fullURL.Host = host + ":" + port
	}

	return fullURL.String(), nil
}
