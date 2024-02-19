---
title: Using Echo's Context with Templ components
summary: How to use Echo's echo.Context as ctx values in Templ components
published: 2024-02-19
updated: 2024-02-19
draft: false
tags:
  - Go
  - Golang
  - Echo
  - Templ
  - context
---

When I was setting up the handlers and views for this site, I wanted to be able to pass a `currentPath` value to the site's [Templ](https://templ.guide/) header component so I could use it to highlight the current page link. According to Templ's [documentation](https://templ.guide/syntax-and-usage/context), you can pass values to components using Go's `context.Context`. The process is simple enough if you're using standard context, and I thought I would be able to easily use [Echo's](https://echo.labstack.com/docs/context) `echo.Context` interchangeably.

I was very wrong.

The problem is that `echo.Context` and `context.Context` are completely different things. Echo's `Context` carries references to all sorts of useful things about the current http request, but it's not a standard Context type. It does, however, carry with it the standard `Request.Context`, which is. So, the trick to passing context to Templ components in Echo is to modify the request's context. We can do that with a middleware.

### Modifying request.Context with middleware

While I was searching for a method to set a `context.Context` value, I came upon [this answer](https://stackoverflow.com/a/69331251) on StackOverflow.  It's a middleware that reimplements the `Get` and `Set` methods on `echo.Context` to update `request.Context` whenever they're called. It's relatively straightforward to add this middleware to your app, so let's put it together.

#### 1. Extend echo.Context

First, we need to implement our own `Get` and `Set` methods so we can modify `request.Context` whenever we update `echo.Context`.

```go
package custommiddleware
// extend echo.Context
type contextValue {
  echo.Context
}

func (c contextValue) Get(key string) interface{} {
  // grab value from echo.Context
  val := c.Context.Get(key)
  // if it's not nil, return it
  if val != nil {
    return val
  }
  // otherwise, return Request.Context
  return c.Request().Context().Value(key)
}

func (c contextValue) Set(key string, val interface{}) {
  // we're replacing the whole Request in echo.Context
  // with a copied request that has the updated context value
  c.SetRequest(
    c.Request().WithContext(
      context.WithValue(c.Request().Context(), key, value)
    )
  )
}
```

Okay, we have an extended context with custom `Get` and `Set` methods. Now we just need a middleware that will run everything on every request.

#### 2. The middleware

Middlewares are pretty simple. They're functions that take a `HandlerFunc` as an argument and return a `HandlerFunc`. They modify the request and pass it along to the next handler. In Echo, these handlers carry the `echo.Context` with them. That means we can just swap out our new `contextValue` before moving on to the next handler.

```go
package custommiddleware

func ContextValueMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
  // this is just an echo.HandlerFunc
	return func(c echo.Context) error {
    // instead of passing next(c) as you usually would,
    // you return it with the extended version
		return next(contextValue{c})
	}
}
```

#### 3. Adding the new middleware

You use this middleware just like you would with any Echo app.

```go
app := echo.New()
// ... routes, etc.
app.Use(custommiddleware.ContextValueMiddleware)
// ... more middleware, etc.
```

That's pretty much it! Now I can update the current path context value and use it in Templ components without the need for ugly prop drilling. Here's how I did it:

```go
// make another middleware that sets `currentPath`
package custommiddleware

func SetCurrentPath(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Set("currentPath", c.Request().URL.Path)
		return next(c)
	}
}
```

```go
// my templ component can now access ctx.CurrentPath
package component

// GetCurrentPath ensures that we're always
// returning a string from `ctx`
func GetCurrentPath(c context.Context) string {
	currentPath := c.Value("currentPath")
	if currentPath != nil {
		return currentPath.(string)
	}
	return ""
}


templ Header() {
  <nav>
    <ul>
      <li>
        if GetCurrentPath(ctx) == "/" {
          <a href="/" class="current nav-link">Home</a>
        } else {
          <a href="/" class="nav-link">Home</a>
        }
      </li>
      <!-- etc., etc. -->
    </ul>
  </nav>
}
```

It's a super simple process that greatly simplifies passing values to my components. I hope it helps you out in your projects. Be sure to give [eudore](https://stackoverflow.com/a/69331251) some props over on StackOverflow.