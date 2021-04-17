package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
)

// mPart holds multipart data
type mPartData struct {
	source      io.Reader
	contentType string
}

func main() {
	if len(os.Args) < 2 {
		// No argument specified
		printHelp()
	}

	filename := os.Args[1]
	mdBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	partData, err := createMultipart(mdBytes)
	err = do(http.DefaultClient, partData)
	if err != nil {
		log.Fatal(err)
	}
}

// do makes a request to the server using the client
func do(client *http.Client, mPart *mPartData) error {
}

func createMultipart(filedata []byte) (*mPartData, error) {
	var b bytes.Buffer
	w := multipart.NewWriter(&b)

	fw, err := w.CreateFormFile("file", "file")
	if err != nil {
		return nil, err
	}

	fw.Write(filedata)
	w.Close()

	return &mPartData{&b, w.FormDataContentType()}, nil
}

func printHelp() {
	fmt.Println(strings.TrimSpace(usageString))
}

const usageString = `
Usage: mark2web FILE
Renders markdown FILE as webpage, returning a URL to the webpage.
`
