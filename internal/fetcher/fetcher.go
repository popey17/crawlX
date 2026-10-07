package fetcher

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

var client = &http.Client{Timeout: 5 * time.Second}

func Fetch(ctx context.Context, method string, url string) (status int, body []byte, err error) {
	request, err := http.NewRequestWithContext(ctx, method, url, nil)
	if err != nil {
		return 0, nil, err
	}
	request.Header.Set("User-Agent", "crawlx/0.1 (+local training)")

	resp, err := client.Do(request)

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
