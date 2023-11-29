package main

import (
	"fmt"

	"github.com/raphael-foliveira/invgen/internal/crawler"
)

func main() {
	if err := crawler.Run(); err != nil {
		panic(err)
	}
	fmt.Println("Done")
}
