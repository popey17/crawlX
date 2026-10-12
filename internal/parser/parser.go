package parser

import (
	"bytes"
	"strings"

	"golang.org/x/net/html"
)

type Meta struct {
	Title       string
	Description string
	Keywords    string
}

func ParseHTML(body []byte) (meta Meta, href []string, err error) {

	z := html.NewTokenizer(bytes.NewReader(body))
	inTitle := false

	for {
		tt := z.Next()
		if tt == html.ErrorToken {
			break
		}

		if tt == html.StartTagToken || tt == html.SelfClosingTagToken {
			t := z.Token()
			if t.Data == "title" {
				inTitle = true
			}
			if t.Data == "meta" {
				var name, content string
				for _, a := range t.Attr {
					switch strings.ToLower(a.Key) {
					case "name":
						name = strings.ToLower(a.Val)
					case "content":
						content = strings.TrimSpace(a.Val)
					}
				}

				switch name {
				case "description":
					meta.Description = content
				case "keywords":
					meta.Keywords = content
				}
			}
			if t.Data == "a" {
				for _, attr := range t.Attr {
					if attr.Key == "href" && attr.Val != "" {
						href = append(href, attr.Val)
					}
				}
			}
		}
		if tt == html.TextToken && inTitle {
			meta.Title = strings.TrimSpace(z.Token().Data)
		}
		if tt == html.EndTagToken {
			if z.Token().Data == "title" {
				inTitle = false
			}
		}
	}
	return meta, href, nil
}
