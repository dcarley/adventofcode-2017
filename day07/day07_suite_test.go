package day07_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestDay07(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Day07 Suite")
}
