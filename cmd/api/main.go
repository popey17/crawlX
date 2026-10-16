package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"crawlx/internal/crawler"
)

func main() {
	fmt.Println("crawlx")
	fmt.Println("============")

	seed := flag.String("seed", "", "seed url")
	max_pages := flag.Int("max-pages", 10, "max page to crawl")
	max_depth := flag.Int("max-depth", 1, "max page to crawl")
	workers := flag.Int("workers", 10, "workers count to work")
	flag.Parse()

	fmt.Printf("seed=%s\nmax-pages=%d max-depth=%d\n\n", *seed, *max_pages, *max_depth)

	if *seed == "" {
		fmt.Println("Seed url required")
		os.Exit(1)
	}

	if *workers < 1 {
		*workers = 1
	}

	crawlerStruct := crawler.NewCrawler(*seed, *max_pages, *max_depth, *workers)

	result, err := crawlerStruct.Crawl(http.MethodGet)
	if err != nil {
		fmt.Println(err)
	}

	errors := 0
	for _, res := range result.PageResult {
		if res.Error != "" || res.Status >= 400 {
			errors++
		}
	}

	fmt.Printf("Done pages= %d edges=%d error=%d\n", len(result.PageResult), len(result.Edges), errors)
	crawler.GetReference(&result)
}
