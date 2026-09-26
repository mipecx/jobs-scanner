package source

import (
	"golang.org/x/net/html"
)

func Walk(node *html.Node, condition func(*html.Node) bool) []*html.Node {
	if node == nil {
		return nil
	}
	var nodes []*html.Node
	if condition(node) {
		nodes = append(nodes, node)
	}
	nodes = append(nodes, Walk(node.FirstChild, condition)...)
	nodes = append(nodes, Walk(node.NextSibling, condition)...)

	return nodes
}

func NormalizeTree(node *html.Node, condition func(*html.Node) bool) {
	if node == nil {
		return
	}
	if condition(node) {
		next := node.NextSibling
		node.Parent.RemoveChild(node)
		NormalizeTree(next, condition)
	} else {
		NormalizeTree(node.FirstChild, condition)
		NormalizeTree(node.NextSibling, condition)
	}
}

func FindNodes(node *html.Node, isTarget func(*html.Node) bool) *html.Node {
	if node == nil {
		return nil
	}

	var targetNode *html.Node
	if isTarget(node) {
		targetNode = node
	}

	if targetNode == nil {
		targetNode = FindNodes(node.FirstChild, isTarget)
	}
	if targetNode == nil {
		targetNode = FindNodes(node.NextSibling, isTarget)
	}

	return targetNode
}
