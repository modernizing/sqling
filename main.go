package main

import (
	"fmt"
	"github.com/pingcap/parser"
	"github.com/pingcap/parser/ast"
	_ "github.com/pingcap/parser/test_driver"
)

func parse(sql string) (*ast.StmtNode, error) {
	p := parser.New()

	stmtNodes, _, err := p.Parse(sql, "", "")
	if err != nil {
		return nil, err
	}

	return &stmtNodes[0], nil
}

type Visitor interface {
	Enter(n ast.Node) (node ast.Node, skipChildren bool)
	Leave(n ast.Node) (node ast.Node, ok bool)
}

type colX struct{
	tabName string
	colNames []string
}

func (v *colX) Enter(in ast.Node) (ast.Node, bool) {
	if name, ok := in.(*ast.ColumnName); ok {
		v.colNames = append(v.colNames, name.Name.O)
	}

	if name, ok := in.(*ast.TableName); ok {
		v.tabName = name.Name.String()
	}

	return in, false
}

func (v *colX) Leave(in ast.Node) (ast.Node, bool) {
	return in, true
}

func extract(rootNode *ast.StmtNode) *colX {
	v := &colX{}
	(*rootNode).Accept(v)
	return v
}

func main() {
	astNode, err := parse("CREATE TABLE `cms_help` (\n  `id` bigint(20) NOT NULL AUTO_INCREMENT,\n  `category_id` bigint(20) DEFAULT NULL,\n  `icon` varchar(500) DEFAULT NULL,\n  `title` varchar(100) DEFAULT NULL,\n  `show_status` int(1) DEFAULT NULL,\n  `create_time` datetime DEFAULT NULL,\n  `read_count` int(1) DEFAULT NULL,\n  `content` text,\n  PRIMARY KEY (`id`)\n) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='帮助表';\n")
	if err != nil {
		fmt.Printf("parse error: %v\n", err.Error())
		return
	}

	fmt.Printf("%v\n", extract(astNode))

}