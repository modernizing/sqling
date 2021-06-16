package main

import (
	"fmt"
	. "github.com/inherd/sqling/converter"
	. "github.com/inherd/sqling/model"
	. "github.com/inherd/sqling/render"
	"github.com/xwb1989/sqlparser"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 1 {
		return
	}

	dat, err := ioutil.ReadFile(args[1])
	Check(err)
	sql := string(dat)

	stmt, err := sqlparser.Parse(sql)
	if err != nil {
		fmt.Println(err)
	}

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

	Write(structs)
}
