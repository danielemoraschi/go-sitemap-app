package main

import (
    "fmt"
    "sync"
    "github.com/danielemoraschi/go-sitemap-common"
    "github.com/danielemoraschi/go-sitemap-common/http/fetcher"
    "github.com/danielemoraschi/go-sitemap-common/policy"
    "github.com/danielemoraschi/go-sitemap-common/http"
    "github.com/danielemoraschi/go-sitemap-common/parser"
)

func removeDuplicatesUnordered(elements []http.HttpResource) []http.HttpResource {
	encountered := map[string]http.HttpResource{}

	// Create a map of all unique elements.
	for v, el := range elements {
		encountered[elements[v].String()] = el
	}

	// Place all keys from the map into a slice.
	result := []http.HttpResource{}
	for _, el := range encountered {
		result = append(result, el)
	}
	return result
}

func main() {

    var wg sync.WaitGroup
	urlToVisit := "http://golang.org/"

    res, _ := http.HttpResourceFactory(urlToVisit, "")

    fmt.Printf("Content: %s\n", string(res.Content()))

	uniqueVisitPolicy := policy.UniqueUrlPolicyFactory()
	sameDomainPolicy := policy.SameDomainPolicyFactory(urlToVisit)

	var policies []policy.PolicyInterface
	policies = append(policies, uniqueVisitPolicy)
	policies = append(policies, sameDomainPolicy)

	var urlList []http.HttpResource
    urlList = append(urlList, res)

    wg.Add(1)
	go crawler.Crawl(&urlList, &wg, urlToVisit, 1, fetcher.HttpFetcher{}, parser.HttpParser{}, policies)

	wg.Wait()

	urlList = removeDuplicatesUnordered(urlList)

    fmt.Printf("TOT: %d\n", len(urlList))
	fmt.Printf("ALL: %s\n", urlList)
}
