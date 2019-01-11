package main

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

func main() {
	url := "https://www.gamersky.com/ent/wp/"

	res, _ := http.Get(url)

	defer res.Body.Close()

	dc, _ := goquery.NewDocumentFromReader(res.Body)

	dc.Find("#pe100_page_contentpage .page_css a").Each(func(i int, selection *goquery.Selection) {
		if selection.Text() ==  "下一页" {
			log.Println(selection.Attr("href"))
		}
	})

}
