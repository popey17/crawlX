package crawler

import (
	"context"
	"crawlx/internal/fetcher"
	"crawlx/internal/parser"
	"fmt"
	"net/url"
	"sync"
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

func (c *Crawler) enqueueLinks(mutex *sync.Mutex, normalizeUrl []string, res *result, URL string, cur queue, jobs chan queue, pending *int) {
	for _, link := range normalizeUrl {
		mutex.Lock()
		res.Edges = append(res.Edges, edge{
			Source: URL,
			Target: link,
			Type:   "anchor",
		})
		mutex.Unlock()

		linkURL, err := url.Parse(link)
		if err != nil {
			continue
		}

		BaseURL, err := url.Parse(URL)
		if err != nil {
			continue
		}

		if linkURL.Hostname() != BaseURL.Hostname() || (c.maxDepth != 0 && cur.Depth >= c.maxDepth) {
			continue
		}

		mutex.Lock()

		if c.seen[link] {
			mutex.Unlock()
			continue
		} else {
			c.seen[link] = true
			*pending = *pending + 1
		}
		mutex.Unlock()
		jobs <- queue{URL: link, Depth: cur.Depth + 1, parent: URL}
	}
}

func appendResult(mutex *sync.Mutex, res *[]pageResult, status int, title string, URL string, depth int, firstParent string, err string, crawled int) {
	mutex.Lock()
	*res = append(*res, pageResult{
		Status:      status,
		Title:       title,
		URL:         URL,
		Depth:       depth,
		FirstParent: firstParent,
		Error:       err,
	})
	mutex.Unlock()
	fmt.Printf("#%d depth=%d status=%d err=%s %s %s\n", crawled, depth, status, err, URL, firstParent)
}

func (c *Crawler) worker(mutex *sync.Mutex, wg *sync.WaitGroup, jobs chan queue, pending *int, crawled *int, method string, res *result) {
	defer wg.Done()
	finish := func() {
		mutex.Lock()
		*pending--
		if *pending == 0 {
			close(jobs)
		}
		mutex.Unlock()
	}
	for job := range jobs {

		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

		mutex.Lock()
		*crawled++
		n := *crawled
		mutex.Unlock()
		status, body, err := fetcher.Fetch(ctx, method, job.URL)
		cancel()

		if err != nil {
			appendResult(mutex, &res.PageResult, status, "", job.URL, job.Depth, job.parent, err.Error(), n)
			finish()
			continue
		}

		// 	//parse html
		meta, href, err := parser.ParseHTML(body)
		if err != nil {
			finish()
			continue
		}
		appendResult(mutex, &res.PageResult, status, meta.Title, job.URL, job.Depth, job.parent, "", n)

		// 	//normalize URL
		normalizeUrl, err := parser.ReadUrl(job.URL, href)
		if err != nil {
			fmt.Println(err)
			finish()
			continue
		}
		c.enqueueLinks(mutex, normalizeUrl, res, job.URL, job, jobs, pending)

		finish()
	}
}

func (c *Crawler) Crawl(method string) (result, error) {
	jobs := make(chan queue, 4000)
	wg := &sync.WaitGroup{}
	mutex := &sync.Mutex{}

	jobs <- queue{
		URL:    c.startURL,
		Depth:  0,
		parent: "",
	}

	c.seen[c.startURL] = true
	pending := 1
	crawled := 0

	res := &result{}

	for range 4 {
		wg.Add(1)
		go c.worker(mutex, wg, jobs, &pending, &crawled, method, res)
	}

	wg.Wait()

	return *res, nil
}
