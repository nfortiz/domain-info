package utils

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

func GetHtmlReader(domain string) *goquery.Document {
	res, err := http.Get("http://" + domain)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("Status code error: %d", res.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	return doc
}

func GetIcon(doc *goquery.Document) string {
	node := doc.Find("link[rel*=\"icon\"]")
	nodeContent, exists := node.Attr("href")
	if !exists {
		log.Println("The webpage doesnt contain an icon")
	}
	return nodeContent
}

func GetTitle(doc *goquery.Document) string {
	node := doc.Find("title")
	nodeContent := node.Text()
	return nodeContent
}