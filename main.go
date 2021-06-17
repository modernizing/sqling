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

type colX struct {
	colNames []string
}

func (v *colX) Enter(in ast.Node) (ast.Node, bool) {
	switch in.(type) {
	case *ast.CreateTableStmt:
		n := in.(*ast.CreateTableStmt)
		tableName := n.Table.Name.String()
		for _, col := range n.Cols {
			fmt.Println(col)
		}
		fmt.Println(tableName)
	}

	return in, false
}

func (v *colX) Leave(in ast.Node) (ast.Node, bool) {
	return in, true
}

func extract(rootNode *ast.StmtNode, v *colX) {
	(*rootNode).Accept(v)
}

func parse(sql string) (*[]ast.StmtNode, error) {
	p := parser.New()

	stmtNodes, _, err := p.Parse(sql, "", "")
	if err != nil {
		return nil, err
	}

	return &stmtNodes, nil
}

func toStruct(sql string, structs []CocoStruct) {
	astNode, err := parse(sql)
	if err != nil {
		fmt.Printf("parse error: %v\n", err.Error())
		return
	}

	v := &colX{}
	for _, node := range *astNode {
		extract(&node, v)
	}

	fmt.Println(v)
}
