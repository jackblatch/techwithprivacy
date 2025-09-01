package main

import (
	"context"
	"log"
	"os"
	"techwithprivacy/web/routes"
)

func main() {
	f, err := os.Create("index.html")
	if err != nil {
		log.Fatalf("failed to create output file: %v", err)
	}

	page, err := routes.GetIndex()

	if err != nil {
		log.Fatalf("failed to get index page: %v", err)
	}

	err = page.Render(context.Background(), f)
	if err != nil {
		log.Fatalf("failed to write output file: %v", err)
	}

}
