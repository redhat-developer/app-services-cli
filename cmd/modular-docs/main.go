package main

import (
	"fmt"
	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/docs"
	"os"
)

func main() {
	err := docs.CreateModularDocs()
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
}
