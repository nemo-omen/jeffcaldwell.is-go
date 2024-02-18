---
title: Using Echo's Context with Templ components
summary: How to use Echo's echo.Context as ctx values in Templ components
published: 2024-02-17
updated: 2024-02-17
draft: true
tags:
  - Go
  - Golang
  - Echo
  - Templ
  - context
---

When I was setting up the handlers and views for this site, I wanted to be able to pass a `currentPath` value to the site's [Templ](https://templ.guide/) header component so I could use it to highlight the current page link. According to Templ's [documentation](https://templ.guide/syntax-and-usage/context), you can pass values to components using Go's `context.Context`. The process is simple enough if you're using standard context, and I thought I would be able to easily use [Echo's](https://echo.labstack.com/docs/context) `echo.Context` interchangeably.

I was very wrong.

The problem is that `echo.Context` and `context.Context` are completely different things. Echo's `Context` carries references to all sorts of useful things about the current http request, while Go's `Context` is 