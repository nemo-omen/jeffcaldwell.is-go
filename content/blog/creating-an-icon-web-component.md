---
title: Creating an Icon Web Component
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

Keep in mind — I know next to nothing about working with web components. This is as instructional for me as I hope it will be for you, reader. In my ideal world I would be able to define and render the component on the server. It should probably not just be stringifying the component's markup, though. One of the most powerful things about Svelte, React, etc. components is the ability to pass data into them. I just want to be able to do that on the server without relying on an external library.

### The baseline

To create a web component, you need two parts — the JavaScript declaration, which is a class that extends `HTMLElement` or some other element, and you need to register the custom element for use somewhere in the DOM. That declaration and registration, at their simplest, look like this:

```js
/* This class extends another class, 
 * so don't forget to call super(), 
 * or it won't work.
 */
class Icon extends HTMLElement {
  constructor() {
    super();
  }
}

/* Register a custom element by providing the name,
 * which must have a dash, and constructor — the 
 * name of the class you just created.
 */
customElements.define('jcis-icon', Icon);
```

The above would technically allow us to plop `<jcis-icon></jcis-icon>` right into an HTML document. It doesn't do anything and there's no HTML inside of it, but it'll be there in the DOM.

So, how do we get it to actually do something?

The easiest way is to just put everything together in `Icon`'s constructor. Here's version 1.

```js
class IconV1 extends HTMLElement {
  constructor() {
    super();
    /* We're mimicking Amelia's pattern of predefining 
     * paths. We'll try out other ways later.
     */
    this.paths = {
      "mastodon": [
        "M3.019 12.0075C2.98744 10.7478 3.00692 9.5598 3.00692 8.56644C3.00692 4.22767 5.84954 2.95597 5.84954 2.95597C7.28286 2.29767 9.74238 2.0209 12.2993 2H12.3621C14.919 2.0209 17.3801 2.29767 18.8134 2.95597C18.8134 2.95597 21.656 4.22767 21.656 8.56644C21.656 8.56644 21.6916 11.7674 21.2596 13.9898C20.9852 15.4007 18.8034 16.9446 16.2974 17.2438C14.9906 17.3999 13.7042 17.5431 12.3322 17.4802C10.0885 17.3775 8.31815 16.9446 8.31815 16.9446C8.31815 17.1631 8.33166 17.3711 8.35853 17.5655C8.44182 18.1978 8.65659 18.6604 8.96296 19C9.72944 19.8497 11.0692 19.9301 12.3577 19.9743C14.178 20.0366 15.7986 19.5254 15.7986 19.5254L15.8735 21.1712C15.8735 21.1712 14.6003 21.8548 12.3322 21.9805C11.0815 22.0493 9.52858 21.9491 7.71969 21.4704C6.18802 21.065 5.15153 20.1804 4.45091 19C3.35714 17.1573 3.08191 14.5938 3.019 12.0075ZM6.31815 16.9446V14.3967L8.79316 15.0018C8.8405 15.0134 8.95098 15.0383 9.11692 15.0723C9.40521 15.1313 9.73416 15.1908 10.0959 15.2467C10.8485 15.3628 11.6341 15.4462 12.4237 15.4823C13.4425 15.529 14.3249 15.4652 16.0603 15.2579C17.7233 15.0594 19.208 14.0622 19.2963 13.6082C19.3783 13.1861 19.4472 12.6858 19.5021 12.1261C19.5714 11.4205 19.6155 10.6558 19.6388 9.88068C19.654 9.37026 19.6582 8.93648 19.6564 8.62452L19.656 8.56644C19.656 7.1368 19.2873 6.12756 18.6928 5.40793C18.5008 5.17553 18.3004 4.99408 18.1087 4.85958C18.0183 4.79617 17.9737 4.77136 17.9787 4.77345C16.9662 4.30844 14.8859 4.02069 12.3621 3.99993H12.3156C9.77596 4.02069 7.6969 4.30836 6.66627 4.78161C6.68919 4.77136 6.64459 4.79617 6.55423 4.85958C6.36257 4.99408 6.16214 5.17553 5.97016 5.40793C5.37568 6.12756 5.00692 7.1368 5.00692 8.56644C5.00692 8.7976 5.00628 8.96339 5.00392 9.44137C4.9981 10.6238 5.00004 11.2256 5.01841 11.9589C5.07185 14.156 5.2822 15.7941 5.71797 17C5.93023 17.5874 6.19005 18.0709 6.49741 18.4507C6.37791 18.0162 6.31815 17.5142 6.31815 16.9446ZM8.08576 6.37135C8.71735 6.37135 9.22924 6.88324 9.22924 7.51482C9.22924 8.14626 8.71735 8.6583 8.08576 8.6583C7.45432 8.6583 6.94229 8.14626 6.94229 7.51482C6.94229 6.88324 7.45432 6.37135 8.08576 6.37135Z",
      ],
    };

    /* @See: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/template */
    const template = document.createElement('template');

    /* We're grabbing the name attribute from
     * the markup, which looks like:
     * <jcis-icon name="mastodon"></jcis-icon>
     */
    const name = this.getAttribute('name');

    let iconPaths;

    // Guard against empty attribute
    if(!name) {
      iconPaths = [];
    } else {
      iconPaths = this.paths[name];
    }

    /* Define styles and markup */
    template.innerHTML = `<style>
    .icon {
      fill: currentColor;
      width: 1em;
      height: 1em;
      overflow: visible;
      transition: all 0.2s ease-in-out;
    }
    </style>
    <svg
      class="icon"
      viewbox="0 0 24 24"
      fill-rule="evenodd"
      clip-rule="evenodd"
      >
      ${iconPaths.map((path) => `<path d="${path}"></path>`)}
    </svg>`;

    /* Clone the template, and slap it directly
     * into this custom element.
     */
    const templateClone = template.content.cloneNode(true);
    this.appendChild(templateClone);
  }
}

customElements.define("jcis-icon-v1", IconV1);
```

If you put `<jcis-icon-v1 name="mastodon"></jcis-icon-v1>` in your HTML, it'll show up just fine.

<figure>
  <jcis-icon-v1 name="mastodon"></jcis-icon-v1>
  <figcaption>IconV1. Go on, inspect it!</figcaption>
</figure>

Okay, so this works, but there are some problems. First, when we create everything in the constructor like we just did, we losr one of the best features of web components — style encapsulation. If an external stylesheet happens to use the same class name things can get super confusing and irritating for whoever's using the component. There are other issues with interactivity that I won't go over here, you can check out [some reference post]() to read up on it.

Web components come with a few built-in lifecycle methods that get called at different points of the components existence on the page. They also come with something called a [shadow dom](https://developer.mozilla.org/en-US/docs/Web/API/Web_components/Using_shadow_DOM) that helps us with encapsulation.

Let's encorporate those into Icon v2.

```js
class IconV2 extends HTMLElement {
  constructor() {
    super();
    this.paths = {
      "mastodon": [
        "M3.019 12.0075C2.98744 10.7478 3.00692 9.5598 3.00692 8.56644C3.00692 4.22767 5.84954 2.95597 5.84954 2.95597C7.28286 2.29767 9.74238 2.0209 12.2993 2H12.3621C14.919 2.0209 17.3801 2.29767 18.8134 2.95597C18.8134 2.95597 21.656 4.22767 21.656 8.56644C21.656 8.56644 21.6916 11.7674 21.2596 13.9898C20.9852 15.4007 18.8034 16.9446 16.2974 17.2438C14.9906 17.3999 13.7042 17.5431 12.3322 17.4802C10.0885 17.3775 8.31815 16.9446 8.31815 16.9446C8.31815 17.1631 8.33166 17.3711 8.35853 17.5655C8.44182 18.1978 8.65659 18.6604 8.96296 19C9.72944 19.8497 11.0692 19.9301 12.3577 19.9743C14.178 20.0366 15.7986 19.5254 15.7986 19.5254L15.8735 21.1712C15.8735 21.1712 14.6003 21.8548 12.3322 21.9805C11.0815 22.0493 9.52858 21.9491 7.71969 21.4704C6.18802 21.065 5.15153 20.1804 4.45091 19C3.35714 17.1573 3.08191 14.5938 3.019 12.0075ZM6.31815 16.9446V14.3967L8.79316 15.0018C8.8405 15.0134 8.95098 15.0383 9.11692 15.0723C9.40521 15.1313 9.73416 15.1908 10.0959 15.2467C10.8485 15.3628 11.6341 15.4462 12.4237 15.4823C13.4425 15.529 14.3249 15.4652 16.0603 15.2579C17.7233 15.0594 19.208 14.0622 19.2963 13.6082C19.3783 13.1861 19.4472 12.6858 19.5021 12.1261C19.5714 11.4205 19.6155 10.6558 19.6388 9.88068C19.654 9.37026 19.6582 8.93648 19.6564 8.62452L19.656 8.56644C19.656 7.1368 19.2873 6.12756 18.6928 5.40793C18.5008 5.17553 18.3004 4.99408 18.1087 4.85958C18.0183 4.79617 17.9737 4.77136 17.9787 4.77345C16.9662 4.30844 14.8859 4.02069 12.3621 3.99993H12.3156C9.77596 4.02069 7.6969 4.30836 6.66627 4.78161C6.68919 4.77136 6.64459 4.79617 6.55423 4.85958C6.36257 4.99408 6.16214 5.17553 5.97016 5.40793C5.37568 6.12756 5.00692 7.1368 5.00692 8.56644C5.00692 8.7976 5.00628 8.96339 5.00392 9.44137C4.9981 10.6238 5.00004 11.2256 5.01841 11.9589C5.07185 14.156 5.2822 15.7941 5.71797 17C5.93023 17.5874 6.19005 18.0709 6.49741 18.4507C6.37791 18.0162 6.31815 17.5142 6.31815 16.9446ZM8.08576 6.37135C8.71735 6.37135 9.22924 6.88324 9.22924 7.51482C9.22924 8.14626 8.71735 8.6583 8.08576 8.6583C7.45432 8.6583 6.94229 8.14626 6.94229 7.51482C6.94229 6.88324 7.45432 6.37135 8.08576 6.37135Z",
      ],
    };
  }

  connectedCallback() {
    /* Here's how we access the shadow dom */
    this.shadow = this.attachShadow({mode: "open"});

    /* Most of this is the same as v1 */
    const template = document.createElement('template');
    const name = this.getAttribute('name');
    let iconPaths;

    if(!name){
      iconPaths = [];
    } else {
      iconPaths = this.paths[name];
    }

    template.innerHTML = `<style>
    .icon {
      fill: currentColor;
      width: 1em;
      height: 1em;
      overflow: visible;
      transition: all 0.2s ease-in-out;
    }
    </style>
    <svg
      class="icon"
      viewbox="0 0 24 24"
      fill-rule="evenodd"
      clip-rule="evenodd"
      >
      ${iconPaths.map((path) => `<path d="${path}"></path>`)}
    </svg>`;

    const templateClone = template.content.cloneNode(true);

    /* 
     * Instead of appending the template to `this`
     * we append it to the shadow root.
     */
    this.shadow.appendChild(templateClone);
  }
}

customElements.define("jcis-icon-v2", IconV2);
```
Here's `IconV2`:

<figure>
  <jcis-icon-v2 name="mastodon"></jcis-icon-v2>
  <figcaption>Icon v2. It looks pretty much the same as v1, but has some benefits.</figcaption>
</figure>

The encapsulation provided by using shadow dom makes the component more reusable. Here's IconV1 and IconV2 together. The checkbox on this figure will set `.icon`'s fill color to `tomato`. Notice anything different about V2?


<style>
  #figure-color-demo {
    gap: var(--space-xs);

    & span {
      line-height: 1;
      display: flex;
      align-items: center;
      justify-content: center;
      gap: var(--space-xs);
    }
  }

  #figure-color-demo >*:not(label) {
    border: 2px solid var(--secondary-alpha);
    padding: var(--space-2xs);
    font-size: var(--step-2);
  }

  #figure-color-demo:has(input:checked) {
    
    & .icon {
      fill: tomato !important;
    }
  }
</style>

<figure id="figure-color-demo">
  <section>
  <span>V1:
  <jcis-icon-v1 name="mastodon"></jcis-icon-v1>
  </span>
  <span>V2:
  <jcis-icon-v2 name="mastodon"></jcis-icon-v2>
  </span>
  </section>
  <label>Click to change `.icon`'s fill to tomato <input type="checkbox" class="make-tomato" name="color-change-input" /></label>
</figure>

There's a lot more we can do with web components that I won't get into in this post. I will, however, talk about `<slot>`s, because they offer a way to add paths to the icon without using predefined path arrays or attributes.

Continue with this answer: [stackoverflow](https://stackoverflow.com/a/68159761)

<jcis-icon-v3>
  <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" slot="paths"><path d="M3.019 12.0075C2.98744 10.7478 3.00692 9.5598 3.00692 8.56644C3.00692 4.22767 5.84954 2.95597 5.84954 2.95597C7.28286 2.29767 9.74238 2.0209 12.2993 2H12.3621C14.919 2.0209 17.3801 2.29767 18.8134 2.95597C18.8134 2.95597 21.656 4.22767 21.656 8.56644C21.656 8.56644 21.6916 11.7674 21.2596 13.9898C20.9852 15.4007 18.8034 16.9446 16.2974 17.2438C14.9906 17.3999 13.7042 17.5431 12.3322 17.4802C10.0885 17.3775 8.31815 16.9446 8.31815 16.9446C8.31815 17.1631 8.33166 17.3711 8.35853 17.5655C8.44182 18.1978 8.65659 18.6604 8.96296 19C9.72944 19.8497 11.0692 19.9301 12.3577 19.9743C14.178 20.0366 15.7986 19.5254 15.7986 19.5254L15.8735 21.1712C15.8735 21.1712 14.6003 21.8548 12.3322 21.9805C11.0815 22.0493 9.52858 21.9491 7.71969 21.4704C6.18802 21.065 5.15153 20.1804 4.45091 19C3.35714 17.1573 3.08191 14.5938 3.019 12.0075ZM6.31815 16.9446V14.3967L8.79316 15.0018C8.8405 15.0134 8.95098 15.0383 9.11692 15.0723C9.40521 15.1313 9.73416 15.1908 10.0959 15.2467C10.8485 15.3628 11.6341 15.4462 12.4237 15.4823C13.4425 15.529 14.3249 15.4652 16.0603 15.2579C17.7233 15.0594 19.208 14.0622 19.2963 13.6082C19.3783 13.1861 19.4472 12.6858 19.5021 12.1261C19.5714 11.4205 19.6155 10.6558 19.6388 9.88068C19.654 9.37026 19.6582 8.93648 19.6564 8.62452L19.656 8.56644C19.656 7.1368 19.2873 6.12756 18.6928 5.40793C18.5008 5.17553 18.3004 4.99408 18.1087 4.85958C18.0183 4.79617 17.9737 4.77136 17.9787 4.77345C16.9662 4.30844 14.8859 4.02069 12.3621 3.99993H12.3156C9.77596 4.02069 7.6969 4.30836 6.66627 4.78161C6.68919 4.77136 6.64459 4.79617 6.55423 4.85958C6.36257 4.99408 6.16214 5.17553 5.97016 5.40793C5.37568 6.12756 5.00692 7.1368 5.00692 8.56644C5.00692 8.7976 5.00628 8.96339 5.00392 9.44137C4.9981 10.6238 5.00004 11.2256 5.01841 11.9589C5.07185 14.156 5.2822 15.7941 5.71797 17C5.93023 17.5874 6.19005 18.0709 6.49741 18.4507C6.37791 18.0162 6.31815 17.5142 6.31815 16.9446ZM8.08576 6.37135C8.71735 6.37135 9.22924 6.88324 9.22924 7.51482C9.22924 8.14626 8.71735 8.6583 8.08576 8.6583C7.45432 8.6583 6.94229 8.14626 6.94229 7.51482C6.94229 6.88324 7.45432 6.37135 8.08576 6.37135Z"></path></svg>
</jcis-icon-v3>