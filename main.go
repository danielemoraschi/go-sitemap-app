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
	"github.com/danielemoraschi/go-sitemap-common/output"
	"sort"
)


func main() {
	concurrency := 100
	sem := make(chan bool, concurrency)

    var wg sync.WaitGroup
	//urlToVisit := "https://golang.org"
	urlToVisit := "http://google.com"

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

	tpl := template.XMLUrlSetFactory()
	//tpl := template.JsonUrlSetFactory()

	out, _ := tpl.Set(urlList.Data()).Generate()
	sort.Float64s()
    fmt.Printf("TOT: %d\n", urlList.Count())
	//fmt.Printf("RESULT: %v\n", string(out))

	//outw := output.FileWriterFactory("./sitemap.xml")
	outw := output.StOutWriterFactory()
	outw.Write(out)
	//fmt.Printf("ALL: %s\n", urlList)
}
