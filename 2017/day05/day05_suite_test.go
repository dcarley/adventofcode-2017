package day05_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestDay05(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Day05 Suite")
}
