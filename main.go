package main
import (
  "github.com/gocolly/colly"
  "fmt"
  )

func main() {
  baseUrl = "https://khbartar.blog.ir"
  var urls []string
  c := colly.NewCollector(
    colly.Async(true),
  )
  c.AllowedDomains = []string{"khbartar.blog.ir"}
  c.Limit(&colly.LimitRule{Parallelism: 3})
  c.OnHTML("a[href]", func(e *colly.HTMLElement) {
    url = baseUrl + e.Attr("href")
    urls = append(urls, url)
    fmt.Println(url)
    e.Request.Visit(baseUrl + url)
})
  c.Visit("https://khbartar.blog.ir")
  fmt.Println(len(urls))
  }
