---
title: Go is changing how I write JavaScript
summary: Learning Go is affecting how I write JavaScript, and I think it's a good thing.
published: 2024-02-20
updated: 2024-02-20
draft: false
tags:
  - JS
  - Go
  - Programming
  - Result type
  - Try/catch
---

One of my favorite libraries is Angus Croll's [just](https://anguscroll.com/just/). If you look at the source for pretty much any of the modules in that library, what you'll find is straightforward, easy to read, easy to reason about code. It's nothing fancy, and that's what makes it so good.

I've been learning Go for the last month or so, and most of the code I've encountered is similar to Croll's JS code. It's almost always very easy to figure out what's happening. In fact, [simplicity is one of the hallmarks of the language](https://www.programmingtalks.org/talk/dotgo-2015-rob-pike-simplicity-is-complicated) and that ethos of simplicity is reflected in the culture of people who write Go. Reading and writing in the language is starting to change my approach to JavaScript, and I think it's helping me to be a better coder.

I like JavaScript. It's the first language I learned to code with and I generally feel comfortable with all of its quirks. It's not the only language I'm familiar with. I just finished a CS degree where the primary language was C++ (_shudder_) and one of my biggest projects was written in Python. I even spent a couple of semesters writing Java. All of those languages can be written with simplicity and clarity in mind, but none of them are engineered for simplicity the way Go is.

While learning Go, I worked on a JS project and found myself adjusting my approach to how I write code and structure applications. Here are a couple of examples of how my habits are changing after learning a little Go.

### Try/catch

Try/catch blocks help to guard against unexpected errors and exceptions. The idea is to `try` a potentially risky procedure, `catch` any errors that arise, and move on. It's a good idea in theory but I've never felt fully comfortable with using it. I always feel like my code gets trapped in the `try` portion of the try/catch block.

Here's a try/catch that'll be familiar to anyone who's coded JavaScript:

```js
function getMyMomsFavoriteWebsite(siteUrl) {
  try {
    const response = await fetch(siteUrl);

    // Wait, there's a potential error when parsing the response
    // I should try/catch that as well
    try {
      const data = response.json()
      // dear lord, I'm in try/catch hell
    } catch (error) {
      // do something if an error pops up
      console.error("lazy error handling, etc", error);
    }
    // okay, both `response` and `data` live in this scope
    // and I can never leave...
  } catch (error) {
    // do something if an error pops up
    console.error("No, don't just log it to the console.", error);
  }
}
```

There's another, simpler approach that's often taught in JS tutorials. I'm not really happy with it either.

### The simple way (AKA: how they do it in tutorials)
Here's the thing about the usual `fetch`, `ok`, `.json()` flow of the next example — there's loads of potential for errors.

```js
function getMyMomsFavoriteWebsite(siteUrl) {
  const response = await fetch(siteUrl);

  if (!response.ok) {
    console.error(`There was an error fetching the resource at ${siteUrl}`);
  }

  const data = await response.json();
  // okay, all sorts of errors can happen here...
  return data;
}
```

I prefer Go's error handling.

### The Go way

```go
// we need to define a type first (usually)
type FavoriteSite struct {
  Url string
  Title string
  Author string
  // ...etc.
}

func GetMyMomsFavoriteSite(siteUrl string) (*FavoriteSite, error) {
  client := http.Client{}
  res, err := client.Get("https://mymomsfavoritesite.com")
  if err != nil {
    return "", fmt.Errorf("error getting your mom's favorite site %v", err)
    // okay, error handled, we can just move on
  }

  var fav := FavoriteSite{}

  defer res.Body().Close()
  err := json.Unmarshal(res.Body, fav)

  if err != nil {
    return "", fmt.Errorf("error parsing the response body", err)
  }
  
  return &fav, nil
}
```

One part of this code is a little more complicated — we need to know the shape of the data the response's json is going to give us. There are other ways to do it, but that's a normal approach. Notice the `if err != nil` part? In Go, errors are values, so you can check whether a function returned a `nil` error and respond accordingly. After getting used to handling errors the Go way, my approach to JS started to change. First, I started using try/catch differently.

### Try/catch — new approach 1
```js
function getMyMomsFavoriteSite(siteUrl) {
  // declare response here to use later
  let response;
  try {
    response = await fetch(siteUrl);
  } catch(error) {
    console.error(error)
    // we'll return something else in the next example
    return {}
  }

  let data;
  try {
    // oh, look! response is available outside it's try block!
    data = await response.json();
  } catch (error) {
    console.error(error)
    return {}
  }

  return data
}
```

To me, the above example is a huge improvement over the approach taken to many `fetch` tutorials you'll see. Declaring the `response` and `data` variables outside the try/catch blocks means I have access to them outside of their block's scope. The try/catch block is now used strictly for exceptions and errors. But what about that `return {}`? I used an empty object in the example because the function is expected to return a JS object. I'm coming around to another approach that's inspired by Rust's [Result](https://doc.rust-lang.org/rust-by-example/std/result.html).

### Adding a `result` object

The `Result` in Rust is basically a way to return both the value that should be returned from a function and an indication of whether the function's operation was successful. This isn't really necessary in Go because you can define multiple return values, and a lot of functions will return a result and an error. There's no such thing in Rust. There's also no such thing in JavaScript. We can still use objects to represent the result of a function that might need to be an error.

It's as simple as:

```js
return {ok: true, value: "Hey, I'm the valid result of a function!"}
```

Or:

```js
return {ok: false, error: "Nope! Didn't work. Enjoy this cryptic message instead"}
```

The examples above don't take TypeScript into account, and [there](https://imhoff.blog/posts/using-results-in-typescript) are [several](https://typescript.wtf/blog/result-type) [good](https://lab.scub.net/understanding-result-pattern-in-typescript-e82934cea096) [examples](https://github.com/badrap/result) of a `Result` type that ensures type safety. But I've been writing these examples in JS to keep this post simple. If you know and love TypeScript, feel free to adapt.

Here's the above example with object results:

```js
function getMyMomsFavoriteWebsite(siteUrl) {
  let response;
  try {
    response = await fetch(siteUrl);
  } catch(err) {
    return {ok: false, error: `error fetching site at ${siteUrl}: ${err}`};
  }

  let data;
  try {
    data = await response.json();
  } catch(err) {
    return {ok: false, error: `error parsing response to json: ${err}`};
  }

  return {ok: true, value: data};
}
```

I haven't used the new approach for long, but it's already helping me think about how I approach functions, especially when they're within the lower layers of an application where I don't want to handle errors.

What do you think? Do you know of a better way to approach try/catch and error handling in JavaScript? Let me know at [@trainingmontage@hachyderm.io](https://hachyderm.io/@trainingmontage).