package crawl

import (
	"fmt"
	"log"
	"regexp"

	"github.com/gocolly/colly"

	"btap/crawler/model"
)


func Crawl(link string) []model.Movie {
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) { //Handle error trong quá trình craw html
		log.Println("Something went wrong:", err)
	})

	var listMovies []model.Movie

	c.OnHTML(".lister-list tr", func(e *colly.HTMLElement) { //Bóc tách dữ liệu từ HTML lấy được
		data := model.Movie{}
		
		data.Title = e.ChildText(".titleColumn") //Tìm đến thẻ con của thẻ tr có class titleColumn và lấy nội dung
		data.Title = regexp.MustCompile(`\s+`).ReplaceAllString(data.Title, " ")
		data.Rating = e.ChildText(".imdbRating") //tìm đến thẻ con có class imdbRating và lấy nội dung
		data.Url = e.ChildAttr("a", "href")      //Tìm đến thẻ con a và lấy nội dung thuộc tính href
		data.Url = "https://www.imdb.com/"+data.Url

		listMovies = append(listMovies, data)
	})

	c.OnScraped(func(r *colly.Response) { //Hoàn thành job craw
		fmt.Println("Finished", r.Request.URL)
	})

	c.Visit(link) //Trình thu thập truy cập URL đó
	return listMovies
}
