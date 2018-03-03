package main

import (
	"fmt"

	// "golang.org/x/net/html"
	"github.com/PuerkitoBio/goquery"
)

func main() {
	// resp, err := http.Get("http://blog.admanmedia.com")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// body := resp.Body
	// defer body.Close()

	// text, err := ioutil.ReadAll(body)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// doc, err := html.Parse(body)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(doc.FirstChild.Data)

	doc, err := goquery.NewDocument("http://blog.admanmedia.com")
	if err != nil {
		fmt.Println(err)
		return
	}

	doc.Find(".entry-title a").Each(func(i int, s *goquery.Selection) {
		href, _ := s.Attr("href")
		fmt.Printf("%s\n(%s)\n\n", s.Text(), href)
	})

}
