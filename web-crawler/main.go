package main

import (
	"fmt"
	"github.com/aditya109/web-crawler/pkg/html_parser"
	"log"
	"os"
	"time"
)

func main() {
	start := time.Now()
	foundUrls := make(map[string]bool)
	seedUrls := os.Args[1:]

	fmt.Println(seedUrls)
	// Channels
	chUrls := make(chan string)
	chFinished := make(chan bool)

	for _, url := range seedUrls {
		go html_parser.Crawl(url, chUrls, chFinished)
	}

	for c := 0; c < len(seedUrls); {
		select {
		case url := <-chUrls:
			foundUrls[url] = true
		case <-chFinished:
			c++
		}
	}

	fmt.Println("\nFound", len(foundUrls), "unique urls:")

	for url := range foundUrls {
		fmt.Println(" - " + url)
	}

	if len(os.Args) < 1 {
		fmt.Println("Please specify start page")
		os.Exit(1)
	}
	close(chUrls)
	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}
