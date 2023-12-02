package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func WriteToFile(filename string, contents []byte) {
	err := os.MkdirAll(filepath.Dir(filename), os.ModePerm)
	if err != nil {
		log.Fatalf("making directory: %s", err)
	}
	err = os.WriteFile(filename, contents, os.FileMode(0644))
	if err != nil {
		log.Fatalf("writing file: %s", err)
	}

	log.Printf("Wrote to File: %s", filename)
}

func GetWithAOCCookie(url string, cookie string) []byte {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		log.Fatalf("making request: %s", err)
	}

	sessionCookie := http.Cookie{
		Name:  "session",
		Value: cookie,
	}

	req.AddCookie(&sessionCookie)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("making request: %s", err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("reading response body: %s", err)
	}
	fmt.Println("response lenght is", len(body))

	if strings.HasPrefix(string(body), "Please don't repeatedly") {
		log.Fatalf("repeated requests to advent-of-code website\n")
	}

	return body
}
