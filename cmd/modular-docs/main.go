package main

import (
	"fmt"
	"github.com/redhat-developer/app-services-cli/internal/docs"
	"os"
)

func main() {
	err := docs.CreateModularDocs()
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
}
