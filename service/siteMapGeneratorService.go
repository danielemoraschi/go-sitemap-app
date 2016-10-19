package service

import (
    "github.com/danielemoraschi/go-sitemap-common/policy"
    "github.com/danielemoraschi/go-sitemap-common/http/fetcher"
    "github.com/danielemoraschi/go-sitemap-common/sitemap/template"
    "github.com/danielemoraschi/go-sitemap-common/parser"
    "github.com/danielemoraschi/go-sitemap-common/output"
)

type SiteMapGeneratorService struct{}

func (SiteMapGeneratorService) Run(urlToVisit string, depth, concurrency int) {

    var policies []policy.PolicyInterface
    policies = append(policies, policy.UniqueUrlPolicyFactory())
    policies = append(policies, policy.SameDomainPolicyFactory(urlToVisit))
    policies = append(policies, policy.ValidExtensionPolicyFactory())

    tpl := template.XMLUrlSetFactory()
    //tpl := template.JsonUrlSetFactory()

    //outw := output.FileWriterFactory("./sitemap.xml")
    outw := output.StOutWriterFactory()

    GenerateSiteMap(
        urlToVisit,
        depth,
        concurrency,
        fetcher.HttpFetcher{},
        parser.HttpParser{},
        policies,
        tpl,
        outw,
    )
}
