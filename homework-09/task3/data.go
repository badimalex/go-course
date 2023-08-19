package homework09

import (
	"fmt"
	"io"
)

func Save(writer io.Writer, args ...interface{}) {
	for _, arg := range args {
		switch s := arg.(type) {
		case string:
			fmt.Fprintln(writer, s)
		}
	}
}
