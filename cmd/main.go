package main

import (
	"fmt"
	"os"

	"github.com/raphael-foliveira/invgen/internal/crawler"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Download path must be provided")
		return
	}
	if err := crawler.Run(); err != nil {
		panic(err)
	}
	fmt.Println("Done")
}
