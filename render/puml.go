package render

import (
	"bufio"
	"fmt"
	"github.com/iancoleman/strcase"
	. "github.com/inherd/sqling/model"
	"os"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func OutputPuml(structs []CocoStruct, refs []CocoRef) {
	fileName := "sqling.puml"
	f, err := os.Create(fileName)
	Check(err)
	defer f.Close()
	w := bufio.NewWriter(f)

	fmt.Fprintln(w, "@startuml")
	for _, cocoStruct := range structs {
		fmt.Fprintln(w, "class "+strcase.ToCamel(cocoStruct.Name)+" {")

		for _, field := range cocoStruct.Fields {
			fmt.Fprintln(w, " - "+strcase.ToCamel(field.Name)+": "+field.FieldType)
		}

		fmt.Fprintln(w, "}")
	}

	for _, ref := range refs {
		fmt.Fprintln(w, strcase.ToCamel(ref.Source)+" --> "+strcase.ToCamel(ref.Target))
	}

	fmt.Fprintln(w, "@enduml")

	w.Flush()

	fmt.Println("write to file: " + fileName)
}
