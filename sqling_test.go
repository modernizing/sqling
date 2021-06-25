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

var humanCarSql = `CREATE TABLE human(
    id VARCHAR(12) PRIMARY KEY,
    sname VARCHAR(12),
    age INT,
    sex CHAR(1)
);

CREATE TABLE car(
    id VARCHAR(12) PRIMARY KEY,
    mark VARCHAR(24),
    price NUMERIC(6,2),
    hid VARCHAR(12),
    CONSTRAINT fk_human FOREIGN KEY(hid) REFERENCES human(id)
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

func Test_MultipleStructs(t *testing.T) {
	g := NewGomegaWithT(t)

	structs, _ := parser.ParseSql(humanCarSql)

	g.Expect(len(structs)).To(Equal(2))
}

func Test_StructRefLen(t *testing.T) {
	g := NewGomegaWithT(t)

	_, refs := parser.ParseSql(humanCarSql)

	g.Expect(len(refs)).To(Equal(1))
}

