package main

import (
	"context"
	"fmt"

	"cloud.google.com/go/bigquery"
)

func main() {
	_, err := bigquery.NewClient(context.Background(), "fake-project-id")
	if err != nil {
		// do nothing
	}
	fmt.Println("Hello, World!")
}
