package fetcher

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

var client = &http.Client{Timeout: 5 * time.Second}

func Fetch(url string) (status int, body []byte, err error) {
	resp, err := client.Get(url)

	if err != nil {
		return 0, nil, err
	}
	defer func() {
		if cerr := resp.Body.Close(); cerr != nil && err == nil {
			err = fmt.Errorf("closing body: %w", cerr)
		}
	}()
	body, err = io.ReadAll(resp.Body)

	if err != nil {
		return 0, nil, err
	}
	return resp.StatusCode, body, nil
}
