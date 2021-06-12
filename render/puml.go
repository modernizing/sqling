package render

import (
	"bufio"
	"fmt"
	"github.com/iancoleman/strcase"
	. "github.com/inherd/sqling/model"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func WritePuml(structs []CocoStruct) {
	f, err := os.Create("sqling.puml")
	check(err)
	defer f.Close()
	w := bufio.NewWriter(f)

	_, err = fmt.Fprintln(w, "@startuml")
	check(err)

	for _, cocoStruct := range structs {
		_, err = fmt.Fprintln(w, "class "+strcase.ToCamel(cocoStruct.Name)+" {")

		for _, field := range cocoStruct.Fields {
			_, err = fmt.Fprintln(w, " - "+strcase.ToCamel(field.Name)+": "+field.FieldType)
		}

		_, err = fmt.Fprintln(w, "}")
	}

	_, err = fmt.Fprintln(w, "@enduml")

	w.Flush()
}
