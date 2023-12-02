package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

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
