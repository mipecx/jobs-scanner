package habr

import (
	"bytes"
	"context"
	"fmt"
	"strings"

	"jobs_scanner/internal/model"
	"jobs_scanner/internal/source"

	"golang.org/x/net/html"
)

type HabrParser struct {
	ByteHTML []byte
}

func (p *HabrParser) Fetch(ctx context.Context) ([]model.Vacancy, error) {
	doc, err := html.Parse(bytes.NewReader(p.ByteHTML))
	if err != nil {
		return nil, fmt.Errorf("error parsing the file")
	}
	source.NormalizeTree(doc, blankNode)

	var vacancies []model.Vacancy
	nodes := source.Walk(doc, isCard)
	for _, n := range nodes {
		titleNode := source.FindNodes(n, isTitle)
		// companyNode := source.FindNodes(n, isCompany)
		company := ""
		title := ""
		url := ""
		// fmt.Println(titleNode)
		// salary := source.FindNodes(n, isSalary)
		// date := source.FindNodes(n, isDate)
		for _, attr := range n.Attr {
			if attr.Key == "href" {
				url = "https://career.habr.com" + attr.Val
			}
		}
		// data := strings.ReplaceAll(n.FirstChild.Data, "\n", "")
		vacancies = append(vacancies, model.Vacancy{Title: title, URL: url, Company: company})
	}
	return vacancies, nil
}

func isCard(node *html.Node) bool {
	if node.Type != html.ElementNode {
		return false
	}
	if node.Data != "div" {
		return false
	}
	for _, v := range node.Attr {
		if v.Key == "class" {
			classes := strings.SplitSeq(v.Val, " ")
			for class := range classes {
				if class == "vacancy-card" {
					return true
				}
			}
		}
	}
	return false
}

func blankNode(node *html.Node) bool {
	if node.Type != html.TextNode {
		return false
	}
	trimmedText := strings.TrimSpace(node.Data)
	return trimmedText == ""
}

func isTitle(node *html.Node) bool {
	for _, attr := range node.Attr {
		if attr.Key == "a" && attr.Val == "vacancy-card__title-link" {
			fmt.Println("Found one!")
			return true
		}
	}
	return false
}

func isCompany(node *html.Node) bool {
	for _, attr := range node.Attr {
		if attr.Key == "class" && attr.Val == "vacancy-card__company" {
			return true
		}
	}
	return false
}

func isSalary(node *html.Node) bool {
	for _, attr := range node.Attr {
		if attr.Key == "class" && attr.Val == "vacancy-card__salary" {
			return true
		}
	}
	return false
}

func isDate(node *html.Node) bool {
	for _, attr := range node.Attr {
		if attr.Key == "class" && attr.Val == "vacancy-card__date" {
			return true
		}
	}
	return false
}
