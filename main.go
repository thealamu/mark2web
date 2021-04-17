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

const serverURL = "http://localhost:8080"

// mPart holds multipart data
type mPartData struct {
	source      io.Reader
	contentType string
}

func main() {
	if len(os.Args) < 2 {
		// No argument specified
		printHelp()
		return
	}

	filename := os.Args[1]
	mdBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	partData, err := createMultipart(mdBytes)
	if err != nil {
		log.Fatal(err)
	}

	url, err := getURLForMarkdown(http.DefaultClient, partData)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(url)
}

// getURLForMarkdown returns a URL for the markdown
func getURLForMarkdown(client *http.Client, mPart *mPartData) (string, error) {
	req, err := http.NewRequest(http.MethodPost, serverURL, mPart.source)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", mPart.contentType)
	// make the request
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	urlBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return formatURL(string(urlBytes)), nil
}

// formatURL appends the URL path to the serverURL
func formatURL(url string) string {
	id := url[strings.LastIndex(url, "/")+1:]
	return fmt.Sprintf("%s/%s", serverURL, id)
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
