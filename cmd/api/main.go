package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"crawlx/internal/crawler"
	"crawlx/internal/ui"
)

func main() {
	ui.Banner()

	seed := flag.String("seed", "", "seed url")
	max_pages := flag.Int("max-pages", 10, "max page to crawl")
	max_depth := flag.Int("max-depth", 1, "max page to crawl")
	workers := flag.Int("workers", 10, "workers count to work")
	flag.Parse()

	fmt.Println(ui.Dim(fmt.Sprintf("seed=%s\nmax-pages=%d max-depth=%d workers=%d\n", *seed, *max_pages, *max_depth, *workers)))

	if *seed == "" {
		fmt.Println(ui.Red("Seed url required"))
		os.Exit(1)
	}

	if *workers < 1 {
		*workers = 1
	}

	crawlerStruct := crawler.NewCrawler(*seed, *max_pages, *max_depth, *workers)

	result, err := crawlerStruct.Crawl(http.MethodGet)
	if err != nil {
		fmt.Println(ui.Red(err.Error()))
	}

	errors := 0
	for _, res := range result.PageResult {
		if res.Error != "" || res.Status >= 400 {
			errors++
		}
	}

	fmt.Println("")
	done := fmt.Sprintf("Done pages= %d edges=%d error=%d", len(result.PageResult), len(result.Edges), errors)
	if errors == 0 {
		fmt.Println(ui.Cyan(done))
	} else {
		fmt.Printf("Done pages= %d edges=%d %s\n", len(result.PageResult), len(result.Edges), ui.Red(fmt.Sprintf("error=%d", errors)))
	}
	crawler.GetReference(&result)
}
