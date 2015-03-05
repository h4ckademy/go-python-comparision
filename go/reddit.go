package main

import (
	"fmt"
	"log"
	"runtime"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/fatih/color"
)

var redditUrl string = "http://reddit.com/r/programming"

type Article struct {
	Title       string
	Link        string
	Description string
}

func NewArticle(title string, link string) Article {
	a := Article{
		Title: title,
		Link:  link,
	}

	a.getDescription()

	return a
}

func (a *Article) getDescription() {
	doc, _ := goquery.NewDocument(a.Link)

	metas := doc.Find("meta[name=description]")

	if len(metas.Nodes) > 0 {
		a.Description, _ = metas.Attr("content")
	}
}

func (a Article) Show() {
	green := color.New(color.FgGreen).SprintFunc()
	blue := color.New(color.FgBlue).SprintFunc()
	fmt.Printf("* [%s] (%s)", green(a.Title), blue(a.Link))
	fmt.Println(a.Description + "\n")
}

func main() {

	numCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPU)

	doc, err := goquery.NewDocument(redditUrl)
	if err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup
	links := doc.Find("a.title")
	wg.Add(len(links.Nodes))

	articlesChan := make(chan Article, len(links.Nodes))

	for i := range links.Nodes {
		s := links.Eq(i)
		go func(wg *sync.WaitGroup, arts chan Article, node *goquery.Selection) {
			defer wg.Done()

			title := s.Text()
			link, _ := s.Attr("href")

			a := NewArticle(title, link)

			fmt.Printf("fetching %s\n", link)

			arts <- a
		}(&wg, articlesChan, s)
	}
	wg.Wait()
	close(articlesChan)

	for article := range articlesChan {
		article.Show()
	}
}
