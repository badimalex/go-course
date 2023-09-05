package index

import (
	"go-course/homework-13/pkg/crawler"
	"strings"
)

type Index struct {
	index map[string][]int
}

func New(data []crawler.Document) *Index {
	idx := &Index{
		index: make(map[string][]int),
	}
	for _, doc := range data {
		idx.AddDocument(doc.ID, doc.Title)
	}
	return idx
}

func (i *Index) AddDocument(docID int, content string) {
	words := strings.Fields(strings.ToLower(content))
	for _, word := range words {
		i.index[word] = append(i.index[word], docID)
	}
}

func (i *Index) Search(word string) []int {
	return i.index[word]
}
