package main

import (
	"github.com/danielemoraschi/go-sitemap-app/service"
)


func main() {
	//urlToVisit := "https://golang.org"
	urlToVisit := "http://google.com"

	concurrency := 100
	depth := 2

	service.SiteMapGeneratorService{}.Run(urlToVisit, concurrency, depth)
}
