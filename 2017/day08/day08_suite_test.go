package day08_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestDay08(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Day08 Suite")
}
