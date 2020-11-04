package kafka

import (
	"os"
	"github.com/landoop/tableprinter"
)

// PrintInstances prints the instances in a formatted table
func PrintInstances(kafkaInstances []Instance) {
	printer := tableprinter.New(os.Stdout)
	printer.Print(kafkaInstances)
}