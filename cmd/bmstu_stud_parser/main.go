package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"
	jww "github.com/spf13/jwalterweatherman"
	"net/http"
)


func parser(g *geziyor.Geziyor, r *client.Response){
	r.HTMLDoc.Find("div.card").Each(func(_ int, s *goquery.Selection) {

		fmt.Println("Заголовок : ", s.Find("span.card-title.grey-text").Text())
		fmt.Println("Описание : ", s.Find("p").Text())

		img, err := s.Find("img").Attr("src")

		if err == false {
			jww.INFO.Println("Cant parse img")
		} else {
			fmt.Println("Картинка : ", img)
		}
	})
}
func main() {
	URLs := []string{}
	str := "http://studsovet.bmstu.ru/page/"

	for i:=1;;i++ {
		resp, err := http.Get(str + fmt.Sprintf("%d",i) +"/")

		if err != nil {
			jww.INFO.Fatalln("Smth went wrong")
		}
		if resp == nil || resp.StatusCode == http.StatusNotFound {

			break
		}
		URLs = append(URLs, str + fmt.Sprintf("%d",i)+"/")
	}
	fmt.Println(URLs)

	geziyor.NewGeziyor(&geziyor.Options{
		StartURLs: URLs,
		ParseFunc: parser,
		Exporters: nil,
	}).Start()


}


