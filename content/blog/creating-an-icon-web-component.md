---
title: Creating an Icon Web Component(?)
summary: Icons are ubiquitous in web development and there are plenty of great examples of icon components for various frameworks. What would building an icon web component from scratch look like?
published: 2024-02-11T22:33:44-06:00
updated: 2024-02-11T22:33:44-06:00
draft: true
tags:
  - Web Components
  - Icons
scripts:
  - creating-an-icon-web-component/icon.js
---

One of my favorite component recipes is the [Icon](https://svelte.recipes/components/icon) recipe shared by [Amelia Wattenberger](https://wattenberger.com/) on her super helpful [svelte.recipes](https://svelte.recipes/) project. It's the kind of component pattern that's come in very handy over and over again in my projects. Each time I implement it, whether in a `.svelte`, `.astro`, `.tsx`, or even in a Django template, I sort of wish I could just use a reusable web component instead. It's a pretty simple thing to put together, though, so I guess I've never really _needed_ to make something more portable. But, hey, it's time to change that! I want to learn more about web components and I need something to write about, so let's think about what we'd want one to look like.

<p class="highlight">
  Amelia's one of the most talented devs around and she's always working on something really cool. Definitely take a minute to <a href="https://wattenberger.com" target="_blank">check out her site</a>.
</p>

### What I want in an Icon web component

- Ability to define a `paths` or `pathsByName` object that defines an icon's paths.
- Ability to include minimal styles that allow the `fill` to be defined by the font color of the icon's parent element
- Server rendering would be pretty sweet (is this declarative shadow DOM?)

Keep in mind — I know next to nothing about working with web components. This is as instructional for me as I hope it will be for you, reader. In my ideal world I would be able to define and render the component on the server. It should probably not just be stringifying the component's markup, though. One of the most powerful things about Svelte, React, etc. components is the ability to pass data into them. I just want to be able to do that on the server without relying on an external library. I may not be able to do that. I really don't know.

### The baseline

To create a web component, you need two parts — the JavaScript declaration, which is a class that extends `HTMLElement` or some other element, and you need to register the custom element for use somewhere in the DOM. That declaration and registration, at their simplest, look like this:

```js
/* This class extends another class, so don't forget super() */
class Icon extends HTMLElement {
  constructor() {
    super();
  }
}

/* Register a custom element by providing the name,
 * which must have a dash, and the new element's class name.
*/
customElements.define('jcis-icon', Icon);
```

The above would technically allow us to plop `<jcis-icon></jcis-icon>` right into an HTML document. It doesn't do anything, but it'll be there and it'll be registered as a valid custom element.

So, how do we get it to actually do something?

Here's where things get a little clunky. You can see that `Icon` is a class that extends `HTMLElement`. That means it comes with methods for setting up the custom element, among other things.

<p style="font-size: var(--step-1); color: var(--foreground);">
  <jcis-icon name="mastodon"></jcis-icon>
</p>