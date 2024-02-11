---
title: Creating a Markdown Editor with In-place Preview
summary: Markdown editors in the browser are nothing new — they're a rite of passage for intermediate frontend devs. Let's make one that has realtime rendering of the markdown to HTML.
published: 2024-02-07T22:33:28-06:00
updated: 2024-02-07T22:35:14-06:00
draft: true
tags:
  - Markdown
  - CSS
  - JS
scripts:
  - creating-a-realtime-markdown-editor/test.js
---
<test-element></test-element>

### What we're building

If you've used markdown editors you've probably noticed all but a few of them have a common split-pane interface — raw markdown on the left, rendered HTML on the right. There are a few, though, like Obsidian, that render the markdown directly in the editor. That's what I want to build.

I'd also like to build it independently of any frontend framework, so that leaves web components. I could also build a whole server and use something like HTMX. We'll see.

This post won't be a tutorial as much as it will be a walkthrough of my exploration. I'm not sure how I'll organize that exploration. Most of the time these sorts of things are just long documents with headings that the reader comes upon when they're already complete. I'd like to keep this kind of post open to ongoing updates.

### Some test code blocks

```js
console.log("This is fun.");
```

```go
package service

import (
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

type MarkdownService struct{}

func (s MarkdownService) ParseMarkdownContent(content []byte) (string, error) {
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(content)
	htmlFlags := html.CommonFlags | html.HrefTargetBlank | html.LazyLoadImages
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	return string(markdown.Render(doc, renderer)), nil
}
```