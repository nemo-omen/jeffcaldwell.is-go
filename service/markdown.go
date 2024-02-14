package service

import (
	"bytes"
	"fmt"

	"github.com/yuin/goldmark"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
	"go.abhg.dev/goldmark/toc"
)

type MarkdownService struct{}

type CustomExtender struct {
	Extender toc.Extender
}

func (s MarkdownService) ParseMarkdownContent(content []byte) (string, error) {
	var buf bytes.Buffer

	md := goldmark.New(
		goldmark.WithExtensions(
			extension.GFM,
			meta.Meta,
			// highlighting.NewHighlighting(
			// highlighting.WithStyle("dracula"),
			// 	highlighting.WithFormatOptions(
			// 		chromahtml.WithLineNumbers(true),
			// 		chromahtml.WithClasses(true),
			// 	),
			// ),
		),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			html.WithUnsafe(),
			html.WithHardWraps(),
			html.WithXHTML(),
		),
	)

	if err := md.Convert(content, &buf); err != nil {
		return "", fmt.Errorf("error parsing md content: %v", err)
	}

	return buf.String(), nil
}

func (s MarkdownService) GetFrontmatter(content []byte) (map[string]interface{}, error) {
	md := goldmark.New(
		goldmark.WithExtensions(meta.Meta),
	)

	var buf bytes.Buffer
	var frontmatter map[string]interface{}
	context := parser.NewContext()

	if err := md.Convert(content, &buf, parser.WithContext(context)); err != nil {
		return frontmatter, fmt.Errorf("error getting frontmatter: %v", err)
	}

	frontmatter = meta.Get(context)

	return frontmatter, nil
}
