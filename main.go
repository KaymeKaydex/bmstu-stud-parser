package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"
	"github.com/geziyor/geziyor/export"
)


func main() {
	geziyor.NewGeziyor(&geziyor.Options{
		StartURLs: []string{"http://studsovet.bmstu.ru/"},
		ParseFunc: func(g *geziyor.Geziyor, r *client.Response) {

			r.HTMLDoc.Find("div.card").Each(func(_ int, s *goquery.Selection) {

				fmt.Println("Заголовок : ", s.Find("span.card-title.grey-text").Text())
				fmt.Println("Описание : ", s.Find("p").Text())
				img, err := s.Find("img").Attr("src")
				if err != false {

				}
				fmt.Println("Картинка : ", img)
			})


		},
		Exporters: []export.Exporter{&export.JSON{}},
	}).Start()


}


