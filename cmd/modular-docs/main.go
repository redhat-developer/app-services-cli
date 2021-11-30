package main

import (
	"fmt"
	"github.com/redhat-developer/app-services-cli/internal/doc"
	"os"
)

func main() {
	err := doc.CreateModularDocs()
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
}
