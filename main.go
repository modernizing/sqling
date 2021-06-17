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

type database struct {
	tables []table
}

type table struct {
	name    string
	comment string
	columns []column
}

type column struct {
	name string
	typ  string
}

func (v *database) Enter(in ast.Node) (ast.Node, bool) {
	switch in.(type) {
	case *ast.CreateTableStmt:
		n := in.(*ast.CreateTableStmt)
		tableName := n.Table.Name.String()
		tabl := table{name: tableName}
		for _, col := range n.Cols {
			tabl.columns = append(tabl.columns, column{
				name: col.Name.String(),
				typ:  col.Tp.String(),
			})
		}
		for _, opt := range n.Options {
			switch opt.Tp {
			case ast.TableOptionComment:
				tabl.comment = opt.StrValue
			}
		}

		if n.Table.TableInfo != nil {
			fmt.Println(n.Table.TableInfo.Comment)
			tabl.comment = n.Table.TableInfo.Comment
		}

		v.tables = append(v.tables, tabl)
	}

	return in, false
}

func (v *database) Leave(in ast.Node) (ast.Node, bool) {
	return in, true
}

func extract(rootNode *ast.StmtNode, v *database) {
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

	v := &database{}
	for _, node := range *astNode {
		extract(&node, v)
	}

}
