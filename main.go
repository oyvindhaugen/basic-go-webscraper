package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gocolly/colly"
)

func main() {
	fName := "nrk_links.xlsx"
	file, err := os.Create(fName)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	w := csv.NewWriter(file)
	defer w.Flush()

	c := colly.NewCollector()

	c.OnHTML("a", func(e *colly.HTMLElement) {
		test := e.Attr("href")
		if strings.HasPrefix(test, "https://www.nrk.no/") {
			fmt.Println(test)
			w.Write([]string{test})
		}
	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting ", r.URL)
	})
	c.Visit("https://nrk.no")

}
