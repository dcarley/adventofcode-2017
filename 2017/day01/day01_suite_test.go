package day01_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestDay01(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Day01 Suite")
}