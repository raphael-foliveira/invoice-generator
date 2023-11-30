package main

import (
	"fmt"
	"log"
	"os"

	"github.com/raphael-foliveira/invoice-generator/internal/cfg"
	"github.com/raphael-foliveira/invoice-generator/internal/crawler"
)

func main() {
	if len(os.Args) == 1 {
		log.Fatalf("Download path must be provided")
		return
	}
	id, err := cfg.ReadInvoiceData()
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting crawler")
	c := crawler.NewCrawler(id, os.Args[1])
	if err := c.Run(); err != nil {
		panic(err)
	}
	fmt.Println("Done")
}
