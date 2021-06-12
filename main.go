package main

import (
	"fmt"
	"github.com/xwb1989/sqlparser"
	"io/ioutil"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := ioutil.ReadFile("mall.sql")
	check(err)
	sql := string(dat)

	stmt, err := sqlparser.Parse(sql)
	if err != nil {
		// Do something with the err
		fmt.Println(err)
	}

	// Otherwise do something with stmt
	var structs []CocoStruct
	switch stmt := stmt.(type) {
	case *sqlparser.DDL:
		switch stmt.Action {
		case "create":
			var fields []CocoField
			for _, column := range stmt.TableSpec.Columns {
				fields = append(fields, CocoField{
					name:  column.Name.String(),
					fType: fromMysqlType(column.Type.Type),
				})
			}

			cocoStruct := CocoStruct{
				name:   stmt.NewName.Name.String(),
				fields: fields,
			}

			structs = append(structs, cocoStruct)
		}
	}

	writePuml(structs)
}

