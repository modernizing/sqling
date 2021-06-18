package render

import (
	"encoding/json"
	"fmt"
	"github.com/inherd/sqling/modeling/model"
	"io/ioutil"
)


type SqlingJson struct {
	Structs []model.CocoStruct
	Refs    []model.CocoRef
}

func OutputJson(structs []model.CocoStruct, refs []model.CocoRef) {
	filename := "sqling.json"

	output := &SqlingJson{
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
