package main

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
	"github.com/tealeg/xlsx"
)

// defining a data structure to store the scraped data
type PokemonProduct struct {
	url, image, name, price string
}

func main() {

	c := colly.NewCollector(
	// colly.AllowedDomains("https://scrapeme.live/"),
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

	var PokemonProducts []PokemonProduct

	c.OnHTML(".product", func(h *colly.HTMLElement) {

		PokemonProduct := PokemonProduct{
			url:   h.ChildAttr("a", "href"),
			image: h.ChildAttr("img", "src"),
			name:  h.ChildText("h2"),
			price: h.ChildText("span.amount"),
		}
		// fmt.Println(PokemonProduct)
		PokemonProducts = append(PokemonProducts, PokemonProduct)

	})

	c.Visit("https://scrapeme.live/shop/")

	// fmt.Println(PokemonProducts)

	createAndSaveFile("sheet1", PokemonProducts)

}

func createAndSaveFile(sheetname string, datas []PokemonProduct) {
	file := xlsx.NewFile()
	sheet, err := file.AddSheet(sheetname)
	if err != nil {
		panic(err)
	}

	row := sheet.AddRow()
	row.AddCell().SetValue("name")
	row.AddCell().SetValue("image")
	row.AddCell().SetValue("price")
	row.AddCell().SetValue("url")
	for _, i := range datas {
		row := sheet.AddRow()
		row.AddCell().SetValue(i.name)
		row.AddCell().SetValue(i.image)
		row.AddCell().SetValue(i.price)
		row.AddCell().SetValue(i.url)
	}
	err = file.Save("pockeymon" + ".xlsx")
	if err != nil {
		log.Println(err)
		panic(err)
	}
}
