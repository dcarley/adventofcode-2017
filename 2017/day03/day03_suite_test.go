package day03_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestDay03(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Day03 Suite")
}
