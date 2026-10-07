package main

import (
	"context"
	"crawlx/internal/fetcher"
	"fmt"
	"net/http"
	"time"
)

func main() {
	fmt.Println("crawlx: ok")

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	status, body, err := fetcher.Fetch(ctx, http.MethodGet, "https://example.com")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("status=%d byte=%d\n", status, len(body))
}
