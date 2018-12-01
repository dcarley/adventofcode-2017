package day09_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestDay09(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Day09 Suite")
}
