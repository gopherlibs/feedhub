# FeedHub Go Module [![Go Reference](https://pkg.go.dev/badge/github.com/gopherlibs/feedhub.svg)](https://pkg.go.dev/github.com/gopherlib/feedhub) [![CircleCI Build Status](https://circleci.com/gh/gopherlibs/feedhub.svg?style=shield)](https://app.circleci.com/pipelines/github/gopherlibs/feedhub) [![GitHub License](https://img.shields.io/badge/license-BSD2-blue.svg)](https://raw.githubusercontent.com/gopherlibs/feedhub/trunk/LICENSE)


*This project is a fork of 'Gorilla's Feeds'. They haven't really worked on theirs since 2019 and archived it in 2022. This project aims to be a successor.*

FeedHub is a web feed generator library for generating RSS, Atom, and JSON feeds from Go applications.


## Table of Contents

- [Goals](#goals)
- [Usage](#usage)
- [Migration](#migration)


## Goals

 * Provide a simple interface to create both Atom & RSS 2.0 feeds
 * Full support for [Atom][atom], [RSS 2.0][rss], and [JSON Feed Version 1][jsonfeed] spec elements
 * Ability to modify particulars for each spec

[atom]: https://tools.ietf.org/html/rfc4287
[rss]: http://www.rssboard.org/rss-specification
[jsonfeed]: https://jsonfeed.org/version/1


## Usage

```go
package main

import (
    "fmt"
    "log"
    "time"

    "github.com/gopherlibs/feedhub/feedhub"
)

func main() {
    now := time.Now()
    feed := &feedhub.Feed{
        Title:       "jmoiron.net blog",
        Link:        &feedhub.Link{Href: "http://jmoiron.net/blog"},
        Description: "discussion about tech, footie, photos",
        Author:      &feeds.Author{Name: "Jason Moiron", Email: "jmoiron@jmoiron.net"},
        Created:     now,
    }

    feed.Items = []*feedhub.Item{
        &feedhub.Item{
            Title:       "Limiting Concurrency in Go",
            Link:        &feedhub.Link{Href: "http://jmoiron.net/blog/limiting-concurrency-in-go/"},
            Description: "A discussion on controlled parallelism in golang",
            Author:      &feeds.Author{Name: "Jason Moiron", Email: "jmoiron@jmoiron.net"},
            Created:     now,
        },
        &feedhub.Item{
            Title:       "Logic-less Template Redux",
            Link:        &feedhub.Link{Href: "http://jmoiron.net/blog/logicless-template-redux/"},
            Description: "More thoughts on logicless templates",
            Created:     now,
        },
        &feedhub.Item{
            Title:       "Idiomatic Code Reuse in Go",
            Link:        &feedhub.Link{Href: "http://jmoiron.net/blog/idiomatic-code-reuse-in-go/"},
            Description: "How to use interfaces <em>effectively</em>",
            Created:     now,
        },
    }

    atom, err := feed.ToAtom()
    if err != nil {
        log.Fatal(err)
    }

    rss, err := feed.ToRss()
    if err != nil {
        log.Fatal(err)
    }

    json, err := feed.ToJSON()
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(atom, "\n", rss, "\n", json)
}
```

Outputs:

```xml
<?xml version="1.0" encoding="UTF-8"?>
<feed xmlns="http://www.w3.org/2005/Atom">
  <title>jmoiron.net blog</title>
  <link href="http://jmoiron.net/blog"></link>
  <id>http://jmoiron.net/blog</id>
  <updated>2013-01-16T03:26:01-05:00</updated>
  <summary>discussion about tech, footie, photos</summary>
  <entry>
    <title>Limiting Concurrency in Go</title>
    <link href="http://jmoiron.net/blog/limiting-concurrency-in-go/"></link>
    <updated>2013-01-16T03:26:01-05:00</updated>
    <id>tag:jmoiron.net,2013-01-16:/blog/limiting-concurrency-in-go/</id>
    <summary type="html">A discussion on controlled parallelism in golang</summary>
    <author>
      <name>Jason Moiron</name>
      <email>jmoiron@jmoiron.net</email>
    </author>
  </entry>
  <entry>
    <title>Logic-less Template Redux</title>
    <link href="http://jmoiron.net/blog/logicless-template-redux/"></link>
    <updated>2013-01-16T03:26:01-05:00</updated>
    <id>tag:jmoiron.net,2013-01-16:/blog/logicless-template-redux/</id>
    <summary type="html">More thoughts on logicless templates</summary>
    <author></author>
  </entry>
  <entry>
    <title>Idiomatic Code Reuse in Go</title>
    <link href="http://jmoiron.net/blog/idiomatic-code-reuse-in-go/"></link>
    <updated>2013-01-16T03:26:01-05:00</updated>
    <id>tag:jmoiron.net,2013-01-16:/blog/idiomatic-code-reuse-in-go/</id>
    <summary type="html">How to use interfaces &lt;em&gt;effectively&lt;/em&gt;</summary>
    <author></author>
  </entry>
</feed>

<?xml version="1.0" encoding="UTF-8"?>
<rss version="2.0">
  <channel>
    <title>jmoiron.net blog</title>
    <link>http://jmoiron.net/blog</link>
    <description>discussion about tech, footie, photos</description>
    <managingEditor>jmoiron@jmoiron.net (Jason Moiron)</managingEditor>
    <pubDate>2013-01-16T03:22:24-05:00</pubDate>
    <item>
      <title>Limiting Concurrency in Go</title>
      <link>http://jmoiron.net/blog/limiting-concurrency-in-go/</link>
      <description>A discussion on controlled parallelism in golang</description>
      <pubDate>2013-01-16T03:22:24-05:00</pubDate>
    </item>
    <item>
      <title>Logic-less Template Redux</title>
      <link>http://jmoiron.net/blog/logicless-template-redux/</link>
      <description>More thoughts on logicless templates</description>
      <pubDate>2013-01-16T03:22:24-05:00</pubDate>
    </item>
    <item>
      <title>Idiomatic Code Reuse in Go</title>
      <link>http://jmoiron.net/blog/idiomatic-code-reuse-in-go/</link>
      <description>How to use interfaces &lt;em&gt;effectively&lt;/em&gt;</description>
      <pubDate>2013-01-16T03:22:24-05:00</pubDate>
    </item>
  </channel>
</rss>

{
  "version": "https://jsonfeed.org/version/1",
  "title": "jmoiron.net blog",
  "home_page_url": "http://jmoiron.net/blog",
  "description": "discussion about tech, footie, photos",
  "author": {
    "name": "Jason Moiron"
  },
  "items": [
    {
      "id": "",
      "url": "http://jmoiron.net/blog/limiting-concurrency-in-go/",
      "title": "Limiting Concurrency in Go",
      "summary": "A discussion on controlled parallelism in golang",
      "date_published": "2013-01-16T03:22:24.530817846-05:00",
      "author": {
        "name": "Jason Moiron"
      }
    },
    {
      "id": "",
      "url": "http://jmoiron.net/blog/logicless-template-redux/",
      "title": "Logic-less Template Redux",
      "summary": "More thoughts on logicless templates",
      "date_published": "2013-01-16T03:22:24.530817846-05:00"
    },
    {
      "id": "",
      "url": "http://jmoiron.net/blog/idiomatic-code-reuse-in-go/",
      "title": "Idiomatic Code Reuse in Go",
      "summary": "How to use interfaces \u003cem\u003eeffectively\u003c/em\u003e",
      "date_published": "2013-01-16T03:22:24.530817846-05:00"
    }
  ]
}
```


## Migration

Coming from [Gorilla Feeds](https://github.com/gorilla/feeds)?
Switching over to FeedHub is fairly straight forward with a few steps.
FeedHub v1.0.0 is identical to Gorilla Feeds v1.1.1.
Only the following changes referring to the module/package need to be done:

1. In Go files where `github.com/gorilla/feeds` is imported, instead import `github.com/gopherlibs/feedhub/feedhub`.
1. In your Go code replace references to the package name `feeds` with `feedhub`.
  - *optional* If you have a lot of references and don't want to do this step now, in the import statement you can alias the package name. You can use the import statement `feeds github.com/gopherlibs/feedhub/feedhub` and then not have to update the references.
1. Run `go get github.com/gopherlibs/feedhub` to download the module locally.
1. Run `go mod tidy` to update your `go.mod` file with the new module.
