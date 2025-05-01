// internal/build/loader.go
package build

import (
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const (
	PATH    = "./content"
	TOPPAGE = "dist"
	POSTS   = "dist/posts"
)

func Run() error {
	fmt.Println("🔧 build stub: nothing to do yet")
	return nil
}

func LoadFile(path string) []byte {
	fmt.Println("Loading file...")
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("File loaded successfully")
	var content []byte
	buffer := make([]byte, 1024)
	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		content = append(content, buffer[:n]...)
	}
	fmt.Println("File read successfully")
	return content
}

func Walk() {
	err := filepath.WalkDir(PATH, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		} else if d.IsDir() {
			fmt.Println("Directory:", path)
		} else if strings.Contains(path, "posts") == false {
			fmt.Println("Top page:")
			fmt.Println("File:", path)
			fileName := d.Name()
			if filepath.Ext(fileName) == ".md" {
				fileData := LoadFile(path)
				Convert(fileData, fileName, TOPPAGE)
			}
		} else {
			fmt.Println("Posts")
			fmt.Println("File:", path)
			fileName := d.Name()
			if filepath.Ext(fileName) == ".md" {
				fileData := LoadFile(path)
				Convert(fileData, fileName, POSTS)
			}
		}
		return nil
	})
	if err != nil {
		fmt.Println("Error walking the path:", err)
	}

}
