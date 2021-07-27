package goq

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func ExampleScrape(rw http.ResponseWriter, r *http.Request)  {
  // Request the HTML page.
//  res, err := http.Get("http://metalsucks.net")
  res, err := http.Get("https://developer.mozilla.org/en-US/")  
  if err != nil {
    log.Fatal(err)
  }
  defer res.Body.Close()
  if res.StatusCode != 200 {
    log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
  }

  // Load the HTML document
  doc, err := goquery.NewDocumentFromReader(res.Body)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Fprintln(rw, "<table>")
  // Find the review items
  //doc.Find(".left-content article .post-title").Each(func(i int, s *goquery.Selection) {
  doc.Find(".readable-line-length").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the title
		title := s.Find("a").Text()
    //fmt.Printf("Review %d: %s\n", i, title)
    fmt.Fprintf(rw, "<tr><td>Review %d: %s</td></tr>\n", i, title)		
	})
  fmt.Fprintln(rw, "</table>")
}
