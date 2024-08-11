package main

import (
	"regexp"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

// converts Markdown to HTML
func mdToHTML(md []byte) []byte {
	// 1. markdown parser
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs
	parser := parser.NewWithExtensions(extensions)

	// 2. html renderer
	opts := html.RendererOptions{
		Flags: html.CommonFlags | html.HrefTargetBlank,
	}
	renderer := html.NewRenderer(opts)

	// 3. convert markdown to html
	doc := parser.Parse(md)
	output := markdown.Render(doc, renderer)

	return output
}

// extracts header links from markdown file, to be shown in right sidebar
func extractHeaders(content []byte) []string {
	var headers []string
	// regex to match only h2
	re := regexp.MustCompile(`(?m)^##\s*(.+)`)
	matches := re.FindAllSubmatch(content, -1)

	for _, match := range matches {
		// match[1] contains header text without the '##'
		headers = append(headers, string(match[1]))
	}

	return headers
}
