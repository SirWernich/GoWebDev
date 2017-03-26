package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	name := "Bob Smith"

	tmplt := `
        <!DOCTYPE html>
        <html lang="en">
        <head>
        <meta charset="UTF-8">
        <title>Hello World!</title>
        </head>
        <body>
        <h1> ` + name + `</h1>
        </body>
        </html>
        `

	// try to create the file
	f, err := os.Create("index.html")

	// if there was an error, log it and give up
	if err != nil {
		log.Fatal("Error creating file:", err)
	}
	// close the file when done, if it was successfully created.
	defer f.Close()

	// io.Copy takes a writer(dst) and a reader (src), so you
	// create a Reader from your string
	io.Copy(f, strings.NewReader(tmplt))

}
