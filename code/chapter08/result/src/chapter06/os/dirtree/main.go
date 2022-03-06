package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {
	ReadAndOutputDir("../../..", 3)
}

func ReadAndOutputDir(rootPath string, deep int) {
	file, err := os.Open(rootPath)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	defer file.Close()

	for {
		fileInfos, err := file.Readdir(100)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("readdir error:", err)
			return
		}

		if len(fileInfos) == 0 {
			break
		}

		for _, fileInfo := range fileInfos {
			if fileInfo.IsDir() {
				if deep > 0 {
					ReadAndOutputDir(filepath.Join(rootPath, string(os.PathSeparator), fileInfo.Name()), deep-1)
				}
			} else {
				fmt.Println("file:", fileInfo.Name(), "in directory:", rootPath)
			}
		}
	}
}
