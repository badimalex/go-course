package main

import (
	"encoding/json"
	"fmt"
	"go-course/homework-11/pkg/crawler"
	"go-course/homework-11/pkg/crawler/spider"
	"go-course/homework-11/pkg/index"
	"go-course/homework-11/pkg/netsrv"
	"go-course/homework-11/pkg/search"
	"io"
	"log"
	"os"
	"sort"
)

const dataFile = "data.json"

func main() {
	scanner := spider.New()

	links := []string{"https://go.dev", "https://golang.org"}
	var data []crawler.Document

	_, err := os.Stat(dataFile)
	if err == nil {
		file, err := os.Open(dataFile)
		if err != nil {
			log.Fatalf("error opening data file: %v", err)
		}
		defer file.Close()

		data, err = loadData(file)
		if err != nil {
			log.Fatalf("error loading data from file: %v", err)
		}
	} else {
		for _, link := range links {
			site, err := scanner.Scan(link, 2)
			if err != nil {
				fmt.Println("error scanning link:", err)
				continue
			}
			data = append(data, site...)
		}

		file, err := os.Create(dataFile)
		if err != nil {
			log.Fatalf("error creating data file: %v", err)
		}
		defer file.Close()

		err = saveData(data, file)
		if err != nil {
			log.Fatalf("error saving data to file: %v", err)
		}
	}

	idx := index.New()
	for _, doc := range data {
		idx.AddDocument(doc.ID, doc.Title)
	}

	sortData(data)

	res := search.New(idx, data)
	server := netsrv.New(res)

	server.Start()
}

func loadData(reader io.Reader) ([]crawler.Document, error) {
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

func saveData(data []crawler.Document, writer io.Writer) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = writer.Write(jsonData)
	return err
}

func sortData(data []crawler.Document) {
	sort.Slice(data, func(i, j int) bool {
		return data[i].ID < data[j].ID
	})
}