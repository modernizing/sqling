package main

import (
	"fmt"
	. "github.com/inherd/sqling/converter"
	. "github.com/inherd/sqling/model"
	. "github.com/inherd/sqling/render"
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
					Name:  column.Name.String(),
					FieldType: FromMysqlType(column.Type.Type),
				})
			}

			cocoStruct := CocoStruct{
				Name:   stmt.NewName.Name.String(),
				Fields: fields,
			}

			structs = append(structs, cocoStruct)
		}
	}

	WritePuml(structs)
}
