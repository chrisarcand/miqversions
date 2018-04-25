package main

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

func main() {
	data := [][]string{
		[]string{"", "5.1.z", "2.0"},
		[]string{"", "5.2.z", "3.0"},
		[]string{"Anand", "5.3.z", "3.1"},
		[]string{"Botvinnik", "5.4.z", "3.2"},
		[]string{"Capablanca", "5.5.z", "4.0"},
		[]string{"Darga", "5.6.z", "4.1"},
		[]string{"Euwe", "5.7.z", "4.2"},
		[]string{"Fine", "5.8.z", "4.5"},
		[]string{"Gaprindashvili", "5.9.z", "4.6"},
		[]string{"Gaprindashvili", "5.10.z", "4.7"},
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ManageIQ", "CloudForms Management Engine", "CloudForms"})

	for _, v := range data {
		table.Append(v)
	}
	table.Render()
}
