package main

import (
	"errors"
	"html/template"
	"regexp"
	"strconv"
	"strings"
)

// takes a markdown file and returns a BlogPost struct, and
// also extracts metadata from the markdown file
func parseMarkdownFile(content []byte) (*BlogPost, error) {
	sections := strings.SplitN(string(content), "---", 2)
	if len(sections) < 2 {
		return nil, errors.New("Invalid markdown file")
	}

	fileMetadata := sections[0]
	fileContent := sections[1]

	metadata := parseMetadata(fileMetadata)
	htmlContent := mdToHTML([]byte(fileContent))
	headers := extractHeaders([]byte(fileContent))

	return &BlogPost{
		Title:                   metadata.Title,
		Slug:                    metadata.Slug,
		Content:                 template.HTML(htmlContent),
		Description:             metadata.Description,
		Order:                   metadata.Order,
		Headers:                 headers,
		MetaDescription:         metadata.MetaDescription,
		MetaPropertyTitle:       metadata.MetaPropertyTitle,
		MetaPropertyDescription: metadata.MetaPropertyDescription,
		MetaOgURL:               metadata.MetaOgURL,
	}, nil
}

type Metadata struct {
	Title                   string
	Slug                    string
	Parent                  string
	Order                   int
	Description             string
	MetaDescription         string
	MetaPropertyTitle       string
	MetaPropertyDescription string
	MetaOgURL               string
}

// parses metadata from a markdown file
func parseMetadata(metadata string) *Metadata {
	re := regexp.MustCompile(`(?m)^(\w+):\s*(.+)`)
	matches := re.FindAllStringSubmatch(metadata, -1)

	metaMap := Metadata{}
	for _, match := range matches {
		if len(match) == 3 {
			switch match[1] {
			case "Title":
				metaMap.Title = match[2]
			case "Slug":
				metaMap.Slug = match[2]
			case "Parent":
				metaMap.Parent = match[2]
			case "Order":
				metaMap.Order, _ = strconv.Atoi(match[2])
			case "Description":
				metaMap.Description = match[2]
			case "MetaDescription":
				metaMap.MetaDescription = match[2]
			case "MetaPropertyTitle":
				metaMap.MetaPropertyTitle = match[2]
			case "MetaPropertyDescription":
				metaMap.MetaPropertyDescription = match[2]
			case "MetaOgURL":
				metaMap.MetaOgURL = match[2]
			}
		}
	}

	return &metaMap
}
