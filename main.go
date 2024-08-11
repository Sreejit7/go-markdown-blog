package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BlogPost struct {
	Title                   string
	Slug                    string
	Content                 template.HTML
	Description             string
	Order                   int
	Headers                 []string
	MetaDescription         string
	MetaPropertyTitle       string
	MetaPropertyDescription string
	MetaOgURL               string
}

type SidebarData struct {
	Categories []Category
}

type Category struct {
	Name  string
	Pages []BlogPost
	Order int
}

func main() {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	// load templates
	r.LoadHTMLGlob("templates/*")

	// static assets
	r.Static("/static", "./static")

	// register the sidebar template as a partial
	r.SetFuncMap(template.FuncMap{
		// "loadSidebar": func() SidebarData {
		// 	return sidebarData
		// },
		"dict": dict,
	})

	// load and parse markdown files
	posts, err := loadMarkdownFiles("./markdown")
	if err != nil {
		log.Fatal(err)
	}

	// home page

	// route for posts
	for _, post := range posts {
		if post.Slug != "" {
			r.GET("/"+post.Slug, func(ctx *gin.Context) {
				sidebarLinks := createSidebarLinks(post.Headers)
				ctx.HTML(http.StatusOK, "layout.html", gin.H{
					"Title":                   post.Title,
					"Content":                 post.Content,
					"Description":             post.Description,
					"Headers":                 post.Headers,
					"MetaDescription":         post.MetaDescription,
					"MetaPropertyTitle":       post.MetaPropertyTitle,
					"MetaPropertyDescription": post.MetaPropertyDescription,
					"MetaOgURL":               post.MetaOgURL,
					"SidebarLinks":            sidebarLinks,
				})
			})
		} else {
			log.Printf("Warning: Post titled '%s' has an empty slug and will not be accessible via a unique URL.\n", post.Title)
		}
	}

	// Handle 404
	r.NoRoute(func(ctx *gin.Context) {
		ctx.HTML(http.StatusNotFound, "404.html", gin.H{
			"Title": "Page Not Found",
		})
	})

	r.Run()
}
