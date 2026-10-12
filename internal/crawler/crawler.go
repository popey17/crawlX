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

func GetReference(edges []edge) {
	refs := map[string][]string{}
	for _, edge := range edges {
		refs[edge.Target] = append(refs[edge.Target], edge.Source)
	}

	// for target, sources := range refs {
	// 	if len(sources) >= 2 {
	// 		// fmt.Println(target, sources)
	// 	}
	// }
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
		defer cancel()

		status, body, err := fetcher.Fetch(ctx, method, URL)
		if err != nil {
			res.PageResult = append(res.PageResult, pageResult{
				Status:      status,
				Title:       "",
				URL:         URL,
				Depth:       cur.Depth,
				FirstParent: cur.parent,
				Error:       err.Error(),
			})
			fmt.Printf("#%d depth=%d status=%d err=%s %s %s\n", crawled, cur.Depth, status, err.Error(), URL, cur.parent)
			continue
		}

		//parse html
		meta, href, err := parser.ParseHTML(body)
		if err != nil {
			continue
		}

		res.PageResult = append(res.PageResult, pageResult{
			Status:      status,
			Title:       meta.Title,
			URL:         URL,
			Depth:       cur.Depth,
			FirstParent: cur.parent,
			Error:       "",
		})
		fmt.Printf("#%d depth=%d status=%d err=%s %s %s\n", crawled, cur.Depth, status, "", URL, cur.parent)

		//normalize URL
		normalizeUrl, err := parser.ReadUrl(URL, href)
		if err != nil {
			fmt.Println(err)
			continue
		}
		// fmt.Println(normalizeUrl)

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
				queueItem = append(queueItem, queue{URL: link, Depth: cur.Depth + 1, parent: URL})
			}
		}
	}
	return res, nil
}
