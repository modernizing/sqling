package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Sqling", func() {
	Context("With more than 300 pages", func() {
		It("should be a novel", func() {
			Expect("NOVEL").To(Equal("NOVEL"))
		})
	})
})
