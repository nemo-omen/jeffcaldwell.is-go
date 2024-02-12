---
title: Todo
updated: 2024-02-12T14:51:30-06:00
---

- [x] Finish /about page
- [x] Add Todo Page
  - [ ] Create a [github.com/gomarkdown](https://github.com/gomarkdown/markdown) hook to render todo lists as lists with checkboxes
- [x] Syntax highlighting
  - [x] Client side with Prism or HLJS (first)
  - [ ] Custom gomarkdown renderer with [github.com/alecthomas/chroma](https://github.com/alecthomas/chroma) (later)
- [x] Find a proper host -> DigitalOcean App
- [x] Find a deployment strategy -> DigitalOcean App
- [x] Deploy!
- [ ] Adjust header styles for mobile (make sub-header a column, enlarge touch targets)
- [ ] Post publishing graph - see [https://github.com/johnwargo/eleventy-plugin-post-stats](https://github.com/johnwargo/eleventy-plugin-post-stats)
- [ ] Sitemap
- [ ] Finish /now page (?)
- [ ] Refactor `PostService` - I need file handling and markdown parsing for other types of content, so they should be handled by their own services.
  - [ ] `FileService`
  - [ ] `MarkdownService`
- [ ] Add projects page
  - [ ] `ProjectService`
- [ ] Analytics
  - [ ] [How bear does analytics with CSS](https://herman.bearblog.dev/how-bear-does-analytics-with-css/)
- [ ] Rewrite server with std lib `net/http`
  - [ ] Use `context.Context` to pass values to .templ components.