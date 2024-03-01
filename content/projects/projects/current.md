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

Things started to slow down a bit when making multiple requests to update feeds. Even with a more dynamic frontend, updates began to significantly affect the experience. One way to handle this would have been to use JavaScript's `Promise.all()` or `Promise.allSettled()`. Another way would have been to create a streaming connection with server-sent events to stream back feed items as they're retrieved and processed. Instead, I decided to start over using Go for the backend.

The process of rewriting the application in a language I'm relatively unfamiliar with has led me to rethink some of the architectural decisions I made while working with TypeScript. There's something about tackling a problem using an unfamiliar tool that helps reveal architectural issues that are separate from the conventions of a language. Also, because this is a learning project, I get to do what I want, and I want to learn Go.
