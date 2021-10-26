package main
import (
  "github.com/gocolly/colly"
  "fmt"
  )

func main() {
  baseUrl := "https://khbartar.blog.ir"
  fmt.Println(baseUrl)
  var urls []string
  c := colly.NewCollector(
    colly.Async(true),
  )
  c.AllowedDomains = []string{"khbartar.blog.ir"}
  c.Limit(&colly.LimitRule{Parallelism: 3})
  c.OnHTML("a[href]", func(e *colly.HTMLElement) {
    
    url := e.Attr("href")
    if string(url[0]) != "/" || string(url[1]) == "/" {
      return
      }
    url = baseUrl + url
    urls = append(urls, url)
    fmt.Println(url)
    e.Request.Visit(baseUrl + url)
})
  c.Visit("https://khbartar.blog.ir")
  c.Wait()
  fmt.Println(len(urls))
  }
