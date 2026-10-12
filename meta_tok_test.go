package main

import (
	"fmt"
	"strings"
	"golang.org/x/net/html"
)

func main() {
	s := `<html><head>
<meta name="description" content="plain desc">
<meta name="description" content="slash desc" />
<meta property="og:description" content="og desc">
<title>T</title></head></html>`
	z := html.NewTokenizer(strings.NewReader(s))
	for {
		tt := z.Next()
		if tt == html.ErrorToken {
			break
		}
		if tt == html.StartTagToken || tt == html.SelfClosingTagToken {
			t := z.Token()
			if t.Data == "meta" {
				fmt.Printf("tokenType=%v attrs=%v\n", tt, t.Attr)
			}
		}
	}
}
