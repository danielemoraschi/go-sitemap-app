package main

import (
	"fmt"
    "sync"
    "github.com/danielemoraschi/go-sitemap-common"
    "github.com/danielemoraschi/go-sitemap-common/http/fetcher"
    "github.com/danielemoraschi/go-sitemap-common/policy"
    "github.com/danielemoraschi/go-sitemap-common/http"
)

func removeDuplicatesUnordered(elements []string) []string {
	encountered := map[string]bool{}

	// Create a map of all unique elements.
	for v := range elements {
		encountered[elements[v]] = true
	}

	// Place all keys from the map into a slice.
	result := []string{}
	for key := range encountered {
		result = append(result, key)
	}
	return result
}

func main() {
    var wg sync.WaitGroup
	urlToVisit := "http://golang.org/"

    res := http.HttpResourceFactory(urlToVisit)

    fmt.Println("Content: %s", string(res.Content()))

	uniqueVisitPolicy := policy.UniqueUrlPolicyFactory()
	sameDomainPolicy := policy.SameDomainPolicyFactory(urlToVisit)

	var policies []policy.PolicyInterface
	policies = append(policies, uniqueVisitPolicy)
	policies = append(policies, sameDomainPolicy)

	var urlList []string
    urlList = append(urlList, urlToVisit)

    wg.Add(1)
	go crawler.Crawl(&urlList, &wg, urlToVisit, 2, fetcher.FakeFetcher, policies)

	wg.Wait()

	urlList = removeDuplicatesUnordered(urlList)
	fmt.Println("ALL: %s", urlList)
}
