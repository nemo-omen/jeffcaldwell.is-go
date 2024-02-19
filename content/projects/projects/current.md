---
name: Current
link: https://github.com/nemo-omen/current
summary: Current is an RSS reader geared toward journalists.
startDate: 2023-12-16
technologies:
  - Go
  - SqLite
  - HTMX
images:
  - src: "/current/featured-1.png"
    alt: ""
  - src: "/current/featured-2.png"
    alt: ""
---

Current is an RSS reader for anyone who wants more than an inbox for new posts. In addition to keeping track of the latest posts, users can annotate posts, create tags and collections, and create notes linked to multiple posts.

The goal of this project is to make an RSS reader that's useful for journalists, but the planned features should make it a great tool for anyone who wants more than basic functionality.

This is an ongoing project. In its first iteration, I used TypeScript with  [Hono](https://hono.dev/) server backed by an SqLite database. Rather than using a framework like React, Vue, or Svelte, I chose to use HTMX for interactive elements in the client. This allowed me to keep application logic simple — all state was kept on the server.

Things started to slow down a bit when making multiple requests to update feeds. Even with a more dynamic frontend, updates began to significantly affect the experience. A big part of the problem is the single-threaded nature of JS/TS. While the server could asynchronously make requests to external RSS feeds, there's no way to make those requests concurrently.

So, I'm in the middle of rewriting the backend in Go with the [Echo](https://echo.labstack.com/) framework. I'm sure that using Go rather than TypeScript on the server will not only decrease response speeds significantly, but will also give me the opportunity to refine the overall application architecture.
