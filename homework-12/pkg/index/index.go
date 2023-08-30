package index

import (
	"strings"
)

type Index struct {
	Words map[string][]int
}

func New() *Index {
	return &Index{
		Words: make(map[string][]int),
	}
}

func (i *Index) AddDocument(docID int, content string) {
	words := strings.Fields(strings.ToLower(content))
	for _, word := range words {
		i.Words[word] = append(i.Words[word], docID)
	}
}

func (i *Index) Search(word string) []int {
	return i.Words[word]
}
