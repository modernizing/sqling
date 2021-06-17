package main

import (
	"fmt"
	//. "github.com/inherd/sqling/converter"
	. "github.com/inherd/sqling/model"
	. "github.com/inherd/sqling/render"
	"github.com/pingcap/parser"
	"github.com/pingcap/parser/ast"
	_ "github.com/pingcap/tidb/types/parser_driver"
	"io/ioutil"
)

func main() {
	//args := os.Args
	//if len(args) < 2 {
	//	return
	//}
	//
	//filename := args[1]
	//
	dat, err := ioutil.ReadFile("_fixtures/mall.sql")
	Check(err)
	sql := string(dat)

	var structs []CocoStruct
	toStruct(sql, structs)

	// Write(structs)
}

type colX struct{
	colNames []string
}

func (v *colX) Enter(in ast.Node) (ast.Node, bool) {
	if name, ok := in.(*ast.ColumnName); ok {
		fmt.Println(name)
		v.colNames = append(v.colNames, name.Name.O)
	}

	fmt.Println(in.Text())
	return in, false
}

func (v *colX) Leave(in ast.Node) (ast.Node, bool) {
	return in, true
}

func extract(rootNode *ast.StmtNode) []string {
	v := &colX{}
	(*rootNode).Accept(v)
	return v.colNames
}


func parse(sql string) (*ast.StmtNode, error) {
	p := parser.New()

	stmtNodes, _, err := p.Parse(sql, "", "")
	if err != nil {
		return nil, err
	}

	return &stmtNodes[0], nil
}


func toStruct(sql string, structs []CocoStruct) {
	astNode, err := parse(sql)
	if err != nil {
		fmt.Printf("parse error: %v\n", err.Error())
		return
	}

	extract(astNode)
}

//
//func parseSql(sql string, structs []CocoStruct) []CocoStruct {
//	stmt, err := sqlparser.Parse(sql)
//	if err != nil {
//		fmt.Println(err)
//	}
//
//	switch stmt := stmt.(type) {
//	case *sqlparser.DDL:
//		switch stmt.Action {
//		case "create":
//			var fields []CocoField
//			for _, column := range stmt.TableSpec.Columns {
//				fields = append(fields, CocoField{
//					Name:      column.Name.String(),
//					FieldType: FromMysqlType(column.Type.Type),
//				})
//			}
//
//			cocoStruct := CocoStruct{
//				Name:   stmt.NewName.Name.String(),
//				Fields: fields,
//			}
//
//			structs = append(structs, cocoStruct)
//		}
//	}
//	return structs
//}
