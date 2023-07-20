package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"go-course/homework-05/pkg/crawler"
	"go-course/homework-05/pkg/crawler/spider"
	"go-course/homework-05/pkg/index"
	"io"
	"os"
	"sort"
	"strings"
)

const dataFile = "data.json"

func main() {
	searchWord := flag.String("s", "", "Search word")
	flag.Parse()

	scanner := spider.New()

	links := []string{"https://go.dev", "https://golang.org"}
	var data []crawler.Document

	if fileExists(dataFile) {
		file, err := os.Open(dataFile)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		data, err = loadDataFromFile(file)
		if err != nil {
			panic(err)
		}
	} else {
		for _, link := range links {
			site, err := scanner.Scan(link, 2)
			if err != nil {
				fmt.Println("Error scanning link:", err)
				continue
			}
			data = append(data, site...)
		}
	}
	file, err := os.Create(dataFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = saveDataToFile(data, file)
	if err != nil {
		panic(err)
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

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func loadDataFromFile(reader io.Reader) ([]crawler.Document, error) {
	jsonData, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	var data []crawler.Document
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func saveDataToFile(data []crawler.Document, writer io.Writer) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = writer.Write(jsonData)
	return err
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
