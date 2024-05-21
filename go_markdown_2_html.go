package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gomarkdown/markdown"
)

func main(){
	fmt.Println("Hi")
	file := "test.md"
	content, err := os.ReadFile(file)

	if err != nil {
		log.Fatalf("%s file not found", file)
	}

	html := markdown.ToHTML(content, nil, nil)

	fmt.Println(string(html))

	fileOut := "test.html"
	errOut := os.WriteFile(fileOut, html, 0644)
	if errOut != nil {
		log.Fatalf("Couldn't create the file %s ", fileOut)
	}

	fmt.Printf("HTML file called %s created", fileOut)
}