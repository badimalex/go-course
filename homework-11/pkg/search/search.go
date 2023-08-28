package search

import (
	"fmt"
	"go-course/homework-11/pkg/crawler"
	"go-course/homework-11/pkg/index"
	"strings"
)

type Search struct {
	idx   *index.Index
	cache []crawler.Document
}

func New(idx *index.Index, cache []crawler.Document) *Search {
	return &Search{
		idx:   idx,
		cache: cache,
	}
}

func (s *Search) Results(phrase string) []string {
	results := s.idx.Search(strings.ToLower(phrase))
	var rows []string

	for _, id := range results {
		i := binSearch(s.cache, id)
		if i == -1 {
			fmt.Println("Document not found for ID:", id)
		} else {
			rows = append(rows, fmt.Sprintf("%s: %s\n", s.cache[i].URL, s.cache[i].Title))
		}
	}

	return rows
}

func binSearch(arr []crawler.Document, target int) int {
	l := 0
	r := len(arr) - 1

	for l <= r {
		mid := (r + l) / 2

		if target == arr[mid].ID {
			return mid
		}

		if target < arr[mid].ID {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return -1
}
