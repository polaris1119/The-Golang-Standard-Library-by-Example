package main

import (
	"archive/tar"
	"log"
	"os"
	"path/filepath"
)

func main() {
	f, err := os.Create("result/polarisxu.tar")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	tw := tar.NewWriter(f)
	defer tw.Close()

	root, err := filepath.Abs("..")
	if err != nil {
		log.Fatal(err)
	}

	err = ScanDir(tw, root, root)
	if err != nil {
		os.Remove("result/polarisxu.tar")
		log.Fatal(err)
	}
}

func ScanDir(tw *tar.Writer, root, dirName string) error {
	dirEntries, err := os.ReadDir(dirName)
	if err != nil {
		return err
	}

	for _, dirEntry := range dirEntries {
		newPath := dirName + "/" + dirEntry.Name()
		if dirEntry.IsDir() {
			if err = ScanDir(tw, root, newPath); err != nil {
				return err
			}
		} else {
			hdr := &tar.Header{
				Name: dirName[len(root):] + "/" + dirEntry.Name(),
				// Mode: int64(dirEntry.Type().Perm()),
			}
			fileInfo, err := dirEntry.Info()
			if err != nil {
				return err
			}
			hdr.Mode = int64(fileInfo.Mode().Perm())
			body, err := os.ReadFile(newPath)
			if err != nil {
				return err
			}
			hdr.Size = int64(len(body))
			if err = tw.WriteHeader(hdr); err != nil {
				return err
			}

			if _, err = tw.Write(body); err != nil {
				return err
			}
		}
	}

	return nil
}
