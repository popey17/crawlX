package crawler

import (
	"context"
	"crawlx/internal/fetcher"
	"crawlx/internal/parser"
	"fmt"
	"net/url"
	"time"
)

type Crawler struct {
	seen     map[string]bool
	startURL string
	maxPage  int
	maxDepth int
}

type queue struct {
	URL    string
	Depth  int
	parent string
}

type edge struct {
	Source string
	Target string
	Type   string
}

type pageResult struct {
	Status      int
	Title       string
	URL         string
	Depth       int
	FirstParent string
	Error       string
}

type result struct {
	PageResult []pageResult
	Edges      []edge
}

func NewCrawler(startURL string, maxPage int, maxDepth int) *Crawler {
	return &Crawler{
		seen:     map[string]bool{},
		startURL: startURL,
		maxPage:  maxPage,
		maxDepth: maxDepth,
	}
}

func GetReference(res *result) {
	refs := map[string][]string{}
	for _, edge := range res.Edges {
		refs[edge.Target] = append(refs[edge.Target], edge.Source)
	}

	// for target, sources := range refs {
	// 	fmt.Println("target:", target, "sources:", sources)
	// }
	fmt.Println("\nIssues")
	for _, finalRes := range res.PageResult {
		if finalRes.Error != "" || finalRes.Status >= 400 {
			fmt.Printf("%d %s\n", finalRes.Status, finalRes.URL)
			fmt.Println("referers:")
			for _, refererURL := range refs[finalRes.URL] {
				fmt.Println("-", refererURL)
			}
		}
	}
}

func (c *Crawler) enqueueLinks(normalizeUrl []string, res *result, URL string, cur queue, queueItem *[]queue) {
	for _, link := range normalizeUrl {
		res.Edges = append(res.Edges, edge{
			Source: URL,
			Target: link,
			Type:   "anchor",
		})
		linkURL, err := url.Parse(link)
		if err != nil {
			continue
		}

		BaseURL, err := url.Parse(URL)
		if err != nil {
			continue
		}

		if !c.seen[link] && linkURL.Hostname() == BaseURL.Hostname() && (c.maxDepth == 0 || cur.Depth < c.maxDepth) {
			c.seen[link] = true
			*queueItem = append(*queueItem, queue{URL: link, Depth: cur.Depth + 1, parent: URL})
		}
	}
}

func appendResult(res *[]pageResult, status int, title string, URL string, depth int, firstParent string, err string, crawled int) {
	*res = append(*res, pageResult{
		Status:      status,
		Title:       title,
		URL:         URL,
		Depth:       depth,
		FirstParent: firstParent,
		Error:       err,
	})
	fmt.Printf("#%d depth=%d status=%d err=%s %s %s\n", crawled, depth, status, err, URL, firstParent)
}

func (c *Crawler) Crawl(method string) (result, error) {
	queueItem := []queue{{
		URL:    c.startURL,
		Depth:  0,
		parent: "",
	}}
	c.seen[c.startURL] = true
	var crawled int
	res := result{}

	for len(queueItem) > 0 && (c.maxPage == 0 || crawled < c.maxPage) {
		cur := queueItem[0]
		URL := cur.URL
		queueItem = queueItem[1:]
		crawled++
		// fmt.Println(URL, crawled)

		//crawl
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		status, body, err := fetcher.Fetch(ctx, method, URL)
		cancel()
		if err != nil {
			appendResult(&res.PageResult, status, "", URL, cur.Depth, cur.parent, err.Error(), crawled)
			continue
		}

		//parse html
		meta, href, err := parser.ParseHTML(body)
		if err != nil {
			continue
		}
		appendResult(&res.PageResult, status, meta.Title, URL, cur.Depth, cur.parent, "", crawled)

		//normalize URL
		normalizeUrl, err := parser.ReadUrl(URL, href)
		if err != nil {
			fmt.Println(err)
			continue
		}

		c.enqueueLinks(normalizeUrl, &res, URL, cur, &queueItem)

	}
	return res, nil
}
