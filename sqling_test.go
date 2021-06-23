package main_test

import (
	"github.com/inherd/sqling/modeling/parser"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Sqling", func() {
	humanSql := `CREATE TABLE human(
    id VARCHAR(12) PRIMARY KEY,
    sname VARCHAR(12),
    age INT,
    sex CHAR(1)
);`

	Context("Structs generate", func() {
		It("structs len & field len", func() {
			Expect("NOVEL").To(Equal("NOVEL"))
			structs, _ := parser.ParseSql(humanSql)
			Expect(len(structs)).To(Equal(1))
			fields := structs[0].Fields
			Expect(len(fields)).To(Equal(4))
		})

		It("field content", func() {
			Expect("NOVEL").To(Equal("NOVEL"))
			structs, _ := parser.ParseSql(humanSql)
			fields := structs[0].Fields

			Expect(fields[0].Name).To(Equal("id"))
			Expect(fields[0].FieldType).To(Equal("String"))
		})
	})
})
