package day04_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestDay04(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "2017 Day04 Suite")
}
