package main

import (
	"fmt"
	"github.com/inherd/sqling/converter"
	"github.com/pingcap/parser/mysql"
	"github.com/pingcap/parser/types"

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
	dat, err := ioutil.ReadFile("_fixtures/platform.sql")
	Check(err)
	sql := string(dat)

	Convert(sql)
}

func Convert(sql string) {
	var structs []CocoStruct
	var refs []CocoRef
	toStruct(sql, &structs, &refs)

	Write(structs, refs)
}

type Database struct {
	Tables []Table
	Refs   []CocoRef
}

type Table struct {
	Name    string
	Comment string
	Columns []Column
}

type Column struct {
	Name string
	Tp   string
}

func (v *Database) Enter(in ast.Node) (ast.Node, bool) {
	switch in.(type) {
	case *ast.CreateTableStmt:
		n := in.(*ast.CreateTableStmt)
		tableName := n.Table.Name.String()
		tabl := Table{Name: tableName}
		for _, col := range n.Cols {
			tabl.Columns = append(tabl.Columns, Column{
				Name: col.Name.String(),
				Tp:   v.getType(col.Tp),
			})
		}

		for _, constraint := range n.Constraints {
			if constraint.Refer != nil {
				target := constraint.Refer.Table.Name.String()
				v.Refs = append(v.Refs, CocoRef{
					Source: tableName,
					Target: target,
				})
			}
		}

		for _, opt := range n.Options {
			switch opt.Tp {
			case ast.TableOptionComment:
				tabl.Comment = opt.StrValue
			}
		}

		if n.Table.TableInfo != nil {
			tabl.Comment = n.Table.TableInfo.Comment
		}

		v.Tables = append(v.Tables, tabl)
	}

	return in, false
}

func (v *Database) getType(ft *types.FieldType) string {
	switch ft.Tp {
	case mysql.TypeTiny, mysql.TypeShort, mysql.TypeInt24, mysql.TypeLong, mysql.TypeLonglong,
		mysql.TypeBit, mysql.TypeYear:
		return "int"
	case mysql.TypeFloat, mysql.TypeDouble:
		return "float"
	case mysql.TypeNewDecimal:
		return "decimal"
	case mysql.TypeDate, mysql.TypeDatetime:
		return "datetime"
	case mysql.TypeTimestamp:
		return "timestamp"
	case mysql.TypeDuration:
		return "duration"
	case mysql.TypeJSON:
		return "json"
	case mysql.TypeVarchar, mysql.TypeString:
		return "varchar"
	}

	return "text"
}

func (v *Database) Leave(in ast.Node) (ast.Node, bool) {
	return in, true
}

func extract(rootNode *ast.StmtNode, v *Database) {
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

func toStruct(sql string, structs *[]CocoStruct, refs *[]CocoRef) {
	astNode, err := parse(sql)
	if err != nil {
		fmt.Printf("parse error: %v\n", err.Error())
		return
	}

	v := &Database{}
	for _, node := range *astNode {
		extract(&node, v)
	}

	*refs = v.Refs

	for _, tab := range v.Tables {
		coco := CocoStruct{
			Name:   tab.Name,
			Comment: tab.Comment,
			Fields: nil,
		}

		for _, col := range tab.Columns {
			coco.Fields = append(coco.Fields, CocoField{
				Name:      col.Name,
				FieldType: converter.FromMysqlType(col.Tp),
			})
		}

		*structs = append(*structs, coco)
	}
}
