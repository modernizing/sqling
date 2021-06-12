package main

import (
	"bufio"
	"fmt"
	"github.com/iancoleman/strcase"
	"os"
)

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
