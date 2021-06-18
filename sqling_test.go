package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/inherd/sqling/modeling/parser"
)

var _ = Describe("Sqling", func() {
	Context("With more than 300 pages", func() {
		It("should be a novel", func() {
			PlSql("create table employees_test (employee_id number primary key, commission_pct number, salary number);")
			Expect("NOVEL").To(Equal("NOVEL"))
		})
	})
})
