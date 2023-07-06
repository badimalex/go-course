package main

import (
	"flag"
	"fmt"
	"log"
	"search/pkg/crawler/spider"
	"strings"
)

func main() {
	searchWord := flag.String("s", "", "Search word")
	flag.Parse()

	scanner := spider.New()

	data1, err := scanner.Scan("https://go.dev", 2)
	if err != nil {
		log.Fatal(err)
	}

	data2, err := scanner.Scan("https://golang.org", 2)
	if err != nil {
		log.Fatal(err)
	}
	data := append(data1, data2...)

	if *searchWord != "" {
		fmt.Printf("Search results for '%s':\n", *searchWord)
		for _, doc := range data {
			if strings.Contains(strings.ToLower(doc.Title), strings.ToLower(*searchWord)) {
				fmt.Printf("URL: %s\n", doc.URL)
				fmt.Printf("Title: %s\n", doc.Title)
				fmt.Println()
			}
		}
	} else {
		fmt.Println("No search word provided.")
	}
}
