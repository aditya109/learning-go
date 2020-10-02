package html_parser

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"strings"
)

// Extract all http** links from a given package
func Crawl(url string, ch chan string, chFinished chan bool) {
	response, err := http.Get(url)

	defer func() {
		// Notify that we're done after this function
		chFinished <- true
	}()

	if err != nil {
		fmt.Println("error: Failed to crawl: ", url)
		return
	}
	b := response.Body
	defer b.Close() // close Body when the function completes

	z := html.NewTokenizer(b)

	for {
		tt := z.Next()
		switch {
		case tt == html.ErrorToken:
			// End of the document, we're done
			return
		case tt == html.StartTagToken:
			t := z.Token()

			// Check if the token is an <a> tag
			isAnchor := t.Data == "a"
			if !isAnchor {
				continue
			}

			// Extract the href value, if there is one
			ok, url := GetHref(t)
			if !ok {
				continue
			}
			// Make sure the url begins in http**
			hasProto := strings.Index(url, "http") == 0
			if hasProto {
				ch <- url
			}
		}
	}

}

// Helper function to pull the href attribute from a Token
func GetHref(t html.Token) (ok bool, href string) {
	// Iterate over token attribute until we find an "href"
	for _, a := range t.Attr {
		if a.Key == "href" {
			href = a.Val
			ok = true
		}
	}
	// "bare" return will return the variables (ok, href) as defined in the function definition
	return
}
