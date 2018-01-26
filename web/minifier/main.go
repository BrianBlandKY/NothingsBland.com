package main

import (
	"io/ioutil"

	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/html"
)

// Needs to traverse templates and minify
// Also traverse /public/assets and minify
func main() {

	content, err := ioutil.ReadFile("../templates/layout.html")

	m := minify.New()
	m.AddFunc("text/html", html.Minify)

	b, err := m.Bytes("text/html", content)
	if err != nil {
		panic(err)
	}

	ioutil.WriteFile("test.html", b, 0777)
}
