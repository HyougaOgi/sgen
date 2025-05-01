package build

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/gomarkdown/markdown"
)

func Convert(html []byte, fileName string, path string) {
	md := []byte(html)

	markdownHTML := markdown.ToHTML(md, nil, nil)
	fileNameWithoutExt := strings.TrimSuffix(fileName, filepath.Ext(fileName))
	outputPath := filepath.Join(path, fileNameWithoutExt+".html")
	file, err := os.Create(outputPath)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	_, err = file.Write(markdownHTML)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("File %s converted to HTML and saved to %s\n", fileName, outputPath)
}
