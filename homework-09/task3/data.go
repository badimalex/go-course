package main

import (
	"fmt"
	"io"
	"os"
)

func Save(writer io.Writer, args ...interface{}) {
	for _, arg := range args {
		switch s := arg.(type) {
		case string:
			fmt.Fprintln(writer, s)
		}
	}
}

func main() {
	file, _ := os.Create("test.txt")

	defer file.Close()
	Save(file, "1", 2, 3, "Hello", []string{"a", "b", "c"})
}
