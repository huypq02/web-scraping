package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func main() {
	url := os.Args[2]
	cmd := os.Args[1]
	// we are interested in getting the file or object name
	// so take the last item from the slice
	var fileName string
	if cmd == "-u" {
		subStringsSlice := strings.Split(url, "/")
		fileName = subStringsSlice[len(subStringsSlice)-1]
	} else if cmd == "-l" {
		log.Fatal("Developing ...")
	} else {
		log.Fatal("Out of scope")
	}

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	// Size of Download file
	new, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer new.Close()

	// Is our request ok?
	resp, _ := http.Head(url)

	if resp.StatusCode != http.StatusOK {
		fmt.Println(resp.Status)
		os.Exit(1)
		// exit if not ok
	}

	size, _ := strconv.Atoi(resp.Header.Get("Content-Length"))
	downloadSize := int64(size)
	fmt.Println(downloadSize)

	n, err := io.Copy(new, response.Body)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Number of bytes copied to STDOUT:", n)
}
