package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	filepath.Walk("../../..", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		fmt.Println("file:", info.Name(), "in directory:", path)

		return nil
	})
}
