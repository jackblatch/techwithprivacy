package routes

import (
	"fmt"
	"os"
	"techwithprivacy/web/components"
	"techwithprivacy/web/pages"

	"github.com/a-h/templ"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
	"github.com/microcosm-cc/bluemonday"
)

func toHTML(md []byte) []byte {
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(md)
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)
	unsafeHtml := markdown.Render(doc, renderer)
	return bluemonday.UGCPolicy().SanitizeBytes(unsafeHtml)
}
func GetIndex() (templ.Component, error) {
	markdownContent, err := os.ReadFile("content.md")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}

	html := toHTML(markdownContent)
	component := pages.Index(string(html))
	page := components.RootLayout("Tech with Privacy", component)
	return page, nil
}
