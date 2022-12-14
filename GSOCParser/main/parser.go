package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func OnPage(link string) string {
	res, err := http.Get(link)
	if err != nil {
		log.Fatal(err)
	}
	content, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}

func main() {
	html := OnPage("https://summerofcode.withgoogle.com/programs/2022/organizations/performance-co-pilot")
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Fatalln(err)
	}

	dom.Find(".tech__content").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(selection.Text())
	})
}
