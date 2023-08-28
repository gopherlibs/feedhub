package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gopherlibs/feedhub/feedhub"
)

func main() {
	now := time.Now()
	feed := &feedhub.Feed{
		Title:       "jmoiron.net blog",
		Link:        &feedhub.Link{Href: "http://jmoiron.net/blog"},
		Description: "discussion about tech, footie, photos",
		Author:      &feedhub.Author{Name: "Jason Moiron", Email: "jmoiron@jmoiron.net"},
		Created:     now,
	}

	feed.Items = []*feedhub.Item{
		&feedhub.Item{
			Title:       "Limiting Concurrency in Go",
			Link:        &feedhub.Link{Href: "http://jmoiron.net/blog/limiting-concurrency-in-go/"},
			Description: "A discussion on controlled parallelism in golang",
			Author:      &feedhub.Author{Name: "Jason Moiron", Email: "jmoiron@jmoiron.net"},
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

	fmt.Println("=======")
	fmt.Println("Running server.... press ctrl-c to exit")

	http.HandleFunc("/feed.rss", func(w http.ResponseWriter, r *http.Request) {
		feed.WriteRss(w)
	})

	err = http.ListenAndServe(":3333", nil)
}
