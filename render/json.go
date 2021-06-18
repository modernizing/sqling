package render

import (
	"encoding/json"
	"fmt"
	"github.com/inherd/sqling/cmd"
	. "github.com/inherd/sqling/model"
	"io/ioutil"
)

func OutputJson(structs []CocoStruct, refs []CocoRef) {
	filename := "sqling.json"

	output := &cmd.SqlingJson{
		Structs: structs,
		Refs:    refs,
	}

	str, err := json.MarshalIndent(output, "", "\t")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(str))

	err = ioutil.WriteFile(filename, str, 0644)
	return
}
