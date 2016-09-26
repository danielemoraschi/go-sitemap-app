package main

import (
    "fmt"
    "sync"
    "github.com/danielemoraschi/go-sitemap-common"
    "github.com/danielemoraschi/go-sitemap-common/http/fetcher"
    "github.com/danielemoraschi/go-sitemap-common/policy"
    "github.com/danielemoraschi/go-sitemap-common/http"
    "github.com/danielemoraschi/go-sitemap-common/parser"
	"github.com/danielemoraschi/go-sitemap-common/sitemap/template"
)


func main() {
	concurrency := 100
	sem := make(chan bool, concurrency)

    var wg sync.WaitGroup
	//urlToVisit := "https://golang.org"
	urlToVisit := "http://asaquattrocento.com"

    res, _ := http.HttpResourceFactory(urlToVisit, "")

    //fmt.Printf("Content: %s\n", string(res.Content()))

	uniqueVisitPolicy := policy.UniqueUrlPolicyFactory()
	sameDomainPolicy := policy.SameDomainPolicyFactory(urlToVisit)
	validExtensionPolicy := policy.ValidExtensionPolicyFactory()

	var policies []policy.PolicyInterface
	policies = append(policies, uniqueVisitPolicy)
	policies = append(policies, sameDomainPolicy)
	policies = append(policies, validExtensionPolicy)

	var urlList crawler.UrlCollection
    urlList.Add(res)

    wg.Add(1)
	sem <- true
	go crawler.Crawl(sem, &urlList, &wg, urlToVisit, 2, fetcher.HttpFetcher{}, parser.HttpParser{}, policies)
	for i := 0; i < cap(sem); i++ {
		sem <- true
	}
	wg.Wait()

	urlList.RemoveDuplicatesUnordered()

	tpl := template.UrlSetFactory()
	xml, _ := tpl.Set(urlList.Data()).Generate()

    fmt.Printf("TOT: %d\n", urlList.Count())
	fmt.Printf("XML: %v\n", string(xml))
	//fmt.Printf("ALL: %s\n", urlList)
}
