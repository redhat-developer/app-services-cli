package main

import (
	"fmt"
	"os"

	"github.com/redhat-developer/app-services-cli/internal/docs"
)

func main() {
	err := docs.CreateModularDocs()
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
}
