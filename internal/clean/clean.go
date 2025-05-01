package clean

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

const (
	PATH = "./dist"
)

func Clean() {
	fmt.Println("Cleaning up...")
	err := filepath.WalkDir(PATH, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		} else if d.IsDir() {
			fmt.Println("Directory:", path)
		} else {
			if filepath.Ext(d.Name()) == ".html" {
				if err := os.Remove(path); err != nil {
					log.Printf("Error deleting file %s: %v\n", path, err)
				} else {
					fmt.Printf("Deleted file: %s\n", path)
				}
			}
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Cleanup completed.")
}
