package day02_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestDay02(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Day02 Suite")
}
