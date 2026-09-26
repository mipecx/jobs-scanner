package main

import (
	"context"
	"fmt"
	"io"
	"os"

	"jobs_scanner/internal/source"
	"jobs_scanner/internal/source/habr"
	"jobs_scanner/internal/source/linkedin"
)

func main() {
	var websites []source.Fetcher
	websites = append(websites, &habr.HabrParser{ByteHTML: byteFromFile("./probe.html")})
	websites = append(websites, &linkedin.LinkedinParser{ByteHTML: byteFromFile("./probe2.html")})

	for _, w := range websites {
		vacancies, err := w.Fetch(context.Background())
		if err != nil {
			fmt.Printf("can not fetch from habr: %s", err)
			return
		}

		for _, v := range vacancies {
			fmt.Println(v)
		}
	}
}

func byteFromFile(path string) []byte {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("error opening file: %v", err)
		return nil
	}
	defer file.Close()

	docHTML, err := io.ReadAll(file)
	if err != nil {
		fmt.Printf("error reading the file: %v", err)
		return nil
	}

	return docHTML
}
