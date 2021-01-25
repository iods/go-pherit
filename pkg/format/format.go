package format

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"reflect"
)

func render(rows [][]string) []byte {

	b := bytes.NewBuffer(nil)

	// set some table options
	table := tablewriter.NewWriter(b)
	table.SetHeader([]string{"FlagSet", "Name", "Value", "Type", "Default", "Usage"})
	table.SetAutoWrapText(false)
	table.SetAutoFormatHeaders(true)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetHeaderLine(true)
	table.SetBorder(true)
	table.SetTablePadding("\t")

	for _, value := range rows {
		table.Append(value)
	}

	table.Render()

	return b.Bytes()
}

// RenderFlagSet Renders a flag.FlagSet as a table.
func RenderFlagSet(sets ...*flag.FlagSet) {

	var rows [][]string

	for _, 	set := range sets {

		set.VisitAll(func(f *flag.Flag) {
			rows = append(rows, []string{
				"",
				f.Name,
				fmt.Sprintf("%v", f.Value),
				reflect.TypeOf(f.Value).String(),
				f.DefValue,
				f.Usage,
			})
		})
	}
	fmt.Printf("%s\n", render(rows))
}