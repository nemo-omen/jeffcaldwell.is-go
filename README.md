# Jeff Caldwell's site

Okay, so I had this mostly finished in [Astro](https://github.com/nemo-omen/jeffcaldwell.is) when I decided I really wanted to write the site in Go. Astro's nice — it's one of my favorite ways to make a website — but I started feeling a little constricted in what I could do when I tried to make vanilla web components outside of an Astro component. Plus, I'm tired of writing everything in JavaScript.

So, I'm starting again in Go with the [Echo](https://echo.labstack.com/) server framework, and I'm using [Templ](https://templ.guide/) for templating and components. Everything's similar enough to minimal Express and JSX that the move is pretty painless and allows me to ease into using Go.

## The plan

For now, I plan to serve regular 'ol routes for most of the site. It's not big or anything. For the blog, I hope to use flat markdown files for the posts, which I should be able to render using [gomarkdown](https://pkg.go.dev/github.com/gomarkdown/markdown/parser#Extensions). I'd also like to use frontmatter to define metadata for each post. I don't know if I should use a library or extend the gomarkdown parser for that. This is a good rundown of using and extending gomarkdown by [Krzysztof Kowalczyk](https://blog.kowalczyk.info/article/cxn3/advanced-markdown-processing-in-go.html) that should help.

Honestly, I'll probably use a module like this one, which parses the frontmatter and returns a struct and the rest of the document: [adrg/frontmatter](https://github.com/adrg/frontmatter). Using that would mean I could even write frontmatter in JSON if I wanted to... and anything's better than YAML.

The question then is whether the slug should map to the filename. That seems easiest. I could just attach it to whatever struct I finally use to represent a post before sending the rendered HTML back to the client.

I'll also need to figure out how to generate an RSS/Atom feed, which could be as simple as using [gorilla/feeds](https://github.com/gorilla/feeds), which would allow me to generate RSS, Atom, and JSON feeds. 

And, for good measure, I'd like to add [webmentions](https://indieweb.org/Webmention) — ideally without relying on one generous person's service. I think implementing them on my own would require another server. I might just be able to build it into this, though. Who knows?

## To Do
- [ ] Home page
  - [ ] Blurb about where users have landed
  - [x] List latest posts
  - [ ] List recent projects
- [ ] Blog page
  - [x] List all blog posts
  - [x] Link to `/subscribe`
  - [ ] Link to `/blogroll`
    - [ ] With OPML download (thanks for the idea [Robb Knight](https://rknight.me/blog/roll/))
- [ ] Subscribe
  - [ ] Generate feeds
    - [ ] RSS
    - [ ] Atom
    - [ ] JSON
- [ ] About
- [ ] Contact
- [ ] Projects
  - [ ] ISS analysis project
  - [ ] Current (RSS reader)
  - [ ] Desk (journalism timeline)
- [ ] Now - see [nownownow](https://nownownow.com/about)
- [ ] Notes 
  - [ ] Syndicate to Mastodon