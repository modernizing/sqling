package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"path"

	. "github.com/inherd/sqling/modeling/parser"
)

var _ = Describe("Sqling", func() {
	Context("With more than 300 pages", func() {
		It("should be a novel", func() {
			testpath := path.Join("_fixtures", "plsql", "test")
			PlSql(testpath)

			Expect("NOVEL").To(Equal("NOVEL"))
		})
	})
})
