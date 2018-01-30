package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/css"
	"github.com/tdewolff/minify/html"
	"github.com/tdewolff/minify/js"
)

func main() {
	minPath := "../web/public_min"
	assetPath := "../web/public"

	// Delete directory
	_, err := os.Stat(minPath)
	if os.IsExist(err) {
		log.Println("Delete the directory")
		os.RemoveAll(minPath)
	}

	// Create directory
	os.Mkdir(minPath, 0777)

	// Build Minifier
	m := minify.New()
	m.AddFunc("text/css", css.Minify)
	m.Add("text/html", &html.Minifier{
		KeepDocumentTags: true,
	})
	m.AddFunc("text/javascript", js.Minify)

	// Loop through assets
	err = filepath.Walk(assetPath, func(path string, f os.FileInfo, err error) error {
		outputPath := strings.Replace(path, assetPath, minPath, -1)
		if f.IsDir() {
			os.MkdirAll(outputPath, 0777)
			return nil
		}

		formats := make(map[string]string)
		formats[".css"] = "text/css"
		formats[".html"] = "text/html"
		formats[".js"] = "text/javascript"

		for ext, format := range formats {
			if strings.LastIndex(f.Name(), ext) > -1 {
				content, err := ioutil.ReadFile(path)
				b, err := m.Bytes(format, content)
				if err != nil {
					panic(err)
				}
				ioutil.WriteFile(outputPath, b, 0777)
				return nil
			}
		}
		// Just move the file
		content, _ := ioutil.ReadFile(path)
		ioutil.WriteFile(outputPath, content, 0777)
		return nil
	})

}
