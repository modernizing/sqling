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
	dat, err := ioutil.ReadFile("_fixtures/constraint.sql")
	Check(err)
	sql := string(dat)

	var structs []CocoStruct
	var refs    []CocoRef
	toStruct(sql, &structs, &refs)

	Write(structs, refs)
}

type database struct {
	tables []table
	refs   []CocoRef
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
				typ:  v.getType(col.Tp),
			})
		}

		for _, constraint := range n.Constraints {
			if constraint.Refer != nil {
				target := constraint.Refer.Table.Name.String()
				v.refs = append(v.refs, CocoRef{
					Source: tableName,
					Target: target,
				})
			}
		}

		for _, opt := range n.Options {
			switch opt.Tp {
			case ast.TableOptionComment:
				tabl.comment = opt.StrValue
			}
		}

		if n.Table.TableInfo != nil {
			tabl.comment = n.Table.TableInfo.Comment
		}

		v.tables = append(v.tables, tabl)
	}

	return in, false
}

func (v *database) getType(ft *types.FieldType) string {
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
	//default:
	//	fmt.Println(ft.Tp)
	}

	return "text"
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

func toStruct(sql string, structs *[]CocoStruct, refs *[]CocoRef) {
	astNode, err := parse(sql)
	if err != nil {
		fmt.Printf("parse error: %v\n", err.Error())
		return
	}

	v := &database{}
	for _, node := range *astNode {
		extract(&node, v)
	}

	*refs = v.refs

	for _, tab := range v.tables {
		coco := CocoStruct{
			Name:   tab.name,
			Comment: tab.comment,
			Fields: nil,
		}

		for _, col := range tab.columns {
			coco.Fields = append(coco.Fields, CocoField{
				Name:      col.name,
				FieldType: converter.FromMysqlType(col.typ),
			})
		}

		*structs = append(*structs, coco)
	}
}
