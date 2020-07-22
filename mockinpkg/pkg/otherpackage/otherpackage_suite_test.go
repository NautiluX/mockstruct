package mypackage_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestMypackage(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Mypackage Suite")
}
