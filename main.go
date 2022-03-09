package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/gocolly/colly"
	"github.com/syndtr/goleveldb/leveldb"
)

type WebPage struct {
	Url  string `json:"url"`
	Text string `json:"title"`
}

func add_to_db(key string, value string) {

	db, err := leveldb.OpenFile("./data.db", nil)
	if err != nil {
		fmt.Println(err)
	}
	err = db.Put([]byte(key), []byte(value), nil)
	// fmt.Println(key)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
}

func return_url(text string) {
	db, err := leveldb.OpenFile("./data.db", nil)

	if err != nil {
		fmt.Println(err)
	}

	iter := db.NewIterator(nil, nil)
	for iter.Next() {
		key := string(iter.Key())
		// fmt.Println(key)
		value := string(iter.Value())
		if strings.Contains(value, text) {
			fmt.Println(key)
		}
	}
	iter.Release()
	_ = iter.Error()
	defer db.Close()
}
func crawler() {

	var webpage WebPage
	c := colly.NewCollector(
		colly.AllowedDomains("wikipedia.org", "en.wikipedia.org", "https://en.wikipedia.org/"),
		colly.MaxDepth(1),
		colly.Async(true),
	)

	c.OnHTML("html", func(e *colly.HTMLElement) {

		webpage.Text = e.ChildText("p")
		webpage.Url = e.Request.URL.String()
		add_to_db(webpage.Url, webpage.Text)
	})

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		c.Visit(e.Request.AbsoluteURL(link))
	})

	c.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		RandomDelay: 1 * time.Second,
	})

	c.OnRequest(func(r *colly.Request) { // url

		fmt.Println("Crawl on Page", r.URL.String())
	})
	url := "https://en.wikipedia.org/wiki/2022_Peshawar_mosque_attack"

	c.Visit(url)
	c.Wait()

}

func main() {

	crawler() // comment to stop crawling
	fmt.Println("==========Crawled pages ==========")
	fmt.Println("==========links where text is found ==========")

	return_url("act")

}
