package service

import (
    "github.com/danielemoraschi/go-sitemap-common/policy"
    "github.com/danielemoraschi/go-sitemap-common"
    "github.com/danielemoraschi/go-sitemap-common/http/fetcher"
    "github.com/danielemoraschi/go-sitemap-common/output"
    "github.com/danielemoraschi/go-sitemap-common/http"
    "github.com/danielemoraschi/go-sitemap-common/sitemap/template"
    "github.com/danielemoraschi/go-sitemap-common/parser"
    "sync"
    "fmt"
)

type ServiceInterface interface {
    Run(urlToVisit string, depth, concurrency int)
}

func GenerateSiteMap(
    urlToVisit string,
    depth, concurrency int,
    fetch fetcher.FetcherInterface,
    parse parser.ParserInterface,
    policies []policy.PolicyInterface,
    tpl template.TemplateInterface,
    out output.OutputInterface,
) {
    sem := make(chan bool, concurrency)

    var wg sync.WaitGroup
    res, _ := http.HttpResourceFactory(urlToVisit, "")

    var urlList crawler.UrlCollection
    urlList.Add(res)

    wg.Add(1)
    sem <- true

    go crawler.Crawl(sem, &urlList, &wg, urlToVisit, depth, fetch, parse, policies)

    for i := 0; i < cap(sem); i++ {
        sem <- true
    }

    wg.Wait()

    urlList.RemoveDuplicatesUnordered()

    content, _ := tpl.Set(urlList.Data()).Generate()

    fmt.Printf("TOT: %d\n", urlList.Count())
    //fmt.Printf("RESULT: %v\n", string(out))

    out.Write(content)
    //fmt.Printf("ALL: %s\n", urlList)
}


