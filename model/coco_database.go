package model

type Column struct {
	Name string
	Tp   string
}

type Table struct {
	Name    string
	Comment string
	Columns []Column
}


