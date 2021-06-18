package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSqling(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Sqling Suite")
}
