package parser

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"
	jww "github.com/spf13/jwalterweatherman"

	"github.com/KaymeKaydex/bmstu-stud-parser/internal/app/model"
)

var ResultBucket []*model.News

type Parser struct{}

func (p *Parser) GetURLs(StartPage string) (URLs []string) {
	for i := 1; ; i++ {
		resp, err := http.Get(StartPage + fmt.Sprintf("%d", i) + "/")

		if err != nil {
			jww.INFO.Fatalln("Smth went wrong")
		}
		if resp == nil || resp.StatusCode == http.StatusNotFound {
			jww.INFO.Printf("Found %d pages", i-1)
			break
		}

		URLs = append(URLs, StartPage+fmt.Sprintf("%d", i)+"/")
	}
	return URLs
}

func (p *Parser) Parse(g *geziyor.Geziyor, r *client.Response) {
	r.HTMLDoc.Find("div.card").Each(func(_ int, s *goquery.Selection) {

		img, err := s.Find("img").Attr("src")
		if err == false {
			jww.INFO.Println("Cant parse img")
		} else {
		}

		news := &model.News{
			Title:       s.Find("span.card-title.grey-text").Text(),
			Description: s.Find("p").Text(),
			ImageURI:    img,
		}

		ResultBucket = append(ResultBucket, news)
	})
}
