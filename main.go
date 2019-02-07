package main

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

func main() {
	data := [][]string{
		[]string{"", "5.1.z", "2.0", "", ""},
		[]string{"", "5.2.z", "3.0", "", ""},
		[]string{"Anand", "5.3.z", "3.1", "", ""},
		[]string{"Botvinnik", "5.4.z", "3.2", "", ""},
		[]string{"Capablanca", "5.5.z", "4.0", "2.2.x", "9.4.x"},
		[]string{"Darga", "5.6.z", "4.1", "2.2.x", "9.4.x"},
		[]string{"Euwe", "5.7.z", "4.2", "2.3.x", "9.5.x"},
		[]string{"Fine", "5.8.z", "4.5", "2.3.x", "9.5.x"},
		[]string{"Gaprindashvili", "5.9.z", "4.6", "2.3.x", "9.5.x", "2.4.y."},
		[]string{"Hammer", "5.10.z", "4.7", "2.4.x", "9.5.x", "2.7.y"},
		[]string{"Ivanchuck", "5.11.z", "4.8", "2.4.x", "9.5.x", "2.?"},
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ManageIQ", "CloudForms Management Engine", "CloudForms", "Ruby", "PostgreSQL", "Ansible"})

	for _, v := range data {
		table.Append(v)
	}
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.Render()
}
