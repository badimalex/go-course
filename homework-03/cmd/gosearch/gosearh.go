package main

import (
	"flag"
	"fmt"
	"go-course/homework-03/pkg/crawler"
	"go-course/homework-03/pkg/crawler/spider"
	"go-course/homework-03/pkg/index"
	"sort"
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

	sort.Slice(data, func(i, j int) bool {
		return data[i].ID < data[j].ID
	})

	if *searchWord != "" {
		fmt.Printf("Search results for '%s':\n", *searchWord)
		results := index.Search(strings.ToLower(*searchWord))

		for _, id := range results {
			i := BinarySearch(data, id)
			fmt.Println(data[i].URL, data[i].Title)
		}
	} else {
		fmt.Println("No search word provided.")
	}
}

func BinarySearch(arr []crawler.Document, target int) int {
	l := 0
	r := len(arr) - 1

	for l <= r {
		mid := (r + l) / 2

		if target == arr[mid].ID {
			return mid
		}

		if target < arr[mid].ID {
			r = mid - 1
		}

		if target > arr[mid].ID {
			l = mid + 1
		}
	}

	return -1
}
