package main

import "fmt"

type Table struct {
	Rows []Row
	Name string
	ColumnNames []string
}

type Row struct {
	Columns []Column
	Id int
}

// Column Class
type Column struct {
	Id int
	Value string
}

func printTable(table Table)  {
	var rows []Row = table.Rows
	fmt.Println(table.Name)
	for _, row := range rows{
		var columns []Column = row.Columns
		for i,column := range columns {
			fmt.Println(table.ColumnNames[i],column.Id,column.Value)
		}
	}
}

func main()  {
	var table = Table{}
	table.Name = "Customer"
	table.ColumnNames = []string{"Id", "Name","SSN"}
	var rows []Row = make([]Row,3)
}