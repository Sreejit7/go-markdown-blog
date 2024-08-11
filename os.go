package main

import (
	"os"
	"strings"
)

func loadMarkdownFiles(dir string) ([]BlogPost, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var posts []BlogPost
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".md") {
			// read file
			content, err := os.ReadFile(dir + "/" + file.Name())
			if err != nil {
				return nil, err
			}

			// parse markdown file
			post, err := parseMarkdownFile(content)
			if err != nil {
				return nil, err
			}
			posts = append(posts, *post)
		}
	}

	return posts, nil
}
