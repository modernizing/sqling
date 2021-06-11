package main

import (
	"bufio"
	"fmt"
	"github.com/iancoleman/strcase"
	"github.com/xwb1989/sqlparser"
	"io/ioutil"
	"os"
)

type CocoStruct struct {
	name   string
	fields []CocoField
}

type CocoField struct {
	name  string
	fType string
}

func mysqlTypeToJava(typ string) string {
	switch typ {
	case "bit":
		return "Boolean"
	case "byte":
		return " Byte"
	case "short":
		return " Short"
	case "int":
		return "Integer"
	case "smallint", "tinyint":
		return "Integer"
	case "bigint":
		return "Long"
	case "float":
		return "Float"
	case "double":
		return "Double"
	case "decimal", "numeric":
		return "BigDecimal"
	case "date":
		return "Date"
	case "datetime", "timestamp":
		return "Timestamp"
	case "time":
		return "Time"
	case "year":
		return "Short"
	case "varchar", "char", "text":
		return "String"
	}
	return typ
}

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
					fType: mysqlTypeToJava(column.Type.Type),
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

func writePuml(structs []CocoStruct) {
	f, err := os.Create("sqling.puml")
	check(err)
	defer f.Close()
	w := bufio.NewWriter(f)

	_, err = fmt.Fprintln(w, "@startuml")
	check(err)

	for _, cocoStruct := range structs {
		_, err = fmt.Fprintln(w, "class "+strcase.ToCamel(cocoStruct.name)+" {")

		for _, field := range cocoStruct.fields {
			_, err = fmt.Fprintln(w, " - "+strcase.ToCamel(field.name)+": "+field.fType)
		}

		_, err = fmt.Fprintln(w, "}")
	}

	_, err = fmt.Fprintln(w, "@enduml")

	w.Flush()
}
