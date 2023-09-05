package retriever

import (
	"fmt"
	"go-course/homework-13/pkg/crawler"
	"go-course/homework-13/pkg/index"
	"strings"
)

type Retriever struct {
	idx   *index.Index
	cache []crawler.Document
}

func New(idx *index.Index, cache []crawler.Document) *Retriever {
	return &Retriever{
		idx:   idx,
		cache: cache,
	}
}

func (r *Retriever) Find(phrase string) []crawler.Document {
	results := r.idx.Search(strings.ToLower(phrase))
	var rows []crawler.Document

	for _, id := range results {
		i := binSearch(r.cache, id)
		if i == -1 {
			fmt.Println("Document not found for ID:", id)
		} else {
			rows = append(rows, r.cache[i])
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
