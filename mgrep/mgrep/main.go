package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func discoverDirs(wl *worklist.Worklist, path string) {
	entires, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("Readdir error:", err)
		return
	}

	for _, entry := range entires {
		if entry.IsDir() {
			nextPath := filepath.Join(path, entry.Name())
		}
	}
}

func main() {

}