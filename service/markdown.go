package service

import (
	"bytes"
	"fmt"

	"github.com/yuin/goldmark"
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
			html.WithHardWraps(),
			html.WithXHTML(),
		),
	)

	if err := md.Convert(content, &buf); err != nil {
		return "", fmt.Errorf("error parsing md content: %v", err)
	}

	return buf.String(), nil
}
