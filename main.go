package main

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

func main() {
	data := [][]string{
		[]string{"", "5.1.x", "2.0"},
		[]string{"", "5.2.x", "3.0"},
		[]string{"Anand", "5.3.x", "3.1"},
		[]string{"Botvinnik", "5.4.x", "3.2"},
		[]string{"Capablanca", "5.5.x", "4.0"},
		[]string{"Darga", "5.6.x", "4.1"},
		[]string{"Euwe", "5.7.x", "4.2"},
		[]string{"Fine", "5.8.x (pending)", "4.5 (pending)"},
		[]string{"Gaprindashvili", "5.9.x (pending)", "TBD (pending)"},
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ManageIQ", "CloudForms Management Engine", "CloudForms"})

	for _, v := range data {
		table.Append(v)
	}
	table.Render()
}
