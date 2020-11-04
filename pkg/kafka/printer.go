package kafka

import (
	"os"

	"github.com/landoop/tableprinter"
)

// PrintToTable prints the clusters in a formatted table
func PrintToTable(clusters []Cluster) {
	printer := tableprinter.New(os.Stdout)
	printer.Print(clusters)
}
