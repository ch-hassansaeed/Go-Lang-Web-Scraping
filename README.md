# Go-Lang-Web-Scraping
using Colly lib we can write any kind of crawler/scraper/spider program
 
Colly can easily extract structured data from any websites, which can be used for a wide range of applications, like data mining, data processing or archiving bots and automated systems.

###Basic Scraping Program 

```
func main() {
 	c := colly.NewCollector()
 
 	// Find and visit all links
 	c.OnHTML("a", func(e *colly.HTMLElement) {
 		e.Request.Visit(e.Attr("href"))
 	})
 
 	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit("http://go-colly.org/")
}
```
