package main

import (
	"flag"
	"fmt"
	"go-course/homework-03/pkg/crawler"
	"go-course/homework-03/pkg/crawler/spider"
	"go-course/homework-03/pkg/index"
	"strings"
)

func main() {
	searchWord := flag.String("s", "", "Search word")
	flag.Parse()

	scanner := spider.New()

	links := []string{"https://go.dev", "https://golang.org"}
	var data []crawler.Document

	for _, link := range links {
		site, _ := scanner.Scan(link, 2)
		data = append(data, site...)
	}

	index := index.New()
	for _, doc := range data {
		index.AddDocument(doc.ID, doc.Title)
	}

	fmt.Println(index)

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
