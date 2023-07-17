package index

import (
	"strings"
)

type Index struct {
	index map[string][]int
}

func New() *Index {
	return &Index{
		index: make(map[string][]int),
	}
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
