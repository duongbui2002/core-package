package main

import (
	"os"
	"path/filepath"
)

func main() {
	dir, _ := os.Getwd()
	print(filepath.Dir(dir))
}
