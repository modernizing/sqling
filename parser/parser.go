package parser

import (
	"fmt"
	"github.com/inherd/sqling/converter"
	"github.com/inherd/sqling/model"
	"github.com/pingcap/parser"
	"github.com/pingcap/parser/ast"
	"github.com/pingcap/parser/mysql"
	"github.com/pingcap/parser/types"
)

type Database struct {
	Tables []model.Table
	Refs   []model.CocoRef
}

func (v *Database) Enter(in ast.Node) (ast.Node, bool) {
	switch in.(type) {
	case *ast.CreateTableStmt:
		n := in.(*ast.CreateTableStmt)
		tableName := n.Table.Name.String()
		tabl := model.Table{Name: tableName}
		for _, col := range n.Cols {
			tabl.Columns = append(tabl.Columns, model.Column{
				Name: col.Name.String(),
				Tp:   v.toMysqlType(col.Tp),
			})
		}

		for _, constraint := range n.Constraints {
			if constraint.Refer != nil {
				target := constraint.Refer.Table.Name.String()
				v.Refs = append(v.Refs, model.CocoRef{
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

func (v *Database) toMysqlType(ft *types.FieldType) string {
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

func parseString(sql string) (*[]ast.StmtNode, error) {
	p := parser.New()

	stmtNodes, _, err := p.Parse(sql, "", "")
	if err != nil {
		return nil, err
	}

	return &stmtNodes, nil
}

func ParseSql(sql string) ([]model.CocoStruct, []model.CocoRef) {
	astNode, err := parseString(sql)
	if err != nil {
		fmt.Printf("parse error: %v\n", err.Error())
		return nil, nil
	}

	v := &Database{}
	for _, node := range *astNode {
		(node).Accept(v)
	}

	return toCocoStructs(v)
}

func toCocoStructs(v *Database) ([]model.CocoStruct, []model.CocoRef) {
	var structs []model.CocoStruct
	var refs []model.CocoRef

	refs = v.Refs

	for _, tab := range v.Tables {
		coco := model.CocoStruct{
			Name:    tab.Name,
			Comment: tab.Comment,
			Fields:  nil,
		}

		for _, col := range tab.Columns {
			coco.Fields = append(coco.Fields, model.CocoField{
				Name:      col.Name,
				FieldType: converter.FromMysqlType(col.Tp),
			})
		}

		structs = append(structs, coco)
	}

	return structs, refs
}
