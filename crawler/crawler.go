package crawler

import (
	"fmt"
	"net/http"
	"os"

	"io/ioutil"

	"bytes"

	"golang.org/x/net/html"
)

// Crawl crawl web site
func Crawl() {
	body := fetchPage("https://golang.org")
	// cover byte to io.reader
	r := bytes.NewReader(body)
	doc, err := html.Parse(r)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	linkList := []string{}
	for _, link := range findLinks(linkList, doc) {
		fmt.Println(link)
	}
	fmt.Printf("%s", linkList)
}

func fetchPage(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("fetch error!")
		os.Exit(1)
	}
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	// fmt.Printf("%s", body)
	return body
}

func findLinks(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, attr := range n.Attr {
			if attr.Key == "href" {
				links = append(links, attr.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		findLinks(links, c)
	}
	return links
}
