package server

import (
	"bytes"
	"embed"
	"html/template"

	"github.com/vcokltfre/snippy/server/database"
)

type basePage struct {
	Title   string
	Content template.HTML
}

type snippetPage struct {
	Language string
	Content  string
	ID       string
}

//go:embed web
var web embed.FS

var (
	base    = template.Must(template.ParseFS(web, "web/base.html"))
	index   = template.Must(template.ParseFS(web, "web/index.html"))
	snippet = template.Must(template.ParseFS(web, "web/snippet.html"))
)

func renderBase(title, content string) string {
	var buf bytes.Buffer
	base.Execute(&buf, basePage{
		Title:   title,
		Content: template.HTML(content),
	})
	return buf.String()
}

func RenderIndex(snippets []database.Snippet) string {
	for i := 0; i < len(snippets); i++ {
		snippets[i].Content = Shorten(snippets[i].Content)
	}

	var buf bytes.Buffer
	index.Execute(&buf, snippets)
	return renderBase("Home", buf.String())
}

func RenderSnippet(language, content, id string) string {
	var buf bytes.Buffer
	snippet.Execute(&buf, snippetPage{
		Language: language,
		Content:  content,
		ID:       id,
	})
	return renderBase("Snippy", buf.String())
}
