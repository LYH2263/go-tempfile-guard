package main

import (
	"fmt"
	"path/filepath"

	"github.com/LYH2263/go-tempfile-guard/internal/writer"
)

func main() {
	p := filepath.Join(".", "out.txt")
	fmt.Println(writer.AtomicWrite(p, []byte("hi")))
}
