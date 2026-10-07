package main

import (
	"crawlx/internal/fetcher"
	"fmt"
)

func main() {
	fmt.Println("crawlx: ok")

	status, body, err := fetcher.Fetch("https://example.com")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("status=%d byte=%d\n", status, len(body))
}
