package linkedin

import (
	"bytes"
	"context"
	"fmt"
	"strings"

	"jobs_scanner/internal/model"
	"jobs_scanner/internal/source"

	"golang.org/x/net/html"
)

type LinkedinParser struct {
	ByteHTML []byte
}

func (p *LinkedinParser) Fetch(ctx context.Context) ([]model.Vacancy, error) {
	doc, err := html.Parse(bytes.NewReader(p.ByteHTML))
	if err != nil {
		return nil, fmt.Errorf("error parsing the file")
	}
	var vacancies []model.Vacancy
	nodes := source.Walk(doc, isCard)
	for _, v := range nodes {
		href := ""
		for _, attr := range v.Parent.Attr {
			if attr.Key == "href" {
				href = attr.Val
			}
		}
		vacancies = append(vacancies, model.Vacancy{Title: strings.TrimSpace(v.FirstChild.Data), URL: href})
	}
	return vacancies, nil
}

func isCard(node *html.Node) bool {
	if node.Type != html.ElementNode {
		return false
	}
	if node.Data != "span" {
		return false
	}
	for _, v := range node.Attr {
		if v.Key == "class" {
			classes := strings.SplitSeq(v.Val, " ")
			for class := range classes {
				if class == "sr-only" {
					return true
				}
			}
		}
	}
	return false
}
