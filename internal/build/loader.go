// internal/build/loader.go
package build

import (
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"


)

const (
    ROOT = "./content"
)

func Run() error {
	fmt.Println("🔧 build stub: nothing to do yet")
	return nil
}

func LoadFile(path string) []byte{
	fmt.Println("Loading file...")
    file, err := os.Open(path)
    defer file.Close()
    if err != nil {
        log.Fatal(err)  
    }
    fmt.Println("File loaded successfully")
    var content []byte
    buffer := make([]byte, 1024)
    for{
        n, err := file.Read(buffer)
        if err == io.EOF{
            break
        }
        if err != nil {
            log.Fatal(err)
        }
        content = append(content, buffer[:n]...)
    }
    
    return content
}   

func Walk() {
    err := filepath.WalkDir(ROOT, func(path string, d fs.DirEntry, err error) error {
        if err != nil { 
            return err
        }else if d.IsDir() {
            fmt.Println("Directory:", path)
        }else {
            fmt.Println("File:", path)
            fileName := d.Name()
            if filepath.Ext(fileName) == ".md" {
                fileData := LoadFile(path)
                Convert(fileData, fileName)
            }
        }
        return nil
    })
    if err != nil {
        fmt.Println("Error walking the path:", err)
    }
}
