package main_test

import (
	"github.com/inherd/sqling/modeling/parser"
	. "github.com/onsi/gomega"
	"testing"
)

var humanSql = `CREATE TABLE human(
    id VARCHAR(12) PRIMARY KEY,
    sname VARCHAR(12),
    age INT,
    sex CHAR(1)
);`

func Test_StructLen(t *testing.T) {
	g := NewGomegaWithT(t)
	structs, _ := parser.ParseSql(humanSql)

	g.Expect(len(structs)).To(Equal(1))
}

func Test_FieldsLen(t *testing.T) {
	g := NewGomegaWithT(t)

	structs, _ := parser.ParseSql(humanSql)

	fields := structs[0].Fields
	g.Expect(len(fields)).To(Equal(4))
}

func Test_StructField(t *testing.T) {
	g := NewGomegaWithT(t)

	structs, _ := parser.ParseSql(humanSql)
	fields := structs[0].Fields

	g.Expect(fields[0].Name).To(Equal("id"))
	g.Expect(fields[0].FieldType).To(Equal("String"))
}

