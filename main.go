package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("quotes.toscrape.com"),
	)

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36")
		fmt.Println("visiting", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("response status code : ", r.StatusCode)
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("error : ", err.Error())
	})

	c.OnHTML(".quote", func(h *colly.HTMLElement) {
		div := h.DOM
		quote := div.Find(".text").Text()
		auther := div.Find(".author").Text()
		fmt.Printf("\n %s\n\tBy %s\n", quote, auther)
	})

	// c.OnHTML("span.text",func(h *colly.HTMLElement) {
	// 	fmt.Println("\n\nQuotes :- ",h.Text)
	// })

	// c.OnHTML("small.author",func(h *colly.HTMLElement) {
	// 	fmt.Println("\n\nAuthor :- ",h.Text)
	// })

	c.Visit("https://quotes.toscrape.com/")

}
