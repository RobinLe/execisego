package crawler

import "testing"

func TestCrawler(t *testing.T) {
	Crawl()
}

func TestCFetch(t *testing.T) {
	fetchPage("https://golang.org")
}
