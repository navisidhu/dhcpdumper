package main

import (
	"os"
	"path/filepath"
	"strings"
)

func glob(directory string, extension []string) ([]string, error) {
	// Use the filepath.Walk function to traverse the directory and find files with the given extension.
	var files []string
	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			for _, ext := range extension {
				if strings.HasSuffix(info.Name(), ext) {
					files = append(files, path)
					break
				}
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}
